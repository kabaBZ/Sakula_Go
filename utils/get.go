package utils

import (
	"fmt"
	"github.com/kabaBZ/Sakula_Go/handle_error"
	"io"
	"net/http"
)

/*
http请求:
· 发起百度搜索的GET请求：“http://www.baidu.com/s?wd=ReganYue”，打印回复的内容
· 对https://httpbin.org/post发起post请求，携带表单数据(application/x-www-form-urlencoded),
建值自拟(strings.NewReader("money=0&hobby=撩妹")),打印回复的内容
· 表单数据类型application/x-www-form-urlencoded，表单读取API：strings.NewReader("money=0&hobby=撩妹")
*/
func Get() {
	req, _ := http.NewRequest("GET", "http://www.baidu.com/s?wd=ReganYue", nil)
	req.Header.Set("token", "d8cdcf8427e")
	// 再设置个json
	req.Header.Set("Content-Type", "application/json")
	resp, err := (&http.Client{}).Do(req)
	HandleError(err, "http.Get")

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	fmt.Println(string(bytes))
	Post()

}
