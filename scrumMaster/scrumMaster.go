package scrumMaster

type ScrumMaster struct {
	position string
	salary   int64
	address  string
}

func (sm *ScrumMaster) GetPosition() string {
	return sm.position
}

func (sm *ScrumMaster) SetPosition(position string) {
	sm.position = position
}

func (sm *ScrumMaster) GetSalary() int64 {
	return sm.salary
}
func (sm *ScrumMaster) SetSalary(salary int64) {
	sm.salary = salary
}

func (sm *ScrumMaster) GetAddress() string {
	return sm.address
}

func (sm *ScrumMaster) SetAddress(address string) {
	sm.address = address
}
