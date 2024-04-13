package gen

import "time"

// SrPayment 交易信息表
type SrPayment struct {
	ID            int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	OrderNo       string     `gorm:"unique;column:order_no;type:char(20);not null;default:'';comment:'订单编号'"`     // 订单编号
	TransactionID *string    `gorm:"unique;column:transaction_id;type:varchar(50);comment:'交易号'"`                 // 交易号
	PaymentType   string     `gorm:"column:payment_type;type:varchar(20);not null;default:'';comment:'支付类型'"`     // 支付类型
	Currency      string     `gorm:"column:currency;type:char(3);not null;default:'';comment:'币种'"`               // 币种
	Amount        float64    `gorm:"column:amount;type:float(10,4);not null;default:0.0000;comment:'金额'"`         // 金额
	Status        int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'状态'"`                  // 状态
	CreatedAt     *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy      *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy      *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrPayment) TableName() string {
	return "sr_payment"
}

// SrPaymentColumns get sql column name.获取数据库列名
var SrPaymentColumns = struct {
	ID            string
	OrderNo       string
	TransactionID string
	PaymentType   string
	Currency      string
	Amount        string
	Status        string
	CreatedAt     string
	UpdatedAt     string
	DeletedAt     string
	CreateBy      string
	UpdateBy      string
}{
	ID:            "id",
	OrderNo:       "order_no",
	TransactionID: "transaction_id",
	PaymentType:   "payment_type",
	Currency:      "currency",
	Amount:        "amount",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	CreateBy:      "create_by",
	UpdateBy:      "update_by",
}
