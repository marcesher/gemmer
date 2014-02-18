//http://golang.org/pkg/os/exec/#Command
//https://gobyexample.com/spawning-processes
//cd gemfilemaker && go run gemfilemaker.go

package main

import (
	"io/ioutil"
	"os/exec"
	"strings"
)

func ListGems() string {
	cmd := exec.Command("gem", "list", "--remote")
	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(out))
}

func GemsTo(filepath string) {
	g := ListGems()
	var gems []string
	for _, item := range strings.Split(g, "\n") {
		gems = append(gems, strings.Split(item, " ")[0])
	}
	err := ioutil.WriteFile(filepath, []byte(strings.Join(gems, "\n")), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	GemsTo("../gemlist.txt")
}
