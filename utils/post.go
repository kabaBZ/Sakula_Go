package utils

import (
	"fmt"
	"github.com/kabaBZ/Sakula_Go/handle_eror"
	"io"
	"net/http"
	"strings"
)

/*
http请求:
· 对https://httpbin.org/post发起post请求，携带表单数据(application/x-www-form-urlencoded),
建值自拟(strings.NewReader("money=0&hobby=撩妹")),打印回复的内容
· 表单数据类型application/x-www-form-urlencoded，表单读取API：strings.NewReader("money=0&hobby=撩妹")
*/
func Post() {
	resp, err := http.Post("https://httpbin.org/post?name=Regan&money=0",
		"application/x-www-form-urlencoded",
		strings.NewReader("姓名=ReganYue&人民币=0&心愿=发财发财"))
	HandleError(err, "http.Post")

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	fmt.Println(string(bytes))
}
