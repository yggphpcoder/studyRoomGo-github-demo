package gen

// WxPayNoticeLog 微信支付回调日志
type WxPayNoticeLog struct {
	ID     int    `gorm:"autoIncrement:true;primaryKey;column:id;type:int;not null"`
	Type   string `gorm:"column:type;type:varchar(20);not null;default:'';comment:'类型'"` // 类型
	Header string `gorm:"column:header;type:text;not null;comment:'header数据'"`           // header数据
	Body   string `gorm:"column:body;type:text;not null;comment:'body数据'"`               // body数据
}

// TableName get sql table name.获取数据库表名
func (m *WxPayNoticeLog) TableName() string {
	return "wx_pay_notice_log"
}

// WxPayNoticeLogColumns get sql column name.获取数据库列名
var WxPayNoticeLogColumns = struct {
	ID     string
	Type   string
	Header string
	Body   string
}{
	ID:     "id",
	Type:   "type",
	Header: "header",
	Body:   "body",
}
