package autospace

import (
	"testing"
)

func assertCases(t *testing.T, cases map[string]string) {
	for source, exptected := range cases {
		actual := Format(source)
		if exptected != actual {
			t.Errorf("\nexptected: %s\nactual   : %s", exptected, actual)
		}
	}
}

func TestFormat(t *testing.T) {
	cases := map[string]string{
		"部署到heroku有问题网页不能显示":                                                             "部署到 heroku 有问题网页不能显示",
		"[北京]美企聘site/web大型应用开发高手-Ruby":                                                   "[北京] 美企聘 site/web 大型应用开发高手-Ruby",
		"[成都](团800)招聘Rails工程师":                                                           "[成都](团 800) 招聘 Rails 工程师",
		"Teahour.fm第18期发布":                                                               "Teahour.fm 第 18 期发布",
		"Yes!升级到了Rails 4":                                                                "Yes! 升级到了 Rails 4",
		"记事本,记事本显示阅读次数#149":                                                              "记事本,记事本显示阅读次数 #149",
		"里面用@foo符号的话后面的变量名会被替换成userN":                                                    "里面用 @foo 符号的话后面的变量名会被替换成 userN",
		"WWDC上讲到的Objective C/LLVM改进":                                                     "WWDC 上讲到的 Objective C/LLVM 改进",
		"在Ubuntu11.10 64位系统安装newrelic出错":                                                 "在 Ubuntu11.10 64 位系统安装 newrelic 出错",
		"升级了10.9 附遇到的bug":                                                                "升级了 10.9 附遇到的 bug",
		"在做ROR 3.2 Tutorial第Chapter 9.4.2遇到一个问题求助！":                                      "在做 ROR 3.2 Tutorial 第 Chapter 9.4.2 遇到一个问题求助！",
		"Mac安装软件新方法：Homebrew-cask":                                                       "Mac 安装软件新方法：Homebrew-cask",
		"without looking like it’s been marked up with tags or formatting instructions.": "without looking like it’s been marked up with tags or formatting instructions.",
	}
	assertCases(t, cases)
}

func TestFormatForDate(t *testing.T) {
	cases := map[string]string{
		"于3月10日开始": "于 3 月 10 日开始",
		"于3月开始":    "于 3 月开始",
		"于2009年开始": "于 2009 年开始",
		"2013年3月10日-Ruby Saturday活动召集": "2013 年 3 月 10 日-Ruby Saturday 活动召集",
		"2013年12月22号开始出发":              "2013 年 12 月 22 号开始出发",
		"12月22号开始出发":                   "12 月 22 号开始出发",
		"22号开始出发":                      "22 号开始出发",
	}
	assertCases(t, cases)
}

func TestFormatForEnglishLetter(t *testing.T) {
	cases := map[string]string{
		"长桥LongBridge App下载": "长桥 LongBridge App 下载",
	}
	assertCases(t, cases)
}

func TestFormatForNumber(t *testing.T) {
	cases := map[string]string{
		"在Ubuntu 11.10 64位系统安装Go出错": "在 Ubuntu 11.10 64 位系统安装 Go 出错",
		"喜欢暗黑2却对 D3不满意的可以看看这个。":     "喜欢暗黑 2 却对 D3 不满意的可以看看这个。",
		"Ruby 2.7版本第3次发布":           "Ruby 2.7 版本第 3 次发布",
	}
	assertCases(t, cases)
}

func TestFormatForSpecialSymbols(t *testing.T) {
	cases := map[string]string{
		"公告:(美股)阿里巴巴[BABA.US]发布2019下半年财报!":           "公告:(美股) 阿里巴巴 [BABA.US] 发布 2019 下半年财报!",
		"消息http://github.com解禁了":                     "消息 http://github.com 解禁了",
		"美股异动|阿帕奇石油(APA.US)盘前涨超15% 在苏里南近海发现大量石油":     "美股异动 | 阿帕奇石油 (APA.US) 盘前涨超 15% 在苏里南近海发现大量石油",
		"美国统计局：美国11月原油出口下降至302.3万桶/日，10月为338.3万桶/日。": "美国统计局：美国 11 月原油出口下降至 302.3 万桶/日，10 月为 338.3 万桶/日。",
	}
	assertCases(t, cases)
}

func TestFormatForFullwidthSymbols(t *testing.T) {
	cases := map[string]string{
		"（美股）市场：发布「最新」100消息【BABA.US】“大涨”50%；同比上涨20%！": "（美股）市场：发布「最新」100 消息【BABA.US】“大涨” 50%；同比上涨 20%！",
		"第3季度财报发布看涨看跌？敬请期待。":                          "第 3 季度财报发布看涨看跌？敬请期待。",
	}
	assertCases(t, cases)
}
