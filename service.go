ckage main

import (
	"flag"
	"fmt"
	"github.com/sidbusy/weixinmp"
	"log"
	"net/http"
	"oilo2o.com/bp/models"
	"oilo2o.com/wxapp/station"
	"strconv"
	//"oilo2o.com/wxapp/tips"
	"oilo2o.com/wxapp/conf"
	"oilo2o.com/wxapp/promotion"
	"oilo2o.com/wxapp/userlocation"
	"oilo2o.com/wxapp/worldcup"
	"os"
)

var (
	listenAddr = flag.String("http", ":10003", "http listen address")
	url        = flag.String("url", "/receiver", "url for http server")

	token = flag.String("token", "oilo2o_com", "token for weixinmp")

	appid  = flag.String("appid", conf.APPID, "appid for weixinmp")
	secret = flag.String("secret", conf.SECRET, "secret for weixinmp")
)

var logger *log.Logger
var (
	storePromo = &[]weixinmp.Article{
		weixinmp.Article{
			"昆仑山矿泉水减赠大优惠！",
			"",
			"http://203.195.192.100/store/promotion_kunlunshan_big.jpg",
			"http://203.195.192.100/store/promotion_kunlunshan.htm",
		},
		weixinmp.Article{
			"泰国湄南河香米30元抵扣券",
			"",
			"http://203.195.192.100/store/promotion_meinanhe_small.png",
			"http://203.195.192.100/store/promotion_meinanhe.htm",
		},
		weixinmp.Article{
			"优途全校三合一汽油发动机清洗剂买四赠一",
			"",
			"http://203.195.192.100/store/promotion_youtu_small.png",
			"http://203.195.192.100/store/promotion_youtu.htm",
		},
	}

	welcomeBPPC = &[]weixinmp.Article{
		weixinmp.Article{
			"中油BP",
			"谢谢您选择了中油BP！我们提供各种实时信息，以及其他便捷服务。\n\n请点击以下获得优惠信息详情",
			"http://203.195.192.100/demo/welcome.jpg",
			"http://203.195.192.100/demo/promotion_1.htm",
		},
	}

	MEMBER_REGISTER_URL = "http://203.195.192.100:8080/bp/member/new?openId=%s"
)

