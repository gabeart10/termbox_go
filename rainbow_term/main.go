package main

import (
	t "github.com/nsf/termbox-go"
	"os"
	"time"
)

func main() {
	var colors = [7]t.Attribute{1, 2, 3, 4, 5, 6, 7}
	var speedChan = make(chan time.Duration, 1)
	go func(defSpeed time.Duration) {
		var speed = defSpeed
		for {
			var exitEvent = t.PollEvent()
			switch exitEvent.Key {
			case t.KeyCtrlC:
				t.Close()
				os.Exit(3)
			case t.KeyArrowLeft:
				if speed != 15 {
					speed++
					speedChan <- speed
				}
			case t.KeyArrowRight:
				if speed != 1 {
					speed--
					speedChan <- speed
				}
			}
		}
	}(5)
	t.Init()
	t.SetOutputMode(t.Output256)
	var w, h = t.Size()
	var set_start = 0
	var speed time.Duration = 5
	for {
		for i := 0; i < h; i++ {
			select {
			case speed = <-speedChan:
			default:
			}
			go func(start, width, height int, colors [7]t.Attribute) {
				currentColor := start
				for n := 0; n < width; n++ {
					t.SetCell(n, height, ' ', colors[currentColor], colors[currentColor])
					if currentColor == 6 {
						currentColor = 0
					} else {
						currentColor++
					}
				}
			}(set_start, w, i, colors)
			time.Sleep(speed * time.Millisecond)
		}
		if set_start == 6 {
			set_start = 0
		} else {
			set_start++
		}

		t.Flush()
	}
}
