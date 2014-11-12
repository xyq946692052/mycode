package main

/*
 * 生成sign签名
 */
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

type MapSorter []Item

type Item struct {
	Key string
	Val string
}

func NewMapSorter(m map[string]string) MapSorter {
	ms := make(MapSorter, 0, len(m))
	for k, v := range m {
		ms = append(ms, Item{k, v})
	}
	return ms
}

func (ms MapSorter) Len() int {
	return len(ms)
}

func (ms MapSorter) Less(i, j int) bool {
	// return ms[i].Val<ms[j].Val  // 按值排序
	return ms[i].Key < ms[j].Key // 按键排序
}

func (ms MapSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}

/*
 *MD5加密
 */
func md5Encry(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	str := hex.EncodeToString(h.Sum(nil))
	return str
}

func main() {
	m := map[string]string{
		"appid":            "wxd930khk43hkhkff5443d",
		"auth_code":        "123456",
		"body":             "test",
		"device_info":      "123",
		"mch_id":           "19000000109",
		"sub_mch_id":       "321",
		"nonce_str":        "jhaksdjhfiyi68768p32ikjg23",
		"total_fee":        "1",
		"spbill_create_ip": "127.0.0.1",
		"key":              "3rgk2j3gkj23g4kj23gkj324g43",
	}

	ms := NewMapSorter(m)
	sort.Sort(ms)

	str1 := ""
	for _, item := range ms {
		// fmt.Printf("%s=%s&",item.Key,item.Val) // 获得map排序结果
		str1 += item.Key + "=" + item.Val + "&"
	}
	key := "99dsas9gds98dfr98dvbdfsvb9a"
	str2 := str1 + "key=" + key

	str3 := md5Encry(str2)
	fmt.Println("md5加密后:", str3)

	//将字符串中的字母转换成大写
	str4 := strings.ToUpper(str3)
	fmt.Println(str4)

}
