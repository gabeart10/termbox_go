package main
import (
	"github.com/nsf/termbox-go" t
)

func main() {
	t.SetOutputMode(t.Output256)
	var colors [6]int{1,202,11,10,19,17,18}
	var current_color = 0
	go func() {
		for {
			var exitEvent = t.PollEvent()
			if exitEvent.Key == t.KeyCtrlC {
				t.Close()
				os.Exit(3) for {
		var w, h = t.Size()
		for i := 0;i < w;i++ {
			
