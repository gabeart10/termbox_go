package main

import (
	t "github.com/nsf/termbox-go"
	"os"
	"time"
)

func main() {
	var colors = [7]t.Attribute{1, 2, 3, 4, 5, 6, 7}
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
	t.SetOutputMode(t.Output256)
	var w, h = t.Size()
	var set_start = 0
	for {
		for i := 0; i < h; i++ {
			go func(start int, width int, height int, colors [7]t.Attribute) {
				currentColor := start
				for n := 0; n < width; n++ {
					t.SetCell(width, height, ' ', colors[currentColor], colors[currentColor])
					if currentColor == 6 {
						currentColor = 0
					} else {
						currentColor++
					}
				}
				t.Flush()
			}(set_start, w, h, colors)
			time.Sleep(10 * time.Millisecond)
		}
		if set_start == 6 {
			set_start = 0
		} else {
			set_start++
		}
	}
}
