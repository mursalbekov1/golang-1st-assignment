package projectManager

type ProjectManager struct {
	name     string
	position string
	salary   int64
	address  string
}

func (pm *ProjectManager) GetPosition() string {
	return pm.position
}

func (pm *ProjectManager) SetPosition(position string) {
	pm.position = position
}

func (pm *ProjectManager) GetSalary() int64 {
	return pm.salary
}
func (pm *ProjectManager) SetSalary(salary int64) {
	pm.salary = salary
}

func (pm *ProjectManager) GetAddress() string {
	return pm.address
}

func (pm *ProjectManager) SetAddress(address string) {
	pm.address = address
}
