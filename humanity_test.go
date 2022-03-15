package humanity

import (
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_BytesFloat64(t *testing.T) {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}

	Convey("Looping through increasingly larger number, the suffixes are accurate, as are the numbers", t, func() {
		var c float64
		for i := 0; i <= 8; i++ { // All 8 because floats scale
			if i == 1 {
				c = 1.0
			}
			c = c * 1024

			// Confirm the suffix
			f := ByteFormat(c)

			//Printf("\t%f %s\n", c, f)
			So(f, ShouldEndWith, units[i])

			// Confirm the digits
			f = strings.TrimSuffix(f, units[i])
			fn, e := strconv.ParseFloat(f, 64)
			So(e, ShouldBeNil)
			if c == 0 {
				So(fn, ShouldAlmostEqual, 0)
			} else {
				So(fn, ShouldAlmostEqual, 1)
			}
		}

	})
}

func Test_BytesInt(t *testing.T) {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}

	Convey("Looping through increasingly larger number, the suffixes are accurate, as are the numbers", t, func() {
		var c int
		for i := 0; i <= 3; i++ { // only the first 4 because possible overflow
			if i == 1 {
				c = 1.0
			}
			c = c * 1024

			// Confirm the suffix
			f := ByteFormat(c)

			//Printf("\t%f %s\n", c, f)
			So(f, ShouldEndWith, units[i])

			// Confirm the digits
			f = strings.TrimSuffix(f, units[i])
			fn, e := strconv.ParseFloat(f, 64)
			So(e, ShouldBeNil)
			if c == 0 {
				So(fn, ShouldAlmostEqual, 0)
			} else {
				So(fn, ShouldAlmostEqual, 1)
			}
		}

	})
}

func Test_BytesInt64(t *testing.T) {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}

	Convey("Looping through increasingly larger number, the suffixes are accurate, as are the numbers", t, func() {
		var c int64
		for i := 0; i <= 6; i++ { // Only the first 7 because overflow
			if i == 1 {
				c = 1.0
			}
			c = c * 1024

			// Confirm the suffix
			f := ByteFormat(c)

			//Printf("\t%f %s\n", c, f)
			So(f, ShouldEndWith, units[i])

			// Confirm the digits
			f = strings.TrimSuffix(f, units[i])
			fn, e := strconv.ParseFloat(f, 64)
			So(e, ShouldBeNil)
			if c == 0 {
				So(fn, ShouldAlmostEqual, 0)
			} else {
				So(fn, ShouldAlmostEqual, 1)
			}
		}

	})
}

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
