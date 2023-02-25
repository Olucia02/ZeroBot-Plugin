// Package zaobao 早报
package zaobao

import (
	"fmt"
	"strings"
	"time"

	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"github.com/wdvxdr1123/ZeroBot/utils/helper"
)

const new60s = "https://api.emoao.com/api/60s"

var (
	yan       = "https://v1.hitokoto.cn/?c=k&encode=text"
	winter    = []string{"20230109", "20230220"}
	summer    = []string{"20230715", "20230907"}
	sem2023Up = []string{"20230226", "20230715"}
)

func init() { // 主函数
	en := control.Register("zaobao", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "早报",
		Help: "- 指令列表\n" +
			"- 西邮早安！" +
			"- 早报",
	})
	en.OnFullMatch("西邮早安！").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		var (
			msg strings.Builder
		)
		t := time.Now()
		msg.WriteString(fmt.Sprintf("%4d年%02d月%02d日，星期%s\n",
			t.Year(), t.Month(), t.Day(), func() string {
				switch t.Weekday() {
				case 1:
					return "一"
				case 2:
					return "二"
				case 3:
					return "三"
				case 4:
					return "四"
				case 5:
					return "五"
				case 6:
					return "六"
				default:
					return "日"
				}
			}()))
		msg.WriteString(fmt.Sprintf("现在是%02d:%02d:%02d，%s！",
			t.Hour(), t.Minute(), t.Second(), func() string {
				switch {
				case t.Hour() < 10:
					return "早上好"
				case t.Hour() < 15:
					return "中午好"
				case t.Hour() < 18:
					return "下午好"
				default:
					return "晚上好"
				}
			}()))
		msg.WriteString("\n新的一天元气满满,学习加油！")
		msg.WriteString(fmt.Sprintf("\n今天剩余时光：%.2f%%！",
			float64((23-t.Hour())*60+(60-t.Minute()))/14.4))
		msg.WriteString(
			func() string {
				if t.Weekday() < 5 && t.Weekday() > 0 {
					return fmt.Sprintf("\n距离周末还有：%.2f天！",
						4-float64(t.Weekday())+float64(24-t.Hour())/24)
				}
				return "\n现在是周末时光！"
			}())
		msg.WriteString(
			func() string {
				if stringtotime(winter[1]).After(t) {
					return fmt.Sprintf("\n距离寒假还有：%.2f天！",
						float64(subDays(stringtotime(winter[0]), t))+float64(24-t.Hour())/24-1)
				} else if stringtotime(winter[1]).After(t) {
					return fmt.Sprintf("\n寒假时光剩余：%.2f天！",
						float64(subDays(stringtotime(winter[1]), t))+float64(24-t.Hour())/24-1)
				} else {
					return "\n寒假时光已结束！"
				}
			}())
		msg.WriteString(
			func() string {
				if stringtotime(summer[0]).After(t) {
					return fmt.Sprintf("\n距离暑假还有：%.2f天！",
						float64(subDays(stringtotime(summer[0]), t))+float64(24-t.Hour())/24-1)
				} else if stringtotime(summer[1]).After(t) {
					return fmt.Sprintf("\n暑假时光剩余：%.2f天！",
						float64(subDays(stringtotime(summer[1]), t))+float64(24-t.Hour())/24-1)
				} else {
					return "\n暑假时光已结束！"
				}
			}())
		msg.WriteString(
			func() string {
				if stringtotime(sem2023Up[1]).After(t) {
					return fmt.Sprintf("\n本学期剩余时光：%.2f%%！",
						(float64(subDays(stringtotime(sem2023Up[1]), t))+float64(24-t.Hour())/24-1)/
							float64(subDays(stringtotime(sem2023Up[1]), stringtotime(sem2023Up[0])))*100)
				}
				return "\n下半学期已经结束啦！"
			}())
		yi, err := web.GetData(yan)
		if err == nil {
			msg.WriteString("\n今日一言：")
			msg.WriteString(helper.BytesToString(yi))
		}
		ctx.SendChain(message.Text(msg.String()))
	})
	en.OnFullMatch("早报").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		data, err := web.GetData(new60s)
		if err != nil {
			ctx.SendChain(message.Text("获取图片失败惹"))
			return
		}
		ctx.SendChain(message.ImageBytes(data))
	})
}

func stringtotime(str string) time.Time {
	duetimecst, _ := time.ParseInLocation("20060102", str, time.Local)
	return duetimecst
}

// 计算日期相差多少天
// 返回值day>0, t1晚于t2; day<0, t1早于t2
func subDays(t1, t2 time.Time) (day int) {
	swap := false
	if t1.Unix() < t2.Unix() {
		t2, t1 = t1, t2
		swap = true
	}

	day = int(t1.Sub(t2).Hours() / 24)

	// 计算在被24整除外的时间是否存在跨自然日
	if int(t1.Sub(t2).Milliseconds())%86400000 > int(86400000-t2.Unix()%86400000) {
		day++
	}

	if swap {
		day = -day
	}

	return
}
