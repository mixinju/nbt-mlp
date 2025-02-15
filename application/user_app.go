package application

import (
    "fmt"
    "strconv"
    "strings"
    "sync"

    "nbt-mlp/Infrastructure/persistence"
    "nbt-mlp/common/util"
    "nbt-mlp/common/util/errno"
    "nbt-mlp/domain/entity"
    "nbt-mlp/domain/repository"

    "github.com/mozillazg/go-pinyin"
    "github.com/pkg/errors"
    "github.com/xuri/excelize/v2"
    "go.uber.org/zap"
    "gorm.io/gorm"
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
    // 数据库持久化
    userRepo repository.UserRepositoryIface
}

func NewUserAppImpl() UserAppIface {
    return &UserAppImpl{userRepo: persistence.NewUserDao()}
}

var _ UserAppIface = &UserAppImpl{}

func (u *UserAppImpl) Delete(id uint64) *errno.Errno {
    // Check if user exists
    _, err := u.userRepo.Query(id)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return errno.ErrUserNotFound
    }

    // Delete the user
    err = u.userRepo.Delete(id)
    if err != nil {
        log.Error("Delete user failed", zap.String("userId", strconv.FormatUint(id, 10)))
        return errno.ErrDatabase
    }

    return nil
}

// QueryUserByIdAndPassword 注意传递的明文密码
func (u *UserAppImpl) QueryUserByIdAndPassword(id uint64, password string) (entity.User, *errno.Errno) {
    user, err := u.userRepo.Query(id)

    if errors.Is(err, gorm.ErrRecordNotFound) {
        return entity.User{}, errno.ErrUserNotFound
    }

    if !util.ComparePassword(user.Password, password) {
        return entity.User{}, errno.ErrPasswordNotMatch
    }

    return user, nil
}

func (u *UserAppImpl) Query(id uint64) (entity.User, *errno.Errno) {
    user, err := u.userRepo.Query(id)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return entity.User{}, errno.ErrUserNotFound
    }
    if err != nil {
        return entity.User{}, errno.ErrDatabase
    }
    return user, nil
}

func (u *UserAppImpl) QueryUsers(ids []uint64) ([]entity.User, *errno.Errno) {
    users, err := u.userRepo.QueryUsers(ids)
    // TODO 错误处理有问题
    if err != nil {
        return nil, errno.ErrDatabase
    }
    return users, nil
}

func (u *UserAppImpl) Update(ut entity.User) *errno.Errno {
    // 检查用户是否存在
    _, err := u.userRepo.Query(ut.ID)
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return errno.ErrUserNotFound
    }

    err = ut.Check()
    if err != nil {
        return errno.ErrUserParamsInvalid
    }

    // 更新用户
    err = u.userRepo.Update(ut)
    if err != nil {
        log.Error("Update user failed", zap.String("userId", strconv.FormatUint(ut.ID, 10)))
        return errno.ErrDatabase
    }

    return nil
}

func (u *UserAppImpl) Save(ut entity.User) *errno.Errno {
    // 先校验基本参数是否缺失
    // 这部分放到 entity.User 处理
    if ut.ID == 0 || len(ut.Name) == 0 || len(ut.ClassName) == 0 {
        log.Error("注册用户失败: 缺少必要的参数")
    }

    // 检查对应的学号是否存在

    // 自动生成密码-默认密码:学号添加姓名拼音
    defaultPassword := generateDefaultPassword(ut.ID, ut.Name)
    ut.Password, _ = util.HashPassword(defaultPassword)
    log.Info("生成默认密码", zap.String("password", ut.Password))

    // $2a$10$MBT36F3XF00SxFb5yfnWL.HnOzfys3ccbyktRbNgzz6y1O6llTaNi
    // 如果密码长度不符合要求,则返回错误

    // 其他注册逻辑...

    err := u.userRepo.Save(ut)
    if err != nil {
        log.Error("Create user failed",
            zap.String("userId", strconv.FormatUint(ut.ID, 10)),
            zap.String("userName", ut.Name),
        )
    }

    return nil
}

func (u *UserAppImpl) BatchSave(us []entity.User) *errno.Errno {
    var wg sync.WaitGroup
    var mu sync.Mutex
    var failedUsers []entity.User

    // TODO 直接使用数据的批量插入即可,无需使用协程
    wg.Add(len(us))
    for _, user := range us {
        go func(user entity.User) {
            defer wg.Done()
            err := u.Save(user)
            if err != nil {
                mu.Lock()
                failedUsers = append(failedUsers, user)
                mu.Unlock()
            }
        }(user)
    }
    wg.Wait()

    if len(failedUsers) != 0 {
        // TODO: Send email notification
        // logger.Info("Batch registration failed")
    }
    return nil
}

func (u *UserAppImpl) ReadFromExcelFile(filePath string) ([]entity.User, *errno.Errno) {
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

func generateDefaultPassword(id uint64, name string) string {
    a := pinyin.NewArgs()
    pinyinName := pinyin.Pinyin(name, a)
    flatPinyinName := flattenPinyin(pinyinName)
    return fmt.Sprintf("%d%s", id, flatPinyinName)
}

func flattenPinyin(p [][]string) string {
    var result []string
    for _, s := range p {
        result = append(result, s...)
    }
    return strings.Join(result, "")
}
