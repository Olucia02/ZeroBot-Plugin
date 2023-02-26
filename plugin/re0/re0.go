// Package re0 重启
package re0

import (
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
	"os"
	"os/exec"
)

func init() { // 主函数
	en := control.Register("re0", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "重启",
		Help: "重启命令大全\n" +
			"- 重启\n" +
			"- 洗白白\n" +
			"- 洗澡澡\n" +
			"- 洗脸脸",
	})
	en.OnFullMatchGroup([]string{"重启", "洗脸脸", "洗澡澡", "洗白白"}, zero.SuperUserPermission).SetBlock(true).
		Handle(func(ctx *zero.Ctx) {
			cmd := exec.Command("powershell.exe", "/c", "start", "run.bat")
			output, err := cmd.CombinedOutput()
			if err != nil {
				ctx.SendChain(message.Text(zero.BotConfig.NickName[0], "重启失败惹", string(output), err))
				return
			}
			ctx.SendChain(message.Text(zero.BotConfig.NickName[0], "去", ctx.State["matched"].(string), "啦~"))
			os.Exit(0)
		})
}
