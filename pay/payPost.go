package main

import(
  "fmt"
  "net/http"
  "io/ioutil"
  "log"
  "bytes"
  "encoding/json"
)

type Datapost struct{
  Appid        string   //公众账号ID
  Mch_id       string   //商户号
  Sub_mch_id   string   //子商户号
  Device_info  string   //设备号
  Nonce_str    string   //随机字符串
  Sign         string   //签名
  Body         string    //商品描述
  Attach       string   //附加数据
  Out_trade_no string   //商品订单号
  Total_fee    int      //总金额
  Spbill_create_ip   string   //终端IP
  Time_start   string   //交易起始时间
  Time_expire  string   //交易结束时间
  Goods_tag    string   //商品标记
  Auth_code    string   //授权码
}

func main(){
  var data Datapost
  data.Appid=""
  data.Mch_id=""
  data.Sub_mch_id=""
  data.Device_info=""
  data.Nonce_str=""
  data.Sign=""
  data.Body=""
  data.Attach=""
  data.Out_trade_no=""
  data.Total_fee=1233
  data.Spbill_create_ip=""
  data.Time_start=""
  data.Time_expire=""
  data.Goods_tag=""
  data.Auth_code=""
  
  b,err:=json.Marshal(data)
  if(err!=nil){
      fmt.Println("json err:",err)
   }
  
  body:=bytes.NewBuffer([]byte(b))
  res,err:=http.Post("https://api.mch.weixin.qq.com/pay/micropay","application/json;charset=utf-8",body)
  if(err!=nil){    
    log.Fatal(err)
    return
   }
  result,err:=ioutil.ReadAll(res.Body)
  res.Body.Close()
  if err!=nil{
      log.Fatal(err)
      return
  }
  fmt.Printf("ok----%s",result)
}

