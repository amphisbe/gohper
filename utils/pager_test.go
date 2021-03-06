package utils

import (
	"testing"

	"github.com/cosiner/gohper/testing2"
)

func TestPageStart(t *testing.T) {
	pg := Pager{
		BeginPage:  1,
		BeginIndex: 0,
		PageSize:   10,
	}
	testing2.
		Expect(0).Arg("abcde").
		Expect(0).Arg("-1").
		Expect(0).Arg("0").
		Expect(0).Arg("1").
		Expect(10).Arg("2").
		Expect(20).Arg("3").
		Run(t, pg.BeginByString)
}
