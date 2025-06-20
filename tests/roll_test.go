package tests

import (
	"testing"

	"vtt_api/dice"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRollDices(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Roll dices Suite")
}

var _ = Describe("RollService", func() {
	Context("Rolagem simples", func() {
		It("Deve rolar um dado de 6 faces", func() {
			result, err := dice.RollDices("/r 1d6")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Total).To(BeNumerically(">=", 1))
			Expect(result.Total).To(BeNumerically("<=", 6))
		})

		It("Deve rolar um dado fudge", func() {
			result, err := dice.RollDices("/r 1df")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Total).To(BeNumerically(">=", -1))
			Expect(result.Total).To(BeNumerically("<=", 1))
		})

		It("Deve rolar um dado de 6 faces sem indicação da quantidade", func() {
			result, err := dice.RollDices("/r d6")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Total).To(BeNumerically(">=", 1))
			Expect(result.Total).To(BeNumerically("<=", 6))
		})
	})
})
