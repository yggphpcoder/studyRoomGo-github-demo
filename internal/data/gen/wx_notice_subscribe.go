package gen

// WxNoticeSubscribe 微信消息订阅
type WxNoticeSubscribe struct {
	MemberID    int64  `gorm:"uniqueIndex:wx_notice_subscribe_pk;column:member_id;type:bigint;not null;default:0;comment:'会员id'"`   // 会员id
	Type        string `gorm:"uniqueIndex:wx_notice_subscribe_pk;column:type;type:varchar(500);not null;default:'';comment:'订阅类型'"` // 订阅类型
	IsSubscribe int    `gorm:"column:is_subscribe;type:int;not null;default:0;comment:'是否订阅'"`                                      // 是否订阅
}

// TableName get sql table name.获取数据库表名
func (m *WxNoticeSubscribe) TableName() string {
	return "wx_notice_subscribe"
}

// WxNoticeSubscribeColumns get sql column name.获取数据库列名
var WxNoticeSubscribeColumns = struct {
	MemberID    string
	Type        string
	IsSubscribe string
}{
	MemberID:    "member_id",
	Type:        "type",
	IsSubscribe: "is_subscribe",
}
