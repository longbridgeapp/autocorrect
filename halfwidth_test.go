package autocorrect

import (
	"testing"
)

func Test_halfWidth(t *testing.T) {
	source := "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ１２３４５６７８９０"
	assertEqual(t, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", halfWidth(source))
	assertEqual(t, "他说：我们将在16:32分出发去CBD中心。", halfWidth("他说：我们将在１６：３２分出发去ＣＢＤ中心。"))
}

func Benchmark_halfWidth(b *testing.B) {
	source := "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ１２３４５６７８９０"
	for i := 0; i < b.N; i++ {
		// about 0.003ms/op
		halfWidth(source)
	}
}
