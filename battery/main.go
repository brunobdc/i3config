package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	status, err := ioutil.ReadFile("/sys/class/power_supply/BAT0/status")
	check(err)

	dat, err := ioutil.ReadFile("/sys/class/power_supply/BAT0/charge_full")
	check(err)

	chargedfull, err := strconv.Atoi(strings.TrimSpace(string(dat)))
	check(err)

	dat, err = ioutil.ReadFile("/sys/class/power_supply/BAT0/charge_now")
	check(err)

	chargeNow, err := strconv.Atoi(strings.TrimSpace(string(dat)))
	check(err)

	charged := (chargeNow * 100) / chargedfull

	label := ""
	color := ""

	setLabel := func(newLabel string) {
		if strings.TrimSpace(string(status)) == "Charging" {
			label = ""
		} else {
			label = newLabel
		}
	}

	switch {
	case charged <= 10:
		setLabel("")
		color = "#FF5555"
	case charged <= 30:
		setLabel("")
		color = "#ffd555"
	case charged <= 60:
		setLabel("")
		color = "#ffff55"
	case charged <= 90:
		setLabel("")
		color = "#ccff55"
	default:
		setLabel("")
		color = "#99ff55"
	}

	fulltext := fmt.Sprintf("<span foreground=\"%s\">%s</span> %v%%", color, label, charged)
	shortext := fmt.Sprintf("<span foreground=\"%s\">%s</span>", color, label)

	fmt.Printf("%s\n%s\n", fulltext, shortext)
}
