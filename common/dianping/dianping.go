package dianping

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	config2 "studyRoomGo/common/config"
	"studyRoomGo/common/dianping/gen"
	"time"

	"github.com/l306287405/dianping"
)

const authCallback = ""

const (
	ERR_1000 = "返回数据错误"

	ERR_1001 = "验券错误，请联系管理员"
	ERR_1006 = "此券号不存在，请检查"
	ERR_1008 = "已使用过的消费券，请检查"
	ERR_1009 = "团购卡未绑定套餐，请联系客服"
)

var config = dianping.Config{
	AppKey: "", //配置美团北极星key，secret
	Secret: "",
}

var sdkService, _ = dianping.NewService(config, nil)

// CheckAuth 检查授权情况
func CheckAuth(appKey string) (err error) {
	var data gen.DianpingAuth
	mgr := gen.DianpingAuthMgr(config2.AppConf.Data.DB)
	data, e := mgr.FetchUniqueByDianpingAuthAppKeyUIndex(appKey)
	if e != nil {
		return e
	}

	query := dianping.NewReqParams()
	query.Set("state", data.AppKey)
	//query.Set("redirect_url", authCallback)
	query.Set("scope", "[\"tuangou\"]")
	url := sdkService.MerchantAuth(query)
	fmt.Println("大众点评团购核销授权：" + url)
	if isRefreshAuthCode(data) {
		err = errors.New(fmt.Sprint("需要重新授权：" + url))
		return err
	}
	if isRefreshAccessToken(data) {
		RefreshAccessToken(data.AppKey)
	}
	return nil
}

// SaveAuthCode 获取authCode
func SaveAuthCode(appKey string, authCode string) (err error) {
	var data gen.DianpingAuth
	mgr := gen.DianpingAuthMgr(config2.AppConf.Data.DB)
	data, e := mgr.FetchUniqueByDianpingAuthAppKeyUIndex(appKey)
	if e != nil {
		return e
	}
	data.AuthCode = authCode
	mgr.Save(&data)
	return nil
}

// SaveAccessToken 获取accessToken
func SaveAccessToken(appKey string) error {
	var data gen.DianpingAuth
	mgr := gen.DianpingAuthMgr(config2.AppConf.Data.DB)
	data, e := mgr.FetchUniqueByDianpingAuthAppKeyUIndex(appKey)
	if e != nil {
		return e
	}

	query := dianping.NewReqParams()
	query.AddKeyAndSecret(&config)
	query.Set("auth_code", data.AuthCode)
	resp, err := sdkService.OauthTokenByCode(query)
	if err != nil {
		return err
	}
	expiresAt := time.Now().Add(time.Second * time.Duration(resp.ExpiresIn))

	data.Bid = resp.Bid
	data.AccessToken = resp.AccessToken
	data.RefreshToken = resp.RefreshToken
	data.ExpiresIn = resp.ExpiresIn
	data.ExpiresAt = &expiresAt
	data.RemainRefreshCount = resp.RemainRefreshCount
	mgr.Save(&data)

	return nil
}

// RefreshAccessToken 刷新accessToken
func RefreshAccessToken(appKey string) error {
	var data gen.DianpingAuth
	mgr := gen.DianpingAuthMgr(config2.AppConf.Data.DB)
	data, e := mgr.FetchUniqueByDianpingAuthAppKeyUIndex(appKey)
	if e != nil {
		return e
	}
	query := dianping.NewReqParams()
	query.AddKeyAndSecret(&config)
	query.Set("refresh_token", data.RefreshToken)
	resp, err := sdkService.OauthTokenByRefresh(query)
	if err != nil {
		return err
	}

	expiresAt := time.Now().Add(time.Second * time.Duration(resp.ExpiresIn))

	data.Bid = resp.Bid
	data.AccessToken = resp.AccessToken
	data.RefreshToken = resp.RefreshToken
	data.ExpiresIn = resp.ExpiresIn
	data.ExpiresAt = &expiresAt
	data.RemainRefreshCount = resp.RemainRefreshCount
	mgr.Save(&data)

	return nil
}

// 检查是否需要刷新accessToken
func isRefreshAccessToken(data gen.DianpingAuth) bool {
	refreshAt := time.Now().Add(time.Hour * 24 * 3) //距离过期时间3天前提示刷新
	if refreshAt.After(*data.ExpiresAt) {
		return true
	}
	return false
}

