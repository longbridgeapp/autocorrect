package autocorrect

import (
	"testing"

	"github.com/longbridgeapp/assert"
)

func Test_haftwidth(t *testing.T) {
	source := "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ１２３４５６７８９０"
	assert.Equal(t, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", haftwidth(source))
	assert.Equal(t, "他说：我们将在16:32分出发去CBD中心。", haftwidth("他说：我们将在１６：３２分出发去ＣＢＤ中心。"))
	// Fullwidth space
	assert.Equal(t, "ジョイフル－後場売り気配 200 店舗を閉鎖へ 7 月以降、不採算店中心に", haftwidth("ジョイフル－後場売り気配　200　店舗を閉鎖へ　7 月以降、不採算店中心に"))
}

func Benchmark_haftwidth(b *testing.B) {
	source := "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ１２３４５６７８９０"
	for i := 0; i < b.N; i++ {
		// about 0.003ms/op
		haftwidth(source)
	}
}
