package test

import (
	"Test_Unit-test/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestEmployee(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`Employee is valid`, func(t *testing.T) {
		employee := entity.Employee{
			EmployeeID: "EM111",
			Name:       "Rattasat",
			Position:   "Admin",
			Salary:     50000.00,
		}
		ok, err := govalidator.ValidateStruct(employee)
		if err != nil {
			t.Logf("Error: %v", err)
		}
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
	t.Run(`EmployeeID invalid format`, func(t *testing.T) {
		employee := entity.Employee{
			EmployeeID: "E111", //error
			Name:       "Rattasat",
			Position:   "Admin",
			Salary:     50000.00,
		}
		ok, err := govalidator.ValidateStruct(employee)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("EmployeeID invalid format"))
	})

	t.Run(`Name must be between 5 and 50 characters`, func(t *testing.T) {
		employee := entity.Employee{
			EmployeeID: "EM111",
			Name:       "Ratt", // error
			Position:   "Admin",
			Salary:     50000.00,
		}
		ok, err := govalidator.ValidateStruct(employee)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Name must be between 5 and 50 characters"))
	})

	t.Run(`Position must be Manager or Staff or Admin`, func(t *testing.T) {
		employee := entity.Employee{
			EmployeeID: "EM111",
			Name:       "Rattasat",
			Position:   "Ant", // error
			Salary:     50000.00,
		}
		ok, err := govalidator.ValidateStruct(employee)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Position must be Manager or Staff or Admin"))
	})

	t.Run(`Salary must be between 9000 and 100000`, func(t *testing.T) {
		employee := entity.Employee{
			EmployeeID: "EM111",
			Name:       "Rattasat",
			Position:   "Staff",
			Salary:     100.00, // error
		}
		ok, err := govalidator.ValidateStruct(employee)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Salary must be between 9000 and 100000"))
	})

}
