package gen

import (
	"time"
)

// DianpingTuangouLog 团购兑换记录
type DianpingTuangouLog struct {
	ID              int64     `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"` // 主键编码
	TuangouID       int64     `gorm:"column:tuangou_id;type:bigint;not null;default:0;comment:'团购表id'"`            // 团购表id
	DealID          int64     `gorm:"column:deal_id;type:bigint;not null;default:0;comment:'套餐id'"`                // 套餐id
	MemberID        int64     `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`              // 会员id
	PaymentDetailID int64     `gorm:"column:payment_detail_id;type:bigint;not null;comment:'支付详情id'"`              // 支付详情id
	Code            int       `gorm:"column:code;type:int;not null;default:0;comment:'状态'"`                        // 状态
	ReceiptCode     string    `gorm:"column:receipt_code;type:varchar(50);not null;comment:'输入券码'"`                // 输入券码
	PaymentAmount   int64     `gorm:"column:payment_amount;type:bigint;not null;comment:'支付金额'"`                   // 支付金额
	Note            string    `gorm:"column:note;type:varchar(2000);not null;comment:'结果记录'"`                      // 结果记录
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp;not null;default:CURRENT_TIMESTAMP"`
	CreateBy        *uint     `gorm:"column:create_by;type:int unsigned"`
	Action          string    `gorm:"column:action;type:varchar(50);not null;comment:'类型'"` // 类型
}

// TableName get sql table name.获取数据库表名
func (m *DianpingTuangouLog) TableName() string {
	return "dianping_tuangou_log"
}

// DianpingTuangouLogColumns get sql column name.获取数据库列名
var DianpingTuangouLogColumns = struct {
	ID              string
	TuangouID       string
	DealID          string
	MemberID        string
	PaymentDetailID string
	Code            string
	ReceiptCode     string
	PaymentAmount   string
	Note            string
	CreatedAt       string
	CreateBy        string
	Action          string
}{
	ID:              "id",
	TuangouID:       "tuangou_id",
	DealID:          "deal_id",
	MemberID:        "member_id",
	PaymentDetailID: "payment_detail_id",
	Code:            "code",
	ReceiptCode:     "receipt_code",
	PaymentAmount:   "payment_amount",
	Note:            "note",
	CreatedAt:       "created_at",
	CreateBy:        "create_by",
	Action:          "action",
}
