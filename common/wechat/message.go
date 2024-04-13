package wechat

import (
	"errors"
	"fmt"
	config2 "studyRoomGo/common/config"
	"studyRoomGo/common/wechat/gen"
	merchant2 "studyRoomGo/common/wechat/merchant"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"

	"github.com/silenceper/wechat/v2/miniprogram/subscribe"
	"gorm.io/gorm/clause"
)

type Message struct {
	AppId      string
	AppSecret  string
	Version    int8
	TemplateId string
	ToOpenId   string
	TmplIds    TmplIds
}

type TmplIds map[string]string

func LoadTmpl(MerchantId int64) TmplIds {
	merchant := make(map[int64]TmplIds)
	merchant[1] = TmplIds{
		//微信小程序通知模版
		"SendSeatBookMessage":         "",
		"SendSeatBookReadyMessage":    "",
		"SendSeatBookCompleteMessage": "",
		"SendCouponReceiveMessage":    "",
	}
	merchant[3] = TmplIds{
		"SendSeatBookMessage":         "",
		"SendSeatBookReadyMessage":    "",
		"SendSeatBookCompleteMessage": "",
	}
	merchant[2] = TmplIds{
		"SendStudentCheckInSuccess":  "", //上课签到
		"SendStudentCheckOutSuccess": "", //上课签退
		"ServiceNotice":              "", //工单通知，通知客服人员
	}
	return merchant[MerchantId]
}

func MerchantToOfficalMerchantId(MerchantId int64) int64 {
	m := make(map[int64]int64)
	m[1] = 2
	return m[MerchantId]
}
func NewMessage(merchantId int64) (*Message, error) {
	merchant := merchant2.NewMerchant(merchantId)

	data, err := merchant.GetMerchantAppSecret()
	if err != nil {
		return nil, err
	}
	return &Message{
		AppId:     data.AppID,
		AppSecret: data.AppSecret,
		Version:   data.Version,
		TmplIds:   LoadTmpl(merchantId),
	}, nil
}

type SeatBookMessage struct {
	ShopName  string
	SeatName  string
	Desc      string
	StartDate string
	TimeRange string
	Tip       string
}

func (m *Message) SendSeatBookMessage(memberId int64, message SeatBookMessage) error {
	m.TemplateId = m.TmplIds["SendSeatBookMessage"]

	if m.TemplateId == "" {
		return nil
	}

	mgr := gen.WxUserinfoMgr(config2.AppConf.Data.DB.Where("member_id = ?", memberId))
	user, err := mgr.Get()
	if err != nil {
		return errors.New("memberId错误")
	}
	m.ToOpenId = user.Openid

	data := make(map[string]*subscribe.DataItem)

	data["thing28"] = &subscribe.DataItem{
		Value: message.ShopName,
	}
	data["thing29"] = &subscribe.DataItem{
		Value: message.SeatName,
	}

	data["thing4"] = &subscribe.DataItem{
		Value: message.Tip,
	}
	data["time36"] = &subscribe.DataItem{
		Value: message.TimeRange,
	}
	data["time32"] = &subscribe.DataItem{
		Value: message.StartDate,
	}

	if m.Version == 1 {
		err = m.sendMessage(data) //定制版，密钥授权
		if err != nil {
			fmt.Printf("sendMessage 【SendSeatBookMessage】error ,memberId:%v", memberId)
			_, _ = SaveWXSubscribe(&gen.WxNoticeSubscribe{
				MemberID:    memberId,
				Type:        m.TemplateId,
				IsSubscribe: 0,
			})
		} else {
			fmt.Printf("sendMessage 【SendSeatBookMessage】success ,memberId:%v", memberId)
			_, _ = SaveWXSubscribe(&gen.WxNoticeSubscribe{
				MemberID:    memberId,
				Type:        m.TemplateId,
				IsSubscribe: 1,
			})
		}
	}
	return nil
}

type SendSeatBookReadyMessage struct {
	ShopName      string
	SeatName      string
	Desc          string
	StartDateTime string
	Countdown     string
	Tip           string
}

