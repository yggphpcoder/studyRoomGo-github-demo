package gen

import "time"

// MeProfitSharingWechatOrderLog 分账订单日志
type MeProfitSharingWechatOrderLog struct {
	ID              int        `gorm:"autoIncrement:true;primaryKey;column:id;type:int;not null"`
	ProfitSharingID int64      `gorm:"column:profit_sharing_id;type:bigint;not null;default:0;comment:'分账单id'"` // 分账单id
	OrderNo         string     `gorm:"column:order_no;type:char(20);not null;default:'';comment:'分账单订单编号'"`     // 分账单订单编号
	Status          int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'状态'"`              // 状态
	Body            *string    `gorm:"column:body;type:text;comment:'响应内容'"`                                    // 响应内容
	Header          *string    `gorm:"column:header;type:text;comment:'响应'"`                                    // 响应
	CreatedAt       *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

// TableName get sql table name.获取数据库表名
func (m *MeProfitSharingWechatOrderLog) TableName() string {
	return "me_profit_sharing_wechat_order_log"
}

// MeProfitSharingWechatOrderLogColumns get sql column name.获取数据库列名
var MeProfitSharingWechatOrderLogColumns = struct {
	ID              string
	ProfitSharingID string
	OrderNo         string
	Status          string
	Body            string
	Header          string
	CreatedAt       string
}{
	ID:              "id",
	ProfitSharingID: "profit_sharing_id",
	OrderNo:         "order_no",
	Status:          "status",
	Body:            "body",
	Header:          "header",
	CreatedAt:       "created_at",
}