// 检查是否需要重新授权
func isRefreshAuthCode(data gen.DianpingAuth) bool {
	if data.RemainRefreshCount == 0 {
		refreshAt := time.Now().Add(time.Hour * 24 * 3) //距离过期时间3天前提示刷新
		if refreshAt.After(*data.ExpiresAt) {
			return true
		}
	}
	return false
}

// SessionScope 获取店铺列表
func SessionScope(appKey string) error {
	var data gen.DianpingAuth
	mgr := gen.DianpingAuthMgr(config2.AppConf.Data.DB)
	data, e := mgr.FetchUniqueByDianpingAuthAppKeyUIndex(appKey)
	if e != nil {
		return e
	}
	query := dianping.NewReqParams()
	query.AddPublicParams(&config)
	query.Set("bid", data.Bid)
	query.Set("session", data.AccessToken)
	resp, err := sdkService.OauthSessionScope(query)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Sprintf("%v", resp.Data)

	return nil
}

// 绑定店铺id
func CreateShopList(appKey string) error {
	var data gen.DianpingAuth
	mgr := gen.DianpingAuthMgr(config2.AppConf.Data.DB)
	data, e := mgr.FetchUniqueByDianpingAuthAppKeyUIndex(appKey)
	if e != nil {
		return e
	}
	query := dianping.NewReqParams()
	query.AddPublicParams(&config)
	query.Set("bid", data.Bid)
	query.Set("session", data.AccessToken)
	resp, err := sdkService.OauthSessionScope(query)
	if err != nil {
		fmt.Println(err)
	}
	for _, res := range resp.Data {
		shopMgr := gen.DianpingShopMgr(config2.AppConf.Data.DB)
		shop, _ := shopMgr.FetchUniqueByDianpingShopOpenShopUUIDUIndex(res.OpenShopUuid)
		shop.MerchantID = data.MerchantID
		shop.OpenShopUUID = res.OpenShopUuid
		shop.ShopName = res.Shopname
		shop.ShopAddress = res.ShopAddress
		shop.BranchName = res.Branchname
		shop.CityName = res.Cityname
		shopMgr.Save(&shop)

	}
	return nil
}

// 绑定团购列表
func CreateTuangouList(appKey string) error {
	var data gen.DianpingAuth
	mgr := gen.DianpingAuthMgr(config2.AppConf.Data.DB)
	data, e := mgr.FetchUniqueByDianpingAuthAppKeyUIndex(appKey)
	if e != nil {
		return e
	}

	var shopData []*gen.DianpingShop
	shopMgr := gen.DianpingShopMgr(config2.AppConf.Data.DB)
	shopData, e = shopMgr.GetBatchFromMerchantID([]int64{data.MerchantID})
	if e != nil {
		return e
	}
	for _, shop := range shopData {
		query := dianping.NewReqParams()
		query.AddPublicParams(&config)
		query.Set("bid", data.Bid)
		query.Set("session", data.AccessToken)
		query.Set("open_shop_uuid", shop.OpenShopUUID)
		resp, err := sdkService.DealQueryshopdeal(query)
		if err != nil {
			fmt.Println(err)
		}
		if resp.Resp.Code != 200 {
			fmt.Println(resp.Resp.Msg)
			return nil
		}
		for _, res := range resp.Data {
			tuangouMgr := gen.DianpingTuangouMgr(config2.AppConf.Data.DB)

			tuangou, _ := tuangouMgr.GetByOption(tuangouMgr.WithDealID(res.DealId))

			tuangou.OpenShopUUID = shop.OpenShopUUID
			tuangou.MerchantID = data.MerchantID
			tuangou.DealID = res.DealId
			tuangou.DealgroupID = res.DealgroupId
			tuangou.Marketprice = *res.Marketprice
			tuangou.Price = res.Price
			tuangou.Title = res.Title
			tuangouMgr.Save(&tuangou)

		}
	}

	return nil
}

