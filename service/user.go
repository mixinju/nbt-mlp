package service

import (
    "fmt"
    "log"
    "strings"

    "nbt-mlp/dao/model"

    "github.com/mozillazg/go-pinyin"

    "go.uber.org/zap"
)

var LOG, _ = zap.NewProduction()

func Login() {
}

func Register(u model.User) error {
    // 先校验基本参数是否缺失
    if u.Id == 0 || len(u.Name) == 0 || len(u.ClassName) == 0 {
        log.Printf("注册用户失败: %s", "缺少必要的参数")
        return fmt.Errorf("缺少必要的参数")
    }

    // 检查对应的学号是否存在

    // 自动生成密码-默认密码:学号添加姓名拼音
    u.Password = generateDefaultPassword(u.Id, u.Name)

    // 如果密码长度不符合要求,则返回错误

    // 其他注册逻辑...

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
    // 批量注册- 使用 go协程的方式
    // 使用 channel 作为通信方式,通知注册结果
    for _, user := range users {
        go Register(user)
        fmt.Println(user)
    }
}
