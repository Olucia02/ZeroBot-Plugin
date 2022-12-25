// Package yuanshen 原神面板
package yuanshen

import (
	"encoding/json"
	"fmt"

	"github.com/Coloured-glaze/gg"
	"github.com/FloatTech/floatbox/img/writer"
	"github.com/FloatTech/floatbox/web"
	ctrl "github.com/FloatTech/zbpctrl"
	"github.com/FloatTech/zbputils/control"
	zero "github.com/wdvxdr1123/ZeroBot"
	"github.com/wdvxdr1123/ZeroBot/message"
)

const (
	url = "https://enka.microgg.cn/u/%v/__data.json"
	tu  = "https://enka.network/ui/%v.png"
)

func init() { // 主函数
	en := control.Register("yuanshen", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "原神相关功能",
		Help: "命令大全\n" +
			"- 神里面板",
	})
	en.OnSuffix("面板").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		str := ctx.State["args"].(string) //获取key
		var wifeid int64
		/*qquid := ctx.Event.UserID
		//获取uid
		//uid := Getuid(qquid)*/
		uid := "113781666" //测试用
		if uid == "" {
			ctx.SendChain(message.Text("未绑定uid"))
			return
		}
		es, err := web.GetData(fmt.Sprintf(url, uid)) // 网站返回结果
		if err != nil {
			ctx.SendChain(message.Text("网站获取信息失败", err))
			return
		}
		//解析
		var alldata Data
		err = json.Unmarshal(es, &alldata)
		if err != nil {
			ctx.SendChain(message.Text("出现错误捏：", err))
			return
		}
		switch str {
		case "全部":
			ctx.SendChain(message.Text())
		default: //角色名解析为id
			var flag bool
			wifeid, flag = Namemap[str]
			if flag == false {
				ctx.SendChain(message.Text("请输入角色全名"))
				return
			}
		}
		var t = -1
		//匹配角色
		for i := 0; i < 8; i++ {
			if wifeid == int64(alldata.PlayerInfo.ShowAvatarInfoList[i].AvatarID) {
				t = i
			}
		}
		if t == -1 { //在返回数据中未找到想要的角色
			ctx.SendChain(message.Text("该角色未展示"))
			return
		}
		/*角色数据
		uid=uid
		游戏昵称a:= alldata.PlayerInfo.Nickname
		深渊层数:b:=alldata.PlayerInfo.TowerFloorIndex
		角色的基本信息:名字:str  等级: a := alldata.PlayerInfo.ShowAvatarInfoList[t].Level
		好感度:a := alldata.AvatarInfoList[t].FetterInfo.ExpLevel
		插画:
		角色属性:生命: a :=(int) alldata.AvatarInfoList[t].FightPropMap.Num2000
		攻击力:a :=(int) alldata.AvatarInfoList[t].FightPropMap.Num2001
		防御力:a :=(int) alldata.AvatarInfoList[t].FightPropMap.Num2002
		元素精通:a :=(int) alldata.AvatarInfoList[t].FightPropMap.Num28
		暴击率:a :=(int) alldata.AvatarInfoList[t].FightPropMap.Num20
		暴击伤害:a :=(int) alldata.AvatarInfoList[t].FightPropMap.Num22
		元素充能:a :=(int) alldata.AvatarInfoList[t].FightPropMap.Num23
		元素加伤:a :=(int) alldata.AvatarInfoList[t].FightPropMap.Num30/40...46
		武器:名称: 等级: 攻击力: 副词条: 精炼等级: 插画:
		圣遗物:
		花:等级: 插画: 主词条: 副词条:1 2 3 4
		羽:
		沙:
		杯:
		冠:
		命之座:数字几命
		天赋:1插画:等级:
		    2
			3
		*/
		//a := alldata.AvatarInfoList[t].FightPropMap.Num2000
		//画图
		dc := gg.NewContext(1920, 1080) // 画布大小
		dc.SetHexColor("#98F5FF")
		dc.Clear()         // 背景
		dc.SetRGB(0, 0, 0) // 换黑色
		//角色立绘
		lihui, err := gg.LoadImage("data/yuanshen/lihui/" + str + "/01.jpg")
		if err != nil {
			ctx.SendChain(message.Text("获取图片失败", err))
			return
		}
		dc.DrawImage(lihui, 10, 0)
		//输出图片

		ff, cl := writer.ToBytes(dc.Image())  //图片放入缓存
		ctx.SendChain(message.ImageBytes(ff)) //输出
		cl()
	})

	//获取json
	en.OnFullMatch("更新面板").SetBlock(true).Handle(func(ctx *zero.Ctx) {

		ctx.SendChain(message.Text("更新完成"))
	})
	//绑定uid
	en.OnPrefix("绑定").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		uid := ctx.State["args"].(string)
		qquid := ctx.Event.UserID
		//存储进数据库
		ctx.SendChain(message.Text("绑定完成"))
	})
}
