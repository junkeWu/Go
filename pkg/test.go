package main

import (
    "fmt"
    "net/http"
	"io/ioutil"
	"time"
)

func main() {
	weather()

}
func weather()  {
	fmt.Println("getting.....")
	fmt.Println("time.Now():",time.Now().Unix())
    res, err := http.Get("https://uaqy.api.storeapi.net/api/74/186")
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