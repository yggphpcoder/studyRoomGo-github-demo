package merchant

import (
	"errors"
	"fmt"
	"github.com/silenceper/wechat/v2/cache"
	config2 "studyRoomGo/common/config"
	"studyRoomGo/common/wechat/gen"
	"time"
)

var memory *cache.Memory

func init() {
	memory = cache.NewMemory()
}

// Merchant struct
type Merchant struct {
	MerchantId int64
	cache      Cache
}

type MerchantApp struct {
	AppID     string
	AppSecret string
	Version   int8
}

// NewMerchant init
func NewMerchant(MerchantId int64) *Merchant {
	return &Merchant{
		MerchantId: MerchantId,
		cache:      memory,
	}
}

func (m *Merchant) GetMerchantAppSecret() (*MerchantApp, error) {
	// 先从cache中取
	merchantCacheKey := fmt.Sprintf("merchant_app_%d", m.MerchantId)
	if val := m.cache.Get(merchantCacheKey); val != nil {
		d, ok := val.(MerchantApp)
		if !ok {
			return nil, errors.New("merchantId获取缓存失败")
		}
		return &MerchantApp{
			AppID:     d.AppID,
			AppSecret: d.AppSecret,
			Version:   d.Version,
		}, nil
	}
	mgr := gen.CoMerchantWxMiniprogramMgr(config2.AppConf.Data.DB.Where("merchant_id = ?", m.MerchantId))
	data, err := mgr.Get()
	if err != nil {
		return nil, errors.New("merchantId错误")
	}
	expires := 86400 * 10
	info := MerchantApp{
		AppID:     data.AppID,
		AppSecret: data.AppSecret,
		Version:   data.Version,
	}
	err = m.cache.Set(merchantCacheKey, info, time.Duration(expires)*time.Second)
	if err != nil {
		return nil, err
	}
	return &info, nil
}

// Cache interface
type Cache interface {
	Get(key string) interface{}
	Set(key string, val interface{}, timeout time.Duration) error
	IsExist(key string) bool
	Delete(key string) error
}
