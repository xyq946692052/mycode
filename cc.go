package main

import (
	"flag"
	"github.com/sidbusy/weixinmp"
	"log"
	"net/http"
	"os"
	"fmt"
)

var (
	listenAddr = flag.String("http", ":10000", "http listen address")
	url        = flag.String("url", "/bppc", "url for http server")

	token = flag.String("token", "bppc_com", "token for weixinmp")

	appid  = flag.String("appid", "wxaa950eb1dad7f423", "appid for weixinmp")
	secret = flag.String("secret", "d6db09db8e8132b58176bf5e7d866fd9", "secret for weixinmp")
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stderr, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
}

func main() {
	flag.Parse()

	http.HandleFunc(*url, wxRequestAndEventHandler)
	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}

func wxRequestAndEventHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
	mp := weixinmp.New(*token, *appid, *secret)
	fmt.Println(mp.Request)
	// check request
	if !mp.Request.IsValid(w, r) {
		return
	}

	// handle request message
	switch mp.Request.MsgType {
	case weixinmp.MsgTypeText:
	case weixinmp.MsgTypeImage:
	case weixinmp.MsgTypeVoice:
	case weixinmp.MsgTypeVideo:
	case weixinmp.MsgTypeLocation:
	case weixinmp.MsgTypeLink:
	case weixinmp.MsgTypeEvent:
		switch mp.Request.Event {
		case weixinmp.EventSubscribe:
		case weixinmp.EventUnsubscribe:
		case weixinmp.EventScan:
		case weixinmp.EventLocation:
		case weixinmp.EventClick:
			switch mp.Request.EventKey {
			case "ABOUT_BPPC":
			case "XXXX":
			default:
			}
		}
	default:
	}

}