// SendSeatBookReadyMessage 准备入座提醒
func (m *Message) SendSeatBookReadyMessage(memberId int64, message SendSeatBookReadyMessage) error {
	m.TemplateId = m.TmplIds["SendSeatBookReadyMessage"]

	if m.TemplateId == "" {
		return nil
	}

	mgr := gen.WxUserinfoMgr(config2.AppConf.Data.DB.Where("member_id = ?", memberId))
	user, err := mgr.Get()
	if err != nil {
		return errors.New("memberId错误")
	}
	m.ToOpenId = user.Openid

	data := make(map[string]*subscribe.DataItem)
	data["thing31"] = &subscribe.DataItem{
		Value: message.ShopName,
	}
	data["short_thing32"] = &subscribe.DataItem{
		Value: message.SeatName,
	}
	data["thing9"] = &subscribe.DataItem{
		Value: message.Tip,
	}
	data["time24"] = &subscribe.DataItem{
		Value: message.StartDateTime,
	}
	data["short_thing33"] = &subscribe.DataItem{
		Value: message.Countdown,
	}
	if m.Version == 1 {
		err = m.sendMessage(data) //定制版，密钥授权
		if err != nil {
			log.Infof("sendMessage 【SendSeatBookReadyMessage】error ,memberId:%v", memberId)
			_, _ = SaveWXSubscribe(&gen.WxNoticeSubscribe{
				MemberID:    memberId,
				Type:        m.TemplateId,
				IsSubscribe: 0,
			})
		} else {
			log.Infof("sendMessage 【SendSeatBookReadyMessage】success ,memberId:%v", memberId)
			_, _ = SaveWXSubscribe(&gen.WxNoticeSubscribe{
				MemberID:    memberId,
				Type:        m.TemplateId,
				IsSubscribe: 1,
			})
		}
	}
	return nil
}

type SendStudentCheckInSuccess struct {
	CourseName  string
	StudentName string
	AddressName string
	CourseTime  string
	CheckIntime string
}

func (m *Message) SendStudentCheckInSuccessMessage(studentId int64, mess SendStudentCheckInSuccess) error {
	m.TemplateId = m.TmplIds["SendStudentCheckInSuccess"]

	if m.TemplateId == "" {
		return nil
	}

	mgr := gen.EdStudentRelateMemberMgr(config2.AppConf.Data.DB.Where("student_id = ?", studentId))
	user, err := mgr.Get()
	if err != nil {
		return errors.New("student_id错误")
	}
	mgr2 := gen.WxOauthMgr(config2.AppConf.Data.DB.Where("member_id = ?", user.MemberID))
	oauth, err := mgr2.Get()
	if err != nil {
		return errors.New("oauth错误")
	}

	m.ToOpenId = oauth.Openid

	data := make(map[string]*message.TemplateDataItem)
	data["thing18"] = &message.TemplateDataItem{ //签到人
		Value: mess.StudentName,
	}
	data["thing12"] = &message.TemplateDataItem{ //签到课程
		Value: mess.CourseName,
	}
	data["thing13"] = &message.TemplateDataItem{
		Value: mess.AddressName,
	}
	data["time10"] = &message.TemplateDataItem{ //签到时间
		Value: mess.CheckIntime,
	}
	data["time19"] = &message.TemplateDataItem{ //课程时间
		Value: mess.CourseTime,
	}
	if m.Version == 1 {
		err = m.sendOfficialMessage(data) //定制版，密钥授权
		fmt.Printf("sendOfficialMessage 【SendStudentCheckInSuccessMessage】%v,studentId:%v", err, studentId)
		return err
	}
	return nil
}

type SendStudentCheckOutSuccess struct {
	StudentName  string
	CheckOutTime string
	CheckOutDate string
}

