package main

import (
	t "github.com/nsf/termbox-go"
	"math/rand"
	"os"
	"time"
)

func remove(array []int, item int) []int {
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
	for i := 0; i < w; i++ {
		widthVals = append(widthVals, i)
	}
	for i := 0; i < h; i++ {
		heightVals = append(heightVals, i)
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
		var currentW = widthVals[rand.Intn(len(widthVals))]
		var currentH = heightVals[rand.Intn(len(heightVals))]
		var randColorNum = rand.Intn(265)
		var randColor = t.Attribute(randColorNum)
		t.SetCell(currentW, currentH, ' ', randColor, randColor)
		widthVals = remove(widthVals, currentW)
		heightVals = remove(heightVals, currentH)
	}
}
