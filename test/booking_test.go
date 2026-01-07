package test

import (
	"Test_Unit-test/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestBooking(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run(`Booking is valid`, func(t *testing.T) {
		booking := entity.Booking{
			UserEmail:    "Rattasat@gmail.com",
			MeetingLink:  "http://www.test.com",
			Topic:        "Test",
			Participants: 25,
		}
		ok, err := govalidator.ValidateStruct(booking)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
	t.Run(`Invalid email format`, func(t *testing.T) {
		booking := entity.Booking{
			UserEmail:    "Rattagmail.com", // error
			MeetingLink:  "http://www.test",
			Topic:        "Test",
			Participants: 25,
		}
		ok, err := govalidator.ValidateStruct(booking)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Invalid email format"))
	})

	t.Run(`Invalid URL`, func(t *testing.T) {
		booking := entity.Booking{
			UserEmail:    "Ratta@gmail.com",
			MeetingLink:  "http//test", // error
			Topic:        "Test",
			Participants: 25,
		}
		ok, err := govalidator.ValidateStruct(booking)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Invalid URL"))
	})

	t.Run(`Topic is required`, func(t *testing.T) {
		booking := entity.Booking{
			UserEmail:    "Ratta@gmail.com",
			MeetingLink:  "https://www.test",
			Topic:        "", // error
			Participants: 25,
		}
		ok, err := govalidator.ValidateStruct(booking)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Topic is required"))
	})

	t.Run(`Participants must be between 2 and 50`, func(t *testing.T) {
		booking := entity.Booking{
			UserEmail:    "Ratta@gmail.com",
			MeetingLink:  "https://www.test",
			Topic:        "test",
			Participants: 100, // error
		}
		ok, err := govalidator.ValidateStruct(booking)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Participants must be between 2 and 50"))
	})

}
