package test

import (
	"myapp/entity"
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func Test1(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("test", func(t *testing.T) {
		approved := entity.Product{
			Title:        "hee5",
			Price:       455,
			Code:  "BK123454",
		}

		ok, err := govalidator.ValidateStruct(approved)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func Test2(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("test", func(t *testing.T) {
		approved := entity.Product{
			Title:        "he",
			Price:       455,
			Code:  "B6612122",
		}

		ok, err := govalidator.ValidateStruct(approved)

		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("Title must be greater than 3"))
	})
}

func Test3(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("test", func(t *testing.T) {
		approved := entity.Product{
			Title:        "heee",
			Price:       45,
			Code:  "B6612122",
		}

		ok, err := govalidator.ValidateStruct(approved)

		g.Expect(ok).To(BeFalse())
		g.Expect(err.Error()).To(Equal("Price must be greater than 0"))
	})
}


