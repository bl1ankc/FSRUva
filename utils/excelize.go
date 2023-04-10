package utils

import (
	"github.com/xuri/excelize/v2"
	"strconv"
)

const maxCharCount = 26

func ExportExcel(sheetName string, headers []string, data [][]interface{}) (*excelize.File, error) {
	//文件生成
	file := excelize.NewFile()

	//改名工作薄名字
	file.SetSheetName("Sheet1", sheetName)

	//工作薄获取
	sheetIndex, err := file.GetSheetIndex(sheetName)
	if err != nil {
		return nil, err
	}

	//单元格长度最大值设置
	maxColumnRowNameLen := 1 + len(strconv.Itoa(len(data))) //单元格长度
	columnCount := len(headers)                             //单元格个数
	if columnCount > maxCharCount {                         //在26以内是A0大于则AA0开始,所以+1
		maxColumnRowNameLen++
	} else if columnCount > maxCharCount*maxCharCount { //同理
		maxColumnRowNameLen += 2
	}

	columnNames := make([][]byte, 0, columnCount)
	for index, header := range headers {
		//先获取字母位置
		columnName := getColumnName(index, maxColumnRowNameLen)
		columnNames = append(columnNames, columnName)
		//再获取字母+数字
		curColumnName := getColumnRowName(columnName, 1)
		err := file.SetCellValue(sheetName, curColumnName, header)
		if err != nil {
			return nil, err
		}
	}
	for rowIndex, row := range data {
		for columnIndex, columnName := range columnNames {
			// 从第二行开始(rowIndex+2,rowIndex从0开始)
			err := file.SetCellValue(sheetName, getColumnRowName(columnName, rowIndex+2), row[columnIndex])
			if err != nil {
				return nil, err
			}
		}
	}
	file.SetActiveSheet(sheetIndex)
	return file, nil
}

// getColumnName 生成列名
// Excel的列名规则是从A-Z往后排;超过Z以后用两个字母表示，比如AA,AB,AC;两个字母不够以后用三个字母表示，比如AAA,AAB,AAC
// 这里做数字到列名的映射：0 -> A, 1 -> B, 2 -> C
// maxColumnRowNameLen 表示名称框的最大长度，假设数据是10行，1000列，则最后一个名称框是J1000(如果有表头，则是J1001),是4位
// 这里根据 maxColumnRowNameLen 生成切片，后面生成名称框的时候可以复用这个切片，而无需扩容
func getColumnName(column, maxColumnRowNameLen int) []byte {
	const A = 'A'
	if column < maxCharCount {
		// 第一次就分配好切片的容量
		slice := make([]byte, 0, maxColumnRowNameLen)
		return append(slice, byte(A+column))
	} else {
		// 递归生成类似AA,AB,AAA,AAB这种形式的列名
		return append(getColumnName(column/maxCharCount-1, maxColumnRowNameLen), byte(A+column%maxCharCount))
	}
}

// getColumnRowName 生成名称框
// Excel的名称框是用A1,A2,B1,B2来表示的，这里需要传入前一步生成的列名切片，然后直接加上行索引来生成名称框，就无需每次分配内存
func getColumnRowName(columnName []byte, rowIndex int) (columnRowName string) {
	l := len(columnName)
	columnName = strconv.AppendInt(columnName, int64(rowIndex), 10)
	columnRowName = string(columnName)
	// 将列名恢复回去
	columnName = columnName[:l]
	return
}