func Prepare(shopId int64, memberId int64, code string) (packageCardId int64, count int, err error) {
	code = strings.Replace(code, " ", "", -1)

	shopMgr := gen.DianpingShopMgr(config2.AppConf.Data.DB)
	shop, err := shopMgr.GetFromShopID(shopId)
	if err != nil {
		return 0, 0, err
	}
	mgr := gen.DianpingAuthMgr(config2.AppConf.Data.DB)
	auth, err := mgr.GetFromMerchantID(shop.MerchantID)
	if err != nil {
		return 0, 0, err
	}
	query := dianping.NewReqParams()
	query.AddPublicParams(&config)
	query.Set("receipt_code", code)
	query.Set("session", auth.AccessToken)
	query.Set("open_shop_uuid", shop.OpenShopUUID)
	resp, err := sdkService.ReceiptPrepare(query)
	if err != nil {
		RecordTuanGouLog(&TuangouLog{
			Action:      "ConsumeError",
			MemberID:    memberId,
			ReceiptCode: code,
			Note:        err.Error(),
		})
		return 0, 0, err
	}
	if resp.Resp.Code != 200 {
		bytes, err := json.Marshal(resp)
		if err != nil {
			return 0, 0, errors.New(ERR_1000)
		}
		RecordTuanGouLog(&TuangouLog{
			Action:      "PrepareError",
			MemberID:    memberId,
			ReceiptCode: code,
			Code:        resp.Resp.Code,
			Note:        string(bytes),
		})

		return 0, 0, returnErr(resp.Resp.Code, resp.Resp.Msg)
	}
	bytes, err := json.Marshal(resp)
	if err != nil {
		return 0, 0, errors.New(ERR_1000)
	}

	tougouMgr := gen.DianpingTuangouMgr(config2.AppConf.Data.DB)
	tougouMgr.Where("open_shop_uuid = ?", shop.OpenShopUUID)
	tougouMgr.Where("deal_id = ?", resp.Data.DealId)
	tougou, err := tougouMgr.Get()
	if err != nil || tougou.PackageCardID == 0 {
		RecordTuanGouLog(&TuangouLog{
			Action:      "PrepareErrorNoPackageCardID",
			TuangouID:   int64(tougou.ID),
			DealID:      *resp.Data.DealId,
			MemberID:    memberId,
			ReceiptCode: code,
			Code:        resp.Resp.Code,
			Note:        string(bytes),
		})
		return 0, 0, errors.New(ERR_1009)
	}
	RecordTuanGouLog(&TuangouLog{
		Action:      "Prepare",
		TuangouID:   int64(tougou.ID),
		DealID:      *resp.Data.DealId,
		MemberID:    memberId,
		ReceiptCode: code,
		Code:        resp.Resp.Code,
		Note:        string(bytes),
	})
	return tougou.PackageCardID, resp.Data.Count, nil

}

func ScanPrepare(shopId int64, memberId int64, code string) (packageCardId []int64, count int, err error) {
	code = strings.Replace(code, " ", "", -1)

	shopMgr := gen.DianpingShopMgr(config2.AppConf.Data.DB)
	shop, err := shopMgr.GetFromShopID(shopId)
	if err != nil {
		return nil, 0, err
	}
	mgr := gen.DianpingAuthMgr(config2.AppConf.Data.DB)
	auth, err := mgr.GetFromMerchantID(shop.MerchantID)
	if err != nil {
		return nil, 0, err
	}
	query := dianping.NewReqParams()
	query.AddPublicParams(&config)
	query.Set("qr_code", code)
	query.Set("session", auth.AccessToken)
	query.Set("open_shop_uuid", shop.OpenShopUUID)
	resp, err := sdkService.ReceiptScanprepare(query)
	if err != nil {
		RecordTuanGouLog(&TuangouLog{
			Action:      "ScanPrepareError",
			MemberID:    memberId,
			ReceiptCode: code,
			Note:        err.Error(),
		})
		return nil, 0, err
	}
	if resp.Resp.Code != 200 {
		bytes, err := json.Marshal(resp)
		if err != nil {
			return nil, 0, errors.New(ERR_1000)
		}
		RecordTuanGouLog(&TuangouLog{
			Action:      "ScanPrepareError",
			MemberID:    memberId,
			ReceiptCode: code,
			Code:        resp.Resp.Code,
			Note:        string(bytes),
		})

		return nil, 0, returnErr(resp.Resp.Code, resp.Resp.Msg)
	}
	bytes, err := json.Marshal(resp)
	if err != nil {
		return nil, 0, errors.New(ERR_1000)
	}
	if len(resp.Data) > 1 { //暂不知什么情况返回多维
		return nil, 0, errors.New(ERR_1001)
	}
	prepareResp := resp.Data[0]
	//for _, prepareResp := range resp.Data {
	tougouMgr := gen.DianpingTuangouMgr(config2.AppConf.Data.DB)
	tougouMgr.Where("open_shop_uuid = ?", shop.OpenShopUUID)
	tougouMgr.Where("deal_id = ?", prepareResp.DealId)
	tougou, _ := tougouMgr.Get()
	RecordTuanGouLog(&TuangouLog{
		Action:      "ScanPrepare",
		TuangouID:   int64(tougou.ID),
		DealID:      *prepareResp.DealId,
		MemberID:    memberId,
		ReceiptCode: code,
		Code:        resp.Resp.Code,
		Note:        string(bytes),
	})
	pId, count, err := Prepare(shopId, memberId, prepareResp.ReceiptCode)
	if err != nil {
		return nil, 0, err
	}
	packageCardId = append(packageCardId, pId)
	//}

	return packageCardId, count, err
}

