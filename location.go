package laction

import (
	"strings"
)

func GetLocation(country, province, city string) string {
	country = strings.ToLower(country)
	if country != "china" {
		if c, ok := country_map[country]; ok {
			return c
		}
		return country
	}

	province = strings.ToLower(province)
	city = strings.ToLower(city)

	if province == "shanxi" && strings.Contains(`"xian", "xianyang", "yanan", "hanzhong", "yulin", "shangnan", "lueyang", "yijun", "linyou", "baihe"`, city) {
		province = "shanxin2"
	}

	if prov, ok := location_en[province]; ok {
		for _, v := range prov {
			if city == v {
				if zh, ok := en_zh_map[province+"_"+city]; ok {
					return zh
				}
			}
		}
	}

	return city
}

var location_en = map[string][]string{
	"beijing":      []string{"xicheng", "dongcheng", "chongwen", "xuanwu", "chaoyang", "haidian", "fengtai", "shijingshan", "mentougou", "fangshan", "tongzhou", "shunyi", "daxing", "changping", "pinggu", "huairou", "miyun", "yanqing"},
	"tianjin":      []string{"qingyang", "hedong", "hexi", "nankai", "hebei", "hongqiao", "tanggu", "hangu", "dagang", "dongli", "xiqing", "beichen", "jinnan", "wuqing", "baodi", "jinghai", "ninghe", "jixian", "kaifaqu"},
	"hebei":        []string{"shijiazhuang", "qinhuangdao", "langfang", "baoding", "handan", "tangshan", "xingtai", "hengshui", "zhangjiakou", "chengde", "cangzhou", "hengshui"},
	"shanxi":       []string{"taiyuan", "datong", "changzhi", "jinzhong", "yangquan", "shuozhou", "yuncheng", "linfen"},
	"namenggu":     []string{"huhehaote", "chifeng", "tongliao", "xilinguole", "xingan"},
	"liaoning":     []string{"dalian", "shenyang", "anshan", "fushun", "yingkou", "jinzhou", "dandong", "chaoyang", "liaoyang", "fuxin", "tieling", "panjin", "benxi", "huludao"},
	"jilin":        []string{"changchun", "jilin", "siping", "liaoyuan", "tonghua", "yanji", "baicheng", "liaoyuan", "songyuan", "linjiang", "huichun"},
	"heilongjiang": []string{"haerbin", "qiqihaer", "daqing", "mudanjiang", "hegang", "jiamusi", "suihua"},
	"shanghai":     []string{"pudong", "yangpu", "xuhui", "jingan", "luwan", "huangpu", "putuo", "zhabei", "hongkou", "changning", "baoshan", "minxing", "jiading", "jinshan", "songjiang", "qingpu", "chongming", "fengxian", "nanhui"},
	"jiangsu":      []string{"nanjing", "suzhou", "wuxi", "changzhou", "yangzhou", "xuzhou", "nantong", "zhenjiang", "taizhou", "huaian", "lianyungang", "suqian", "yancheng", "huaiyin", "muyang", "zhangjiagang"},
	"zhejiang":     []string{"hangzhou", "jinhua", "ningbo", "wenzhou", "jiaxing", "shaoxing", "lishui", "huzhou", "taizhou", "zhoushan", "quzhou"},
	"anhui":        []string{"hefei", "maanshan", "bangbu", "huangshan", "wuhu", "huainan", "tongling", "fuyang", "xuancheng", "anqing"},
	"fujian":       []string{"fuzhou", "xiamen", "quanzhou", "zhangzhou", "nanping", "longyan", "putian", "sanming", "ningde"},
	"jiangxi":      []string{"nanchang", "jingdezhen", "shangrao", "pingxiang", "jiujiang", "jian", "yichun", "yingtan", "xinyu", "ganzhou"},
	"shandong":     []string{"qingdao", "jinan", "zibo", "yantai", "taian", "linyi", "rizhao", "dezhou", "weihai", "dongying", "heze", "jining", "weifang", "zaozhuang", "liaocheng"},
	"henan":        []string{"zhengzhou", "luoyang", "kaifeng", "pingdingshan", "puyang", "anyang", "xuchang", "nanyang", "xinyang", "zhoukou", "xinxiang", "jiaozuo", "sanmenxia", "shangqiu"},
	"hubei":        []string{"wuhan", "xiangfan", "xiaogan", "shiyan", "jingzhou", "huangshi", "yichang", "huanggang", "enshi", "ezhou", "jianghan", "suizao", "jingsha", "xianning"},
	"hunan":        []string{"changsha", "xiangtan", "yueyang", "zhuzhou", "huaihua", "yongzhou", "yiyang", "zhangjiajie", "changde", "hengyang", "xiangxi", "shaoyang", "loudi", "chenzhou"},
	"guangdong":    []string{"guangzhou", "shenzhen", "dongwan", "foshan", "zhuhai", "shantou", "shaoguan", "jiangmen", "meizhou", "jieyang", "zhongshan", "heyuan", "huizhou", "maoming", "zhanjiang", "yangjiang", "chaozhou", "yunfu", "shanwei", "chaoyang", "zhaoqing", "shunde", "qingyuan"},
	"guangxi":      []string{"nanning", "guilin", "liuzhou", "wuzhou", "laibin", "guigang", "yulin", "hezhou"},
	"hainan":       []string{"haikou", "sanya"},
	"zhongqing":    []string{"yuzhong", "dadukou", "jiangbei", "shapingba", "jiulongpo", "nanan", "beibei", "wansheng", "shuangqiao", "yubei", "banan", "wanzhou", "fuling", "qianjiang", "changshou"},
	"sichuan":      []string{"chengdou", "dazhou", "nanchong", "leshan", "mianyang", "deyang", "najiang", "suining", "yibin", "bazhong", "zigong", "kangding", "panzhihua"},
	"guizhou":      []string{"guiyang", "zunyi", "anshun", "qianxinan", "douyun"},
	"yunnan":       []string{"kunming", "lijiang", "zhaotong", "yuxi", "lincang", "wenshan", "honghe", "chuxiong", "dali"},
	"xicang":       []string{"lasa", "linzhi", "rikaze", "changdou"},
	"shanxi2":      []string{"xian", "xianyang", "yanan", "hanzhong", "yulin", "shangnan", "lueyang", "yijun", "linyou", "baihe"},
	"gansu":        []string{"lanzhou", "jinchang", "tianshui", "wuwei", "zhangye", "pingliang", "jiuquan"},
	"qinghai":      []string{"huangnan", "hainan", "xining", "haidong", "haixi", "haibei", "guoluo", "yushu"},
	"ningxia":      []string{"yinchuan", "wuzhong"},
	"xinjiang":     []string{"wulumuqi", "hami", "kashi", "bayinguoleng", "changji", "yili", "aletai", "kelamayi", "boertala"},
	"xianggang":    []string{"zhongxiqu", "wanziqu", "dongqu", "nanqu", "jiulong-youjianwangqu", "jiulong-shenshuibuqu", "jiulong-jiulongchengqu", "jiulong-huangdaxianqu", "jiulong-guantangqu", "xinjie-beiqu", "xinjie-dapuqu", "xinjie-shatianqu", "xinjie-xigongqu", "xinjie-quanwanqu", "xinjie-tunmenqu", "xinjie-yuanlangqu", "xinjie-kuiqingqu", "xinjie-lidaoqu"},
	"aomen":        []string{"huadimatangqu", "shenganduonitangqu", "datangqu", "wangdetangqu", "fengshuntangqu", "jiamotangqu", "shengfangjigetangqu", "ludangcheng"},
}