func (m *Message) SendStudentCheckOutSuccessMessage(studentId int64, mess SendStudentCheckOutSuccess) error {
	m.TemplateId = m.TmplIds["SendStudentCheckOutSuccess"]

	if m.TemplateId == "" {
		return nil
	}

	mgr := gen.EdStudentRelateMemberMgr(config2.AppConf.Data.DB.Where("student_id = ?", studentId))
	user, err := mgr.Get()
	if err != nil {
		return errors.New("student_id错误")
	}
	mgr2 := gen.WxOauthMgr(config2.AppConf.Data.DB.Where("member_id = ?", user.MemberID))
	oauth, err := mgr2.Get()
	if err != nil {
		return errors.New("oauth错误")
	}

	m.ToOpenId = oauth.Openid

	data := make(map[string]*message.TemplateDataItem)
	data["thing1"] = &message.TemplateDataItem{ //签退人
		Value: mess.StudentName,
	}
	data["time2"] = &message.TemplateDataItem{ //签退日期
		Value: mess.CheckOutDate,
	}
	data["time3"] = &message.TemplateDataItem{ //签退时间
		Value: mess.CheckOutTime,
	}

	if m.Version == 1 {
		err = m.sendOfficialMessage(data) //定制版，密钥授权
		fmt.Printf("sendOfficialMessage 【SendStudentCheckOutSuccessMessage】%v,studentId:%v", err, studentId)
		return err
	}
	return nil
}

type CouponReceiveMessage struct {
	ShopName    string
	CouponName  string
	EndDateTime string
	Tip         string
}

func (m *Message) SendCouponReceiveMessage(memberId int64, message CouponReceiveMessage) error {

	m.TemplateId = m.TmplIds["SendCouponReceiveMessage"]

	if m.TemplateId == "" {
		return nil
	}
	mgr := gen.WxUserinfoMgr(config2.AppConf.Data.DB.Where("member_id = ?", memberId))
	user, err := mgr.Get()
	if err != nil {
		return errors.New("memberId错误")
	}
	m.ToOpenId = user.Openid

	data := make(map[string]*subscribe.DataItem)

	data["thing7"] = &subscribe.DataItem{
		Value: message.ShopName,
	}
	data["thing1"] = &subscribe.DataItem{
		Value: message.CouponName,
	}

	data["thing3"] = &subscribe.DataItem{
		Value: message.Tip,
	}
	data["time2"] = &subscribe.DataItem{
		Value: message.EndDateTime,
	}

	if m.Version == 1 {
		err = m.sendMessage(data) //定制版，密钥授权
		if err != nil {
			fmt.Printf("sendMessage 【SendCouponReceiveMessage】error ,memberId:%v \n", memberId)
			_, _ = SaveWXSubscribe(&gen.WxNoticeSubscribe{
				MemberID:    memberId,
				Type:        m.TemplateId,
				IsSubscribe: 0,
			})
		} else {
			fmt.Printf("sendMessage 【SendCouponReceiveMessage】success ,memberId:%v \n", memberId)
			_, _ = SaveWXSubscribe(&gen.WxNoticeSubscribe{
				MemberID:    memberId,
				Type:        m.TemplateId,
				IsSubscribe: 1,
			})
		}
	}
	return nil
}

type SendSeatBookCompleteMessage struct {
	ShopName      string
	SeatName      string
	Desc          string
	CloseDateTime string
	Tip           string
}

func (m *Message) SendSeatBookCompleteMessage(memberId int64, message SendSeatBookCompleteMessage) error {
	m.TemplateId = m.TmplIds["SendSeatBookCompleteMessage"]

	if m.TemplateId == "" {
		return nil
	}

	mgr := gen.WxUserinfoMgr(config2.AppConf.Data.DB.Where("member_id = ?", memberId))
	user, err := mgr.Get()
	if err != nil {
		return errors.New("memberId错误")
	}
	m.ToOpenId = user.Openid

	data := make(map[string]*subscribe.DataItem)
	data["thing6"] = &subscribe.DataItem{
		Value: message.ShopName,
	}
	data["character_string7"] = &subscribe.DataItem{
		Value: message.SeatName,
	}
	data["thing3"] = &subscribe.DataItem{
		Value: message.Tip,
	}
	data["time2"] = &subscribe.DataItem{
		Value: message.CloseDateTime,
	}
	if m.Version == 1 {
		err = m.sendMessage(data) //定制版，密钥授权
		if err != nil {
			log.Infof("sendMessage 【SendSeatBookCompleteMessage】error ,memberId:%v \n", memberId)
			_, _ = SaveWXSubscribe(&gen.WxNoticeSubscribe{
				MemberID:    memberId,
				Type:        m.TemplateId,
				IsSubscribe: 0,
			})
		} else {
			log.Infof("sendMessage 【SendSeatBookCompleteMessage】success ,memberId:%v \n", memberId)
			_, _ = SaveWXSubscribe(&gen.WxNoticeSubscribe{
				MemberID:    memberId,
				Type:        m.TemplateId,
				IsSubscribe: 1,
			})
		}
	}
	return nil
}

