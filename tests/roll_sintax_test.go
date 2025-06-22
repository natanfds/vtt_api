package tests

import (
	"testing"

	"vtt_api/dice"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRollSintax(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Roll Sintax Suite")
}

var _ = Describe("RollSintaxService", func() {

	Context("Validando sintaxe com rolagens válidas", func() {
		It("Um dado simles", func() {
			err := dice.ValidateRollSintax([]string{"1", "d", "20"})
			Expect(err).ToNot(HaveOccurred())
		})

		It("Um dado fudge", func() {
			err := dice.ValidateRollSintax([]string{"1", "d", "20"})
			Expect(err).ToNot(HaveOccurred())
		})

		It("Com kh", func() {
			err := dice.ValidateRollSintax([]string{"2", "d", "6", "kh", "1"})
			Expect(err).ToNot(HaveOccurred())
		})

		It("Explosiva", func() {
			err := dice.ValidateRollSintax([]string{"3", "d", "8", "ex"})
			Expect(err).ToNot(HaveOccurred())
		})

		It("Crítico em valor", func() {
			err := dice.ValidateRollSintax([]string{"4", "d", "10", "csv", "5"})
			Expect(err).ToNot(HaveOccurred())
		})

		It("Crítico em quantidade", func() {
			err := dice.ValidateRollSintax([]string{"5", "d", "12", "km", "2"})
			Expect(err).ToNot(HaveOccurred())
		})

		It("Reroll", func() {
			err := dice.ValidateRollSintax([]string{"6", "d", "4", "re", "7"})
			Expect(err).ToNot(HaveOccurred())
		})

		It("Ying Yang", func() {
			err := dice.ValidateRollSintax([]string{"7", "d", "6", "yy"})
			Expect(err).ToNot(HaveOccurred())
		})

		It("Crítico caso igual", func() {
			err := dice.ValidateRollSintax([]string{"8", "d", "8", "cse", "5"})
			Expect(err).ToNot(HaveOccurred())
		})

		It("Sucesso", func() {
			err := dice.ValidateRollSintax([]string{"9", "d", "10", "su", "7"})
			Expect(err).ToNot(HaveOccurred())
		})

		It("Red or Blue", func() {
			err := dice.ValidateRollSintax([]string{"10", "d", "12", "rb"})
			Expect(err).ToNot(HaveOccurred())
		})
	})

	Context("Validando alertas de rolagens inválidas", func() {

		It("Token repetido", func() {
			err := dice.ValidateRollSintax([]string{"1", "d", "20", "d"})
			Expect(err).To(HaveOccurred())
		})

		It("Token inválido", func() {
			err := dice.ValidateRollSintax([]string{"1", "d", "20", "invalid"})
			Expect(err).To(HaveOccurred())
		})

		It("Fudge fora de ordem", func() {
			err := dice.ValidateRollSintax([]string{"1", "d", "20", "f", "kh"})
			Expect(err).To(HaveOccurred())
		})

		It("Quantidade não especificada", func() {
			err := dice.ValidateRollSintax([]string{"d", "20", "kh"})
			Expect(err).To(HaveOccurred())
		})

		It("Mais de um Keep", func() {
			err := dice.ValidateRollSintax([]string{"1", "d", "20", "kh", "1", "kl", "2"})
			Expect(err).To(HaveOccurred())
		})

		It("Tipos de Crítico Misturados", func() {
			err := dice.ValidateRollSintax([]string{"1", "d", "20", "csv", "5", "cse", "2"})
			Expect(err).To(HaveOccurred())
		})

		It("Ordem errada de tokens", func() {
			err := dice.ValidateRollSintax([]string{"1", "f", "d", "20", "ex", "kh", "1"})
			Expect(err).To(HaveOccurred())
		})
	})
})
