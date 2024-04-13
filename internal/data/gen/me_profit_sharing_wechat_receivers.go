package gen

import "time"

// MeProfitSharingWechatReceivers 微信分账接受账号
type MeProfitSharingWechatReceivers struct {
	ID               int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"`                                       // 主键编码
	MerchantID       int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商户id'"`                                                  // 商户id
	ShopID           int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                                                      // 店铺id
	ReceiversType    string     `gorm:"column:receivers_type;type:varchar(32);not null;default:MERCHANT_ID;comment:'分账接收方类型:MERCHANT_ID,PERSONAL_OPENID'"` // 分账接收方类型:MERCHANT_ID,PERSONAL_OPENID
	ReceiversAccount string     `gorm:"column:receivers_account;type:varchar(64);not null;default:'';comment:'分账接收方账号'"`                                   // 分账接收方账号
	ReceiversName    string     `gorm:"column:receivers_name;type:varchar(255);not null;default:0;comment:'接受方名字'"`                                        // 接受方名字
	RelateOrderType  int8       `gorm:"column:relate_order_type;type:tinyint;not null;default:0;comment:'关联订单类型 1:订座单'"`                                   // 关联订单类型 1:订座单
	Ratio            int8       `gorm:"column:ratio;type:tinyint;not null;default:0;comment:'分账比例'"`                                                       // 分账比例
	Status           int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'状态'"`                                                        // 状态
	CreatedAt        *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt        *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt        *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy         *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy         *uint      `gorm:"column:update_by;type:int unsigned"`
	Remark           string     `gorm:"column:remark;type:varchar(500);not null;default:'';comment:'备注'"` // 备注
}

// TableName get sql table name.获取数据库表名
func (m *MeProfitSharingWechatReceivers) TableName() string {
	return "me_profit_sharing_wechat_receivers"
}

// MeProfitSharingWechatReceiversColumns get sql column name.获取数据库列名
var MeProfitSharingWechatReceiversColumns = struct {
	ID               string
	MerchantID       string
	ShopID           string
	ReceiversType    string
	ReceiversAccount string
	ReceiversName    string
	RelateOrderType  string
	Ratio            string
	Status           string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
	CreateBy         string
	UpdateBy         string
	Remark           string
}{
	ID:               "id",
	MerchantID:       "merchant_id",
	ShopID:           "shop_id",
	ReceiversType:    "receivers_type",
	ReceiversAccount: "receivers_account",
	ReceiversName:    "receivers_name",
	RelateOrderType:  "relate_order_type",
	Ratio:            "ratio",
	Status:           "status",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
	CreateBy:         "create_by",
	UpdateBy:         "update_by",
	Remark:           "remark",
}
