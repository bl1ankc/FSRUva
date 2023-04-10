package utils

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
	"io"
	"net/http"
	"time"
)

/*
	xlsx的包更为方便,但
*/

// ResponseXls 导出Excel
func ResponseXls(c *gin.Context, content io.ReadSeeker, fileTag string) {
	fileName := fmt.Sprintf("%s%s%s.xlsx", NowString(), `-`, fileTag)
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fileName))
	c.Writer.Header().Add("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	http.ServeContent(c.Writer, c.Request, fileName, time.Now(), content)

}

// ToExcel 生成io.ReadSeeker  参数 titleList 为Excel表头，dataList 为数据
func ToExcel(titleList []string, dataList []interface{}, SheetName string) (content io.ReadSeeker) {
	// 生成一个新的文件
	file := xlsx.NewFile()
	// 添加sheet页
	sheet, _ := file.AddSheet(SheetName)
	// 插入表头
	titleRow := sheet.AddRow()
	for _, v := range titleList {
		cell := titleRow.AddCell()
		cell.Value = v
	}
	// 插入内容
	for _, v := range dataList {
		row := sheet.AddRow()
		//row.WriteSlice(v, -1)
		row.WriteStruct(v, -1)
	}

	var buffer bytes.Buffer
	_ = file.Write(&buffer)
	content = bytes.NewReader(buffer.Bytes())
	return
}
