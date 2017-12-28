package humanity

import (
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"strings"
	"testing"
)

func Test_DurationFormatNS(t *testing.T) {
	Convey("Looping through 0-999 the suffixes are accurate, as are the numbers", t, func() {
		var (
			c      int64  // to force type
			offset int64  = 0
			s      string = "ns"
		)

		for c = offset; c < 1000; c++ {
			// Confirm the suffix
			f := DurationFormat(c)
			So(f, ShouldEndWith, s)

			// Confirm the digits
			f = strings.TrimSuffix(f, s)
			fn, e := strconv.ParseFloat(f, 64)
			So(e, ShouldBeNil)
			So(fn, ShouldAlmostEqual, c-offset)
		}
	})
}

func Test_DurationFormatUS(t *testing.T) {
	Convey("Looping through 1000-9999 the suffixes are accurate, as are the numbers", t, func() {
		var (
			c       int64   // to force type
			offset  int64   = 1000
			upset   int64   = 1000 * offset
			offsetf float64 = float64(offset)
			s       string  = "us"
		)

		for c = offset; c < upset; c += offset {
			// Confirm the suffix
			f := DurationFormat(c)
			So(f, ShouldEndWith, s)

			// Confirm the digits
			f = strings.TrimSuffix(f, s)
			fn, e := strconv.ParseFloat(f, 64)
			So(e, ShouldBeNil)
			So(fn*offsetf, ShouldAlmostEqual, c)
		}
	})
}

func Test_DurationFormatMS(t *testing.T) {
	Convey("Looping through 10000-99999 the suffixes are accurate, as are the numbers", t, func() {
		var (
			c       int64   // to force type
			offset  int64   = 1000000
			upset   int64   = 1000 * offset
			offsetf float64 = float64(offset)
			s       string  = "ms"
		)

		for c = offset; c < upset; c += offset {
			// Confirm the suffix
			f := DurationFormat(c)
			So(f, ShouldEndWith, s)

			// Confirm the digits
			f = strings.TrimSuffix(f, s)
			fn, e := strconv.ParseFloat(f, 64)
			So(e, ShouldBeNil)
			So(fn*offsetf, ShouldAlmostEqual, c)
		}
	})
}

func Test_DurationFormatS(t *testing.T) {
	Convey("Looping through 100000-999999 the suffixes are accurate, as are the numbers", t, func() {
		var (
			c       int64   // to force type
			offset  int64   = 1000000000
			upset   int64   = 1000 * offset
			offsetf float64 = float64(offset)
			s       string  = "s"
		)

		for c = offset; c < upset; c += offset {
			// Confirm the suffix
			f := DurationFormat(c)
			So(f, ShouldEndWith, s)

			// Confirm the digits
			f = strings.TrimSuffix(f, s)
			fn, e := strconv.ParseFloat(f, 64)
			So(e, ShouldBeNil)
			So(fn*offsetf, ShouldAlmostEqual, c)
		}
	})
}
