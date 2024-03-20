package main

import (
	"github.com/KKGo-Software-engineering/coaching-session/week-3/cat"
	"github.com/KKGo-Software-engineering/coaching-session/week-3/speaker"
)

func main() {
	speaker.Speak(cat.New())
}
