package main

import (
	"strings"
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCoverage(t *testing.T) {
	Convey("Coverage should show diff between lyrics and gem names and calculate percent coverage", t, func() {
		lyrics := strings.Fields("never gonna give you up never gonna let you down")
		gems := strings.Fields("blurp never three yard-amp yapper up")
		exp := strings.Fields("gonna give you let down")
		diff, _, cov := Coverage(lyrics, gems)

		//	Forgive me
		expectedSet := NewSetFromStrings(exp)
		diffSet := NewSetFromStrings(diff)
		testdiff := expectedSet.Difference(diffSet)

		So(len(testdiff), ShouldEqual, 0)
		So(cov, ShouldBeBetween, 28, 29)
	})

	Convey("PrepareText should lowercase and remove all punctuation", t, func() {
		text := "Don't, Don't, Don't you... forget about me!"
		exp := "dont dont dont you forget about me"
		clean := PrepareText(text)
		So(exp, ShouldEqual, clean)
	})

}
