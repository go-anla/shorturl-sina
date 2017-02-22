package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// Result is ...
type Result struct {
	// URLShort
	URLShort string `json:"url_short"`
}

func main() {
	argNum := len(os.Args)
	if argNum != 2 {
		fmt.Println("Useage: shorturl https://anla.io")
		return
	}

	urlLong := os.Args[1]
	if !strings.HasPrefix(urlLong, "http://") && !strings.HasPrefix(urlLong, "https://") {
		fmt.Println("The url should start with http:// or https://")
		return
	}

	// urlLong := "https://anla.io"

	resp, err := http.Get("http://api.t.sina.com.cn/short_url/shorten.json?source=3271760578&url_long=" + urlLong)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	res := &[]Result{}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		fmt.Println(err)
	} else {
		a := *res
		fmt.Println(a[0].URLShort)
	}
	return
}
