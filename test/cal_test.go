package test

import (
	"Test_Unit-test/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"

)

func TestCal(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`Cal is valid`, func(t *testing.T) {
		cal := entity.Cal{
			Number2: "123456",
			Number3: "abc",
			Number4: "def",
		}
		ok, err := govalidator.ValidateStruct(cal)
		if err != nil {
			t.Logf("Error: %v", err)
		}
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`Number2 must be negative`, func(t *testing.T) {
		cal := entity.Cal{
			Number2: "abc", // error
			Number3: "abc",
			Number4: "def",
		}
		ok, err := govalidator.ValidateStruct(cal)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Number2 must be negative"))
	})
	t.Run(`Number3 is not zero`, func(t *testing.T) {
		cal := entity.Cal{
			Number2: "123456",
			Number3: "123", // error
			Number4: "def",
		}
		ok, err := govalidator.ValidateStruct(cal)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Number3 is not zero"))
	})
	t.Run(`Number4 must be non-positive`, func(t *testing.T) {
		cal := entity.Cal{	
			Number2: "123456",
			Number3: "abc",
			Number4: "DEF", // error
		}
		ok, err := govalidator.ValidateStruct(cal)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Number4 must be non-positive"))
	})
}