package gen

import "time"

// SrPaymentLog 支付日志
type SrPaymentLog struct {
	ID            int        `gorm:"autoIncrement:true;primaryKey;column:id;type:int;not null"`
	PaymentID     int64      `gorm:"column:payment_id;type:bigint;not null;default:0;comment:'支付表id'"` // 支付表id
	OrderNo       string     `gorm:"column:order_no;type:char(20);not null;default:'';comment:'订单编号'"` // 订单编号
	TransactionID *string    `gorm:"column:transaction_id;type:char(32);comment:'微信支付订单号'"`            // 微信支付订单号
	Status        int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'状态'"`       // 状态
	Note          string     `gorm:"column:note;type:varchar(500);not null;default:'';comment:'其他记录'"` // 其他记录
	CreatedAt     *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

// TableName get sql table name.获取数据库表名
func (m *SrPaymentLog) TableName() string {
	return "sr_payment_log"
}

// SrPaymentLogColumns get sql column name.获取数据库列名
var SrPaymentLogColumns = struct {
	ID            string
	PaymentID     string
	OrderNo       string
	TransactionID string
	Status        string
	Note          string
	CreatedAt     string
}{
	ID:            "id",
	PaymentID:     "payment_id",
	OrderNo:       "order_no",
	TransactionID: "transaction_id",
	Status:        "status",
	Note:          "note",
	CreatedAt:     "created_at",
}
