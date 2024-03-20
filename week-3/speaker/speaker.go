package speaker

type Speaker interface {
	Speak()
}

func Speak(speaker Speaker) {
	speaker.Speak()
}
