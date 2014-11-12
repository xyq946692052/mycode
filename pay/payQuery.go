package main
/*
 *被扫订单查询
 */
import(
  "fmt"
  "net/http"
  "io/ioutil"
  "log"
  "bytes"
  "encoding/json"
)

type DataPost struct{
  Appid        string   //公众账号ID
  Mch_id       string   //商户号
  Sub_mch_id   string   //子商户号
  Transaction_id  string   //微信订单号
  Nonce_str    string   //随机字符串
  Sign         string   //签名
  Out_trade_no string   //商品订单号
 }

func main(){
  var data DataPost
  data.Appid=""
  data.Mch_id=""
  data.Sub_mch_id=""
  data.Transaction_id=""
  data.Nonce_str=""
  data.Sign=""
  data.Out_trade_no=""
                                                                
  b,err:=json.Marshal(data)
  if(err!=nil){
      fmt.Println("json err:",err)
   }

  body:=bytes.NewBuffer([]byte(b))
  res,err:=http.Post("https://api.mch.weixin.qq.com/pay/orderquery","application/json;charset=utf-8",body)
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

