package main

import (
	t "github.com/nsf/termbox-go"
	"os"
	"time"
)

func main() {
	var colors = [7]t.Attribute{1, 2, 3, 4, 5, 6, 7}
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
	t.SetOutputMode(t.Output256)
	for {
		var h, w = t.Size()
		for i := 0; i < h; i++ {
			for n := 0; n < w; n++ {
				t.SetCell(i, n, ' ', colors[current_color], colors[current_color])
				if current_color == 6 {
					current_color = 0
				}
				current_color++
				t.Flush()
				time.Sleep(1 * time.Millisecond)
			}
		}
	}
}
