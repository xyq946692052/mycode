package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("https://api.weixin.qq.com/cgi-bin/user/info?access_token=9IXwRvfstXADjnBA3P7J5X4KTp9mkdpqYUYDYte7TFyRjeru0RE6oSXDjjoN155fEf2votp67Tnn00Bj-1KsDA&openid=o8jUTuACQ2_8HeuDLGvMORXpETeU")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var dat map[string]interface{}
	err := json.Unmarshal([]byte(string(body)), &dat)
	if err == nil {
		fmt.Println(dat)
	}
}
