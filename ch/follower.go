package ch

import "fmt"

type Follower struct {
	Name string
}

func (f *Follower) UpDate(cname string) {
	fmt.Printf("send post to %s from %s\n", f.GetName(), cname)

}

func (f *Follower) GetName() string {
	return f.Name
}
