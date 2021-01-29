package gom

import (
	"fmt"
	"reflect"

	"github.com/tealeg/xlsx/v3"
)

// excel行的demo结构体，tag中必须含有xlsx,从0开始、以及comment标题
type ExcelRowDemo struct {
	Name      string `xlsx:"0" comment:"名称"`
	ExpiresIn string `xlsx:"1" comment:"到期时间"`
}

// excel类
type Excel struct {
	file   *xlsx.File //文件资源句柄
	titles []string   // 标题
}

// 创建excel
func NewExcel() *Excel {
	return &Excel{}
}

// 导出单个sheet数据
type ExportParams struct {
	FilePath  string        // 导出文件全路径
	SheetName string        // sheet名
	Rows      []interface{} // 导出数据
}

func (e *Excel) getTitles(stru interface{}) {
	t := reflect.TypeOf(stru).Elem()
	for i := 0; i < t.NumField(); i++ {
		e.titles = append(e.titles, t.Field(i).Tag.Get("comment")) // 标题的tag必须含有title
	}
}

// 导出excel
func (e *Excel) Export(params *ExportParams) (err error) {
	// 检测导入数据是否空
	if len(params.Rows) == 0 {
		return fmt.Errorf("list is empty")
	}

	// 创建excel
	if e.file == nil {
		e.file = xlsx.NewFile()
	}
	file := e.file

	// sheet
	sheet, err := file.AddSheet(params.SheetName)
	if err != nil {
		return
	}

	//title
	e.getTitles(params.Rows[0])
	row := sheet.AddRow()
	row.WriteSlice(e.titles, -1)

	// data
	for _, vorRow := range params.Rows {
		row := sheet.AddRow()
		row.WriteStruct(vorRow, -1)
	}

	//save file
	err = file.Save(params.FilePath)
	if err != nil {
		return
	}
	return
}

// 导入excel
type ImportParams struct {
	FilePath  string                  //导出的文件全路径
	SheetName string                  //页签
	RowHandle func(r *xlsx.Row) error //页签处理函数
}

func (e *Excel) Import(params *ImportParams) (err error) {
	// 创建excel
	if e.file == nil {
		e.file, err = xlsx.OpenFile(params.FilePath)
		if err != nil {
			return
		}
	}
	file := e.file

	// sheet
	sh, ok := file.Sheet[params.SheetName]
	if !ok {
		return fmt.Errorf("sheet not found")
	}

	// 处理data
	err = sh.ForEachRow(params.RowHandle)
	if err != nil {
		return
	}
	return
}

// 导入所有页签的excel
type ImportAllParams struct {
	FilePath   string                             //导出的文件全路径
	RowHandles map[string]func(r *xlsx.Row) error // map:页签=>处理数据函数
}

func (e *Excel) ImportAll(params *ImportAllParams) (err error) {
	// 创建excel
	if e.file == nil {
		e.file, err = xlsx.OpenFile(params.FilePath)
		if err != nil {
			return
		}
	}
	file := e.file

	// 遍历所有页签
	for _, sh := range file.Sheets {
		rowHandle, ok := params.RowHandles[sh.Name]
		if !ok {
			return fmt.Errorf("no sheetname:%s", sh.Name)
		}
		err = sh.ForEachRow(rowHandle)
		if err != nil {
			return
		}
	}
	return
}
