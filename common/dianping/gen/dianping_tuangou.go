package gen

// DianpingTuangou 大众点评团购关联信息
type DianpingTuangou struct {
	ID            int     `gorm:"autoIncrement:true;primaryKey;column:id;type:int;not null"`
	MerchantID    int64   `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`                 // 商家id
	PackageCardID int64   `gorm:"column:package_card_id;type:bigint;not null;default:0;comment:'系统套餐id'"`           // 系统套餐id
	OpenShopUUID  string  `gorm:"column:open_shop_uuid;type:varchar(255);not null;default:'';comment:'门店id的唯一标识 '"` // 门店id的唯一标识
	DealID        int64   `gorm:"column:deal_id;type:bigint;not null;comment:'套餐id	'"`                              // 套餐id
	DealgroupID   int64   `gorm:"column:dealgroup_id;type:bigint;not null;comment:'团购id'"`                          // 团购id
	Title         string  `gorm:"column:title;type:varchar(255);not null;default:'';comment:'套餐名称'"`                // 套餐名称
	Price         float64 `gorm:"column:price;type:float(10,4);not null;default:0.0000;comment:'套餐价格'"`             // 套餐价格
	Marketprice   float64 `gorm:"column:marketprice;type:float(10,4);not null;default:0.0000;comment:'套餐原价'"`       // 套餐原价
}

// TableName get sql table name.获取数据库表名
func (m *DianpingTuangou) TableName() string {
	return "dianping_tuangou"
}

// DianpingTuangouColumns get sql column name.获取数据库列名
var DianpingTuangouColumns = struct {
	ID            string
	MerchantID    string
	PackageCardID string
	OpenShopUUID  string
	DealID        string
	DealgroupID   string
	Title         string
	Price         string
	Marketprice   string
}{
	ID:            "id",
	MerchantID:    "merchant_id",
	PackageCardID: "package_card_id",
	OpenShopUUID:  "open_shop_uuid",
	DealID:        "deal_id",
	DealgroupID:   "dealgroup_id",
	Title:         "title",
	Price:         "price",
	Marketprice:   "marketprice",
}
