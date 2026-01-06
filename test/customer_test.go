package test

import (
	"Test_Unit-test/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)
func TestCustomer(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Customer is valid`, func(t *testing.T) {
		customer := entity.Customer{
			Username: "Rattasat",
			Email:    "B6640583@Example.com",
			Gender:   "Male",
			Age:      20,
		}

		ok, err := govalidator.ValidateStruct(customer)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
	t.Run(`Username must be betweent 5 and 20 characters`,func(t *testing.T){
		customer := entity.Customer{
			Username: "Rat", // ผิดตรงนี้
			Email:    "B6640583@Example.com",
			Gender:   "Male",
			Age:      20,
		}
		ok,err := govalidator.ValidateStruct(customer)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Username must be between 5 and 20 characters"))
	})
	t.Run(`Invalid email format`,func(t *testing.T){
		customer := entity.Customer{
			Username: "Rattasat",
			Email:    "B6640583", // ผิดตรงนี้
			Gender:   "Male",
			Age:      20,
		}
		ok,err := govalidator.ValidateStruct(customer)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Invalid email format"))
	})
	t.Run(`Gender must be Male or Female`,func(t *testing.T){
		customer := entity.Customer{
			Username: "Rattasat",
			Email:    "B6640583@Example.com",
			Gender:   "Other",  // ผิดตรงนี้
			Age:      20,
		}
		ok,err := govalidator.ValidateStruct(customer)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Gender must be Male or Female"))
	})
	t.Run(`Age must be between 18 and 99`,func(t *testing.T){
		customer := entity.Customer{
			Username: "Rattasat",
			Email:    "B6640583@Example.com",
			Gender:   "Male",
			Age:      17, // ผิดตรงนี้
		}
		ok,err := govalidator.ValidateStruct(customer)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Age must be between 18 and 99"))
	})
}
