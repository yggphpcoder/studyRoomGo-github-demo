package gen

import "time"

// SrMemberPackageCard 会员-套餐关系表
type SrMemberPackageCard struct {
	ID           int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'会员-套餐id'"` // 会员-套餐id
	MerchantID   int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`               // 商家id
	ShopID       int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                   // 店铺id
	MemberID     int64      `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`                 // 会员id
	ActualAmount float64    `gorm:"column:actual_amount;type:float(10,4);not null;default:0.0000;comment:'实收价格'"`   // 实收价格
	RemainMinute int        `gorm:"column:remain_minute;type:int;not null;default:0;comment:'剩余分钟'"`                // 剩余分钟
	RemainDay    int        `gorm:"column:remain_day;type:int;not null;default:0;comment:'剩余天数'"`                   // 剩余天数
	MinutePer    int        `gorm:"column:minute_per;type:int;not null;default:0;comment:'每天可用分钟'"`                 // 每天可用分钟
	MinuteMax    int        `gorm:"column:minute_max;type:int;not null;default:0;comment:'封顶可用分钟'"`                 // 封顶可用分钟
	MinuteMin    int        `gorm:"column:minute_min;type:int;not null;default:0;comment:'最低可用分钟'"`                 // 最低可用分钟
	Status       int8       `gorm:"column:status;type:tinyint;not null;default:0;comment:'状态'"`                     // 状态
	TypeChannel  int8       `gorm:"column:type_channel;type:tinyint;not null;default:0;comment:'渠道类型'"`             // 渠道类型
	ActiveAt     *time.Time `gorm:"column:active_at;type:timestamp;comment:'激活时间'"`                                 // 激活时间
	InvalidAt    *time.Time `gorm:"column:invalid_at;type:datetime;comment:'失效时间'"`                                 // 失效时间
	PackageID    *int64     `gorm:"column:package_id;type:bigint;comment:'套餐id'"`                                   // 套餐id
	PackageName  string     `gorm:"column:package_name;type:varchar(255);not null;default:'';comment:'套餐名'"`        // 套餐名
	TypePackage  int8       `gorm:"column:type_package;type:tinyint;not null;default:0;comment:'套餐类型'"`             // 套餐类型
	TypeCard     int8       `gorm:"column:type_card;type:tinyint;not null;default:0;comment:'卡类型'"`                 // 卡类型
	TypeSale     int8       `gorm:"column:type_sale;type:tinyint;not null;default:0;comment:'销售类型'"`                // 销售类型
	Price        float64    `gorm:"column:price;type:float(10,4);not null;default:0.0000;comment:'售价'"`             // 售价
	OriPrice     float64    `gorm:"column:ori_price;type:float(10,4);not null;default:0.0000;comment:'原价'"`         // 原价
	UseShop      string     `gorm:"column:use_shop;type:varchar(500);not null;default:'';comment:'适用店铺'"`           // 适用店铺
	UseSeat      string     `gorm:"column:use_seat;type:varchar(500);not null;default:'';comment:'适用座位'"`           // 适用座位
	UseTime      string     `gorm:"column:use_time;type:varchar(500);not null;default:'';comment:'使用时段'"`           // 使用时段
	UseBalance   *int8      `gorm:"column:use_balance;type:tinyint;comment:'使用余额支付'"`                               // 使用余额支付
	ValidDay     int        `gorm:"column:valid_day;type:int;not null;default:0;comment:'有效天数'"`                    // 有效天数
	ActiveLimit  int        `gorm:"column:active_limit;type:int;not null;default:0;comment:'激活期限天数'"`               // 激活期限天数
	Sort         int        `gorm:"column:sort;type:int;not null;default:0;comment:'排序'"`                           // 排序
	Remark       *string    `gorm:"column:remark;type:varchar(500);comment:'备注'"`                                   // 备注
	CreatedAt    *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt    *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy     *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy     *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrMemberPackageCard) TableName() string {
	return "sr_member_package_card"
}

// SrMemberPackageCardColumns get sql column name.获取数据库列名
var SrMemberPackageCardColumns = struct {
	ID           string
	MerchantID   string
	ShopID       string
	MemberID     string
	ActualAmount string
	RemainMinute string
	RemainDay    string
	MinutePer    string
	MinuteMax    string
	MinuteMin    string
	Status       string
	TypeChannel  string
	ActiveAt     string
	InvalidAt    string
	PackageID    string
	PackageName  string
	TypePackage  string
	TypeCard     string
	TypeSale     string
	Price        string
	OriPrice     string
	UseShop      string
	UseSeat      string
	UseTime      string
	UseBalance   string
	ValidDay     string
	ActiveLimit  string
	Sort         string
	Remark       string
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
	CreateBy     string
	UpdateBy     string
}{
	ID:           "id",
	MerchantID:   "merchant_id",
	ShopID:       "shop_id",
	MemberID:     "member_id",
	ActualAmount: "actual_amount",
	RemainMinute: "remain_minute",
	RemainDay:    "remain_day",
	MinutePer:    "minute_per",
	MinuteMax:    "minute_max",
	MinuteMin:    "minute_min",
	Status:       "status",
	TypeChannel:  "type_channel",
	ActiveAt:     "active_at",
	InvalidAt:    "invalid_at",
	PackageID:    "package_id",
	PackageName:  "package_name",
	TypePackage:  "type_package",
	TypeCard:     "type_card",
	TypeSale:     "type_sale",
	Price:        "price",
	OriPrice:     "ori_price",
	UseShop:      "use_shop",
	UseSeat:      "use_seat",
	UseTime:      "use_time",
	UseBalance:   "use_balance",
	ValidDay:     "valid_day",
	ActiveLimit:  "active_limit",
	Sort:         "sort",
	Remark:       "remark",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
	CreateBy:     "create_by",
	UpdateBy:     "update_by",
}
