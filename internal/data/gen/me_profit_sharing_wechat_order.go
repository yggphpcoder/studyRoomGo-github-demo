package gen

import "time"

// MeProfitSharingWechatOrder 微信分账订单表
type MeProfitSharingWechatOrder struct {
	ID               int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"`                                  // 主键编码
	OrderNo          string     `gorm:"column:order_no;type:char(20);not null;default:'';comment:'分账单编号'"`                                            // 分账单编号
	TransactionID    string     `gorm:"column:transaction_id;type:varchar(50);not null;default:'';comment:'交易号'"`                                     // 交易号
	WechatOrderID    string     `gorm:"column:wechat_order_id;type:varchar(50);not null;default:'';comment:'微信分账单号'"`                                 // 微信分账单号
	OrderDesc        string     `gorm:"column:order_desc;type:varchar(255);not null;default:'';comment:'订单描述'"`                                       // 订单描述
	ReceiversType    string     `gorm:"column:receivers_type;type:varchar(32);not null;default:wechat;comment:'分账接收方类型:MERCHANT_ID,PERSONAL_OPENID'"` // 分账接收方类型:MERCHANT_ID,PERSONAL_OPENID
	ReceiversAccount string     `gorm:"column:receivers_account;type:varchar(64);not null;default:'';comment:'分账接收方账号'"`                              // 分账接收方账号
	ReceiversName    string     `gorm:"column:receivers_name;type:varchar(255);not null;default:'';comment:'接受方名字'"`                                  // 接受方名字
	MerchantID       int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商户id'"`                                             // 商户id
	ShopID           int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                                                 // 店铺id
	RelateOrderType  int8       `gorm:"column:relate_order_type;type:tinyint;not null;default:0;comment:'关联订单类型 1:订座单'"`                              // 关联订单类型 1:订座单
	RelateOrderNo    string     `gorm:"column:relate_order_no;type:char(20);not null;default:'';comment:'分账关联单号'"`                                    // 分账关联单号
	Status           int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'分账订单状态'"`                                               // 分账订单状态
	UnfreezeUnsplit  int8       `gorm:"column:unfreeze_unsplit;type:tinyint;not null;default:1;comment:'是否解冻剩余未分资金'"`                                 // 是否解冻剩余未分资金
	TotalAmount      int        `gorm:"column:total_amount;type:int;not null;default:0;comment:'订单总金额'"`                                              // 订单总金额
	Ratio            int8       `gorm:"column:ratio;type:tinyint;not null;default:0;comment:'分账比例'"`                                                  // 分账比例
	Amount           int        `gorm:"column:amount;type:int;not null;default:0;comment:'分账金额'"`                                                     // 分账金额
	CreatedAt        *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt        *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt        *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy         *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy         *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *MeProfitSharingWechatOrder) TableName() string {
	return "me_profit_sharing_wechat_order"
}

// MeProfitSharingWechatOrderColumns get sql column name.获取数据库列名
var MeProfitSharingWechatOrderColumns = struct {
	ID               string
	OrderNo          string
	TransactionID    string
	WechatOrderID    string
	OrderDesc        string
	ReceiversType    string
	ReceiversAccount string
	ReceiversName    string
	MerchantID       string
	ShopID           string
	RelateOrderType  string
	RelateOrderNo    string
	Status           string
	UnfreezeUnsplit  string
	TotalAmount      string
	Ratio            string
	Amount           string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
	CreateBy         string
	UpdateBy         string
}{
	ID:               "id",
	OrderNo:          "order_no",
	TransactionID:    "transaction_id",
	WechatOrderID:    "wechat_order_id",
	OrderDesc:        "order_desc",
	ReceiversType:    "receivers_type",
	ReceiversAccount: "receivers_account",
	ReceiversName:    "receivers_name",
	MerchantID:       "merchant_id",
	ShopID:           "shop_id",
	RelateOrderType:  "relate_order_type",
	RelateOrderNo:    "relate_order_no",
	Status:           "status",
	UnfreezeUnsplit:  "unfreeze_unsplit",
	TotalAmount:      "total_amount",
	Ratio:            "ratio",
	Amount:           "amount",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
	CreateBy:         "create_by",
	UpdateBy:         "update_by",
}
