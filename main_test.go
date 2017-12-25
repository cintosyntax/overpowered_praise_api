package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrue(t *testing.T) {
	// Fake tests that should always be true
	Convey("Given some integer with a starting value", t, func() {
		x := 1

		Convey("When the integer is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 2)
			})
		})
	})

}
