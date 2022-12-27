// Package kokomi  原神面板v1
package kokomi

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

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
	en := control.Register("kokomi", &ctrl.Options[*zero.Ctx]{
		DisableOnDefault: false,
		Brief:            "原神相关功能",
		Help: "命令大全,需要依次执行\n" +
			"- 绑定......(uid)\n" +
			"- 更新面板\n" +
			"- 全部面板\n" +
			"- XX面板",
	})
	en.OnSuffix("面板").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		str := ctx.State["args"].(string) // 获取key
		var wifeid int64
		qquid := ctx.Event.UserID
		// 获取uid
		uid := Getuid(qquid)
		// uid := 113781666 //测试用
		suid := strconv.Itoa(uid)
		if uid == 0 {
			ctx.SendChain(message.Text("未绑定uid"))
			return
		}
		// 获取本地缓存数据
		txt, err := os.ReadFile("data/kokomi/js/" + suid + ".txt")
		if err != nil {
			ctx.SendChain(message.Text("错误,本地未找到账号信息", err))
		}

		// 解析
		var alldata Data
		err = json.Unmarshal(txt, &alldata)
		if err != nil {
			ctx.SendChain(message.Text("出现错误捏：", err))
			return
		}
		switch str {
		case "全部":
			var msg strings.Builder
			msg.WriteString("您的展示角色为:\n")
			for i := 0; i < 8; i++ {
				mmm, _ := Uidmap[int64(alldata.PlayerInfo.ShowAvatarInfoList[i].AvatarID)]
				msg.WriteString(mmm)
				if i < 7 {
					msg.WriteByte('\n')
				}
			}
			ctx.SendChain(message.Text(msg.String()))
			return
		default: // 角色名解析为id
			var flag bool
			wifeid, flag = Namemap[str]
			if !flag {
				ctx.SendChain(message.Text("请输入角色全名"))
				return
			}
		}
		var t = -1
		// 匹配角色
		for i := 0; i < 8; i++ {
			if wifeid == int64(alldata.PlayerInfo.ShowAvatarInfoList[i].AvatarID) {
				t = i
			}
		}
		if t == -1 { // 在返回数据中未找到想要的角色
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
		// a := alldata.AvatarInfoList[t].FightPropMap.Num2000
		// 画图
		dc := gg.NewContext(1920, 1080) // 画布大小
		dc.SetHexColor("#98F5FF")
		dc.Clear() // 背景
		pro, flg := Promap[wifeid]
		if !flg {
			ctx.SendChain(message.Text("匹配角色元素失败"))
			return
		}
		beijing, err := gg.LoadImage("data/kokomi/pro/" + pro + ".jpg")
		if err != nil {
			ctx.SendChain(message.Text("获取背景失败", err))
			return
		}
		dc.DrawImage(beijing, 0, 0)
		dc.SetRGB(1, 1, 1) // 换白色
		// 角色立绘565*935
		//	lihui, err := gg.LoadImage("data/kokomi/lihui/" + str + "/01.jpg")
		lihui, err := gg.LoadImage("data/kokomi/character/" + str + "/imgs/splash.webp")
		if err != nil {
			ctx.SendChain(message.Text("获取立绘失败", err))
			return
		}
		dc.Scale(0.8, 0.8)
		dc.DrawImage(lihui, -300, 0)
		dc.Scale(5.0/4, 5.0/4)
		// 好感度,uid
		FontFile := "data/kokomi/font/sakura.ttf" // 字体
		if err := dc.LoadFontFace(FontFile, 25); err != nil {
			panic(err)
		}
		// 版本号
		dc.DrawString("ZeroBot-Plugin v1.6.1-beta2&kokomi v1", 630, 1070)
		if err := dc.LoadFontFace(FontFile, 40); err != nil {
			panic(err)
		}
		dc.DrawString("好感度"+strconv.Itoa(alldata.AvatarInfoList[t].FetterInfo.ExpLevel), 35, 980)
		dc.DrawString("昵称:"+alldata.PlayerInfo.Nickname, 35, 1030)
		dc.DrawString("uid:"+suid, 35, 1075)
		// 角色等级,武器精炼
		dc.DrawString("LV"+strconv.Itoa(alldata.PlayerInfo.ShowAvatarInfoList[t].Level), 630, 130) // 角色等级
		ming := len(alldata.AvatarInfoList[t].TalentIDList)
		dc.DrawString(strconv.Itoa(ming)+"命", 765, 130)
		// 角色名字630/75,字55
		if err := dc.LoadFontFace(FontFile, 55); err != nil { // 字体大小
			panic(err)
		}
		dc.DrawString(str, 630, 75)

		//新建图层,实现阴影400*510
		bg := Yinying(400, 510, 16)
		//字图层
		one := gg.NewContext(400, 500)
		// 属性630*370,字50
		if err := one.LoadFontFace(FontFile, 50); err != nil { // 字体大小
			panic(err)
		}
		one.SetRGB(1, 1, 1) //白色
		one.DrawString("生命值:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num2000)), 5, 65)
		one.DrawString("攻击力:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num2001)), 5, 125)
		one.DrawString("防御力:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num2002)), 5, 185)
		one.DrawString("元素精通:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num28)), 5, 245)
		one.DrawString("暴击率:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num20*100))+"%", 5, 305)
		one.DrawString("暴击伤害:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num22*100))+"%", 5, 365)
		one.DrawString("元素充能:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num23*100))+"%", 5, 425)
		// 元素加伤判断
		switch {
		case alldata.AvatarInfoList[t].FightPropMap.Num30*100 > 0:
			one.DrawString("物理加伤:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num30*100))+"%", 5, 485)
		case alldata.AvatarInfoList[t].FightPropMap.Num40*100 > 0:
			one.DrawString("火元素加伤:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num40*100))+"%", 5, 485)
		case alldata.AvatarInfoList[t].FightPropMap.Num41*100 > 0:
			one.DrawString("雷元素加伤:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num41*100))+"%", 5, 485)
		case alldata.AvatarInfoList[t].FightPropMap.Num42*100 > 0:
			one.DrawString("水元素加伤:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num42*100))+"%", 5, 485)
		case alldata.AvatarInfoList[t].FightPropMap.Num44*100 > 0:
			one.DrawString("风元素加伤:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num44*100))+"%", 5, 485)
		case alldata.AvatarInfoList[t].FightPropMap.Num45*100 > 0:
			one.DrawString("岩元素加伤:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num45*100))+"%", 5, 485)
		case alldata.AvatarInfoList[t].FightPropMap.Num46*100 > 0:
			one.DrawString("冰元素加伤:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num46*100))+"%", 5, 485)
		default: //草或者无
			one.DrawString("元素加伤:"+strconv.Itoa(int(alldata.AvatarInfoList[t].FightPropMap.Num43*100))+"%", 5, 485)
		}
		dc.DrawImage(bg, 710, 320)
		dc.DrawImage(one.Image(), 710, 310)

		// 天赋等级
		if err := dc.LoadFontFace(FontFile, 65); err != nil { // 字体大小
			panic(err)
		}
		var link = []int{0, 0, 0}
		var i = 0
		for k, _ := range alldata.AvatarInfoList[t].SkillLevelMap {
			link[i] = k
			i++
		}
		sort.Ints(link)
		lin1, _ := alldata.AvatarInfoList[t].SkillLevelMap[link[0]]
		lin2, _ := alldata.AvatarInfoList[t].SkillLevelMap[link[1]]
		lin3, _ := alldata.AvatarInfoList[t].SkillLevelMap[link[2]]
		dc.DrawString("天赋等级:"+strconv.Itoa(lin1)+"--"+strconv.Itoa(lin2)+"--"+strconv.Itoa(lin3), 630, 900)
		//武器名字
		if err := dc.LoadFontFace(FontFile, 50); err != nil { // 字体大小
			panic(err)
		}
		wq, _ := IdforNamemap[alldata.AvatarInfoList[t].EquipList[5].Flat.NameTextHash]
		dc.DrawString(wq, 890, 85)

		//详细
		if err := dc.LoadFontFace(FontFile, 40); err != nil { // 字体大小
			panic(err)
		}
		dc.DrawString("精炼:"+strconv.Itoa(int(alldata.AvatarInfoList[t].EquipList[5].Flat.RankLevel)), 890, 145)
		//wq攻击力
		if err := dc.LoadFontFace(FontFile, 45); err != nil { // 字体大小
			panic(err)
		}
		dc.DrawString("攻击力:"+strconv.FormatFloat(alldata.AvatarInfoList[t].EquipList[5].Flat.WeaponStat[0].Value, 'f', 1, 32), 820, 200)
		//Lv
		dc.DrawString("Lv:"+strconv.Itoa(alldata.AvatarInfoList[t].EquipList[5].Weapon.Level), 1110, 200)
		//副词条
		fucitiao, _ := IdforNamemap[alldata.AvatarInfoList[t].EquipList[5].Flat.WeaponStat[1].SubPropId] //名称
		var baifen = "%"
		if fucitiao == "元素精通" {
			baifen = ""
		}
		dc.DrawString(fucitiao+":"+strconv.Itoa(int(alldata.AvatarInfoList[t].EquipList[5].Flat.WeaponStat[1].Value))+baifen, 820, 270)
		//图片
		tuwq, err := gg.LoadPNG("data/kokomi/wq/" + wq + ".png")
		if err != nil {
			ctx.SendChain(message.Text("获取武器图标", err))
			return
		}
		dc.Scale(1.5, 1.5)
		dc.DrawImage(tuwq, 400, 90)
		dc.Scale(1/1.5, 1/1.5)
		//圣遗物
		//缩小
		dc.Scale(0.5, 0.5)
		for i := 0; i < 5; i++ {
			sywname, _ := IdforNamemap[alldata.AvatarInfoList[t].EquipList[i].Flat.SetNameTextHash]
			tusyw, err := gg.LoadImage("data/kokomi/syw/" + sywname + "/" + strconv.Itoa(i+1) + ".webp")
			if err != nil {
				ctx.SendChain(message.Text("获取圣遗物图标", err))
				return
			}
			//圣遗物图标坐标
			var x, y int
			switch i {
			case 0:
				x = 1920 - 310
				y = 35
			case 1:
				x = 1920 - 620
				y = 200
			case 2:
				x = 1920 - 310
				y = 200
			case 3:
				x = 1920 - 620
				y = 365
			case 4:
				x = 1920 - 310
				y = 365
			}
			dc.DrawImage(tusyw, x*2, y*2-5)
		}
		//恢复大小
		dc.Scale(2, 2)
		//圣遗物属性
		for i := 0; i < 5; i++ {
			var x, y int //基轴
			switch i {
			case 0:
				x = 1920 - 310
				y = 35
			case 1:
				x = 1920 - 630
				y = 200
			case 2:
				x = 1920 - 310
				y = 200
			case 3:
				x = 1920 - 630
				y = 365
			case 4:
				x = 1920 - 310
				y = 365
			}
			if err := dc.LoadFontFace(FontFile, 35); err != nil { // 字体大小
				panic(err)
			}
			zhuci := StoS(alldata.AvatarInfoList[t].EquipList[i].Flat.ReliquaryMainStat.MainPropId)
			dc.DrawString(zhuci, float64(x+135), float64(y+35))                                                                                  //主词条
			dc.DrawString(strconv.Itoa(int(alldata.AvatarInfoList[t].EquipList[i].Flat.ReliquaryMainStat.Value)), float64(x+135), float64(y+72)) //主词条
			if err := dc.LoadFontFace(FontFile, 30); err != nil {                                                                                // 字体大小
				panic(err)
			}
			for k := 0; k < 4; k++ {
				var xx, yy int
				switch k {
				case 0:
					xx = x
					yy = y + 115
				case 1:
					xx = x + 150
					yy = y + 115
				case 2:
					xx = x
					yy = y + 150
				case 3:
					xx = x + 150
					yy = y + 150
				}
				dc.DrawString(StoS(alldata.AvatarInfoList[t].EquipList[i].Flat.ReliquarySubStats[k].SubPropId)+":"+strconv.FormatFloat(alldata.AvatarInfoList[t].EquipList[i].Flat.ReliquarySubStats[k].Value, 'f', 1, 64), float64(xx), float64(yy))
			}

		}

		// 输出图片
		ff, cl := writer.ToBytes(dc.Image())  // 图片放入缓存
		ctx.SendChain(message.ImageBytes(ff)) // 输出
		cl()
	})

	// 获取json
	en.OnFullMatch("更新面板").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		qquid := ctx.Event.UserID
		uid := Getuid(qquid)
		// uid := 113781666
		suid := strconv.Itoa(uid)
		es, err := web.GetData(fmt.Sprintf(url, uid)) // 网站返回结果
		if err != nil {
			ctx.SendChain(message.Text("网站获取信息失败", err))
			return
		}
		// 创建存储文件,路径data/kokomi/js
		file, _ := os.OpenFile("data/kokomi/js/"+suid+".txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		_, _ = file.Write(es)
		ctx.SendChain(message.Text("更新成功"))
		file.Close()
	})
	// 绑定uid
	en.OnPrefix("绑定").SetBlock(true).Handle(func(ctx *zero.Ctx) {
		uid := ctx.State["args"].(string)
		sqquid := strconv.Itoa(int(ctx.Event.UserID))
		file, _ := os.OpenFile("data/kokomi/uid/"+sqquid+".txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		_, _ = file.Write([]byte(uid))
		file.Close()
		// 存储进数据库
		ctx.SendChain(message.Text("绑定成功"))
	})
}
