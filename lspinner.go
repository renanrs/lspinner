package lspinner

import (
	"fmt"
	"time"
)

// Spinner shows waiting message with spinner
type Spinner struct {
	done    chan struct{}
	message string
}

// Stop function stops waiting message
func (w *Spinner) Stop() {
	if w.done != nil {
		w.done <- struct{}{}
	}
}

// Wait function shows waiting message with spinner
func (w *Spinner) Wait() {
	w.done = make(chan struct{})
	if w.message == "" {
		w.message = "Please, wait..."
	}

	go func() {
		characters := [...]string{"-", `\`, `|`, `/`}
		index := 0
		for {
			select {
			case <-w.done:
				fmt.Printf("\r \r")
				close(w.done)
				return
			default:
				if index == len(characters) {
					index = 0
				}

				fmt.Printf("\r%v%v ", w.message, characters[index])
				time.Sleep(time.Millisecond * 150)
				index++
			}
		}
	}()
}
