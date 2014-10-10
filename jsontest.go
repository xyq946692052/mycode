package main

import (
  "fmt"
  "encoding/json"
  "os"
)

type ConfigStruct struct{
  Host               string  `json:"host"`
  Port               int     `json:"port"`
  AnalyticsFile      string  `json:"analytics_file"`
  StaticFileVersion  int     `json:"static_file_version"`
  staticDir          string  `json:"static_dir"`
  TemplatesDir       string  `json:"templates_dir"`
  SerTcpSocketHost   string  `json:"serTcpSocketHost"`
  SerTcpSocketPort   string  `json:"serTcpSocketPort"`
  Fruits           []string  `json:"fruits"`
}

type Other struct{
  SerTcpSocketHost   string  `json:"serTcpSocketHost"`
  SerTcpSocketPort   string  `json:"serTcpSocketPort"`
  Fruits           []string  `json:"fruits"`
}

func main(){
  jsonStr := `{"host":"http://localhost:9090","port":9090,"analytics_file":"","static_file_version":1,
              "static_dir":"E:/home/xyq/go","templates_dir": "E:/Project/goTest/src/templates/","serTcpSocketHost": ":12340","serTcpSocketPort": 12340,"fruits":              ["apple", "peach"]}`
  //json str 转 map
  var dat map[string]interface{}
  if err:=json.Unmarshal([]byte(jsonStr),&dat);err==nil{
    fmt.Println("====================json str 转 map======================")
    fmt.Println(dat)
    fmt.Println(dat["host"]) 
  }

  //json str 转 struct
  var config ConfigStruct
  if err:=json.Unmarshal([]byte(jsonStr),&config);err==nil{
    fmt.Println("====================json str 转 struct======================")
    fmt.Println(config)
    fmt.Println(config.Host)
  }

 //struct 转  json str
  if b,err:=json.Marshal(config);err==nil{
       fmt.Println("====================struct 转 json str======================")
       fmt.Println(string(b))
  }
 
 //map 转 json str
    fmt.Println("====================map 转 json str======================")
    enc:=json.NewEncoder(os.Stdout)
    enc.Encode(dat)
 //array 转 json str
   arr:=[]string{"c","c++","java","golang","python","javascript","shell"}
   lang,err:=json.Marshal(arr)
   if err==nil{
     fmt.Println("================array 转 json str===================")
     fmt.Println(string(lang))
   }
  var wo []string
  if err:=json.Unmarshal(lang,&wo);err==nil{
    fmt.Println("===============json 转 []string================")
    fmt.Println(wo)
  }
}
