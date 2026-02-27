package excelPg

import "github.com/xuri/excelize/v2"

// readTableHead 为是否读取表头，true为读取，false为不读取
//
// ReadXlsx 获得批量导入Excel表中的数据列表
func ReadXlsx(excelFilePath string, readTableHead bool) (err error, dataList [][]string) {
	excelFile, errOpen := excelize.OpenFile(excelFilePath)
	if errOpen != nil {
		return errOpen, nil
	}
	sheetList := excelFile.GetSheetList()
	for _, sheet := range sheetList {
		rows, errSheet := excelFile.GetRows(sheet)
		cols, _ := excelFile.GetCols(sheet)
		colLength := len(cols)
		if errSheet != nil {
			return errSheet, nil
		}
		for rowIndex, row := range rows {
			// 若设置了不读第一行则跳过
			if rowIndex == 0 && !readTableHead {
				continue
			}
			var dataRow []string
			for _, colCell := range row {
				dataRow = append(dataRow, colCell)
			}
			// 为了解决一个Bug
			if len(dataRow) < colLength {
				needPlusCount := colLength - len(dataRow)
				for i := 0; i < needPlusCount; i++ {
					dataRow = append(dataRow, "")
				}
			}
			dataList = append(dataList, dataRow)
		}
	}
	return err, dataList
}
