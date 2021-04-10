package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func main() {

	out, err := exec.Command("amixer", "sget", "Master").Output()
	check(err)

	volume := 0
	muted := false

	for _, v := range strings.Split(string(out), "\n") {
		if strings.Contains(v, "Mono: Playback") {

			if strings.Contains(v, "[off]") {
				//muted = true
			} else {
				fi := strings.Index(v, "[") + 1
				li := strings.Index(v, "%")
				if fi == (li - 1) {
					volume, err = strconv.Atoi(string(v[fi]))
					check(err)
				} else {
					volume, err = strconv.Atoi(v[fi:li])
					check(err)
				}
			}
		}
	}

	label := ""

	switch {
	case muted:
		label = ""
	case volume == 0:
		label = ""
	case volume <= 50:
		label = ""
	default:
		label = ""
	}

	fulltext := fmt.Sprintf("<span foreground=\"#55ffee\">%s</span> %v", label, volume)
	shorttext := fmt.Sprintf("<span foreground=\"#55ffee\">%s</span>", label)

	fmt.Printf("%s\n%s\n", fulltext, shorttext)
}