func Consume(shopId int64, memberId int64, code string, count int) (ok bool, err error) {
	code = strings.Replace(code, " ", "", -1)

	shopMgr := gen.DianpingShopMgr(config2.AppConf.Data.DB)
	shop, err := shopMgr.GetFromShopID(shopId)
	if err != nil {
		return false, err
	}
	mgr := gen.DianpingAuthMgr(config2.AppConf.Data.DB)
	auth, err := mgr.GetFromMerchantID(shop.MerchantID)
	if err != nil {
		return false, err
	}
	query := dianping.NewReqParams()
	query.AddPublicParams(&config)
	query.Set("receipt_code", code)
	query.Set("count", strconv.Itoa(count))
	query.Set("requestid", dianping.RandStr(12))
	query.Set("session", auth.AccessToken)
	query.Set("open_shop_uuid", shop.OpenShopUUID)
	query.Set("app_shop_account", strconv.FormatInt(memberId, 10))
	query.Set("app_shop_accountname", strconv.FormatInt(memberId, 10))
	resp, err := sdkService.ReceiptConsume(query)

	if err != nil {
		RecordTuanGouLog(&TuangouLog{
			Action:      "ConsumeError",
			MemberID:    memberId,
			ReceiptCode: code,
			Note:        err.Error(),
		})
		return false, err
	}
	if resp.Resp.Code != 200 {
		bytes, err := json.Marshal(resp)
		if err != nil {
			return false, errors.New(ERR_1000)
		}
		RecordTuanGouLog(&TuangouLog{
			Action:      "ConsumeError",
			MemberID:    memberId,
			ReceiptCode: code,
			Code:        resp.Resp.Code,
			Note:        string(bytes),
		})

		return false, returnErr(resp.Resp.Code, resp.Resp.Msg)
	}
	bytes, err := json.Marshal(resp)
	if err != nil {
		return false, errors.New(ERR_1000)
	}

	for _, receiptConsumeResp := range resp.Data {
		tougouMgr := gen.DianpingTuangouMgr(config2.AppConf.Data.DB)
		tougouMgr.Where("open_shop_uuid = ?", shop.OpenShopUUID)
		tougouMgr.Where("deal_id = ?", receiptConsumeResp.DealId)
		tougou, _ := tougouMgr.Get()
		RecordTuanGouLog(&TuangouLog{
			Action:      "Consume",
			TuangouID:   int64(tougou.ID),
			DealID:      *receiptConsumeResp.DealId,
			MemberID:    memberId,
			ReceiptCode: code,
			Code:        resp.Resp.Code,
			Note:        string(bytes),
		})
	}
	if isRefreshAccessToken(auth) { //刷新token
		err := RefreshAccessToken(auth.AppKey)
		if err != nil {
			RecordTuanGouLog(&TuangouLog{
				Action: "RefreshAccessTokenERR",
			})
		}
	}
	return true, nil

}

type TuangouLog struct {
	ID              int64
	Action          string
	TuangouID       int64
	DealID          int64
	MemberID        int64
	PaymentDetailID int64
	ReceiptCode     string
	Code            int
	PaymentAmount   int64
	Note            string
	CreateBy        int
}

func RecordTuanGouLog(log *TuangouLog) {
	mgr := gen.DianpingTuangouLogMgr(config2.AppConf.Data.DB)
	createBy := uint(0)
	mgr.Save(&gen.DianpingTuangouLog{
		TuangouID:       log.TuangouID,
		DealID:          log.DealID,
		MemberID:        log.MemberID,
		PaymentDetailID: log.PaymentDetailID,
		Code:            log.Code,
		ReceiptCode:     log.ReceiptCode,
		PaymentAmount:   log.PaymentAmount,
		Note:            log.Note,
		CreatedAt:       time.Time{},
		CreateBy:        &createBy,
		Action:          log.Action,
	})
}

func returnErr(errCode int, defaultStr string) (err error) {
	switch errCode {
	case 1006:
		err = errors.New(ERR_1006)
	case 1008:
		err = errors.New(ERR_1008)
	default:
		err = errors.New(defaultStr)
	}
	return err
}
