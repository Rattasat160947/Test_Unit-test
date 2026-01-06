package test

import(
	"testing"
	"Test_Unit-test/entity"
	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestProduct(t *testing.T){
	g := NewGomegaWithT(t)

	t.Run(`product is valid1`,func(t *testing.T){
		product :=entity.Product{
			Name: "Test",
			Price: 100,
			SKU: "P1234",
		}
		ok,err := govalidator.ValidateStruct(product)
		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run(`Name is required`,func(t *testing.T){
		product :=entity.Product{
			Name: "", // ผิดตรงนี้
			Price: 100,
			SKU: "P1234",
		}
		ok,err := govalidator.ValidateStruct(product)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Name is required"))
	})


	t.Run(`SKU invalid format`,func(t *testing.T){
		product :=entity.Product{
			Name: "test",
			Price: 100,
			SKU: "P12", // ผิดตรงนี้
		}
		ok,err := govalidator.ValidateStruct(product)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("SKU invalid format"))
	})
	t.Run(`SKU invalid format`,func(t *testing.T){
		product :=entity.Product{
			Name: "test",
			Price: 100,
			SKU: "ABCD", // ผิดตรงนี้
		}
		ok,err := govalidator.ValidateStruct(product)
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("SKU invalid format"))
	})
}

func TestPrice(t *testing.T){
	g := NewGomegaWithT(t)
	t.Run(`Price is required`, func(t *testing.T) {

		product := entity.Product{
			Name:  "test",
			Price: 0, // ผิดตรงนี้
			SKU:   "P1234",
		}
		ok, err := govalidator.ValidateStruct(product)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Price is required"))
	})
	t.Run(`Price must between 1 and 10000`, func(t *testing.T) {

		product := entity.Product{
			Name:  "test",
			Price: 200000, // ผิดตรงนี้
			SKU:   "P1234",
		}
		ok, err := govalidator.ValidateStruct(product)
		
		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Price must between 1 and 10000"))
	})
	
}