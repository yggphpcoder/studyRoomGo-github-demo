package gen

import "time"

// EdSalemanSalary 教务-销售员-工资表
type EdSalemanSalary struct {
	ID               int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"`    // 主键编码
	MerchantID       int64      `gorm:"column:merchant_id;type:bigint;not null;default:0;comment:'商家id'"`               // 商家id
	ShopID           int64      `gorm:"column:shop_id;type:bigint;not null;default:0;comment:'店铺id'"`                   // 店铺id
	SalemanID        int64      `gorm:"column:saleman_id;type:bigint;not null;default:0;comment:'销售员id'"`               // 销售员id
	SalaryAmount     float64    `gorm:"column:salary_amount;type:decimal(10,2);not null;default:0.00;comment:'课程工资总额'"` // 课程工资总额
	SalaryTime       *time.Time `gorm:"column:salary_time;type:timestamp;comment:'发放工资时间'"`                             // 发放工资时间
	Bonus            float64    `gorm:"column:bonus;type:decimal(10,2);not null;default:0.00;comment:'奖金'"`             // 奖金
	Subsidies        float64    `gorm:"column:subsidies;type:decimal(10,2);not null;default:0.00;comment:'补贴'"`         // 补贴
	Deductions       float64    `gorm:"column:deductions;type:decimal(10,2);not null;default:0.00;comment:'扣款'"`        // 扣款
	TotalSalary      float64    `gorm:"column:total_salary;type:decimal(10,2);not null;default:0.00;comment:'总工资'"`     // 总工资
	SalaryStatus     int8       `gorm:"column:salary_status;type:tinyint;not null;default:0;comment:'工资状态'"`            // 工资状态
	SalaryRemark     *string    `gorm:"column:salary_remark;type:varchar(500);comment:'工资备注'"`                          // 工资备注
	BonusRemark      *string    `gorm:"column:bonus_remark;type:varchar(500);comment:'奖金备注'"`                           // 奖金备注
	SubsidiesRemark  *string    `gorm:"column:subsidies_remark;type:varchar(500);comment:'补贴备注'"`                       // 补贴备注
	DeductionsRemark *string    `gorm:"column:deductions_remark;type:varchar(500);comment:'扣款备注'"`                      // 扣款备注
	Remark           *string    `gorm:"column:remark;type:varchar(500);comment:'备注'"`                                   // 备注
	CreatedAt        *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt        *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt        *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy         *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy         *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *EdSalemanSalary) TableName() string {
	return "ed_saleman_salary"
}

// EdSalemanSalaryColumns get sql column name.获取数据库列名
var EdSalemanSalaryColumns = struct {
	ID               string
	MerchantID       string
	ShopID           string
	SalemanID        string
	SalaryAmount     string
	SalaryTime       string
	Bonus            string
	Subsidies        string
	Deductions       string
	TotalSalary      string
	SalaryStatus     string
	SalaryRemark     string
	BonusRemark      string
	SubsidiesRemark  string
	DeductionsRemark string
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
	SalemanID:        "saleman_id",
	SalaryAmount:     "salary_amount",
	SalaryTime:       "salary_time",
	Bonus:            "bonus",
	Subsidies:        "subsidies",
	Deductions:       "deductions",
	TotalSalary:      "total_salary",
	SalaryStatus:     "salary_status",
	SalaryRemark:     "salary_remark",
	BonusRemark:      "bonus_remark",
	SubsidiesRemark:  "subsidies_remark",
	DeductionsRemark: "deductions_remark",
	Remark:           "remark",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
	CreateBy:         "create_by",
	UpdateBy:         "update_by",
}
