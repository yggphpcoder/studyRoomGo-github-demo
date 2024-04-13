package gen

import "time"

// SrPackageCardMemberBuyLog 套餐-会员购买记录
type SrPackageCardMemberBuyLog struct {
	ID               int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'套餐-会员购买记录id'"` // 套餐-会员购买记录id
	MerchantID       int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`                   // 商家id
	ShopID           int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                       // 店铺id
	OrderNo          string     `gorm:"column:order_no;type:char(20);not null;default:'';comment:'订单编号'"`                   // 订单编号
	MemberID         int64      `gorm:"column:member_id;type:bigint;not null;default:0;comment:'会员id'"`                     // 会员id
	PackageCardID    int64      `gorm:"column:package_card_id;type:bigint;not null;default:0;comment:'套餐id'"`               // 套餐id
	MpID             int64      `gorm:"column:mp_id;type:bigint;not null;default:0;comment:'会员-套餐id'"`                      // 会员-套餐id
	ActualAmount     float64    `gorm:"column:actual_amount;type:float(10,4);not null;default:0.0000;comment:'实收金额'"`       // 实收金额
	Price            float64    `gorm:"column:price;type:float(10,4);not null;default:0.0000;comment:'售价'"`                 // 售价
	OriPrice         float64    `gorm:"column:ori_price;type:float(10,4);not null;default:0.0000;comment:'原价'"`             // 原价
	Name             string     `gorm:"column:name;type:varchar(255);not null;default:'';comment:'套餐名'"`                    // 套餐名
	TypePackage      int8       `gorm:"column:type_package;type:tinyint;not null;default:0;comment:'套餐类型'"`                 // 套餐类型
	TypeCard         int8       `gorm:"column:type_card;type:tinyint;not null;default:0;comment:'卡类型'"`                     // 卡类型
	UseSeat          string     `gorm:"column:use_seat;type:varchar(500);not null;default:0;comment:'适用座位类型'"`              // 适用座位类型
	UseShop          string     `gorm:"column:use_shop;type:varchar(500);not null;default:0;comment:'适用店铺'"`                // 适用店铺
	TypeSale         int8       `gorm:"column:type_sale;type:tinyint;not null;default:0;comment:'销售类型'"`                    // 销售类型
	ValidDay         int        `gorm:"column:valid_day;type:int;not null;default:0;comment:'有效天数'"`                        // 有效天数
	ActiveLimit      int        `gorm:"column:active_limit;type:int;not null;default:0;comment:'激活期限天数'"`                   // 激活期限天数
	BuyLimit         int        `gorm:"column:buy_limit;type:int;not null;default:0;comment:'每人限购次数'"`                      // 每人限购次数
	UseTime          string     `gorm:"column:use_time;type:varchar(1000);not null;default:'';comment:'使用时段'"`              // 使用时段
	BookingDay       int        `gorm:"column:booking_day;type:int;not null;default:0;comment:'总可用天数'"`                     // 总可用天数
	BookingMinute    int        `gorm:"column:booking_minute;type:int;not null;default:0;comment:'总可用分钟'"`                  // 总可用分钟
	BookingMinutePer int        `gorm:"column:booking_minute_per;type:int;not null;default:0;comment:'每天可用分钟'"`             // 每天可用分钟
	BookingMinuteMax int        `gorm:"column:booking_minute_max;type:int;not null;default:0;comment:'封顶可用分钟'"`             // 封顶可用分钟
	BookingMinuteMin int        `gorm:"column:booking_minute_min;type:int;not null;default:0;comment:'最低可用分钟'"`             // 最低可用分钟
	CanBalance       *int8      `gorm:"column:can_balance;type:tinyint;comment:'可否用余额购买'"`                                  // 可否用余额购买
	Remark           *string    `gorm:"column:remark;type:varchar(255);comment:'备注'"`                                       // 备注
	CreatedAt        *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt        *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt        *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy         *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy         *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrPackageCardMemberBuyLog) TableName() string {
	return "sr_package_card_member_buy_log"
}

// SrPackageCardMemberBuyLogColumns get sql column name.获取数据库列名
var SrPackageCardMemberBuyLogColumns = struct {
	ID               string
	MerchantID       string
	ShopID           string
	OrderNo          string
	MemberID         string
	PackageCardID    string
	MpID             string
	ActualAmount     string
	Price            string
	OriPrice         string
	Name             string
	TypePackage      string
	TypeCard         string
	UseSeat          string
	UseShop          string
	TypeSale         string
	ValidDay         string
	ActiveLimit      string
	BuyLimit         string
	UseTime          string
	BookingDay       string
	BookingMinute    string
	BookingMinutePer string
	BookingMinuteMax string
	BookingMinuteMin string
	CanBalance       string
	Remark           string
	CreatedAt        string
	UpdatedAt        string
	DeletedAt        string
	CreateBy         string
	UpdateBy         string
}{
	ID:               "id",
	MerchantID:       "merchant_id",
	ShopID:           "shop_id",
	OrderNo:          "order_no",
	MemberID:         "member_id",
	PackageCardID:    "package_card_id",
	MpID:             "mp_id",
	ActualAmount:     "actual_amount",
	Price:            "price",
	OriPrice:         "ori_price",
	Name:             "name",
	TypePackage:      "type_package",
	TypeCard:         "type_card",
	UseSeat:          "use_seat",
	UseShop:          "use_shop",
	TypeSale:         "type_sale",
	ValidDay:         "valid_day",
	ActiveLimit:      "active_limit",
	BuyLimit:         "buy_limit",
	UseTime:          "use_time",
	BookingDay:       "booking_day",
	BookingMinute:    "booking_minute",
	BookingMinutePer: "booking_minute_per",
	BookingMinuteMax: "booking_minute_max",
	BookingMinuteMin: "booking_minute_min",
	CanBalance:       "can_balance",
	Remark:           "remark",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
	CreateBy:         "create_by",
	UpdateBy:         "update_by",
}
