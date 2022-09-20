package model

type Employee struct {
	ID       int
	Name     string
	Age      int
	Division string
}

var Emps = []Employee{
	{
		ID:       1,
		Name:     "User 1",
		Age:      18,
		Division: "Tech",
	},
}

func GetEmployee() []Employee {
	return Emps
}

func CreateEmployee(emp *Employee) []Employee {
	Emps = append(Emps, *emp)
	return Emps
}
