package gen

import "time"

// SrPaymentWechat 交易信息表
type SrPaymentWechat struct {
	ID                  int64      `gorm:"autoIncrement:true;primaryKey;column:id;type:bigint;not null;comment:'主键编码'"`                                            // 主键编码
	OutTradeNo          string     `gorm:"index:sr_payment_wechat_out_trade_no_index;column:out_trade_no;type:char(20);not null;default:'';comment:'商户订单号'"`       // 商户订单号
	TransactionID       string     `gorm:"index:sr_payment_wechat_transaction_id_index;column:transaction_id;type:char(32);not null;default:'';comment:'微信支付订单号'"` // 微信支付订单号
	TradeType           string     `gorm:"column:trade_type;type:varchar(20);not null;default:'';comment:'交易类型'"`                                                  // 交易类型
	TradeState          string     `gorm:"column:trade_state;type:varchar(32);not null;default:'';comment:'交易状态'"`                                                 // 交易状态
	TradeStateDesc      string     `gorm:"column:trade_state_desc;type:varchar(256);not null;default:'';comment:'交易状态描述	'"`                                        // 交易状态描述
	BankType            string     `gorm:"column:bank_type;type:varchar(32);not null;default:'';comment:'付款银行'"`                                                   // 付款银行
	SuccessTime         string     `gorm:"column:success_time;type:varchar(64);not null;default:'';comment:'支付完成时间'"`                                              // 支付完成时间
	PayerOpenid         string     `gorm:"column:payer_openid;type:varchar(128);not null;default:'';comment:'用户标识'"`                                               // 用户标识
	AmountTotal         int        `gorm:"column:amount_total;type:int;not null;default:0;comment:'总金额,分'"`                                                        // 总金额,分
	AmountPayerTotal    int        `gorm:"column:amount_payer_total;type:int;not null;default:0;comment:'用户支付金额,分'"`                                               // 用户支付金额,分
	AmountCurrency      string     `gorm:"column:amount_currency;type:varchar(16);not null;default:'';comment:'货币类型'"`                                             // 货币类型
	AmountPayerCurrency string     `gorm:"column:amount_payer_currency;type:varchar(16);not null;default:'';comment:'用户支付币种'"`                                     // 用户支付币种
	Attach              string     `gorm:"column:attach;type:varchar(128);not null;default:'';comment:'附加数据'"`                                                     // 附加数据
	CreatedAt           *time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt           *time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt           *time.Time `gorm:"column:deleted_at;type:timestamp;comment:'删除时间'"` // 删除时间
	CreateBy            *uint      `gorm:"column:create_by;type:int unsigned"`
	UpdateBy            *uint      `gorm:"column:update_by;type:int unsigned"`
}

// TableName get sql table name.获取数据库表名
func (m *SrPaymentWechat) TableName() string {
	return "sr_payment_wechat"
}

// SrPaymentWechatColumns get sql column name.获取数据库列名
var SrPaymentWechatColumns = struct {
	ID                  string
	OutTradeNo          string
	TransactionID       string
	TradeType           string
	TradeState          string
	TradeStateDesc      string
	BankType            string
	SuccessTime         string
	PayerOpenid         string
	AmountTotal         string
	AmountPayerTotal    string
	AmountCurrency      string
	AmountPayerCurrency string
	Attach              string
	CreatedAt           string
	UpdatedAt           string
	DeletedAt           string
	CreateBy            string
	UpdateBy            string
}{
	ID:                  "id",
	OutTradeNo:          "out_trade_no",
	TransactionID:       "transaction_id",
	TradeType:           "trade_type",
	TradeState:          "trade_state",
	TradeStateDesc:      "trade_state_desc",
	BankType:            "bank_type",
	SuccessTime:         "success_time",
	PayerOpenid:         "payer_openid",
	AmountTotal:         "amount_total",
	AmountPayerTotal:    "amount_payer_total",
	AmountCurrency:      "amount_currency",
	AmountPayerCurrency: "amount_payer_currency",
	Attach:              "attach",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
	DeletedAt:           "deleted_at",
	CreateBy:            "create_by",
	UpdateBy:            "update_by",
}
