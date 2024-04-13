package gen

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type _SrCouponsApplicableMgr struct {
	*_BaseMgr
}

// SrCouponsApplicableMgr open func
func SrCouponsApplicableMgr(db *gorm.DB) *_SrCouponsApplicableMgr {
	if db == nil {
		panic(fmt.Errorf("SrCouponsApplicableMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_SrCouponsApplicableMgr{_BaseMgr: &_BaseMgr{DB: db.Table("sr_coupons_applicable"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_SrCouponsApplicableMgr) GetTableName() string {
	return "sr_coupons_applicable"
}

// Reset 重置gorm会话
func (obj *_SrCouponsApplicableMgr) Reset() *_SrCouponsApplicableMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_SrCouponsApplicableMgr) Get() (result SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("sr_coupons").Where("id = ?", result.CouponID).Find(&result.SrCoupons).Error; err != nil { // 优惠券总表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// Gets 获取批量结果
func (obj *_SrCouponsApplicableMgr) Gets() (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_SrCouponsApplicableMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCouponID coupon_id获取 优惠券ID
func (obj *_SrCouponsApplicableMgr) WithCouponID(couponID int64) Option {
	return optionFunc(func(o *options) { o.query["coupon_id"] = couponID })
}

// WithApplicableOrderType applicable_order_type获取 适用订单类型
func (obj *_SrCouponsApplicableMgr) WithApplicableOrderType(applicableOrderType int8) Option {
	return optionFunc(func(o *options) { o.query["applicable_order_type"] = applicableOrderType })
}

// WithApplicableShop applicable_shop获取 适用店铺列表
func (obj *_SrCouponsApplicableMgr) WithApplicableShop(applicableShop string) Option {
	return optionFunc(func(o *options) { o.query["applicable_shop"] = applicableShop })
}

// WithApplicableShopType applicable_shop_type获取 适用店铺类型列表
func (obj *_SrCouponsApplicableMgr) WithApplicableShopType(applicableShopType string) Option {
	return optionFunc(func(o *options) { o.query["applicable_shop_type"] = applicableShopType })
}

// WithApplicableSeat applicable_seat获取 适用座位列表
func (obj *_SrCouponsApplicableMgr) WithApplicableSeat(applicableSeat string) Option {
	return optionFunc(func(o *options) { o.query["applicable_seat"] = applicableSeat })
}

// WithApplicableSeatType applicable_seat_type获取 适用座位类型列表
func (obj *_SrCouponsApplicableMgr) WithApplicableSeatType(applicableSeatType string) Option {
	return optionFunc(func(o *options) { o.query["applicable_seat_type"] = applicableSeatType })
}

// WithApplicablePackageCard applicable_package_card获取 适用套餐卡列表
func (obj *_SrCouponsApplicableMgr) WithApplicablePackageCard(applicablePackageCard string) Option {
	return optionFunc(func(o *options) { o.query["applicable_package_card"] = applicablePackageCard })
}

// WithApplicablePackageCardType applicable_package_card_type获取 适用套餐卡类型列表
func (obj *_SrCouponsApplicableMgr) WithApplicablePackageCardType(applicablePackageCardType string) Option {
	return optionFunc(func(o *options) { o.query["applicable_package_card_type"] = applicablePackageCardType })
}

// WithApplicableMemberType applicable_member_type获取 适用用户类型列表
func (obj *_SrCouponsApplicableMgr) WithApplicableMemberType(applicableMemberType string) Option {
	return optionFunc(func(o *options) { o.query["applicable_member_type"] = applicableMemberType })
}

// GetByOption 功能选项模式获取
func (obj *_SrCouponsApplicableMgr) GetByOption(opts ...Option) (result SrCouponsApplicable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where(options.query).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("sr_coupons").Where("id = ?", result.CouponID).Find(&result.SrCoupons).Error; err != nil { // 优惠券总表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_SrCouponsApplicableMgr) GetByOptions(opts ...Option) (results []*SrCouponsApplicable, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where(options.query).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromCouponID 通过coupon_id获取内容 优惠券ID
func (obj *_SrCouponsApplicableMgr) GetFromCouponID(couponID int64) (result SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`coupon_id` = ?", couponID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("sr_coupons").Where("id = ?", result.CouponID).Find(&result.SrCoupons).Error; err != nil { // 优惠券总表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// GetBatchFromCouponID 批量查找 优惠券ID
func (obj *_SrCouponsApplicableMgr) GetBatchFromCouponID(couponIDs []int64) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`coupon_id` IN (?)", couponIDs).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromApplicableOrderType 通过applicable_order_type获取内容 适用订单类型
func (obj *_SrCouponsApplicableMgr) GetFromApplicableOrderType(applicableOrderType int8) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_order_type` = ?", applicableOrderType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromApplicableOrderType 批量查找 适用订单类型
func (obj *_SrCouponsApplicableMgr) GetBatchFromApplicableOrderType(applicableOrderTypes []int8) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_order_type` IN (?)", applicableOrderTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromApplicableShop 通过applicable_shop获取内容 适用店铺列表
func (obj *_SrCouponsApplicableMgr) GetFromApplicableShop(applicableShop string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_shop` = ?", applicableShop).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromApplicableShop 批量查找 适用店铺列表
func (obj *_SrCouponsApplicableMgr) GetBatchFromApplicableShop(applicableShops []string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_shop` IN (?)", applicableShops).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromApplicableShopType 通过applicable_shop_type获取内容 适用店铺类型列表
func (obj *_SrCouponsApplicableMgr) GetFromApplicableShopType(applicableShopType string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_shop_type` = ?", applicableShopType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromApplicableShopType 批量查找 适用店铺类型列表
func (obj *_SrCouponsApplicableMgr) GetBatchFromApplicableShopType(applicableShopTypes []string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_shop_type` IN (?)", applicableShopTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromApplicableSeat 通过applicable_seat获取内容 适用座位列表
func (obj *_SrCouponsApplicableMgr) GetFromApplicableSeat(applicableSeat string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_seat` = ?", applicableSeat).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromApplicableSeat 批量查找 适用座位列表
func (obj *_SrCouponsApplicableMgr) GetBatchFromApplicableSeat(applicableSeats []string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_seat` IN (?)", applicableSeats).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromApplicableSeatType 通过applicable_seat_type获取内容 适用座位类型列表
func (obj *_SrCouponsApplicableMgr) GetFromApplicableSeatType(applicableSeatType string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_seat_type` = ?", applicableSeatType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromApplicableSeatType 批量查找 适用座位类型列表
func (obj *_SrCouponsApplicableMgr) GetBatchFromApplicableSeatType(applicableSeatTypes []string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_seat_type` IN (?)", applicableSeatTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromApplicablePackageCard 通过applicable_package_card获取内容 适用套餐卡列表
func (obj *_SrCouponsApplicableMgr) GetFromApplicablePackageCard(applicablePackageCard string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_package_card` = ?", applicablePackageCard).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromApplicablePackageCard 批量查找 适用套餐卡列表
func (obj *_SrCouponsApplicableMgr) GetBatchFromApplicablePackageCard(applicablePackageCards []string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_package_card` IN (?)", applicablePackageCards).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromApplicablePackageCardType 通过applicable_package_card_type获取内容 适用套餐卡类型列表
func (obj *_SrCouponsApplicableMgr) GetFromApplicablePackageCardType(applicablePackageCardType string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_package_card_type` = ?", applicablePackageCardType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromApplicablePackageCardType 批量查找 适用套餐卡类型列表
func (obj *_SrCouponsApplicableMgr) GetBatchFromApplicablePackageCardType(applicablePackageCardTypes []string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_package_card_type` IN (?)", applicablePackageCardTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetFromApplicableMemberType 通过applicable_member_type获取内容 适用用户类型列表
func (obj *_SrCouponsApplicableMgr) GetFromApplicableMemberType(applicableMemberType string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_member_type` = ?", applicableMemberType).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

// GetBatchFromApplicableMemberType 批量查找 适用用户类型列表
func (obj *_SrCouponsApplicableMgr) GetBatchFromApplicableMemberType(applicableMemberTypes []string) (results []*SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`applicable_member_type` IN (?)", applicableMemberTypes).Find(&results).Error
	if err == nil && obj.isRelated {
		for i := 0; i < len(results); i++ {
			if err = obj.NewDB().Table("sr_coupons").Where("id = ?", results[i].CouponID).Find(&results[i].SrCoupons).Error; err != nil { // 优惠券总表
				if err != gorm.ErrRecordNotFound { // 非 没找到
					return
				}
			}
		}
	}
	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_SrCouponsApplicableMgr) FetchByPrimaryKey(couponID int64) (result SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`coupon_id` = ?", couponID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("sr_coupons").Where("id = ?", result.CouponID).Find(&result.SrCoupons).Error; err != nil { // 优惠券总表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}

// FetchUniqueBySrCouponsApplicableCouponIDUIndex primary or index 获取唯一内容
func (obj *_SrCouponsApplicableMgr) FetchUniqueBySrCouponsApplicableCouponIDUIndex(couponID int64) (result SrCouponsApplicable, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(SrCouponsApplicable{}).Where("`coupon_id` = ?", couponID).First(&result).Error
	if err == nil && obj.isRelated {
		if err = obj.NewDB().Table("sr_coupons").Where("id = ?", result.CouponID).Find(&result.SrCoupons).Error; err != nil { // 优惠券总表
			if err != gorm.ErrRecordNotFound { // 非 没找到
				return
			}
		}
	}

	return
}
