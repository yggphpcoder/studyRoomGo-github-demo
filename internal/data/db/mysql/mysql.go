package mysql

import (
	"errors"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"reflect"
	"strings"
	"time"

	gMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm(source string, loggerWriter io.Writer) (db *gorm.DB, err error) {

	db, err = gorm.Open(gMysql.Open(source), &gorm.Config{
		Logger: logger.New(log.New(loggerWriter, "", log.LstdFlags), logger.Config{
			SlowThreshold:             300 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		}),
	})

	return db, err
}

// 组装查询字段
func SelectField(model interface{}, filter []string, alias string) (selectField []string) {
	if len(filter) == 0 {
		if alias != "" {
			return []string{alias + ".*"}
		}
		return []string{"*"}
	}
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Struct {
		var fieldName = make(map[string]string)
		for k := 0; k < t.NumField(); k++ {
			f := t.Field(k)
			n := f.Name
			TagSettings := schema.ParseTagSetting(f.Tag.Get("gorm"), ";")
			if alias != "" {
				if TagSettings["COLUMN"] != "" {
					n = alias + "." + TagSettings["COLUMN"]
				}
			} else {
				if TagSettings["COLUMN"] != "" {
					n = TagSettings["COLUMN"]
				}
			}
			fieldName[strings.ToUpper(f.Name)] = n
		}
		for _, v := range filter {
			v = strings.ReplaceAll(v, "_", "")
			if n, ok := fieldName[strings.ToUpper(v)]; ok {

				selectField = append(selectField, n)
			}
		}
	}
	return selectField
}

// SelectFieldCount 返回字段数量
// exclude 排除计算
func SelectFieldCount(fields []string, exclude []string) (count int) {
	count = len(fields)
	if len(exclude) > 0 {
		result := make([]string, 0)
		// 去重
		flagMap := make(map[string]bool, 0)
		for _, a := range fields {
			if _, ok := flagMap[a]; ok {
				continue
			}
			flagMap[a] = true
			allow := true
			for _, b := range exclude {
				if a == b { //包含字符
					allow = false
				}
			}
			if allow {
				result = append(result, a)
			}
		}
		count = len(result)
	}

	return count
}

/*
BuildWhere 构建where条件
["字段名","操作符","查询值","与前一个条件的关系[默认and]"]

1.如果是等于，可以省略"操作符" [默认and，如果是and，可以不写]:
[]interface{}{"username", "chen"}或[]interface{}{"username","=" , "chen"}

2.大于:
[]interface{}{"createtime", ">", "2019-1-1"}

3.如果为or，那就必须写全，4个部分都不能少：
[]interface{}{"username", "=", "chen", "or"}

4.其它的where兼容gorm的where方法

5.特殊说明
[]interface{}{"username = ? or nickname = ?", "chen", "yond"}这种拼接参数，在封装时做了特殊处理
*/
func BuildWhere(db *gorm.DB, where interface{}) (*gorm.DB, error) {
	//var err error
	t := reflect.TypeOf(where).Kind()
	if t == reflect.Struct || t == reflect.Map {
		db = db.Where(where)
	} else if t == reflect.Slice {
		for _, item := range where.([]interface{}) {
			item := item.([]interface{})
			column := item[0]
			if reflect.TypeOf(column).Kind() == reflect.String {
				count := len(item)
				if count == 1 {
					return nil, errors.New("切片长度不能小于2")
				}
				columnstr := column.(string)
				// 拼接参数形式
				if strings.Index(columnstr, "?") > -1 {
					db = db.Where(column, item[1:]...)
				} else {
					cond := "and" //cond
					opt := "="
					_opt := " = "
					var val interface{}
					if count == 2 {
						opt = "="
						val = item[1]
					} else {
						opt = strings.ToLower(item[1].(string))
						_opt = " " + strings.ReplaceAll(opt, " ", "") + " "
						val = item[2]
					}

					if count == 4 {
						cond = strings.ToLower(strings.ReplaceAll(item[3].(string), " ", ""))
					}

					/*
					   '=', '<', '>', '<=', '>=', '<>', '!=', '<=>',
					   'like', 'like binary', 'not like', 'ilike',
					   '&', '|', '^', '<<', '>>',
					   'rlike', 'regexp', 'not regexp',
					   '~', '~*', '!~', '!~*', 'similar to',
					   'not similar to', 'not ilike', '~~*', '!~~*',
					*/

					if strings.Index(" in notin ", _opt) > -1 {
						// val 是数组类型
						column = columnstr + " " + opt + " (?)"
					} else if strings.Index(" = < > <= >= <> != <=> like likebinary notlike ilike rlike regexp notregexp", _opt) > -1 {
						column = columnstr + " " + opt + " ?"
					}

					if cond == "and" {
						db = db.Where(column, val)
					} else {
						db = db.Or(column, val)
					}
				}
			}
			//else if t == reflect.Map /*Map*/ {
			//	db = db.Where(item)
			//} else {
			//	/*
			//		// 解决and 与 or 混合查询，但这种写法有问题，会抛出 invalid query condition
			//		db = db.Where(func(db *gorm.DB) *gorm.DB {
			//			db, err = BuildWhere(db, item)
			//			if err != nil {
			//				panic(err)
			//			}
			//			return db
			//		})*/
			//
			//	db, err = BuildWhere(db, item)
			//	if err != nil {
			//		return nil, err
			//	}
			//}
		}
	} else {
		return nil, errors.New("参数有误")
	}
	return db, nil
}

func BuildQueryList(db *gorm.DB, wheres interface{}, columns interface{}, orderBy interface{}, page, rows int) (*gorm.DB, error) {
	var err error
	db, err = BuildWhere(db, wheres)
	if err != nil {
		return nil, err
	}
	db = db.Select(columns)
	if orderBy != nil && orderBy != "" {
		db = db.Order(orderBy)
	}
	if page > 0 && rows > 0 {
		db = db.Limit(rows).Offset((page - 1) * rows)
	}
	return db, err
}

func BuildSort(list map[string]string, alias string) (str string) {
	for field, sort := range list {
		if alias != "" {
			field = alias + "." + field
		}
		n := field + " " + sort
		if str == "" {
			str = n
		} else {
			str = str + "," + n
		}
	}
	return str
}
