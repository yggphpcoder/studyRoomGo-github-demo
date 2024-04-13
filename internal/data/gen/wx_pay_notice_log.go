package gen

import "time"

// WxPayNoticeLog 微信支付回调日志
type WxPayNoticeLog struct {
	ID            int        `gorm:"autoIncrement:true;primaryKey;column:id;type:int;not null"`
	Type          string     `gorm:"column:type;type:varchar(20);not null;default:'';comment:'类型'"`             // 类型
	OrderNo       string     `gorm:"column:order_no;type:char(20);not null;default:'';comment:'订单编号'"`          // 订单编号
	TransactionID string     `gorm:"column:transaction_id;type:char(32);not null;default:'';comment:'微信支付订单号'"` // 微信支付订单号
	Status        string     `gorm:"column:status;type:varchar(10);not null;default:'';comment:'状态'"`           // 状态
	Header        string     `gorm:"column:header;type:text;not null;comment:'header数据'"`                       // header数据
	Body          string     `gorm:"column:body;type:text;not null;comment:'body数据'"`                           // body数据
	CreatedAt     *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

// TableName get sql table name.获取数据库表名
func (m *WxPayNoticeLog) TableName() string {
	return "wx_pay_notice_log"
}

// WxPayNoticeLogColumns get sql column name.获取数据库列名
var WxPayNoticeLogColumns = struct {
	ID            string
	Type          string
	OrderNo       string
	TransactionID string
	Status        string
	Header        string
	Body          string
	CreatedAt     string
}{
	ID:            "id",
	Type:          "type",
	OrderNo:       "order_no",
	TransactionID: "transaction_id",
	Status:        "status",
	Header:        "header",
	Body:          "body",
	CreatedAt:     "created_at",
}
