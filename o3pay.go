sckage main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "log"
  "bytes"
  "encoding/json"
)
   
type PubData struct{
   Appid        string  //公众账号ID
   Mch_id       string  //商户号
   Sub_mch_id   string  //子商户号
   Nonce_str    string  //随机字符串
   Sign         string  //签名
}
   
/*
 *提交被扫支付
 */
type SweptPostData struct{  
   PubData                  //公共数据
   Device_info      string  //设备号
   Body             string  //商品描述  
   Attach           string  //附加数据
   Out_trade_no     string  //商户订单号
   Total_fee        int     //总金额
   Spbill_create_ip string  //终端IP
   Time_start       string  //交易起始时间
   Time_expire      string  //交易结束时间
   Goods_tag        string  //商品标记
   Auth_code        string  //授权码
}

/*
 *被扫订单查询
 */

type OrderQueryData struct{
  PubData                    //公共数据
  Transaction_id     string  //微信订单号
  Out_trade_no       string  //商户订单号
}

/*
 *关闭订单
 */

type CloseOrderData struct{
  PubData                    //公共数据
  Out_trade_no       string  //商户订单号
}

/*
 *退款申请
 */

type RefundApplyData struct{
  PubData                    //公共数据
  Device_info        string  //设备号
  Transaction_id     string  //微信订单号
  Out_trade_no       string  //商户订单号
  Out_refund_no      string  //商户退款单号
  Total_fee          int     //总金额
  Refund_fee         int     //退款金额
  Op_user_id         string  //操作员
  Refund_channel     string  //退款渠道
}

/*
 *退款查询
 */

type RefundQueryData struct{
  PubData                    //公共数据
  Device_info        string  //设备号
  Transaction_id     string  //微信订单号
  Out_trade_no       string  //商户订单号
  Out_refund_no      string  //商户退款单号
  Total_fee          int     //总金额
  Refund_fee         int     //退款金额
}

/*
 *冲正
 */
   type ReverseDepositData struct{
  PubData                    //公共数据
  Out_trade_no       string  //商户订单号
}

/*
 *对账单
 */

type BillData struct{
  PubData                    //公共数据
  Device_info        string  //设备号
  Bill_date          string  //对账单日期
  Bill_type          string  //账单类型
}
  
/*
 *被扫支付
 */
func SweptPost(device_info,body,attach,out_trade_no,spbill_create_ip,time_start,time_expire,goods_tag,auth_code string,total_fee int,pubData PubData) (err error){
    var swp SweptPostData
    swp.Device_info=device_info
    swp.Body=body
    swp.Attach=attach
    swp.Out_trade_no=out_trade_no
    swp.Spbill_create_ip=spbill_create_ip
    swp.Time_start=time_start
    swp.Time_expire=time_expire
    swp.Goods_tag=goods_tag
    swp.Auth_code=auth_code
    swp.Total_fee=total_fee
    swp.PubData=pubData

   err= RequestData(swp,"https://api.mch.weixin.qq.com/pay/micropay")
   if err!=nil{
      fmt.Println("SweptPost error")
   }
    return
}
/* 
 *被扫订单查询
 */

func OrderQuery(transaction_id,out_trade_no string,pubData PubData)(err error){
      var order OrderQueryData
      order.Transaction_id=transaction_id
      order.Out_trade_no=out_trade_no
      order.PubData=pubData
      
      err= RequestData(order,"https://api.mch.weixin.qq.com/pay/orderquery")
      if err!=nil{
          fmt.Println("orderQuery error")
       }
      return
}
          
/*
 *关闭订单
 */

func CloseOrder(out_trade_no string,pubData PubData)(err error){
      var clsOrder CloseOrderData
      clsOrder.Out_trade_no=out_trade_no
      clsOrder.PubData=pubData
     
      err= RequestData(clsOrder,"https://api.mch.weixin.qq.com/pay/closeorder")
      if err!=nil{
          fmt.Println("closeOrder error")
       }
      return
}

/*
 *退款申请
 */
func RefundApply(device_info,transaction_id,out_trade_no,out_refund_no,op_user_id,refund_channel string,total_fee,refund_fee int,pubData PubData)(err error){
     var refApply RefundApplyData
     refApply.Device_info=device_info
     refApply.Transaction_id=transaction_id
     refApply.Out_trade_no=out_trade_no
     refApply.Out_refund_no=out_refund_no
     refApply.Op_user_id=op_user_id
     refApply.Refund_channel=refund_channel
     refApply.Total_fee=total_fee
     refApply.Refund_fee=refund_fee
     refApply.PubData=pubData

     err= RequestData(refApply,"https://api.mch.weixin.qq.com/secapi/pay/refund")
     if err!=nil{
         fmt.Println("refundApply error")
      }
     return
}

/*
 *退款查询
 */

func RefundQuery(device_info,transaction_id,out_trade_no,out_refund_no string,total_fee,refund_fee int,pubData PubData)(err error){
     var refQuery RefundQueryData
     refQuery.Device_info=device_info
     refQuery.Transaction_id=transaction_id
     refQuery.Out_trade_no=out_trade_no
     refQuery.Out_refund_no=out_refund_no
     refQuery.Total_fee=total_fee
     refQuery.Refund_fee=refund_fee
     refQuery.PubData=pubData

     err= RequestData(refQuery,"https://api.mch.weixin.qq.com/pay/refundquery")
     if err!=nil{
         fmt.Println("refundQuery error")
      }
     return
}
/*
 *冲正
 */ 

func ReverseDeposit(out_trade_no string,pubData PubData)(err error){
      var rd ReverseDepositData
      rd.Out_trade_no=out_trade_no
      rd.PubData=pubData

      err= RequestData(rd,"https://api.mch.weixin.qq.com/secapi/pay/reverse")
      if err!=nil{
         fmt.Println("reverseDeposit error")
      }
      return
}

/*
 *对账单
 */

func Bill(device_info,bill_date,bill_type string,pubData PubData)(err error){
      var bill BillData
      bill.Device_info=device_info
      bill.Bill_date=bill_date
      bill.Bill_type=bill_type
      bill.PubData=pubData

      err= RequestData(bill,"https://api.mch.weixin.qq.com/pay/downloadbill")
      if err!=nil{
         fmt.Println("bill error")
      }
      return
}
func RequestData(data interface{},url string)(err error){
     b,err:=json.Marshal(data)
  if(err!=nil){
      fmt.Println("json err:",err)
   } 
     
  
  body:=bytes.NewBuffer([]byte(b))
  res,err:=http.Post(url,"application/json;charset=utf-8",body)
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
  return
}     
  
func main(){
 pubData:=PubData{"wxd930ea5d5a258f4f","123456","123456789","4564dfsdf","FASDF423FFSDF"}
   SweptPost("","","","","","","","","",1,pubData)
// OrderQuery("","",pubData)
// CloseOrder("",pubData)
// RefundApply("","","","","","",1,1,pubdata)
// RefundQuery("","","","",1,1,pubData)
// ReverseDeposit("",pubData)
// Bill("",pubData)
}


                                                                                                                                   1,14         顶端

