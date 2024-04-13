package gen

// DianpingShop 大众点评门店关联信息
type DianpingShop struct {
	ID           int    `gorm:"autoIncrement:true;primaryKey;column:id;type:int;not null"`
	MerchantID   int64  `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'系统商家id'"`          // 系统商家id
	ShopID       int64  `gorm:"unique;column:shop_id;type:bigint;not null;default:0;comment:'系统shop_id'"`    // 系统shop_id
	OpenShopUUID string `gorm:"unique;column:open_shop_uuid;type:varchar(255);not null;comment:'门店id的唯一标识'"` // 门店id的唯一标识
	ShopName     string `gorm:"column:shop_name;type:varchar(255);not null;comment:'门店名称'"`                  // 门店名称
	BranchName   string `gorm:"column:branch_name;type:varchar(255);not null;default:'';comment:'分店名称'"`     // 分店名称
	ShopAddress  string `gorm:"column:shop_address;type:varchar(255);not null;default:'';comment:'门店地址'"`    // 门店地址
	CityName     string `gorm:"column:city_name;type:varchar(255);not null;default:'';comment:'所在城市'"`       // 所在城市
}

// TableName get sql table name.获取数据库表名
func (m *DianpingShop) TableName() string {
	return "dianping_shop"
}

// DianpingShopColumns get sql column name.获取数据库列名
var DianpingShopColumns = struct {
	ID           string
	MerchantID   string
	ShopID       string
	OpenShopUUID string
	ShopName     string
	BranchName   string
	ShopAddress  string
	CityName     string
}{
	ID:           "id",
	MerchantID:   "merchant_id",
	ShopID:       "shop_id",
	OpenShopUUID: "open_shop_uuid",
	ShopName:     "shop_name",
	BranchName:   "branch_name",
	ShopAddress:  "shop_address",
	CityName:     "city_name",
}