var en_zh_map = map[string]string{
	"beijing_xicheng":                  "西城",
	"beijing_dongcheng":                "东城",
	"beijing_chongwen":                 "崇文",
	"beijing_xuanwu":                   "宣武",
	"beijing_chaoyang":                 "朝阳",
	"beijing_haidian":                  "海淀",
	"beijing_fengtai":                  "丰台",
	"beijing_shijingshan":              "石景山",
	"beijing_mentougou":                "门头沟",
	"beijing_fangshan":                 "房山",
	"beijing_tongzhou":                 "通州",
	"beijing_shunyi":                   "顺义",
	"beijing_daxing":                   "大兴",
	"beijing_changping":                "昌平",
	"beijing_pinggu":                   "平谷",
	"beijing_huairou":                  "怀柔",
	"beijing_miyun":                    "密云",
	"beijing_yanqing":                  "延庆",
	"tianjin_qingyang":                 "青羊",
	"tianjin_hedong":                   "河东",
	"tianjin_hexi":                     "河西",
	"tianjin_nankai":                   "南开",
	"tianjin_hebei":                    "河北",
	"tianjin_hongqiao":                 "红桥",
	"tianjin_tanggu":                   "塘沽",
	"tianjin_hangu":                    "汉沽",
	"tianjin_dagang":                   "大港",
	"tianjin_dongli":                   "东丽",
	"tianjin_xiqing":                   "西青",
	"tianjin_beichen":                  "北辰",
	"tianjin_jinnan":                   "津南",
	"tianjin_wuqing":                   "武清",
	"tianjin_baodi":                    "宝坻",
	"tianjin_jinghai":                  "静海",
	"tianjin_ninghe":                   "宁河",
	"tianjin_jixian":                   "蓟县",
	"tianjin_kaifaqu":                  "开发区",
	"hebei_shijiazhuang":               "石家庄",
	"hebei_qinhuangdao":                "秦皇岛",
	"hebei_langfang":                   "廊坊",
	"hebei_baoding":                    "保定",
	"hebei_handan":                     "邯郸",
	"hebei_tangshan":                   "唐山",
	"hebei_xingtai":                    "邢台",
	"hebei_zhangjiakou":                "张家口",
	"hebei_chengde":                    "承德",
	"hebei_cangzhou":                   "沧州",
	"hebei_hengshui":                   "衡水",
	"shanxi_taiyuan":                   "太原",
	"shanxi_datong":                    "大同",
	"shanxi_changzhi":                  "长治",
	"shanxi_jinzhong":                  "晋中",
	"shanxi_yangquan":                  "阳泉",
	"shanxi_shuozhou":                  "朔州",
	"shanxi_yuncheng":                  "运城",
	"shanxi_linfen":                    "临汾",
	"namenggu_huhehaote":               "呼和浩特",
	"namenggu_chifeng":                 "赤峰",
	"namenggu_tongliao":                "通辽",
	"namenggu_xilinguole":              "锡林郭勒",
	"namenggu_xingan":                  "兴安",
	"liaoning_dalian":                  "大连",
	"liaoning_shenyang":                "沈阳",
	"liaoning_anshan":                  "鞍山",
	"liaoning_fushun":                  "抚顺",
	"liaoning_yingkou":                 "营口",
	"liaoning_jinzhou":                 "锦州",
	"liaoning_dandong":                 "丹东",
	"liaoning_chaoyang":                "朝阳",
	"liaoning_liaoyang":                "辽阳",
	"liaoning_fuxin":                   "阜新",
	"liaoning_tieling":                 "铁岭",
	"liaoning_panjin":                  "盘锦",
	"liaoning_benxi":                   "本溪",
	"liaoning_huludao":                 "葫芦岛",
	"jilin_changchun":                  "长春",
	"jilin_jilin":                      "吉林",
	"jilin_siping":                     "四平",
	"jilin_tonghua":                    "通化",
	"jilin_yanji":                      "延吉",
	"jilin_baicheng":                   "白城",
	"jilin_liaoyuan":                   "辽源",
	"jilin_songyuan":                   "松原",
	"jilin_linjiang":                   "临江",
	"jilin_huichun":                    "珲春",
	"heilongjiang_haerbin":             "哈尔滨",
	"heilongjiang_qiqihaer":            "齐齐哈尔",
	"heilongjiang_daqing":              "大庆",
	"heilongjiang_mudanjiang":          "牡丹江",
	"heilongjiang_hegang":              "鹤岗",
	"heilongjiang_jiamusi":             "佳木斯",
	"heilongjiang_suihua":              "绥化",
	"shanghai_pudong":                  "浦东",
	"shanghai_yangpu":                  "杨浦",
	"shanghai_xuhui":                   "徐汇",
	"shanghai_jingan":                  "静安",
	"shanghai_luwan":                   "卢湾",
	"shanghai_huangpu":                 "黄浦",
	"shanghai_putuo":                   "普陀",
	"shanghai_zhabei":                  "闸北",
	"shanghai_hongkou":                 "虹口",
	"shanghai_changning":               "长宁",
	"shanghai_baoshan":                 "宝山",
	"shanghai_minxing":                 "闵行",
	"shanghai_jiading":                 "嘉定",
	"shanghai_jinshan":                 "金山",
	"shanghai_songjiang":               "松江",
	"shanghai_qingpu":                  "青浦",
	"shanghai_chongming":               "崇明",
	"shanghai_fengxian":                "奉贤",
	"shanghai_nanhui":                  "南汇",
	"jiangsu_nanjing":                  "南京",
	"jiangsu_suzhou":                   "苏州",
	"jiangsu_wuxi":                     "无锡",
	"jiangsu_changzhou":                "常州",
	"jiangsu_yangzhou":                 "扬州",
	"jiangsu_xuzhou":                   "徐州",
	"jiangsu_nantong":                  "南通",
	"jiangsu_zhenjiang":                "镇江",
	"jiangsu_taizhou":                  "泰州",
	"jiangsu_huaian":                   "淮安",
	"jiangsu_lianyungang":              "连云港",
	"jiangsu_suqian":                   "宿迁",
	"jiangsu_yancheng":                 "盐城",
	"jiangsu_huaiyin":                  "淮阴",
	"jiangsu_muyang":                   "沐阳",
	"jiangsu_zhangjiagang":             "张家港",
	"zhejiang_hangzhou":                "杭州",
	"zhejiang_jinhua":                  "金华",
	"zhejiang_ningbo":                  "宁波",
	"zhejiang_wenzhou":                 "温州",
	"zhejiang_jiaxing":                 "嘉兴",
	"zhejiang_shaoxing":                "绍兴",
	"zhejiang_lishui":                  "丽水",
	"zhejiang_huzhou":                  "湖州",
	"zhejiang_taizhou":                 "台州",
	"zhejiang_zhoushan":                "舟山",
	"zhejiang_quzhou":                  "衢州",
	"anhui_hefei":                      "合肥",
	"anhui_maanshan":                   "马鞍山",
	"anhui_bangbu":                     "蚌埠",
	"anhui_huangshan":                  "黄山",
	"anhui_wuhu":                       "芜湖",
	"anhui_huainan":                    "淮南",
	"anhui_tongling":                   "铜陵",
	"anhui_fuyang":                     "阜阳",
	"anhui_xuancheng":                  "宣城",
	"anhui_anqing":                     "安庆",
	"fujian_fuzhou":                    "福州",
	"fujian_xiamen":                    "厦门",
	"fujian_quanzhou":                  "泉州",
	"fujian_zhangzhou":                 "漳州",
	"fujian_nanping":                   "南平",
	"fujian_longyan":                   "龙岩",
	"fujian_putian":                    "莆田",
	"fujian_sanming":                   "三明",
	"fujian_ningde":                    "宁德",
	"jiangxi_nanchang":                 "南昌",
	"jiangxi_jingdezhen":               "景德镇",
	"jiangxi_shangrao":                 "上饶",
	"jiangxi_pingxiang":                "萍乡",
	"jiangxi_jiujiang":                 "九江",
	"jiangxi_jian":                     "吉安",
	"jiangxi_yichun":                   "宜春",
	"jiangxi_yingtan":                  "鹰潭",
	"jiangxi_xinyu":                    "新余",
	"jiangxi_ganzhou":                  "赣州",
	"shandong_qingdao":                 "青岛",
	"shandong_jinan":                   "济南",
	"shandong_zibo":                    "淄博",
	"shandong_yantai":                  "烟台",
	"shandong_taian":                   "泰安",
	"shandong_linyi":                   "临沂",
	"shandong_rizhao":                  "日照",
	"shandong_dezhou":                  "德州",
	"shandong_weihai":                  "威海",
	"shandong_dongying":                "东营",
	"shandong_heze":                    "荷泽",
	"shandong_jining":                  "济宁",
	"shandong_weifang":                 "潍坊",
	"shandong_zaozhuang":               "枣庄",
	"shandong_liaocheng":               "聊城",
	"henan_zhengzhou":                  "郑州",
	"henan_luoyang":                    "洛阳",
	"henan_kaifeng":                    "开封",
	"henan_pingdingshan":               "平顶山",
	"henan_puyang":                     "濮阳",
	"henan_anyang":                     "安阳",
	"henan_xuchang":                    "许昌",
	"henan_nanyang":                    "南阳",
	"henan_xinyang":                    "信阳",
	"henan_zhoukou":                    "周口",
	"henan_xinxiang":                   "新乡",
	"henan_jiaozuo":                    "焦作",
	"henan_sanmenxia":                  "三门峡",
	"henan_shangqiu":                   "商丘",
	"hubei_wuhan":                      "武汉",
	"hubei_xiangfan":                   "襄樊",
	"hubei_xiaogan":                    "孝感",
	"hubei_shiyan":                     "十堰",
	"hubei_jingzhou":                   "荆州",
	"hubei_huangshi":                   "黄石",
	"hubei_yichang":                    "宜昌",
	"hubei_huanggang":                  "黄冈",
	"hubei_enshi":                      "恩施",
	"hubei_ezhou":                      "鄂州",
	"hubei_jianghan":                   "江汉",
	"hubei_suizao":                     "随枣",
	"hubei_jingsha":                    "荆沙",
	"hubei_xianning":                   "咸宁",
	"hunan_changsha":                   "长沙",
	"hunan_xiangtan":                   "湘潭",
	"hunan_yueyang":                    "岳阳",
	"hunan_zhuzhou":                    "株洲",
	"hunan_huaihua":                    "怀化",
	"hunan_yongzhou":                   "永州",
	"hunan_yiyang":                     "益阳",
	"hunan_zhangjiajie":                "张家界",
	"hunan_changde":                    "常德",
	"hunan_hengyang":                   "衡阳",
	"hunan_xiangxi":                    "湘西",
	"hunan_shaoyang":                   "邵阳",
	"hunan_loudi":                      "娄底",
	"hunan_chenzhou":                   "郴州",
	"guangdong_guangzhou":              "广州",
	"guangdong_shenzhen":               "深圳",
	"guangdong_dongwan":                "东莞",
	"guangdong_foshan":                 "佛山",
	"guangdong_zhuhai":                 "珠海",
	"guangdong_shantou":                "汕头",
	"guangdong_shaoguan":               "韶关",
	"guangdong_jiangmen":               "江门",
	"guangdong_meizhou":                "梅州",
	"guangdong_jieyang":                "揭阳",
	"guangdong_zhongshan":              "中山",
	"guangdong_heyuan":                 "河源",
	"guangdong_huizhou":                "惠州",
	"guangdong_maoming":                "茂名",
	"guangdong_zhanjiang":              "湛江",
	"guangdong_yangjiang":              "阳江",
	"guangdong_chaozhou":               "潮州",
	"guangdong_yunfu":                  "云浮",
	"guangdong_shanwei":                "汕尾",
	"guangdong_chaoyang":               "潮阳",
	"guangdong_zhaoqing":               "肇庆",
	"guangdong_shunde":                 "顺德",
	"guangdong_qingyuan":               "清远",
	"guangxi_nanning":                  "南宁",
	"guangxi_guilin":                   "桂林",
	"guangxi_liuzhou":                  "柳州",
	"guangxi_wuzhou":                   "梧州",
	"guangxi_laibin":                   "来宾",
	"guangxi_guigang":                  "贵港",
	"guangxi_yulin":                    "玉林",
	"guangxi_hezhou":                   "贺州",
	"hainan_haikou":                    "海口",
	"hainan_sanya":                     "三亚",
	"zhongqing_yuzhong":                "渝中",
	"zhongqing_dadukou":                "大渡口",
	"zhongqing_jiangbei":               "江北",
	"zhongqing_shapingba":              "沙坪坝",
	"zhongqing_jiulongpo":              "九龙坡",
	"zhongqing_nanan":                  "南岸",
	"zhongqing_beibei":                 "北碚",
	"zhongqing_wansheng":               "万盛",
	"zhongqing_shuangqiao":             "双桥",
	"zhongqing_yubei":                  "渝北",
	"zhongqing_banan":                  "巴南",
	"zhongqing_wanzhou":                "万州",
	"zhongqing_fuling":                 "涪陵",
	"zhongqing_qianjiang":              "黔江",
	"zhongqing_changshou":              "长寿",
	"sichuan_chengdou":                 "成都",
	"sichuan_dazhou":                   "达州",
	"sichuan_nanchong":                 "南充",
	"sichuan_leshan":                   "乐山",
	"sichuan_mianyang":                 "绵阳",
	"sichuan_deyang":                   "德阳",
	"sichuan_najiang":                  "内江",
	"sichuan_suining":                  "遂宁",
	"sichuan_yibin":                    "宜宾",
	"sichuan_bazhong":                  "巴中",
	"sichuan_zigong":                   "自贡",
	"sichuan_kangding":                 "康定",
	"sichuan_panzhihua":                "攀枝花",
	"guizhou_guiyang":                  "贵阳",
	"guizhou_zunyi":                    "遵义",
	"guizhou_anshun":                   "安顺",
	"guizhou_qianxinan":                "黔西南",
	"guizhou_douyun":                   "都匀",
	"yunnan_kunming":                   "昆明",
	"yunnan_lijiang":                   "丽江",
	"yunnan_zhaotong":                  "昭通",
	"yunnan_yuxi":                      "玉溪",
	"yunnan_lincang":                   "临沧",
	"yunnan_wenshan":                   "文山",
	"yunnan_honghe":                    "红河",
	"yunnan_chuxiong":                  "楚雄",
	"yunnan_dali":                      "大理",
	"xicang_lasa":                      "拉萨",
	"xicang_linzhi":                    "林芝",
	"xicang_rikaze":                    "日喀则",
	"xicang_changdou":                  "昌都",
	"shanxi2_xian":                     "西安",
	"shanxi2_xianyang":                 "咸阳",
	"shanxi2_yanan":                    "延安",
	"shanxi2_hanzhong":                 "汉中",
	"shanxi2_yulin":                    "榆林",
	"shanxi2_shangnan":                 "商南",
	"shanxi2_lueyang":                  "略阳",
	"shanxi2_yijun":                    "宜君",
	"shanxi2_linyou":                   "麟游",
	"shanxi2_baihe":                    "白河",
	"gansu_lanzhou":                    "兰州",
	"gansu_jinchang":                   "金昌",
	"gansu_tianshui":                   "天水",
	"gansu_wuwei":                      "武威",
	"gansu_zhangye":                    "张掖",
	"gansu_pingliang":                  "平凉",
	"gansu_jiuquan":                    "酒泉",
	"qinghai_huangnan":                 "黄南",
	"qinghai_hainan":                   "海南",
	"qinghai_xining":                   "西宁",
	"qinghai_haidong":                  "海东",
	"qinghai_haixi":                    "海西",
	"qinghai_haibei":                   "海北",
	"qinghai_guoluo":                   "果洛",
	"qinghai_yushu":                    "玉树",
	"ningxia_yinchuan":                 "银川",
	"ningxia_wuzhong":                  "吴忠",
	"xinjiang_wulumuqi":                "乌鲁木齐",
	"xinjiang_hami":                    "哈密",
	"xinjiang_kashi":                   "喀什",
	"xinjiang_bayinguoleng":            "巴音郭楞",
	"xinjiang_changji":                 "昌吉",
	"xinjiang_yili":                    "伊犁",
	"xinjiang_aletai":                  "阿勒泰",
	"xinjiang_kelamayi":                "克拉玛依",
	"xinjiang_boertala":                "博尔塔拉",
	"xianggang_zhongxiqu":              "中西区",
	"xianggang_wanziqu":                "湾仔区",
	"xianggang_dongqu":                 "东区",
	"xianggang_nanqu":                  "南区",
	"xianggang_jiulong-youjianwangqu":  "九龙-油尖旺区",
	"xianggang_jiulong-shenshuibuqu":   "九龙-深水埗区",
	"xianggang_jiulong-jiulongchengqu": "九龙-九龙城区",
	"xianggang_jiulong-huangdaxianqu":  "九龙-黄大仙区",
	"xianggang_jiulong-guantangqu":     "九龙-观塘区",
	"xianggang_xinjie-beiqu":           "新界-北区",
	"xianggang_xinjie-dapuqu":          "新界-大埔区",
	"xianggang_xinjie-shatianqu":       "新界-沙田区",
	"xianggang_xinjie-xigongqu":        "新界-西贡区",
	"xianggang_xinjie-quanwanqu":       "新界-荃湾区",
	"xianggang_xinjie-tunmenqu":        "新界-屯门区",
	"xianggang_xinjie-yuanlangqu":      "新界-元朗区",
	"xianggang_xinjie-kuiqingqu":       "新界-葵青区",
	"xianggang_xinjie-lidaoqu":         "新界-离岛区",
	"aomen_huadimatangqu":              "花地玛堂区",
	"aomen_shenganduonitangqu":         "圣安多尼堂区",
	"aomen_datangqu":                   "大堂区",
	"aomen_wangdetangqu":               "望德堂区",
	"aomen_fengshuntangqu":             "风顺堂区",
	"aomen_jiamotangqu":                "嘉模堂区",
	"aomen_shengfangjigetangqu":        "圣方济各堂区",
	"aomen_ludangcheng":                "路氹城",
}

