package lspinner

import (
	"fmt"
	"time"
)

type Options struct {
	waitingMessage string
}

// New receives message to show while waiting.
func New(opt Options) *Lspinner {
	spinner := &Lspinner{}
	spinner.message = opt.waitingMessage
	return spinner
}

// Lspinner shows waiting message with spinner
type Lspinner struct {
	done    chan struct{}
	message string
}

// Stop function stops waiting message
func (w *Lspinner) Stop() {
	if w.done != nil {
		w.done <- struct{}{}
	}
}

// Wait function shows waiting message with spinner
func (w *Lspinner) Wait() {
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
