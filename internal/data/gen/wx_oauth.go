package gen

// WxOauth 微信网页oauth授权
type WxOauth struct {
	ID         int    `gorm:"autoIncrement:true;primaryKey;column:id;type:int;not null"`
	MerchantID int64  `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"` // 商家id
	Openid     string `gorm:"unique;column:openid;type:varchar(255);not null;comment:'微信用户唯一标识符

微信用户唯一标识符'"` // 微信用户唯一标识符,,微信用户唯一标识符
	AccessToken  string `gorm:"column:access_token;type:varchar(255);not null"`
	RefreshToken string `gorm:"column:refresh_token;type:varchar(255)"`
	UnionID      string `gorm:"column:union_id;type:varchar(255)"`
	Expires      string `gorm:"column:expires;type:varchar(255);not null;default:''"`
	Phone        string `gorm:"column:phone;type:varchar(255);comment:'用户绑定的手机号'"`              // 用户绑定的手机号
	MemberID     int64  `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"` // 会员id
}

// TableName get sql table name.获取数据库表名
func (m *WxOauth) TableName() string {
	return "wx_oauth"
}

// WxOauthColumns get sql column name.获取数据库列名
var WxOauthColumns = struct {
	ID           string
	MerchantID   string
	Openid       string
	AccessToken  string
	RefreshToken string
	UnionID      string
	Expires      string
	Phone        string
	MemberID     string
}{
	ID:           "id",
	MerchantID:   "merchant_id",
	Openid:       "openid",
	AccessToken:  "access_token",
	RefreshToken: "refresh_token",
	UnionID:      "union_id",
	Expires:      "expires",
	Phone:        "phone",
	MemberID:     "member_id",
}
