package gen

import "time"

// SrOrderLog 订单日志
type SrOrderLog struct {
	ID        int        `gorm:"autoIncrement:true;primaryKey;column:id;type:int;not null"`
	OrderID   int64      `gorm:"column:order_id;type:bigint;not null;default:0;comment:'订单id'"`    // 订单id
	OrderNo   string     `gorm:"column:order_no;type:char(20);not null;default:'';comment:'订单编号'"` // 订单编号
	Status    int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'状态'"`       // 状态
	Note      string     `gorm:"column:note;type:varchar(500);not null;default:'';comment:'其他记录'"` // 其他记录
	CreatedAt *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

// TableName get sql table name.获取数据库表名
func (m *SrOrderLog) TableName() string {
	return "sr_order_log"
}

// SrOrderLogColumns get sql column name.获取数据库列名
var SrOrderLogColumns = struct {
	ID        string
	OrderID   string
	OrderNo   string
	Status    string
	Note      string
	CreatedAt string
}{
	ID:        "id",
	OrderID:   "order_id",
	OrderNo:   "order_no",
	Status:    "status",
	Note:      "note",
	CreatedAt: "created_at",
}
