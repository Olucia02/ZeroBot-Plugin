// Package draw 画图测试
package draw

import (
	"encoding/json"
	"strconv"
	// "os"
	"github.com/Coloured-glaze/gg"
	// "github.com/FloatTech/floatbox/file"
	"github.com/FloatTech/floatbox/img/writer"
	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

type li struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Top []struct {
			Title string `json:"title"`
			URL   string `json:"url"`
			Icon  string `json:"icon"`
		} `json:"top"`
		Hot []struct {
			Title string `json:"title"`
			Hot   string `json:"hot"`
			URL   string `json:"url"`
			Icon  string `json:"icon"`
		} `json:"hot"`
	} `json:"data"`
}

const (
	url2 = "http://ovooa.com/API/xz/api.php?msg=%v"
	url  = "http://api.a20safe.com/api.php?api=18&key=7d06a110e9e20a684e02934549db1d3d"
)

func init() {
	en := control.Register("draw", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "画图练习",
		Help:             "-/画图测试",
		//	PublicDataFolder: "draw",
	})
	// cachePath := en.DataFolder() + "cache/" // 缓存文件夹
	//	_ = os.RemoveAll(cachePath)             // 创建
	//	_ = os.MkdirAll(cachePath, 0755)
	en.OnFullMatch("微博热搜").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		data, err := web.GetData(url)
		//	drawedFile := cachePath + "draw.png"
		if err != nil {
			ctx.SendChain(message.Text("获取失败惹", err))
			return
		}
		var jiexi li                       // li结构
		err = json.Unmarshal(data, &jiexi) // 填api返回结果，struct地址
		if err != nil {
			ctx.SendChain(message.Text("出现错误捏：", err))
			return
		}
		top := (jiexi.Data[0].Top[0].Title)
		ctx.SendChain(message.Text(top))
		dc := gg.NewContext(1500, 1000) // 画布大小
		dc.SetRGB(1, 1, 1)
		dc.Clear()         // 白背景
		dc.SetRGB(0, 0, 0) // 换黑色
		FontFile := "data/Font/regular-bold.ttf"
		if err := dc.LoadFontFace(FontFile, 50); err != nil {
			panic(err)
		}
		dc.DrawString("微博今日热搜", 0, 50)
		dc.DrawString("No1:"+top, 10, 150)
		var str string
		ctx.SendChain(message.Text(jiexi.Data[0].Hot[0].Title)) // jiexi.Data[0].Hot[i].Title != ""
		for i := 0; i < 15; i++ {
			str = "No" + strconv.Itoa(i+2) + ":" + jiexi.Data[0].Hot[i].Title
			dc.DrawString(str, 10, float64(250+i*100))
		}
		ff, cl := writer.ToBytes(dc.Image())
		ctx.SendChain(message.ImageBytes(ff))
		cl()
		// ctx.SendChain(message.Text(top)) // 输出提取后的结果
	})
	/*en.OnSuffix("运势").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		xing := ctx.State["args"].(string)
	data, err := web.GetData(fmt.Sprintf(url,xing))
	//	drawedFile := cachePath + "draw.png"
	if err != nil {
		ctx.SendChain(message.Text("获取失败惹", err))
		return
	}
	var jiexi li                       // li结构
	err = json.Unmarshal(data, &jiexi) // 填api返回结果，struct地址
	if err != nil {
		ctx.SendChain(message.Text("出现错误捏：", err))
		return
	}
	top := (jiexi.Data[0].Top[0].Title)
	ctx.SendChain(message.Text(top))
	dc := gg.NewContext(1500, 1000) // 画布大小
	dc.SetRGB(1, 1, 1)
	dc.Clear()         // 白背景
	dc.SetRGB(0, 0, 0) // 换黑色
	FontFile := "data/Font/regular-bold.ttf"
	if err := dc.LoadFontFace(FontFile, 50); err != nil {
		panic(err)
	}
	dc.DrawString("微博今日热搜", 0, 50)
	dc.DrawString("No1:"+top, 10, 150)
	var str string
	ctx.SendChain(message.Text(jiexi.Data[0].Hot[0].Title))
	for i := 0; jiexi.Data[0].Hot[i].Title != ""; i++ {
		str = "No" + strconv.Itoa(i+2) + ":" + jiexi.Data[0].Hot[i].Title
		dc.DrawString(str, 10, float64(250+i*100))
	}
	ff, cl := writer.ToBytes(dc.Image())
	ctx.SendChain(message.ImageBytes(ff))
	cl()
	//ctx.SendChain(message.Text(top)) // 输出提取后的结果*/
}
