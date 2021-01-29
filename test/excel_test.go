package test

import (
	"log"
	"testing"

	"github.com/hjd919/gom"
	"github.com/tealeg/xlsx/v3"
)

type ExcelRow struct {
	Name      string `xlsx:"0" comment:"达人账号名称"`
	ExpiresIn string `xlsx:"1" comment:"授权到期时间"`
}

// 导出
func TestExcelExport(t *testing.T) {
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

// 导入
func TestExcelImport(t *testing.T) {
	excelRows := []ExcelRow{}
	err := gom.NewExcel().Import(&gom.ImportParams{
		SheetName: "生活",
		FilePath:  "./text2.xlsx",
		RowHandle: func(r *xlsx.Row) error { //handle each row
			excelRow := ExcelRow{}
			err := r.ReadStruct(&excelRow)
			if err != nil {
				t.Error(err)
				return err
			}
			excelRows = append(excelRows, excelRow)
			return nil
		},
	})
	if err != nil {
		t.Error(err)
	}

	// 第一行是标题，从第二行开始为数据
	excelRows = excelRows[1:]

	log.Println(excelRows)
}

// 导入多个sheet
func TestExcelImportAll(t *testing.T) {
	excelRows := []ExcelRow{}
	var rowHandles = make(map[string]func(r *xlsx.Row) error)
	rowHandles["生活"] = func(r *xlsx.Row) error { //handle each row
		excelRow := ExcelRow{}
		err := r.ReadStruct(&excelRow)
		if err != nil {
			t.Error(err)
			return err
		}
		excelRows = append(excelRows, excelRow)
		return nil
	}
	err := gom.NewExcel().ImportAll(&gom.ImportAllParams{
		FilePath:   "./text2.xlsx",
		RowHandles: rowHandles,
	})
	if err != nil {
		t.Error(err)
	}

	// 第一行是标题，从第二行开始为数据
	excelRows = excelRows[1:]

	log.Println(excelRows)
}
