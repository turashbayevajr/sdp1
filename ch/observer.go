package ch

type Observer interface {
	UpDate(cname string)
	GetName() string
}
