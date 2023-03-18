// Package baike 百度百科
package baike

import (
	"encoding/json"
	"fmt"

	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const (
	baidu = "https://api.a20safe.com/api.php?api=21&key=7d06a110e9e20a684e02934549db1d3d&text=%s" // api地址
	weiji = "https://zh.wikipedia.org/w/api.php?action=query&prop=extracts&titles=%s&format=json&exintro=1"
)

type ba struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Content string `json:"content"`
	} `json:"data"`
}
type we struct {
	Batchcomplete string `json:"batchcomplete"`
	Warnings      struct {
		Extracts struct {
			NAMING_FAILED string `json:"*"`
		} `json:"extracts"`
	} `json:"warnings"`
	Query struct {
		Pages map[string]struct {
			Pageid  int    `json:"pageid"`
			Ns      int    `json:"ns"`
			Title   string `json:"title"`
			Extract string `json:"extract"`
		} `json:"pages"`
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
		if r.Code == 0 {
			ctx.SendChain(message.Text(r.Data[0].Content)) // 输出提取后的结果
		} else {
			ctx.SendChain(message.Text("百度百科未找到信息"))
		}
	})
	en.OnRegex(`^维基\s*(.+)$`).SetBlock(true).Handle(func(ctx *zero.Ctx) {
		es, err := web.GetData(fmt.Sprintf(weiji, ctx.State["regex_matched"].([]string)[1])) // 将网站返回结果赋值
		if err != nil {
			ctx.SendChain(message.Text("出现错误捏：", err))
		}
		var r we                     // r数组
		err = json.Unmarshal(es, &r) // 填api返回结果，struct地址
		if err != nil {
			ctx.SendChain(message.Text("出现错误捏：", err))
		}
		for k, v := range r.Query.Pages {
			if k != "-1" {
				ctx.SendChain(message.Text(v.Extract)) // 输出提取后的结果
			} else {
				ctx.SendChain(message.Text("维基百科未找到信息"))
			}
		}
	})
}
