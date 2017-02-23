package main

import (
	t "github.com/nsf/termbox-go"
	"os"
)

func main() {
	t.SetOutputMode(t.Output256)
	var colors = [7]int{1, 202, 11, 10, 19, 17, 18}
	var current_color = 0
	go func() {
		for {
			var exitEvent = t.PollEvent()
			if exitEvent.Key == t.KeyCtrlC {
				t.Close()
				os.Exit(3)
			}
		}
	}()
	t.Init()
	for {
		var w, h = t.Size()
		for i := 0; i < h; i++ {
			for n := 0; n < w; n++ {
				t.SetCell(i, n, ' ', 0, colors[current_color])
				if current_color == 6 {
					current_color = 0
				}
				current_color++
			}
		}
	}
}
