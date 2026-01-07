package test

import (
	"Test_Unit-test/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestLibrary(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`Library is valid`, func(t *testing.T) {
		library := entity.Library{
			Title:  "Onepiece",
			Author: "ODA",
			Page:   500,
			ISBN:   "ISBN-1234",
		}
		ok, err := govalidator.ValidateStruct(library)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`Title is required`, func(t *testing.T) {
		library := entity.Library{
			Title:  "", // error
			Author: "ODA",
			Page:   500,
			ISBN:   "ISBN-1234",
		}
		ok, err := govalidator.ValidateStruct(library)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Title is required"))
	})
	t.Run(`Author is required`, func(t *testing.T) {
		library := entity.Library{
			Title:  "Onepiece",
			Author: "", // error
			Page:   500,
			ISBN:   "ISBN-1234",
		}
		ok, err := govalidator.ValidateStruct(library)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Author is required"))
	})
	t.Run(`Page is required`, func(t *testing.T) {
		library := entity.Library{
			Title:  "Onepiece",
			Author: "ODA",
			Page:   0, // error
			ISBN:   "ISBN-1234",
		}
		ok, err := govalidator.ValidateStruct(library)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Page is required"))
	})
	t.Run(`Page must be between 1 and 1000`, func(t *testing.T) {
		library := entity.Library{
			Title:  "Onepiece",
			Author: "ODA",
			Page:   1001, // error
			ISBN:   "ISBN-1234",
		}
		ok, err := govalidator.ValidateStruct(library)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Page must be between 1 and 1000"))
	})
	t.Run(`ISBN invalid format`, func(t *testing.T) {
		library := entity.Library{
			Title:  "Onepiece",
			Author: "ODA",
			Page:   500,
			ISBN:   "I-1234",
		}
		ok, err := govalidator.ValidateStruct(library)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("ISBN invalid format"))
	})
}