var country_map = map[string]string{
	"afghanistan":              "阿富汗",
	"aland islands":            "奥兰群岛",
	"albania":                  "阿尔巴尼亚",
	"algeria":                  "阿尔及利亚",
	"american samoa":           "美属萨摩亚",
	"andorra":                  "安道尔",
	"angola":                   "安哥拉",
	"anguilla":                 "安圭拉",
	"antigua and barbuda":      "安提瓜和巴布达",
	"argentina":                "阿根廷",
	"armenia":                  "亚美尼亚",
	"aruba":                    "阿鲁巴",
	"australia":                "澳大利亚",
	"austria":                  "奥地利",
	"azerbaijan":               "阿塞拜疆",
	"bangladesh":               "孟加拉",
	"bahrain":                  "巴林",
	"bahamas":                  "巴哈马",
	"barbados":                 "巴巴多斯",
	"belarus":                  "白俄罗斯",
	"belgium":                  "比利时",
	"belize":                   "伯利兹",
	"benin":                    "贝宁",
	"bermuda":                  "百慕大",
	"bhutan":                   "不丹",
	"bolivia":                  "玻利维亚",
	"bosnia and herzegovina":   "波斯尼亚和黑塞哥维那",
	"botswana":                 "博茨瓦纳",
	"bouvet island":            "布维岛",
	"brazil":                   "巴西",
	"brunei":                   "文莱",
	"bulgaria":                 "保加利亚",
	"burkina faso":             "布基纳法索",
	"burundi":                  "布隆迪",
	"cambodia":                 "柬埔寨",
	"cameroon":                 "喀麦隆",
	"canada":                   "加拿大",
	"cape verde":               "佛得角",
	"central african republic": "中非",
	"chad":                             "乍得",
	"chile":                            "智利",
	"christmas islands":                "圣诞岛",
	"cocos (keeling) islands":          "科科斯（基林）群岛",
	"colombia":                         "哥伦比亚",
	"comoros":                          "科摩罗",
	"congo (congo-kinshasa)":           "刚果（金）",
	"congo":                            "刚果",
	"cook islands":                     "库克群岛",
	"costa rica":                       "哥斯达黎加",
	"cote d'ivoire":                    "科特迪瓦",
	"china":                            "中国",
	"croatia":                          "克罗地亚",
	"cuba":                             "古巴",
	"czech":                            "捷克",
	"cyprus":                           "塞浦路斯",
	"denmark":                          "丹麦",
	"djibouti":                         "吉布提",
	"dominica":                         "多米尼加",
	"east timor":                       "东帝汶",
	"ecuador":                          "厄瓜多尔",
	"egypt":                            "埃及",
	"equatorial guinea":                "赤道几内亚",
	"eritrea":                          "厄立特里亚",
	"estonia":                          "爱沙尼亚",
	"ethiopia":                         "埃塞俄比亚",
	"faroe islands":                    "法罗群岛",
	"fiji":                             "斐济",
	"finland":                          "芬兰",
	"france":                           "法国",
	"metropolitanfrance":               "法国大都会",
	"french guiana":                    "法属圭亚那",
	"french polynesia":                 "法属波利尼西亚",
	"gabon":                            "加蓬",
	"gambia":                           "冈比亚",
	"georgia":                          "格鲁吉亚",
	"germany":                          "德国",
	"ghana":                            "加纳",
	"gibraltar":                        "直布罗陀",
	"greece":                           "希腊",
	"grenada":                          "格林纳达",
	"guadeloupe":                       "瓜德罗普岛",
	"guam":                             "关岛",
	"guatemala":                        "危地马拉",
	"guernsey":                         "根西岛",
	"guinea-bissau":                    "几内亚比绍",
	"guinea":                           "几内亚",
	"guyana":                           "圭亚那",
	"haiti":                            "海地",
	"honduras":                         "洪都拉斯",
	"hungary":                          "匈牙利",
	"iceland":                          "冰岛",
	"india":                            "印度",
	"indonesia":                        "印度尼西亚",
	"iran":                             "伊朗",
	"iraq":                             "伊拉克",
	"ireland":                          "爱尔兰",
	"isle of man":                      "马恩岛",
	"israel":                           "以色列",
	"italy":                            "意大利",
	"jamaica":                          "牙买加",
	"japan":                            "日本",
	"jersey":                           "泽西岛",
	"jordan":                           "约旦",
	"kazakhstan":                       "哈萨克斯坦",
	"kenya":                            "肯尼亚",
	"kiribati":                         "基里巴斯",
	"korea (south)":                    "韩国",
	"korea (north)":                    "朝鲜",
	"kuwait":                           "科威特",
	"kyrgyzstan":                       "吉尔吉斯斯坦",
	"laos":                             "老挝",
	"latvia":                           "拉脱维亚",
	"lebanon":                          "黎巴嫩",
	"lesotho":                          "莱索托",
	"liberia":                          "利比里亚",
	"libya":                            "利比亚",
	"liechtenstein":                    "列支敦士登",
	"lithuania":                        "立陶宛",
	"luxembourg":                       "卢森堡",
	"macedonia":                        "马其顿",
	"malawi":                           "马拉维",
	"malaysia":                         "马来西亚",
	"madagascar":                       "马达加斯加",
	"maldives":                         "马尔代夫",
	"mali":                             "马里",
	"malta":                            "马耳他",
	"marshall islands":                 "马绍尔群岛",
	"martinique":                       "马提尼克岛",
	"mauritania":                       "毛里塔尼亚",
	"mauritius":                        "毛里求斯",
	"mayotte":                          "马约特",
	"mexico":                           "墨西哥",
	"micronesia":                       "密克罗尼西亚",
	"moldova":                          "摩尔多瓦",
	"monaco":                           "摩纳哥",
	"mongolia":                         "蒙古",
	"montenegro":                       "黑山",
	"montserrat":                       "蒙特塞拉特",
	"morocco":                          "摩洛哥",
	"mozambique":                       "莫桑比克",
	"myanmar":                          "缅甸",
	"namibia":                          "纳米比亚",
	"nauru":                            "瑙鲁",
	"nepal":                            "尼泊尔",
	"netherlands":                      "荷兰",
	"new caledonia":                    "新喀里多尼亚",
	"new zealand":                      "新西兰",
	"nicaragua":                        "尼加拉瓜",
	"niger":                            "尼日尔",
	"nigeria":                          "尼日利亚",
	"niue":                             "纽埃",
	"norfolk island":                   "诺福克岛",
	"norway":                           "挪威",
	"oman":                             "阿曼",
	"pakistan":                         "巴基斯坦",
	"palau":                            "帕劳",
	"palestine":                        "巴勒斯坦",
	"panama":                           "巴拿马",
	"papua new guinea":                 "巴布亚新几内亚",
	"peru":                             "秘鲁",
	"philippines":                      "菲律宾",
	"pitcairn islands":                 "皮特凯恩群岛",
	"poland":                           "波兰",
	"portugal":                         "葡萄牙",
	"puerto rico":                      "波多黎各",
	"qatar":                            "卡塔尔",
	"reunion":                          "留尼汪岛",
	"romania":                          "罗马尼亚",
	"rwanda":                           "卢旺达",
	"russian federation":               "俄罗斯联邦",
	"saint helena":                     "圣赫勒拿",
	"saint kitts-nevis":                "圣基茨和尼维斯",
	"saint lucia":                      "圣卢西亚",
	"saint vincent and the grenadines": "圣文森特和格林纳丁斯",
	"el salvador":                      "萨尔瓦多",
	"samoa":                            "萨摩亚",
	"san marino":                       "圣马力诺",
	"sao tome and principe":            "圣多美和普林西比",
	"saudi arabia":                     "沙特阿拉伯",
	"senegal":                          "塞内加尔",
	"seychelles":                       "塞舌尔",
	"sierra leone":                     "塞拉利昂",
	"singapore":                        "新加坡",
	"serbia":                           "塞尔维亚",
	"slovakia":                         "斯洛伐克",
	"slovenia":                         "斯洛文尼亚",
	"solomon islands":                  "所罗门群岛",
	"somalia":                          "索马里",
	"south africa":                     "南非",
	"spain":                            "西班牙",
	"sri lanka":                        "斯里兰卡",
	"sudan":                            "苏丹",
	"suriname":                         "苏里南",
	"swaziland":                        "斯威士兰",
	"sweden":                           "瑞典",
	"switzerland":                      "瑞士",
	"syria":                            "叙利亚",
	"tajikistan":                       "塔吉克斯坦",
	"tanzania":                         "坦桑尼亚",
	"thailand":                         "泰国",
	"trinidad and tobago":              "特立尼达和多巴哥",
	"timor-leste":                      "东帝汶",
	"togo":                             "多哥",
	"tokelau":                          "托克劳",
	"tonga":                            "汤加",
	"tunisia":                          "突尼斯",
	"turkey":                           "土耳其",
	"turkmenistan":                     "土库曼斯坦",
	"tuvalu":                           "图瓦卢",
	"uganda":                           "乌干达",
	"ukraine":                          "乌克兰",
	"united arab emirates":             "阿拉伯联合酋长国",
	"united kingdom":                   "英国",
	"united states":                    "美国",
	"uruguay":                          "乌拉圭",
	"uzbekistan":                       "乌兹别克斯坦",
	"vanuatu":                          "瓦努阿图",
	"vatican city":                     "梵蒂冈",
	"venezuela":                        "委内瑞拉",
	"vietnam":                          "越南",
	"wallis and futuna":                "瓦利斯群岛和富图纳群岛",
	"western sahara":                   "西撒哈拉",
	"yemen":                            "也门",
	"yugoslavia":                       "南斯拉夫",
	"zambia":                           "赞比亚",
	"zimbabwe":                         "津巴布韦",
}
