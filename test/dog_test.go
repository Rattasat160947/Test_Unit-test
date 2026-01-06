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
			Weight: 10,
		}
		ok, err := govalidator.ValidateStruct(dog)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
	t.Run(`Name is required`, func(t *testing.T) {
		dog := entity.Dog{
			Name:   "" ,// ผิดตรงนี้
			Weight: 10,	
		}																								
		ok, err := govalidator.ValidateStruct(dog)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Name is required"))
	})

	t.Run(`Weight is required`, func(t *testing.T) {
		dog := entity.Dog{
			Name:   "Tong",
			Weight: 0,
		}																								
		ok, err := govalidator.ValidateStruct(dog)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Weight is required"))
	})
	t.Run(`Weight must be between 1 and 20`, func(t *testing.T) {
		dog := entity.Dog{
			Name:   "Tong",
			Weight: 40, // ผิดตรงนี้
		}																								
		ok, err := govalidator.ValidateStruct(dog)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Weight must be between 1 and 20"))
	})
}
