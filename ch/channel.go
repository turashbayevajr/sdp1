package ch

type Channel struct {
	Name      string
	Followers map[string]Observer
}

func (c *Channel) Subcribe(observer Observer) {
	c.Followers[observer.GetName()] = observer
}

func (c *Channel) UnSubcribe(observer Observer) {
	delete(c.Followers, observer.GetName())
}

func (c *Channel) SendAll() {
	for _, follower := range c.Followers {
		follower.UpDate(c.Name)
	}
}
