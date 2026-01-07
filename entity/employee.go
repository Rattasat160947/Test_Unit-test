package entity

type Employee struct{
	EmployeeID string `valid:"matches(^EM\\d{3}$)~EmployeeID invalid format"`
	Name string `valid:"stringlength(5|50)~Name must be between 5 and 50 characters"`
	Position string `valid:"in(Manager|Staff|Admin)~Position must be Manager or Staff or Admin"`
	Salary float64 `valid:"range(9000|100000)~Salary must be between 9000 and 100000"`
}