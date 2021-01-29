package gom

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"time"

	"github.com/jinzhu/gorm"
)

// 文档
// http://gorm.book.jasperxu.com/database.html#dbc
// 生成结构体
// https://github.com/xxjwxc/gormt

var __mysql *gomMysql

// 获取mysql连接
func DB() *gorm.DB {
	return __mysql.db
}

// 连接mysql
func Connect(conf *MysqlConfig) {
	__mysql = newGomMysql(conf)
}

// 获取mysql排序
func Order(sortType, sortField string) string {
	if sortType != "" && sortField != "" {
		return sortType + " " + sortField
	} else if sortType != "" && sortField == "" {
		return sortType
	}
	return "id desc"
}

type MysqlConfig struct {
	Type        string
	User        string
	Password    string
	Host        string
	Port        string
	Database    string
	Charset     string
	TablePrefix string
}

// -inner
type gomMysql struct {
	db   *gorm.DB
	conf *MysqlConfig
}

func newGomMysql(conf *MysqlConfig) *gomMysql {
	self := &gomMysql{
		conf: conf,
	}
	var err error
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.Charset))

	if err != nil {
		panic(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:create_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	self.db = db

	// heath
	// go self.ping()

	return self
}

// func (self *gomMysql) ping() {
// 	for {
// 		select {
// 		case <-self.ctx.Done():
// 			// 程序退出，断开bulkProcessor
// 			log.Println("程序退出，关闭db")
// 			self.db.Close()
// 			return
// 		}
// 	}
// }

// updateTimeStampForCreateCallback will set `CreateTime`, `UpdateTime` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreateTime"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdateTime"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `UpdateTime` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdateTime", time.Now().Unix())
	}
}

// deleteCallback will set `DeleteTime` where deleting
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deleteTimeField, hasDeleteTimeField := scope.FieldByName("DeleteTime")

		if !scope.Search.Unscoped && hasDeleteTimeField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deleteTimeField.DBName),
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
