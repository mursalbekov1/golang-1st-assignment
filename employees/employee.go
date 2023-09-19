package employees

type Employees interface {
	GetPosition() string
	SetPosition(position string)
	GetSalary() int64
	SetSalary(salary int64)
	GetAddress() string
	SetAddress(address string)
}
