package entity_test

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"

	"example/exam/entity"
)

func TestSuccessCase(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	product := entity.Product{
		Name:     "Test Product",
		Price:    100.0,
		SKU:      "SKU12345",
		Quantity: 10,
	}

	ok, err := govalidator.ValidateStruct(product)
	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())

}

func TestMissingName(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	product := entity.Product{
		Name:     "",
		Price:    100.00,
		SKU:      "SKU12345",
		Quantity: 10,
	}
	ok, err := govalidator.ValidateStruct(product)
	g.Expect(ok).To(gomega.BeFalse())
	g.Expect(err).NotTo(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Name is required"))
}

func TestPriceOutOfRange(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	product := entity.Product{
		Name:     "kakaka",
		Price:    -1.00,
		SKU:      "SKU99999",
		Quantity: 10,
	}
	ok, err := govalidator.ValidateStruct(product)
	g.Expect(ok).To(gomega.BeFalse())
	g.Expect(err).NotTo(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Price must be between 1.00 and 100000.00"))
}

func TestInvalidSKU(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	product := entity.Product{
		Name:     "kakaka",
		Price:    100.00,
		SKU:      "SwU99999",
		Quantity: 10,
	}
	ok, err := govalidator.ValidateStruct(product)
	g.Expect(ok).To(gomega.BeFalse())
	g.Expect(err).NotTo(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("SKU is invalid"))
}

func TestQuantityOutOfRange(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	product := entity.Product{
		Name:     "kakaka",
		Price:    100.00,
		SKU:      "SKU99999",
		Quantity: 501,
	}
	ok, err := govalidator.ValidateStruct(product)
	g.Expect(ok).To(gomega.BeFalse())
	g.Expect(err).NotTo(gomega.BeNil())
	g.Expect(err.Error()).To(gomega.Equal("Quantity must be between 1 and 500"))
}
