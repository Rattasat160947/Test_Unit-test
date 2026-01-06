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
			Email:    "B6640583@gmail.com",
			Gender:   "Male",
			Age:      20,
		}

		ok, err := govalidator.ValidateStruct(customer)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
