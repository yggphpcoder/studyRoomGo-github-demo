package gen

import "time"

// SrOrder 订单表
type SrOrder struct {
	ID                int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"`         // 主键编码
	OrderNo           string     `gorm:"unique;column:order_no;type:char(20);not null;default:'';comment:'订单编号'"`             // 订单编号
	OrderDesc         string     `gorm:"column:order_desc;type:varchar(255);not null;default:'';comment:'订单描述'"`              // 订单描述
	MemberID          int64      `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`                      // 会员id
	MerchantID        int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商户id'"`                    // 商户id
	ShopID            int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                        // 店铺id
	OrderType         int8       `gorm:"column:order_type;type:tinyint;not null;default:0;comment:'订单类型 1:购买套餐，2:直接订座'"`      // 订单类型 1:购买套餐，2:直接订座
	OrderTypeRelateID string     `gorm:"column:order_type_relate_id;type:varchar(255);not null;default:0;comment:'订单类型关联信息'"` // 订单类型关联信息
	PayProvider       string     `gorm:"column:pay_provider;type:varchar(50);not null;default:'';comment:'支付方式：wxpay'"`       // 支付方式：wxpay
	Status            int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'订单状态'"`                        // 订单状态
	Price             float64    `gorm:"column:price;type:float(10,4);not null;default:0.0000;comment:'订单金额'"`                // 订单金额
	Quantity          int        `gorm:"column:quantity;type:int;not null;default:0;comment:'数量'"`                            // 数量
	CreatedAt         *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt         *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt         *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy          *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy          *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrOrder) TableName() string {
	return "sr_order"
}

// SrOrderColumns get sql column name.获取数据库列名
var SrOrderColumns = struct {
	ID                string
	OrderNo           string
	OrderDesc         string
	MemberID          string
	MerchantID        string
	ShopID            string
	OrderType         string
	OrderTypeRelateID string
	PayProvider       string
	Status            string
	Price             string
	Quantity          string
	CreatedAt         string
	UpdatedAt         string
	DeletedAt         string
	CreateBy          string
	UpdateBy          string
}{
	ID:                "id",
	OrderNo:           "order_no",
	OrderDesc:         "order_desc",
	MemberID:          "member_id",
	MerchantID:        "merchant_id",
	ShopID:            "shop_id",
	OrderType:         "order_type",
	OrderTypeRelateID: "order_type_relate_id",
	PayProvider:       "pay_provider",
	Status:            "status",
	Price:             "price",
	Quantity:          "quantity",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	DeletedAt:         "deleted_at",
	CreateBy:          "create_by",
	UpdateBy:          "update_by",
}
