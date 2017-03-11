package main

import (
	t "github.com/nsf/termbox-go"
	"math/rand"
	"os"
	"time"
)

func remove(array [][2]int, item [2]int) [][2]int {
	for i := 0; i < len(array); i++ {
		if array[i] == item {
			return append(array[:i], array[i+1:]...)
		}
	}
	return array
}
func main() {
	t.Init()
	t.SetOutputMode(t.Output256)
	var w, h = t.Size()
	var widthVals = make([]int, w)
	var heightVals = make([]int, h)
	var allCords = make([][2]int, h*w)
	for i := 0; i < w; i++ {
		widthVals = append(widthVals, i)
	}
	for i := 0; i < h; i++ {
		heightVals = append(heightVals, i)
	}
	for i := 0; i < h; i++ {
		for n := 0; n < w; n++ {
			allCords = append(allCords, [2]int{n, i})
		}
	}

	go func() {
		for {
			var event = t.PollEvent()
			if event.Key == t.KeyCtrlC {
				t.Close()
				os.Exit(3)
			}
		}
	}()
	for i := 0; i < w*h; i++ {
		rand.Seed(time.Now().UnixNano())
		var currentCord = allCords[rand.Intn(len(allCords))]
		var randColorNum = rand.Intn(265)
		var randColor = t.Attribute(randColorNum)
		t.SetCell(currentCord[0], currentCord[1], ' ', randColor, randColor)
		allCords = remove(allCords, currentCord)
		t.Flush()
		time.Sleep(1 * time.Millisecond)
	}
	for {
	}
}
