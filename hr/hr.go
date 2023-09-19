package hr

type HR struct {
	position string
	salary   int64
	address  string
}

func (hr *HR) GetPosition() string {
	return hr.position
}

func (hr *HR) SetPosition(position string) {
	hr.position = position
}

func (hr HR) GetSalary() int64 {
	return hr.salary
}
func (hr *HR) SetSalary(salary int64) {
	hr.salary = salary
}

func (hr *HR) GetAddress() string {
	return hr.address
}

func (hr *HR) SetAddress(address string) {
	hr.address = address
}
