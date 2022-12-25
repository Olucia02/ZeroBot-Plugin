package yuanshen // 导入yuan-shen模块
import ()

const (
// url = "https://enka.microgg.cn/u/%v/__data.json"
)

var Uidmap = map[int64]string{
	10000036: "重云",
	10000050: "托马",
	10000051: "优菈",
	10000066: "神里绫人",
	10000067: "柯莱",
	10000016: "迪卢克",
	10000025: "行秋",
	10000030: "钟离",
	10000053: "早柚",
	10000071: "赛诺",
	10000002: "神里绫华",
	10000003: "琴",
	10000005: "空",
	10000068: "多莉",
	10000070: "妮露",
	10000072: "坎蒂丝",
	10000001: "凯特",
	10000055: "五郎",
	10000060: "夜兰",
	10000023: "香菱",
	10000042: "刻晴",
	10000039: "迪奥娜",
	10000057: "荒泷一斗",
	10000069: "提纳里",
	10000075: "流浪者",
	10000078: "艾尔海森",
	10000006: "丽莎",
	10000014: "芭芭拉",
	10000049: "宵宫",
	10000056: "九条裟罗",
	10000058: "八重神子",
	10000065: "久岐忍",
	10000044: "辛焱",
	10000047: "枫原万叶",
	10000046: "胡桃",
	10000048: "烟绯",
	10000063: "申鹤",
	10000076: "珐露珊",
	10000015: "凯亚",
	10000020: "雷泽",
	10000074: "莱依拉",
	10000022: "温迪",
	10000064: "云堇",
	10000034: "诺艾尔",
	10000077: "瑶瑶",
	10000021: "安柏",
	10000033: "达达利亚",
	10000043: "砂糖",
	10000029: "可莉",
	10000045: "罗莎莉亚",
	10000062: "埃洛伊",
	10000035: "七七",
	10000052: "雷电将军",
	10000031: "菲谢尔",
	10000037: "甘雨",
	10000038: "阿贝多",
	10000041: "莫娜",
	10000007: "荧",
	10000024: "北斗",
	10000054: "珊瑚宫心海",
	10000026: "魈",
	10000027: "凝光",
	10000073: "纳西妲",
	10000032: "班尼特",
	10000059: "鹿野院平藏",
}

var Namemap = map[string]int64{
	"多莉":    10000068,
	"凯亚":    10000015,
	"宵宫":    10000049,
	"夜兰":    10000060,
	"莱依拉":   10000074,
	"艾尔海森":  10000078,
	"重云":    10000036,
	"辛焱":    10000044,
	"班尼特":   10000032,
	"胡桃":    10000046,
	"云堇":    10000064,
	"久岐忍":   10000065,
	"空":     10000005,
	"香菱":    10000023,
	"早柚":    10000053,
	"神里绫人":  10000066,
	"达达利亚":  10000033,
	"刻晴":    10000042,
	"钟离":    10000030,
	"迪奥娜":   10000039,
	"优菈":    10000051,
	"九条裟罗":  10000056,
	"安柏":    10000021,
	"北斗":    10000024,
	"申鹤":    10000063,
	"纳西妲":   10000073,
	"砂糖":    10000043,
	"荒泷一斗":  10000057,
	"坎蒂丝":   10000072,
	"雷电将军":  10000052,
	"鹿野院平藏": 10000059,
	"诺艾尔":   10000034,
	"甘雨":    10000037,
	"莫娜":    10000041,
	"八重神子":  10000058,
	"柯莱":    10000067,
	"妮露":    10000070,
	"神里绫华":  10000002,
	"丽莎":    10000006,
	"流浪者":   10000075,
	"迪卢克":   10000016,
	"烟绯":    10000048,
	"雷泽":    10000020,
	"阿贝多":   10000038,
	"魈":     10000026,
	"芭芭拉":   10000014,
	"菲谢尔":   10000031,
	"托马":    10000050,
	"珊瑚宫心海": 10000054,
	"提纳里":   10000069,
	"赛诺":    10000071,
	"琴":     10000003,
	"荧":     10000007,
	"珐露珊":   10000076,
	"罗莎莉亚":  10000045,
	"五郎":    10000055,
	"埃洛伊":   10000062,
	"瑶瑶":    10000077,
	"温迪":    10000022,
	"凝光":    10000027,
	"可莉":    10000029,
	"七七":    10000035,
	"枫原万叶":  10000047,
	"凯特":    10000001,
	"行秋":    10000025,
}

