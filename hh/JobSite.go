package hh

type JobSite struct {
	Name      string
	Followers map[string]JobSite
}

func (j *JobSite) Subcribe(observer Observer) {
	j.Followers :=
}

func (j *JobSite) UnSubcribe(observer Observer) {


}

func (j *JobSite) SendAll() {

}
