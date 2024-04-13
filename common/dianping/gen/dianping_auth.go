package gen

import "time"

// DianpingAuth 大众点评授权信息
type DianpingAuth struct {
	ID                 int        `gorm:"autoIncrement:true;primaryKey;column:id;type:int;not null"`
	MerchantID         int64      `gorm:"unique;column:merchant_id;type:bigint;not null;default:0;comment:'系统商家id'"`            // 系统商家id
	AppKey             string     `gorm:"unique;column:app_key;type:varchar(255);not null;comment:'北极星开放平台分配给应用的AppKey'"`       // 北极星开放平台分配给应用的AppKey
	AppSecret          string     `gorm:"column:app_secret;type:varchar(255);not null;comment:'密钥'"`                            // 密钥
	AuthCode           string     `gorm:"column:auth_code;type:varchar(255);not null;default:'';comment:'授权码'"`                 // 授权码
	AccessToken        string     `gorm:"column:access_token;type:varchar(255);not null;default:'';comment:'session'"`          // session
	ExpiresIn          int64      `gorm:"column:expires_in;type:bigint;not null;default:0;comment:'过期时间'"`                      // 过期时间
	ExpiresAt          *time.Time `gorm:"column:expires_at;type:timestamp;comment:'过期时间'"`                                      // 过期时间
	RefreshToken       string     `gorm:"column:refresh_token;type:varchar(255);not null;default:'';comment:'refresh_session'"` // refresh_session
	RemainRefreshCount int        `gorm:"column:remain_refresh_count;type:int;not null;default:0;comment:'剩余刷新次数'"`             // 剩余刷新次数
	Bid                string     `gorm:"column:bid;type:varchar(255);not null;default:'';comment:'客户Id'"`                      // 客户Id
}

// TableName get sql table name.获取数据库表名
func (m *DianpingAuth) TableName() string {
	return "dianping_auth"
}

// DianpingAuthColumns get sql column name.获取数据库列名
var DianpingAuthColumns = struct {
	ID                 string
	MerchantID         string
	AppKey             string
	AppSecret          string
	AuthCode           string
	AccessToken        string
	ExpiresIn          string
	ExpiresAt          string
	RefreshToken       string
	RemainRefreshCount string
	Bid                string
}{
	ID:                 "id",
	MerchantID:         "merchant_id",
	AppKey:             "app_key",
	AppSecret:          "app_secret",
	AuthCode:           "auth_code",
	AccessToken:        "access_token",
	ExpiresIn:          "expires_in",
	ExpiresAt:          "expires_at",
	RefreshToken:       "refresh_token",
	RemainRefreshCount: "remain_refresh_count",
	Bid:                "bid",
}
