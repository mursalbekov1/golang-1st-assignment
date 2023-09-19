package developer

type Developer struct {
	position string
	salary   int64
	address  string
}

func (d *Developer) GetPosition() string {
	return d.position
}

func (d *Developer) SetPosition(position string) {
	d.position = position
}

func (d *Developer) GetSalary() int64 {
	return d.salary
}
func (d *Developer) SetSalary(salary int64) {
	d.salary = salary
}

func (d *Developer) GetAddress() string {
	return d.address
}

func (d *Developer) SetAddress(address string) {
	d.address = address
}
