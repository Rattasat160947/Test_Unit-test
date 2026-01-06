package test

import (
	"Test_Unit-test/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestDog(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`Dog is valid`, func(t *testing.T) {
		dog := entity.Dog{
			Name:   "Tong",
			Weight: 0,
		}
		ok, err := govalidator.ValidateStruct(dog)
		if err != nil {
			t.Logf("validation : %v", err)
		}
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}
