package extend_test

import (
	"onion-architecrure-go/extend"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHelperExtend_Hash(t *testing.T) {
	Convey("md5 hash", t, func() {
		str := "fiiewl"
		testStr := extend.Helper.Hash(str)
		hashStr := "1e1543e12484ce0db6275fe1d2948ef2"
		So(testStr, ShouldEqual, hashStr)
	})
}

func TestHelperExtend_FormatToTimeString(t *testing.T) {
	Convey("date tinme string -> YYYY-MM-DD HH:mm:ss", t, func() {
		str := "2024-02-05 15:17:32"
		dateTime, _ := time.Parse("2006-01-02 15:04:05", str)
		testStr := extend.Helper.FormatToTimeString(dateTime)
		So(testStr, ShouldEqual, str)
	})
}
