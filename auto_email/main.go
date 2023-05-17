package main

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func main() {
	mailHeader := map[string][]string{
		"From":    {"2131646378@qq.com"},
		"To":      {"yangbaixin200211@163.com"},
		"Subject": {"标题"},
	}

	m := gomail.NewMessage()
	m.SetHeaders(mailHeader)
	//m.SetAddressHeader("Cc", "shengang@pingcap.com", "shengang")
	m.SetBody("text/html", "尊敬的客户")
	//m.Attach("../../go.mod")  // 添加附件

	d := gomail.NewDialer("smtp.qq.com", 465, "2131646378@qq.com", "xnhnubgwugpxfbah")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	fmt.Println("send email success!")
}
