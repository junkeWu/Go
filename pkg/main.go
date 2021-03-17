package main

import (
    "fmt"
    "net/http"
	"io/ioutil"
	"flag"
	"os"
	"strings"
)

func main() {
	const number = 9
	var cityCode string
    flag.StringVar(&cityCode,"cityCode","","cityCode's value can't null")
    flag.Parse()

    for _, a := range os.Args[1:] {
		var b []string
		b = strings.Split(a,"=")
		if b[0] != "-cityCode" {
			fmt.Println("请按格式书写命令参数")
			break
		}
		n := len(b[1])
		if n != number {
			fmt.Println("请输入正确的城市代码")
			break
		}
		code := b[1] + ".html"
		weather(code)
	}
	

}
func weather(code string)  {
	fmt.Println("getting.....")
    res, err := http.Get("http://www.weather.com.cn/data/sk/" + code)
    fmt.Println("getting.....")
    if err != nil {
        fmt.Println("some err")
        fmt.Println(err)

    } else {
        getRes, gerErr := ioutil.ReadAll(res.Body)
        if gerErr != nil {
            fmt.Println(gerErr)
        } else {
                fmt.Println("some res")
                fmt.Println(string(getRes))
                }

    }
}