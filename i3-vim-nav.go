package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	xdo "github.com/aep/xdo-go"
	"github.com/proxypoke/i3ipc"
)

func main() {
	dir := string(os.Args[1])

	if match, _ := regexp.MatchString(`left|down|up|right`, dir); !match {
		fmt.Println("must have an argument left, down, up, or right")
		return
	}

	xdot := xdo.NewXdo()
	window := xdot.GetActiveWindow()
	name := strings.ToLower(window.GetName())

	r, _ := regexp.Compile(`\bn?vim?$`)

	if r.MatchString(name) {
		keycmd := exec.Command("xdotool", "key", "--clearmodifiers", "Ctrl+"+strings.Title(dir))
		out, _ := keycmd.Output()
		if len(out) > 0 {
			fmt.Println(out)
		}
	} else {
		conn, err := i3ipc.GetIPCSocket()
		if err != nil {
			fmt.Println("could not connect to i3")
			return
		}
		conn.Command("focus " + dir)
	}
}
