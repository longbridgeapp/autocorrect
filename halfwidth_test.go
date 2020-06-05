package autocorrect

import (
	"fmt"
	"testing"
)

func Test_halfWidth(t *testing.T) {
	source := "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ１２３４５６７８９０："
	fmt.Println(halfWidth(source))
	// Example:
	// abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890:
}
