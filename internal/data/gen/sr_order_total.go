package gen

// SrOrderTotal 订单金额明细表
type SrOrderTotal struct {
	ID          int64   `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'id'"` // id
	OrderID     int64   `gorm:"column:order_id;type:bigint;not null;default:0;comment:'订单id'"`             // 订单id
	Type        string  `gorm:"column:type;type:varchar(20);not null;default:'';comment:'金额类型'"`           // 金额类型
	Value       float64 `gorm:"column:value;type:decimal(10,2);not null;default:0.00;comment:'金额值'"`       // 金额值
	RelateType  int8    `gorm:"column:relate_type;type:tinyint;not null;default:0;comment:'关联类型'"`         // 关联类型
	RelateValue string  `gorm:"column:relate_value;type:varchar(50);not null;default:'';comment:'关联值'"`    // 关联值
}

// TableName get sql table name.获取数据库表名
func (m *SrOrderTotal) TableName() string {
	return "sr_order_total"
}

// SrOrderTotalColumns get sql column name.获取数据库列名
var SrOrderTotalColumns = struct {
	ID          string
	OrderID     string
	Type        string
	Value       string
	RelateType  string
	RelateValue string
}{
	ID:          "id",
	OrderID:     "order_id",
	Type:        "type",
	Value:       "value",
	RelateType:  "relate_type",
	RelateValue: "relate_value",
}
