package hh

type Observable interface {
	Subcribe(observer Observer)
	UnSubcribe(observer Observer)
	SendAll()
}
