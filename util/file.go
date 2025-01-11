package util

import (
	"fmt"
	"strconv"

	"nbt-mlp/dao/model"
	"nbt-mlp/util/errno"

	"github.com/xuri/excelize/v2"
	"go.uber.org/zap"
)

var log, _ = zap.NewProduction()

func ReadUserInfoFromExcel(filePath string) ([]model.User, errno.Errno) {
	// Open the Excel file
	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return nil, *errno.ErrFileOpen
	}
	defer f.Close()

	// Read the rows from the first sheet
	rows, err := f.GetRows(f.GetSheetName(0))
	if err != nil {
		return nil, *errno.ErrFileOpen
	}

	// 文件内容为空
	if len(rows) <= 1 {
		return nil, *errno.ErrFileContentEmpty
	}

	// Process the rows

	result := make([]model.User, 0, len(rows)-1)
	for index, row := range rows {
		u, e := resoloveUser(row, index)
		if e != nil {
			log.Sugar().Errorf("解析单个用户信息失败:文件路径{};原始数据 {};错误信息:{}", filePath, row, e.Error())
		}
		result = append(result, u)
	}

	return result, *errno.OK
}

func resoloveUser(row []string, rowNum int) (model.User, error) {
	var u model.User
	className := row[0]
	name := row[2]

	if len(className) >= 20 || len(name) >= 20 {
		log.Sugar().Errorf("上传数据出错")
		return u, fmt.Errorf("行号:{%d};错误信息:{%s}", rowNum, "数据列内容过长")
	}
	id, err := strconv.ParseUint(row[1], 10, 64)
	if err != nil {
		return u, fmt.Errorf("行号:{%d};错误信息:{%s}", rowNum, "学号转换失败")
	}

	u.Id = id
	u.ClassName = className
	u.Name = name

	return u, nil
}
