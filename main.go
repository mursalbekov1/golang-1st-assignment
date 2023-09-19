package main

import (
	developer "assignment1/developer"
	hr1 "assignment1/hr"
	"assignment1/projectManager"
	"assignment1/recruiter"
	"assignment1/scrumMaster"
	"fmt"
)

func main() {
	//developer
	dev := developer.Developer{}
	dev.SetPosition("Junior")
	dev.SetSalary(220000)
	dev.SetAddress("Google")
	fmt.Println("Developer:")
	fmt.Println(dev.GetAddress())
	fmt.Println(dev.GetPosition())
	fmt.Println(dev.GetSalary())

	//hr manager
	hr := hr1.HR{}
	hr.SetPosition("Senior")
	hr.SetSalary(600000)
	hr.SetAddress("Amazon")
	fmt.Println("HR:")
	fmt.Println(hr.GetAddress())
	fmt.Println(hr.GetPosition())
	fmt.Println(hr.GetSalary())

	//projectManager
	projM := projectManager.ProjectManager{}
	projM.SetPosition("Head")
	projM.SetSalary(400000)
	projM.SetAddress("Microsoft")
	fmt.Println("Project Manager:")
	fmt.Println(projM.GetAddress())
	fmt.Println(projM.GetPosition())
	fmt.Println(projM.GetSalary())

	//recruiter
	recr := recruiter.Recruiter{}
	recr.SetPosition("Senior")
	recr.SetSalary(350000)
	recr.SetAddress("Yandex")
	fmt.Println("Recruiter")
	fmt.Println(recr.GetAddress())
	fmt.Println(recr.GetPosition())
	fmt.Println(recr.GetSalary())

	//scrumMaster
	scrum := scrumMaster.ScrumMaster{}
	scrum.SetPosition("Master")
	scrum.SetSalary(600000)
	scrum.SetAddress("Toyota")
	fmt.Println("SCRUM master")
	fmt.Println(scrum.GetAddress())
	fmt.Println(scrum.GetPosition())
	fmt.Println(scrum.GetSalary())
}
