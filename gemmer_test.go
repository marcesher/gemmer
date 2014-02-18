package main

import (
	"fmt"
	//	"io/ioutil"
	"reflect"
	"strings"
	"testing"
)

func TestCoverage(t *testing.T) {
	lyrics := strings.Fields("never gonna give you up never gonna let you down")
	gems := strings.Fields("blurp never three yard-amp yapper up")
	exp := strings.Fields("gonna give you let down")
	diff, intersect, cov := Coverage(lyrics, gems)
	if !reflect.DeepEqual(exp, diff) {
		t.Errorf("Expected %v but received %v.", exp, diff)
	}
	if cov < 28 || cov > 29 {
		t.Errorf("Expected %v but received %v", "between 28 and 29", cov)
	}
	fmt.Printf("Intersect is %v", intersect)
}

//func TestCoverageWithGemFile(t *testing.T) {
//	lyrics := strings.Fields("never gonna give you up never gonna let you down never gonna run around and desert you never gonna make you cry never gonna say goodbye never gonna tell a lie and hurt you")
//	f, err := ioutil.ReadFile("gemlist.txt")
//	if err != nil {
//		panic(err)
//	}
//
//	gems := strings.Fields(string(f))
//	exp := NewSetFromStrings([]string{"never", "gonna", "and", "you", "down", "lie", "make", "cry", "hurt"})
//	diff, cov := Coverage(lyrics, gems)
//	fmt.Println(diff)
//	if !diff.Equal(exp) {
//		t.Errorf("Expected %v but received %v. \n\nDiff: ", exp, diff, exp.SymmetricDifference(diff))
//	}
//	if cov < 50 || cov > 60 {
//		t.Errorf("Expected cov to be between 50 and 60 but was %v", cov)
//	}
//}

func TestPrepareText(t *testing.T) {
	text := "Don't, Don't, Don't you... forget about me!"
	exp := "dont dont dont you forget about me"
	clean := PrepareText(text)
	if clean != exp {
		t.Errorf("Expected %v but received %v", exp, clean)
	}
}
