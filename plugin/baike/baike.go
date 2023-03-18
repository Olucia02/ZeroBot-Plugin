// Package baike 百度百科
package baike

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const (
	baidu = "http://ovooa.caonm.net/API/bdbk/?Msg=%s" // api地址
	weiji = "http://asoiaf.huijiwiki.com/api.php?action=query&format=json&formatversion=2&list=search&srsearch=%s&srnamespace=0&srlimit=%s"
)

type ba struct {
	Code int    `json:"code"`
	Text string `json:"text"`
	Data struct {
		Msg   string `json:"Msg"`
		Info  string `json:"info"`
		Image string `json:"image"`
		URL   string `json:"url"`
	} `json:"data"`
}
type we struct {
	Batchcomplete bool `json:"batchcomplete"`
	Continue      struct {
		Sroffset int    `json:"sroffset"`
		Continue string `json:"continue"`
	} `json:"continue"`
	Query struct {
		Searchinfo struct {
			Totalhits int `json:"totalhits"`
		} `json:"searchinfo"`
		Search []struct {
			Ns        int    `json:"ns"`
			Title     string `json:"title"`
			Pageid    int    `json:"pageid"`
			Size      int    `json:"size"`
			Wordcount int    `json:"wordcount"`
			Snippet   string `json:"snippet"`
		} `json:"search"`
	} `json:"query"`
}

func init() { // 主函数
	en := control.Register("baike", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Help: "百科\n" +
			"- 百度+[关键字]" +
			"- 维基+[关键字]",
	})
	en.OnRegex(`^百度\s*(.+)$`).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		es, err := web.GetData(fmt.Sprintf(baidu, ctx.State["regex_matched"].([]string)[1])) // 将网站返回结果赋值
		if err != nil {
			ctx.SendChain(message.Text("出现错误捏：", err))
		}
		var r ba                     // r数组
		err = json.Unmarshal(es, &r) // 填api返回结果，struct地址
		if err != nil {
			ctx.SendChain(message.Text("出现错误捏：", err))
		}
		if r.Code == 1 {
			ctx.SendChain(message.Text(r.Data.Info+"\n详情查看:", r.Data.URL)) // 输出提取后的结果
		} else {
			ctx.SendChain(message.Text(r.Text))
		}
	})
	en.OnRegex(`^维基\s*(\S+)\s*(\d+)?$`).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		fig := ctx.State["regex_matched"].([]string)[2]
		if fig == "" {
			fig = "1"
		}
		es, err := web.GetData(fmt.Sprintf(weiji, ctx.State["regex_matched"].([]string)[1], fig)) // 将网站返回结果赋值
		if err != nil {
			ctx.SendChain(message.Text("出现错误捏：", err))
		}
		var r we                     // r数组
		err = json.Unmarshal(es, &r) // 填api返回结果，struct地址
		if err != nil {
			ctx.SendChain(message.Text("出现错误捏：", err))
		}
		if len(r.Query.Search) > 0 {
			ctx.SendChain(message.Text(r.Query.Search[func(fig string) int {
				f, _ := strconv.Atoi(fig)
				return f
			}(fig)].Title, "\n", r.Query.Search[func(fig string) int {
				f, _ := strconv.Atoi(fig)
				return f
			}(fig)].Snippet)) // 输出提取后的结果
		} else {
			ctx.SendChain(message.Text("维基百科未找到信息"))
		}
	})
}
