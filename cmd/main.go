package main

import (
	"fmt"
	"time"

	"github.com/hjd919/gom"
)

type ExcelRow struct {
	Name      string `xlsx:"0" comment:"达人账号名称"`
	ExpiresIn string `xlsx:"1" comment:"授权到期时间"`
}

// 导出
func ExcelExport() {
	var exportRows = make([]interface{}, 0, 2) // 导入的数据
	exportRows = append(exportRows, &ExcelRow{Name: "a", ExpiresIn: "b"})
	exportRows = append(exportRows, &ExcelRow{Name: "c", ExpiresIn: "v"})
	err := gom.NewExcel().Export(&gom.ExportParams{
		FilePath:  "./text2.xlsx", // 导入文件名
		SheetName: "生活",           // 页签名
		Rows:      exportRows,     // 导入数据，二维数组
	})
	if err != nil {
		return
	}
}
func timerfc() {
	fmt.Println(time.Now().Format(time.RFC3339))
}

type captchaCache struct{}

// 设置验证码
func (c *captchaCache) Set(key string, randText string, expire int) {
}

// 获取验证码
func (c *captchaCache) Get(key string) (randText string) {
	return
}

// 删除验证码
func (c *captchaCache) Del(key string) {
}

// func gincaptcha() {
// 	// 设置图片验证码字体目录
// 	captcha.SetFontPath("../captcha/fonts")
// 	r := gin.Default()
// 	var cache captcha.Cache = &captchaCache{}
// 	r.GET("/ping", func(c *gin.Context) {
// 		captcha.New("register-uid", cache).Output(c.Writer, 150, 50)
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080
// }

func main() {
	// ExcelExport()
	// timerfc()

	// gincaptcha()
}
