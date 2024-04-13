package gen

// WxUserinfo 微信用户信息
type WxUserinfo struct {
	ID     int    `gorm:"autoIncrement:true;primaryKey;column:id;type:int;not null"`
	Openid string `gorm:"unique;column:openid;type:varchar(255);not null;comment:'微信用户唯一标识符

微信用户唯一标识符'"` // 微信用户唯一标识符,,微信用户唯一标识符
	NickName  *string `gorm:"column:nick_name;type:varchar(255);comment:'用户昵称'"`                                                                                       // 用户昵称
	AvatarURL *string `gorm:"column:avatar_url;type:varchar(255);comment:'用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。'"` // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	Gender    *int    `gorm:"column:gender;type:int;comment:'用户的性别，值为1时是男性，值为2时是女性，值为0时是未知'"`                                                                          // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City      *string `gorm:"column:city;type:varchar(255);comment:'用户所在城市'"`                                                                                          // 用户所在城市
	Province  *string `gorm:"column:province;type:varchar(255);comment:'用户所在省份'"`                                                                                      // 用户所在省份
	Country   *string `gorm:"column:country;type:varchar(255);comment:'用户所在国家'"`                                                                                       // 用户所在国家
	Language  []byte  `gorm:"column:language;type:varbinary(16);comment:'用户的语言，简体中文为zh_CN'"`                                                                           // 用户的语言，简体中文为zh_CN
	Phone     *string `gorm:"column:phone;type:varchar(255);comment:'用户绑定的手机号'"`                                                                                       // 用户绑定的手机号
	MemberID  int64   `gorm:"unique;column:member_id;type:bigint;not null;default:0;comment:'会员id'"`                                                                   // 会员id
}

// TableName get sql table name.获取数据库表名
func (m *WxUserinfo) TableName() string {
	return "wx_userinfo"
}

// WxUserinfoColumns get sql column name.获取数据库列名
var WxUserinfoColumns = struct {
	ID        string
	Openid    string
	NickName  string
	AvatarURL string
	Gender    string
	City      string
	Province  string
	Country   string
	Language  string
	Phone     string
	MemberID  string
}{
	ID:        "id",
	Openid:    "openid",
	NickName:  "nick_name",
	AvatarURL: "avatar_url",
	Gender:    "gender",
	City:      "city",
	Province:  "province",
	Country:   "country",
	Language:  "language",
	Phone:     "phone",
	MemberID:  "member_id",
}
