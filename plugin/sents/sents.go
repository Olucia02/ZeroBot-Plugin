// Package sents 群发消息
package sents

import (
	"time"

	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"

	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

func init() {
	en := control.Register("sents", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "群发消息，慎用",
		Help: "群发消息\n" +
			" - 群发消息xx",
	})
	en.OnCommand("群发消息", zero.SuperUserPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			next := zero.NewFutureEvent("message", 999, false, zero.OnlyGroup, ctx.CheckSession())
			recv, stop := next.Repeat()
			defer stop()
			ctx.SendChain(message.Text("请输入想要群发的内容:"))
			var mastergid = ctx.Event.GroupID
			var step int
			var origin string
			i := 0
			for {
				select {
				case <-time.After(time.Second * 60):
					ctx.SendChain(message.Text("时间太久啦！不发了喵！"))
					return
				case c := <-recv:
					switch step {
					case 0:
						origin = c.Event.RawMessage
						ctx.SendChain(message.Text("请输入\"确定\"或者\"取消\"来决定是否发送此消息喵~"))
						step++
					case 1:
						msg := c.Event.Message.ExtractPlainText()
						if msg != "确定" && msg != "取消" {
							ctx.SendChain(message.Text("请输入\"确定\"或者\"取消\"喵~"))
							continue
						}
						if msg == "确定" {
							ctx.SendChain(message.Text("正在发送..."))
							zero.RangeBot(func(id int64, ctx *zero.Ctx) bool {
								for _, g := range ctx.GetGroupList().Array() {
									gid := g.Get("group_id").Int()
									ctx.SendGroupMessage(gid, origin)
									i++
									time.Sleep(1000 * time.Microsecond) //1s
								}
								ctx.SendGroupMessage(mastergid, message.Text("共计发送", i, "条消息,\n", "任务完成了喵~"))
								return true
							})

							return
						}
						ctx.SendChain(message.Text("已经取消发送了喵~"))
						return
					}
				}
			}
		})
}
