package util

import "fmt"

// 邮箱服务初始化
func init() {

}

func SendTo(body string, e string) {
	fmt.Println("Sending email to " + e)

}
