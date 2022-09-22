package ch

type Observable interface {
	Subcribe(Observer)
	UnSubcribe(Observer)
	SendAll()
}
