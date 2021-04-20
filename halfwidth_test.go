package autocorrect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_halfwidth(t *testing.T) {
	source := "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ１２３４５６７８９０"
	assertEqual(t, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", halfwidth(source))
	assertEqual(t, "他说：我们将在16:32分出发去CBD中心。", halfwidth("他说：我们将在１６：３２分出发去ＣＢＤ中心。"))
	// Fullwidth space
	assert.Equal(t, "ジョイフル－後場売り気配 200 店舗を閉鎖へ 7 月以降、不採算店中心に", halfwidth("ジョイフル－後場売り気配　200　店舗を閉鎖へ　7 月以降、不採算店中心に"))
}

func Benchmark_halfwidth(b *testing.B) {
	source := "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ１２３４５６７８９０"
	for i := 0; i < b.N; i++ {
		// about 0.003ms/op
		halfwidth(source)
	}
}
