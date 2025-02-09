package application

import (
    "fmt"
    "strconv"

    "github.com/pkg/errors"
    "github.com/xuri/excelize/v2"
    "go.uber.org/zap"
    "gorm.io/gorm"
    "nbt-mlp/common/util"
    "nbt-mlp/common/util/errno"
    "nbt-mlp/domain/entity"
    "nbt-mlp/domain/repository"
)

var log, _ = zap.NewProduction()

type UserAppIface interface {
    Query(id uint64) (entity.User, *errno.Errno)
    Delete(id uint64) *errno.Errno
    QueryUsers(id []uint64) ([]entity.User, *errno.Errno)
    Update(u entity.User) *errno.Errno
    Save(u entity.User) *errno.Errno
    BatchSave(us []entity.User) *errno.Errno
    ReadFromExcelFile(filePath string) ([]entity.User, *errno.Errno)
    QueryUserByIdAndPassword(id uint64, password string) (entity.User, *errno.Errno)
}

type UserAppImpl struct {
    userRepo repository.UserRepositoryIface
}

func (u UserAppImpl) Delete(id uint64) *errno.Errno {
    //TODO implement me
    panic("implement me")
}

func (u UserAppImpl) QueryUserByIdAndPassword(id uint64, password string) (entity.User, *errno.Errno) {
    user, err := u.userRepo.Query(id)

    if errors.Is(err, gorm.ErrRecordNotFound) {
        return entity.User{}, errno.ErrUserNotFound
    }
    hashPassword, _ := util.HashPassword(password)
    if util.ComparePassword(user.Password, hashPassword) {
        return user, nil
    }

    return entity.User{}, errno.ErrPasswordNotMatch

}

func NewUserAppImpl(up repository.UserRepositoryIface) UserAppIface {
    return &UserAppImpl{userRepo: up}
}

var _ UserAppIface = &UserAppImpl{}

func (u UserAppImpl) Query(id uint64) (entity.User, *errno.Errno) {
    //TODO implement me
    panic("implement me")
}

//func (u UserAppImpl) Delete(id uint64) *errno.Errno {
//    //TODO implement me
//    panic("implement me")
//}

func (u UserAppImpl) QueryUsers(id []uint64) ([]entity.User, *errno.Errno) {
    //TODO implement me
    panic("implement me")
}

func (u UserAppImpl) Update(ut entity.User) *errno.Errno {
    //TODO implement me
    panic("implement me")
}

func (u UserAppImpl) Save(ut entity.User) *errno.Errno {
    //TODO implement me
    panic("implement me")
}

func (u UserAppImpl) BatchSave(us []entity.User) *errno.Errno {
    //TODO implement me
    panic("implement me")
}

func (u UserAppImpl) ReadFromExcelFile(filePath string) ([]entity.User, *errno.Errno) {
    // Open the Excel file
    f, err := excelize.OpenFile(filePath)
    if err != nil {
        return nil, errno.ErrFileOpen
    }
    defer func(f *excelize.File) {
        err := f.Close()
        if err != nil {
            log.Sugar().Errorf("关闭文件失败: %v", err)
        }
    }(f)

    // Read the rows from the first sheet
    rows, err := f.GetRows(f.GetSheetName(0))
    if err != nil {
        return nil, errno.ErrFileOpen
    }

    // 文件内容为空
    if len(rows) <= 1 {
        return nil, errno.ErrFileContentEmpty
    }

    // Process the rows

    result := make([]entity.User, 0, len(rows)-1)
    for index, row := range rows {
        u, e := resolveUser(row, index)
        if e != nil {
            log.Sugar().Errorf("解析单个用户信息失败:文件路径{%s};原始数据 {%s};错误信息:{%s}", filePath, row, e.Error())
        }
        result = append(result, u)
    }

    return result, nil
}

func resolveUser(row []string, rowNum int) (entity.User, error) {
    var u entity.User
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

    u.ID = id
    u.ClassName = className
    u.Name = name

    return u, nil
}
