package test

import (
	"fmt"
	"testing"

	"github.com/hjd919/gom"
)

// Category 分类
type Category struct {
	ID         uint   `gorm:"primaryKey;column:id;type:int unsigned;not null"`
	ParentID   int    `gorm:"column:parent_id;type:int;not null"`                      // 分类父id
	Grade      int    `gorm:"column:grade;type:int;not null"`                          // 级别
	Name       string `gorm:"index:index_name;column:name;type:varchar(255);not null"` // 分类名称
	Icon       string `gorm:"column:icon;type:varchar(255);not null"`                  // 分类图标
	Slug       string `gorm:"index:index_slug;column:slug;type:varchar(128);not null"` // 简称
	Type       string `gorm:"column:type;type:varchar(64);not null"`                   // 类型
	Sort       int    `gorm:"column:sort;type:int;not null"`                           // 排序
	RoleID     string `gorm:"column:role_id;type:varchar(64);not null"`                // 角色
	Status     int16  `gorm:"column:status;type:smallint;not null"`                    // 状态
	CreateTime int    `gorm:"column:create_time;type:int"`                             // 创建时间
}

// init index tmysqlt
func TestMysql_Conn(t *testing.T) {
	// 连接数据库
	gom.Connect(&gom.MysqlConfig{
		User:     "root",
		Password: "Yisai726",
		Host:     "101.200.41.141",
		Port:     "7306",
		Database: "hjd",
		Charset:  "utf8mb4",
	})
	// 操作表
	c := &Category{}
	gom.DB().Last(c)
	fmt.Println(gom.JsonEncode(c))
}

// 生成表结构
// gormt -H=101.200.41.140 -d=hjd -p=Yisai726 -u=root --port=3306 -F=false

// aggs tmysqlt