func init() {
	logger = log.New(os.Stderr, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	flag.Parse()
}

func main() {
	http.HandleFunc(*url, receiver)
	http.ListenAndServe(*listenAddr, nil)
}

func receiver(w http.ResponseWriter, r *http.Request) {
	mp := weixinmp.New(*token, *appid, *secret)
	if !mp.Request.IsValid(w, r) {
		return
	}
	logger.Println(mp.Request)
	openId := mp.Request.FromUserName
	picurl := mp.Request.PicUrl
	voiceId := mp.Request.MediaId
	locationx := mp.Request.LocationX
	locationy := mp.Request.LocationY
	// ticket :=mp.Request.Ticket
	key := mp.Request.EventKey
	switch mp.Request.MsgType {
	case weixinmp.MsgTypeText:
		logger.Println("Text")
		reqText := mp.Request.Content
		if len(reqText) == 11 {
			err := models.UpdateCellphone(openId, reqText)
			logger.Println("cellphone upload success", err)
		}
		switch reqText {

		case "1":
			mp.ReplyTextMsg(w, worldcup.WorldCupRule)
		case "2":
			mp.ReplyTextMsg(w, worldcup.Teams)
		case "3":
			msg := "点击查看你组建的冠军队\n"
			coach := openId
			teams, err := models.GetCoachVotedTeams(coach)
			if err != nil {
				logger.Println(err)
			}
			for i := 0; i < len(teams); i++ {
				name, err := models.GetTeamNameWithCode(teams[i].Voted)
				if err != nil {
					logger.Println(err)
					return
				}
				link := fmt.Sprintf("http://203.195.192.100:8080/bp/worldcup/teamview?coach=%s&code=%s", coach, teams[i].Voted)
				msg += teams[i].Voted + "<a href=\"" + link + "\">" + name + "</a>\n"
			}
			logger.Println(msg)
			mp.ReplyTextMsg(w, msg)
		case "4":
			msg := "你参与的足球队：\n"
			player := openId
			joinTeams, err := models.GetJoinTeams(player)
			if err != nil {
				logger.Println(err)
				return
			}

			for i := 0; i < len(joinTeams); i++ {
				logger.Println(joinTeams[i].Coach)
				userInfo, err := mp.GetUserInfo(joinTeams[i].Coach)
				if err != nil {
					logger.Println(err)
					return
				}
				logger.Printf("%v\n", userInfo)
				coachNickname := userInfo.Nickname

				teamName, err := models.GetTeamNameWithCode(joinTeams[i].Voted)
				if err != nil {
					logger.Println(err)
					return
				}

				msg += fmt.Sprintf("%d) ", i+1) + "球队：" + joinTeams[i].Voted + teamName + "\n教练： " + coachNickname + "\n加入时间：" + fmt.Sprintf("%v", joinTeams[i].Joined) + "\n"

			}
			logger.Println(msg)
			mp.ReplyTextMsg(w, msg)
		default:
			if len(reqText) == 2 {
				code := reqText
				name, err := models.GetTeamNameWithCode(code)
				if err != nil {
					logger.Println(err)
					return
				}
				logger.Println(name)
				coach := openId
				voted := code
				err = models.CoachFirstVotedTeam(coach, voted)
				if err != nil {
					logger.Println(err)
					return
				}

				invitationLink := fmt.Sprintf("http://203.195.192.100:8080/bp/worldcup/invite?coach=%s&code=%s", coach, code)
				logger.Println(invitationLink)
				inviteToBPPCWorldCup := &[]weixinmp.Article{
					weixinmp.Article{
						"你已经组建朋友圈⾜球队,点击查看详情",
						"",
						"http://203.195.192.100/worldcup/worldcup_bppc.jpg",
						invitationLink,
					},
				}
				mp.ReplyNewsMsg(w, inviteToBPPCWorldCup)
			}

		}
	case weixinmp.MsgTypeImage:
		logger.Println("Image")
		logger.Println("++++openid:", openId)
		logger.Println("++++picurl:", picurl)

		err := models.UpdateImg(openId, picurl)
		logger.Println("++++", err)

		msg := "图片上传成功"
		mp.ReplyTextMsg(w, msg)
	case weixinmp.MsgTypeLocation:

		logger.Println("Location")
		logger.Println("++++openid:", openId)
		logger.Println("++++X and Y:", locationx, locationy)

		err := models.UpdateLocation(openId, locationx, locationy)
		logger.Println("++++", err)

		msg := "位置信息上传成功"
		mp.ReplyTextMsg(w, msg)

	case weixinmp.MsgTypeVoice:
		logger.Println("Voice")
		logger.Println("++++openid:", openId)
		logger.Println("++++voiceId:", voiceId)

		err := models.UpdateVoice(openId, voiceId)
		logger.Println("++++", err)

		msg := "语音上传成功"
		mp.ReplyTextMsg(w, msg)

	case weixinmp.MsgTypeVideo:
		logger.Println("Viedo")
	case weixinmp.MsgTypeLink:
		logger.Println("Link")
	case weixinmp.MsgTypeEvent:
		switch mp.Request.Event {
		case weixinmp.EventUnsubscribe:
			logger.Println("EventUnsubscribe")
			models.SetUnSubscribed(openId)
			logger.Println("SetUnSubscribed", openId)
		case weixinmp.EventSubscribe:
			logger.Println("EventSubscribe")
			models.SetSubscribed(openId)
			logger.Println("SetSubscribed", openId)
			if len(mp.Request.EventKey) > 0 {
				// scan to subscribe

				logger.Println("scan to subscribe")
				// id := mp.Request.EventKey[8:] // qrscene_XXXXXXX
				card, err := models.BecomeMember(openId)
				if err != nil {
					logger.Println(err)
					return
				}
				welcomeBPPC = &[]weixinmp.Article{
					weixinmp.Article{
						"中油BP",
						fmt.Sprintf("感谢您选择了中油BP！中油BP油品优，商品更优。\n中油BP世界杯加油有惊喜！非凡夏日与中油BP一起为世界杯加油吧！如想了解更多门店及优惠信息，请在菜单栏点击<附近油站>和<促销信息>。\n你的卡号为：%d,请点击以下获得详细信息", card),
						"http://203.195.192.100/demo/welcome.jpg",
						fmt.Sprintf("http://203.195.192.100:8080/bp/member/cardview?card=%d", card),
					},
				}
				mp.ReplyNewsMsg(w, welcomeBPPC)
			} else {
				// normal subscribe
				logger.Println("normal subscribe")

				// generate card
				card, err := models.BecomeMember(openId)
				if err != nil {
					logger.Println(err)
					return
				}

				welcomeBPPC = &[]weixinmp.Article{
					weixinmp.Article{
						"中油BP",
						fmt.Sprintf("感谢您选择了中油BP！中油BP油品优，商品更优。\n中油BP世界杯加油有惊喜！非凡夏日与中油BP一起为世界杯加油吧！如想了解更多门店及优惠信息，请在菜单栏点击<附近油站>和<促销信息>。\n你的卡号为：%d,请点击以下获得详细信息", card),
						"http://203.195.192.100/demo/welcome.jpg",
						fmt.Sprintf("http://203.195.192.100:8080/bp/member/cardview?card=%d", card),
					},
				}

				mp.ReplyNewsMsg(w, welcomeBPPC)
			}
		case weixinmp.EventScan:
			logger.Println("EventScan")
			numid, _ := strconv.Atoi(key)
                        err1:=models.GetCustomer(openId,numid)
                        var msg string    
                        game, _ := models.GetGameWithNumid(numid)
                        if err1 !=nil{
                             logger.Println("***************22")
                          msg = "员工编号："+key+",所属油站："+game.Station+"\n投票成功，非常感谢您的支持！"
                          models.Addvote(openId,numid)
                          mp.ReplyTextMsg(w, msg)
                        } else{
                         logger.Println("***************11")
                          msg="尊敬的客户，您已经给编号为"+key+"的员工投过票，不可重复投票"
                          mp.ReplyTextMsg(w, msg)
                        }
                        
                    	logger.Println("========", game.Name, game.Numid, game.Station, game.Hits, game.Score)
		
		case weixinmp.EventLocation:
			userlocation.UpdateUserLocation(openId, mp, w, logger)
		case weixinmp.EventClick:
			switch mp.Request.EventKey {
			case "STATION_NEARBY":
				station.SearchNearby(openId, mp, w, logger)
			case "STATION_NEARBY_GD":
				station.SearchNearbyGD(openId, mp, w, logger)
			case "GO_WORLDCUP":
				logger.Println("GO_WORLDUP")
				mp.ReplyNewsMsg(w, worldcup.WelcomeWorldCup)

			case "MEMBERCARD":
				logger.Println("MEMBERCARD")
				member, err := models.GetMember(openId)

				if err == nil {
					MEMBERCARD_VIEW := "http://203.195.192.100:28080/bp/member/cardview?card=%d"
					link_v := fmt.Sprintf(MEMBERCARD_VIEW, member.Card)
					MEMBERCARD_EDIT := "http://203.195.192.100:28080/bp/member/cardedit?card=%d"
					link_e := fmt.Sprintf(MEMBERCARD_EDIT, member.Card)

					msg := &[]weixinmp.Article{

						weixinmp.Article{
							"",
							"",
							"",
							"",
						},
						weixinmp.Article{
							"查看会员卡信息",
							"",
							"",
							link_v,
						},
						weixinmp.Article{
							"编辑会员卡信息",
							"",
							"",
							link_e,
						},
					}

					mp.ReplyNewsMsg(w, msg)
				} else {
					_, err := models.BecomeMember(openId)
					if err != nil {
						logger.Println("---", err)
						return
					}
					member, err := models.GetMember(openId)
					if err == nil {
						MEMBERCARD_VIEW := "http://203.195.192.100:28080/bp/member/cardview?card=%d"

						link_v := fmt.Sprintf(MEMBERCARD_VIEW, member.Card)

						MEMBERCARD_EDIT := "http://203.195.192.100:8080/bp/member/cardedit?card=%d"
						link_e := fmt.Sprintf(MEMBERCARD_EDIT, member.Card)

						msg := "<a href=\"" + link_v + "\">查看会员信息</a>    " + "<a href=\"" + link_e + "\">编辑会员信息</a>    "
						mp.ReplyTextMsg(w, msg)
					}

				}

			case "ABOUT_BPPC":
				logger.Println("APPLY")
				_, err := models.GetApply(openId)
				if err != nil {
					logger.Println(err)
					return
				}
				if err == nil {
					APPLY_VIEW := "http://203.195.192.100:28080/bp/applystay/viewImg?openId=%s "
					APPLY_INCELL := "http://203.195.192.100:28080/bp/apply/applyarray?openid=%s"
					APPLY_GETHOTEL := "http://203.195.192.100:28080/bp/apply/getHotelinfo?openid=%s"
					APPLY_GETCAR := "http://203.195.192.100:28080/bp/apply/getCarinfo?openid=%s"

					STATION_SEARCH := "http://203.195.192.100:28080/bp/apply/searchstation?openid=%s"
					link_s := fmt.Sprintf(STATION_SEARCH, openId)

					link_i := fmt.Sprintf(APPLY_INCELL, openId)
					link_vg := fmt.Sprintf(APPLY_GETCAR, openId)
					link_v := fmt.Sprintf(APPLY_VIEW, openId)
					link_va := fmt.Sprintf(APPLY_GETHOTEL, openId)
					msg := &[]weixinmp.Article{
						weixinmp.Article{
							"",
							"",
							"",
							"",
						},
						weixinmp.Article{
							"中油碧辟",
							"",
							"http://203.195.192.100/worldcup/worldcup_bppc.jpg",
							link_v,
						},

						weixinmp.Article{
							"请发送您的手机号码和身份证图片申请住宿",
							"",
							"http://203.195.192.100/worldcup/worldcup_bppc.jpg",
							"",
						},
						weixinmp.Article{
							"请发送您的手机号码和语音或位置为您安排接送",
							"",
							"http://203.195.192.100/demo/welcome.jpg",
							"",
						},
						weixinmp.Article{
							"客服安排",
							"",
							"http://203.195.192.100/demo/welcome.jpg",
							link_i,
						},

						weixinmp.Article{
							"查看住宿安排",
							"",
							"http://203.195.192.100/demo/welcome.jpg",
							link_va,
						},
						weixinmp.Article{
							"查看接送安排",
							"",
							"http://203.195.192.100/demo/welcome.jpg",
							link_vg,
						},
						weixinmp.Article{
							"查找油站",
							"",
							"http://203.195.192.100/demo/welcome.jpg",
							link_s,
						},
					}
					mp.ReplyNewsMsg(w, msg)
				}
			case "FUEL_PROMOTION":
				promotion.ShowPromotionNearby(openId, mp, w, logger)

			case "STORE_PROMOTION":
				mp.ReplyNewsMsg(w, storePromo)

			case "SKILL":
				msg := &[]weixinmp.Article{
					weixinmp.Article{
						"2014油站员工技能大赛",
						"",
						"http://203.195.159.145/skill/skill.png",
						"http://mp.weixin.qq.com/s?__biz=MzA3MjY0NTEyNA==&mid=200664737&idx=1&sn=598bc63dce40d1cedde388792ce39559#rd",
					},
					weixinmp.Article{
						"我的大赛",
						"",
						"http://203.195.159.145/skill/my.png",
						"http://mp.weixin.qq.com/s?__biz=MzA3MjY0NTEyNA==&mid=200673878&idx=1&sn=ea406597178a8403c346e27982b79482#rd",
					},
					weixinmp.Article{
						"实时播报",
						"",
						"http://203.195.159.145/skill/realtime.png",
						"http://mp.weixin.qq.com/s?__biz=MzA3MjY0NTEyNA==&mid=200673890&idx=1&sn=f9c23ed21b9cd3c5a687bd57c5bd7667#rd",
					},
					/*
					weixinmp.Article{
						"决赛直击",
						"",
						"http://203.195.192.100/demo/welcome.jpg",
						"",
					},
					*/
				}
				mp.ReplyNewsMsg(w, msg)

			case "CONFERENCE":
				_, err := models.GetGame(openId)
				//   GAME_ABOUT:="http://203.195.192.100:28080/bp/game/about?openid=%s"
				// link_g:=fmt.Sprintf(GAME_ABOUT,openId)

				if err != nil {
					logger.Println(err)
					return
				}
				if err == nil {
					msg := &[]weixinmp.Article{
						weixinmp.Article{
							"2014年员工技能大赛",
							"",
							"http://203.195.159.145/skill/skills.png",
							"http://mp.weixin.qq.com/s?__biz=MzA3MjY0NTEyNA==&mid=200664737&idx=1&sn=598bc63dce40d1cedde388792ce39559#rd",
						},
						weixinmp.Article{
							"实时播报",
							"",
							"http://203.195.192.100/demo/welcome.jpg",
							"",
						},
						weixinmp.Article{
							"我的大赛",
							"",
							"http://203.195.192.100/demo/welcome.jpg",
							"",
						},
						weixinmp.Article{
							"决赛直击",
							"",
							"http://203.195.192.100/demo/welcome.jpg",
							"",
						},
					}
					mp.ReplyNewsMsg(w, msg)

				}
			default:
				logger.Println("Unknown EventClick")
			}
		case weixinmp.EventView:
			logger.Println("EventView")

		default:
			logger.Println("Unknown Event")
		}
	default:
		logger.Println("Unknown Msg")
	}
}

