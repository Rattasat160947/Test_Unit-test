package test

import (
	"testing"
	"Test_Unit-test/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestExample2(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`ExampleSpec valid`, func(t *testing.T){
		example := entity.ExampleSpec{
			Age : 25,
			Username : "tester",
			RawData : "1234567890",
			Nickname : "ทดสอบ",
			StudentID : "S12345678",
			Role : "user",
			Password : "mypassword",
			Description : "This is a valid description.",
		}
		ok, err := govalidator.ValidateStruct(example)
		if err != nil {
            t.Logf("Validation Error: %v", err) 
        }

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
		g.Expect(err.Error()).To(Equal(""))
	})
}