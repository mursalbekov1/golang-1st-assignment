package recruiter

type Recruiter struct {
	position string
	salary   int64
	address  string
}

func (r *Recruiter) GetPosition() string {
	return r.position
}

func (r *Recruiter) SetPosition(position string) {
	r.position = position
}

func (r *Recruiter) GetSalary() int64 {
	return r.salary
}
func (r *Recruiter) SetSalary(salary int64) {
	r.salary = salary
}

func (r *Recruiter) GetAddress() string {
	return r.address
}

func (r *Recruiter) SetAddress(address string) {
	r.address = address
}
