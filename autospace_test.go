package autospace

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_strategies(t *testing.T) {
	assert.Equal(t, true, len(strategies) > 2)
}

func TestFormat(t *testing.T) {
	cases := map[string]string{
		"部署到heroku有问题网页不能显示":           "部署到 heroku 有问题网页不能显示",
		"[北京]美企聘site/web大型应用开发高手-Ruby": "[北京] 美企聘 site/web 大型应用开发高手-Ruby",
		"[成都](团800)招聘Rails工程师":         "[成都](团 800) 招聘 Rails 工程师",
		"Teahour.fm第18期发布":             "Teahour.fm 第 18 期发布",
		"Yes!升级到了Rails 4":              "Yes! 升级到了 Rails 4",
	}

	for source, exptected := range cases {
		actual := Format(source)
		if exptected != actual {
			t.Errorf("\nexptected: %s\nactual   : %s", exptected, actual)
		}
	}
}
