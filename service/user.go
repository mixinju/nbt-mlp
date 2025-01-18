package service

import (
    "fmt"
    "log"
    "strconv"
    "strings"
    "sync"

    "nbt-mlp/dao"
    "nbt-mlp/middleware"
    "nbt-mlp/util"
    "nbt-mlp/util/errno"

    "nbt-mlp/dao/model"

    "github.com/mozillazg/go-pinyin"

    "go.uber.org/zap"
)

var logger, _ = zap.NewProduction()

func Login(userId uint64, password string) (string, errno.Errno) {

    condition := map[string]interface{}{"userid": userId}
    fields := []string{"password", "name", "class_name"}
    u := dao.GetUserDaoInstance().QueryById(condition, fields)
    hashPassword, err := util.HashPassword(password)
    if err != nil {
        logger.Error("Hash password failed", zap.Error(err))
        return "", errno.Errno{}
    }

    ok := util.ComparePassword(u.Password, hashPassword)
    if !ok {
        return "", *errno.ErrPasswordIncorrect
    }

    // 为登录用户设置 token
    token, err := middleware.SetUpToken(int64(userId))
    if err != nil {
        return "", *errno.ErrTokenSetUpFail
    }

    return token, *errno.OK
}

func Register(u model.User) error {
    // 先校验基本参数是否缺失
    if u.ID == 0 || len(u.Name) == 0 || len(u.ClassName) == 0 {
        log.Printf("注册用户失败: %s", "缺少必要的参数")
        return fmt.Errorf("缺少必要的参数")
    }

    // 检查对应的学号是否存在

    // 自动生成密码-默认密码:学号添加姓名拼音
    defaultPassword := generateDefaultPassword(u.ID, u.Name)
    u.Password, _ = util.HashPassword(defaultPassword)

    // $2a$10$MBT36F3XF00SxFb5yfnWL.HnOzfys3ccbyktRbNgzz6y1O6llTaNi
    // 如果密码长度不符合要求,则返回错误

    // 其他注册逻辑...

    ok := dao.GetUserDaoInstance().Create(u)
    if !ok {
        logger.Error("Create user failed",
            zap.String("userId", strconv.FormatUint(u.ID, 10)),
            zap.String("userName", u.Name),
        )
    }

    return nil
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

// BatchRegister 批量注册

func BatchRegister(users []model.User) {
    var wg sync.WaitGroup
    var mu sync.Mutex
    var failedUsers []model.User

    wg.Add(len(users))
    for _, user := range users {
        go func(user model.User) {
            defer wg.Done()
            err := Register(user)
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
        logger.Info("Batch registration failed")
    }
}

func ModifyPassWord(userId uint64, oldPassWord string, newPassWord string) errno.Errno {
    condition := map[string]interface{}{"userid": userId}
    fields := []string{"password"}
    u := dao.GetUserDaoInstance().QueryById(condition, fields)
    ok := util.ComparePassword(u.Password, oldPassWord)
    if !ok {
        return *errno.ErrPasswordNotMatch
    }
    hashPassword, err := util.HashPassword(newPassWord)
    if err != nil {
        logger.Error("Hash password failed", zap.Error(err))
        return *errno.ErrPasswordGenerate
    }
    u.Password = hashPassword
    ok = dao.GetUserDaoInstance().Update(u)
    if !ok {
        logger.Error("Update user password failed",
            zap.String("userId", strconv.FormatUint(u.ID, 10)),
            zap.String("userName", u.Name),
        )
    }
    return *errno.OK

}
