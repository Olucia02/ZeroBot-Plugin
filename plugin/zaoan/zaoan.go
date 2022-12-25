// Package draw 画图测试
package draw

import (
	"bytes"
	"fmt"
	"image"
	"strconv"
	"strings"

	"github.com/Coloured-glaze/gg"
	"github.com/FloatTech/floatbox/img/writer"
	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"github.com/wdvxdr1123/ZeroBot/utils/helper"

	"time"
)

const (
	tu  = "http://api.iw233.cn/api.php?sort=pc"
	yan = "https://v1.hitokoto.cn/?c=k&encode=text"
)

func init() {
	en := control.Register("zaoan", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "早安/晚安图",
		Help: "- 记录在\"6 30 * * *\"触发的指令\n" +
			"   - /早安||晚安",
	})
	en.OnFullMatchGroup([]string{"/早安", "/晚安"}).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			now := time.Now()
			year := now.Year()        // 年
			month := int(now.Month()) // 月
			day := now.Day()          // 日
			hour := now.Hour()        // 小时
			wen := int(now.Weekday()) // 星期
			var wen1 string
			var yingwen string
			var img1 image.Image
			var err error
			switch wen {
			case 1:
				wen1 = "一"
			case 2:
				wen1 = "二"
			case 3:
				wen1 = "三"
			case 4:
				wen1 = "四"
			case 5:
				wen1 = "五"
			case 6:
				wen1 = "六"
			default:
				wen1 = "日"
			}
			if hour <= 11 { // 早安
				img1, err = gg.LoadImage("data/zaoan/zao.jpg")
				yingwen = "Ciallo～(∠ ·ω< )⌒★"
				if err != nil {
					fmt.Println(err)
					return
				}
			} else if hour <= 18 { // 午安
				img1, err = gg.LoadImage("data/zaoan/wu.jpg")
				yingwen = "Good afternoon"
				if err != nil {
					fmt.Println(err)
					return
				}
			} else { // 晚安
				img1, err = gg.LoadImage("data/zaoan/wan.jpg")
				yingwen = "Good night"
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			img2, err := gg.LoadImage("data/zaoan/an.jpg") // 安
			if err != nil {
				fmt.Println(err)
				return
			}
			pic, err := web.GetData(tu)
			if err != nil {
				ctx.SendChain(message.Text("错误：获取插图失败", err))
				return
			}
			dst, _, err := image.Decode(bytes.NewReader(pic))
			if err != nil {
				ctx.SendChain(message.Text("错误：获取插图失败", err))
				return
			}
			yi, err := web.GetData(yan)
			if err != nil {
				ctx.SendChain(message.Text("获取失败惹", err))
				return
			}
			yiyan := helper.BytesToString(yi)
			yiyanslice := strings.Split(yiyan, "，")
			result := ""
			for _, element := range yiyanslice {
				result += element + "， "
			}
			S := 1080
			sx := float64(S) / float64(dst.Bounds().Size().X) // 计算缩放倍率（宽）
			//	sy := float64(S) / float64(pic.Bounds().Size().Y) // 计算缩放倍率（长）
			// 设置背景
			Y2 := float64(dst.Bounds().Size().Y) * sx // 计算图片宽度
			Y3 := int(Y2)
			// 下方450像素，上方200
			// 开始画图
			dc := gg.NewContext(1080, Y3+600+200) // 画布大小//Y2+450+200
			dc.SetRGB(1, 1, 1)
			dc.Clear()                           // 白背景
			dc.SetRGB(0, 0, 0)                   // 换黑色
			FontFile := "data/zaoan/regular.ttf" // 日期字体
			if err := dc.LoadFontFace(FontFile, 50); err != nil {
				panic(err)
			}
			daily := strconv.Itoa(year) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(day)
			dc.DrawString(daily, 100, 100)        // 日期
			dc.DrawString("星期"+wen1, 450, 120)    // 星期
			dc.Scale(sx, sx)                      // 使画笔按倍率缩放
			dc.DrawImage(dst, 0, int(200*(1/sx))) // 贴图（会受上述缩放倍率影响）
			dc.Scale(1/sx, 1/sx)
			dc.DrawImage(img1, 10, Y3+200+200)      // 早
			dc.DrawImage(img2, 200, Y3+300+200)     // 安
			dc.DrawString(yingwen, 400, Y2+150+300) // 英文字符串
			FontFile = "data/zaoan/regular.ttf"     // 一言字体
			if err := dc.LoadFontFace(FontFile, 50); err != nil {
				panic(err)
			}
			//	iiii, err := text.RenderToBase64(yiyan, FontFile, 300, 10)
			//	if err != nil {
			//		ctx.SendChain(message.Text("ERROR: ", err))
			//		return
			//	}

			/*ii, _, err := image.Decode(bytes.NewBuffer(iii))
			if err != nil {
				ctx.SendChain(message.Text("错误：获取插图失败", err))
				return
			}*/
			//	dc.DrawImage(ii, 500, 200+Y3+340)
			dc.DrawStringWrapped(result, 400, 200+Y2+400, 0.5, 0.5, 100, 2, gg.AlignLeft)
			ff, cl := writer.ToBytes(dc.Image())
			ctx.SendChain(message.ImageBytes(ff))
			cl()
		})
}