// SendMessage 微信小程序登录
func (m *Message) sendMessage(data map[string]*subscribe.DataItem) error {

	wc := wechat.NewWechat()
	cfg := &miniConfig.Config{
		AppID:     m.AppId,
		AppSecret: m.AppSecret,
		Cache:     cache.NewMemory(),
	}
	mini := wc.GetMiniProgram(cfg)
	a := mini.GetSubscribe()

	sub := &subscribe.Message{
		ToUser:           m.ToOpenId,
		TemplateID:       m.TemplateId,
		Page:             "pages/index/bootstrap",
		Data:             data,
		MiniprogramState: "developer",
		Lang:             "zh_CN",
	}
	if config2.AppConf.Environment == "pro" {
		sub.MiniprogramState = "formal"
	}
	err := a.Send(sub)

	if err != nil {
		return err
	}
	return nil
}

// sendOfficialMessage 发送微信公众号 服务号 模版消息
func (m *Message) sendOfficialMessage(data map[string]*message.TemplateDataItem) error {

	wc := wechat.NewWechat()
	cfg := &offConfig.Config{
		AppID:     m.AppId,
		AppSecret: m.AppSecret,
		Cache:     cache.NewMemory(),
	}
	mini := wc.GetOfficialAccount(cfg)
	a := mini.GetTemplate()
	sub := &message.TemplateMessage{
		ToUser:     m.ToOpenId,
		TemplateID: m.TemplateId,
		Data:       data,
	}
	if config2.AppConf.Environment == "pro" {

	}
	_, err := a.Send(sub)

	if err != nil {
		return err
	}
	return nil
}

type ServiceNotice struct {
	ServiceName string
	MemberName  string
	ProductName string
	CreateAt    string
}

func (m *Message) SendServiceNoticeMessage(member_id []int64, mess ServiceNotice) error {
	m.TemplateId = m.TmplIds["ServiceNotice"]

	if m.TemplateId == "" {
		return nil
	}

	mgr2 := gen.WxOauthMgr(config2.AppConf.Data.DB.Where("member_id in ?", member_id))
	oauth, err := mgr2.Gets()
	if err != nil {
		return errors.New("oauth错误")
	}
	for _, v := range oauth {
		m.ToOpenId = v.Openid

		data := make(map[string]*message.TemplateDataItem)
		data["thing8"] = &message.TemplateDataItem{ //工单名称
			Value: mess.ServiceName,
		}
		data["thing5"] = &message.TemplateDataItem{ //姓名

			Value: mess.MemberName,
		}
		data["thing15"] = &message.TemplateDataItem{
			Value: mess.ProductName, //商品名称
		}

		data["time6"] = &message.TemplateDataItem{ //创建时间
			Value: mess.CreateAt,
		}
		if m.Version == 1 {
			err = m.sendOfficialMessage(data) //定制版，密钥授权
			fmt.Printf("ServiceNotice 【ServiceNotice】%v,studentId:%v", err, v.MemberID)
		}
	}

	return nil
}

func SaveWXSubscribe(info *gen.WxNoticeSubscribe) (bool, error) {
	mgr := gen.WxNoticeSubscribeMgr(config2.AppConf.Data.DB)
	mgr.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "member_id"}, {Name: "type"}},
		DoUpdates: clause.AssignmentColumns([]string{"is_subscribe"}),
	})
	mgr.Create(info)
	if mgr.Error != nil {
		return false, mgr.Error
	}
	return true, nil
}

func FindWXSubscribe(memberId int64) (map[string]int, error) {
	sub := make(map[string]int)
	mgr := gen.WxNoticeSubscribeMgr(config2.AppConf.Data.DB)
	list, _ := mgr.GetFromMemberID(memberId)
	for _, d := range list {
		sub[d.Type] = d.IsSubscribe
	}
	return sub, nil
}
