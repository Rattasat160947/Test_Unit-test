package test

import (
	"testing"
	"Test_Unit-test/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestUser(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`user is required`, func(t *testing.T) {
		user := entity.User{
			Name: "", // ผิดตรงนี้
		}
		ok, err := govalidator.ValidateStruct(user)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Name is required"))
	})

	t.Run(`user is valid`, func(t *testing.T) {
		user := entity.User{
			Name: "Test",
		}
		ok, err := govalidator.ValidateStruct(user)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}