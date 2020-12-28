package internal

import (
	"github.com/kami-zh/go-capturer"
	. "github.com/smartystreets/goconvey/convey"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {

	Convey("Test print result in interface data when verbos is false", t, func() {
		data := []interface{}{
			"Email 1 successfully sent",
			"Email 2 successfully sent",
			"Email 3 successfully sent",
			"Email 4 failed to sent",
		}
		opt := new(Option)
		opt.Verbose = false
		out := capturer.CaptureStdout(func() {
			opt.print(data)
		})
		So(out, ShouldEqual, `Successfully sent all email
`)
	})

	Convey("Test print result in interface data when verbos is true", t, func() {
		data := []interface{}{
			"Email 1 successfully sent",
			"Email 2 successfully sent",
			"Email 3 successfully sent",
			"Email 4 failed to sent",
		}
		opt := new(Option)
		opt.Verbose = true
		opt.Count = 4
		out := capturer.CaptureStdout(func() {
			opt.print(data)
		})
		So(len(strings.Split(out, "\n")), ShouldEqual, len(data)+1)
	})
}