type Data struct {
	PlayerInfo struct {
		Nickname             string `json:"nickname"`
		Level                int    `json:"level"`
		Signature            string `json:"signature"`
		WorldLevel           int    `json:"worldLevel"`
		NameCardID           int    `json:"nameCardId"`
		FinishAchievementNum int    `json:"finishAchievementNum"`
		TowerFloorIndex      int    `json:"towerFloorIndex"`
		TowerLevelIndex      int    `json:"towerLevelIndex"`
		ShowAvatarInfoList   []struct {
			AvatarID  int `json:"avatarId"`
			Level     int `json:"level"`
			CostumeID int `json:"costumeId,omitempty"`
		} `json:"showAvatarInfoList"`
		ShowNameCardIDList []int `json:"showNameCardIdList"`
		ProfilePicture     struct {
			AvatarID int `json:"avatarId"`
		} `json:"profilePicture"`
	} `json:"playerInfo"`
	AvatarInfoList []struct {
		AvatarID int `json:"avatarId"`
		PropMap  struct {
			Num1001 struct {
				Type int    `json:"type"`
				Ival string `json:"ival"`
			} `json:"1001"`
			Num1002 struct {
				Type int    `json:"type"`
				Ival string `json:"ival"`
				Val  string `json:"val"`
			} `json:"1002"`
			Num1003 struct {
				Type int    `json:"type"`
				Ival string `json:"ival"`
			} `json:"1003"`
			Num1004 struct {
				Type int    `json:"type"`
				Ival string `json:"ival"`
			} `json:"1004"`
			Num4001 struct {
				Type int    `json:"type"`
				Ival string `json:"ival"`
				Val  string `json:"val"`
			} `json:"4001"`
			Num10010 struct {
				Type int    `json:"type"`
				Ival string `json:"ival"`
				Val  string `json:"val"`
			} `json:"10010"`
		} `json:"propMap"`
		FightPropMap struct {
			Num1    float64 `json:"1"`
			Num2    float64 `json:"2"`
			Num3    float64 `json:"3"`
			Num4    float64 `json:"4"`
			Num5    float64 `json:"5"`
			Num6    float64 `json:"6"`
			Num7    float64 `json:"7"`
			Num8    float64 `json:"8"`
			Num20   float64 `json:"20"`
			Num21   float64 `json:"21"`
			Num22   float64 `json:"22"`
			Num23   float64 `json:"23"`
			Num26   float64 `json:"26"`
			Num27   float64 `json:"27"`
			Num28   float64 `json:"28"`
			Num29   float64 `json:"29"`
			Num30   float64 `json:"30"`
			Num40   float64 `json:"40"`
			Num41   float64 `json:"41"`
			Num42   float64 `json:"42"`
			Num43   float64 `json:"43"`
			Num44   float64 `json:"44"`
			Num45   float64 `json:"45"`
			Num46   float64 `json:"46"`
			Num50   float64 `json:"50"`
			Num51   float64 `json:"51"`
			Num52   float64 `json:"52"`
			Num53   float64 `json:"53"`
			Num54   float64 `json:"54"`
			Num55   float64 `json:"55"`
			Num56   float64 `json:"56"`
			Num70   float64 `json:"70"`
			Num80   float64 `json:"80"`
			Num1000 float64 `json:"1000"`
			Num1010 float64 `json:"1010"`
			Num2000 float64 `json:"2000"`
			Num2001 float64 `json:"2001"`
			Num2002 float64 `json:"2002"`
			Num2003 float64 `json:"2003"`
			Num3007 float64 `json:"3007"`
			Num3008 float64 `json:"3008"`
			Num3015 float64 `json:"3015"`
			Num3016 float64 `json:"3016"`
			Num3017 float64 `json:"3017"`
			Num3018 float64 `json:"3018"`
			Num3019 float64 `json:"3019"`
			Num3020 float64 `json:"3020"`
			Num3021 float64 `json:"3021"`
			Num3022 float64 `json:"3022"`
			Num3045 float64 `json:"3045"`
			Num3046 float64 `json:"3046"`
		} `json:"fightPropMap"`
		SkillDepotID           int   `json:"skillDepotId"`
		InherentProudSkillList []int `json:"inherentProudSkillList"`
		SkillLevelMap          struct {
			Num10461 int `json:"10461"`
			Num10462 int `json:"10462"`
			Num10463 int `json:"10463"`
		} `json:"skillLevelMap"`
		EquipList []struct {
			ItemID    int `json:"itemId"`
			Reliquary struct {
				Level            int   `json:"level"`
				MainPropID       int   `json:"mainPropId"`
				AppendPropIDList []int `json:"appendPropIdList"`
			} `json:"reliquary,omitempty"`
			Flat struct {
				NameTextMapHash    string `json:"nameTextMapHash"`
				SetNameTextMapHash string `json:"setNameTextMapHash"`
				RankLevel          int    `json:"rankLevel"`
				ReliquaryMainstat  struct {
					MainPropID string  `json:"mainPropId"`
					StatValue  float64 `json:"statValue"`
				} `json:"reliquaryMainstat"`
				ReliquarySubstats []struct {
					AppendPropID string  `json:"appendPropId"`
					StatValue    float64 `json:"statValue"`
				} `json:"reliquarySubstats"`
				ItemType  string `json:"itemType"`
				Icon      string `json:"icon"`
				EquipType string `json:"equipType"`
			} `json:"flat"`
			Weapon struct {
				Level        int `json:"level"`
				PromoteLevel int `json:"promoteLevel"`
				AffixMap     struct {
					Num113405 int `json:"113405"`
				} `json:"affixMap"`
			} `json:"weapon,omitempty"`
		} `json:"equipList"`
		FetterInfo struct {
			ExpLevel int `json:"expLevel"`
		} `json:"fetterInfo"`
		TalentIDList            []int `json:"talentIdList,omitempty"`
		ProudSkillExtraLevelMap struct {
			Num4239 int `json:"4239"`
		} `json:"proudSkillExtraLevelMap,omitempty"`
		CostumeID int `json:"costumeId,omitempty"`
	} `json:"avatarInfoList"`
	TTL int    `json:"ttl"`
	UID string `json:"uid"`
}

func Getuid(qquid int64) (uid string) { //获取对应游戏uid
	/*for _, r := range uidList {
		if r.Qquid == qquid {
			uid = r.UID
			return
		}
	}*/
	return ""
}
