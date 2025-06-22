package tests

import (
	"testing"

	"vtt_api/dice"
	"vtt_api/models"

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

		It("Deve rolar um dado fudge sem indicação da quantidade", func() {
			result, err := dice.RollDices("/r df")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Total).To(BeNumerically(">=", -1))
			Expect(result.Total).To(BeNumerically("<=", 1))
		})

		It("Deve rolar 10 dados fudge", func() {
			result, err := dice.RollDices("/r 10df")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Total).To(BeNumerically(">=", -10))
			Expect(result.Total).To(BeNumerically("<=", 10))
		})

		It("Deve rolar 10 dados de 6 faces", func() {
			result, err := dice.RollDices("/r 10d6")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Total).To(BeNumerically(">=", 10))
			Expect(result.Total).To(BeNumerically("<=", 60))
		})

		It("Deve retornar 10", func() {
			result, err := dice.RollDices("/r 10")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Total).To(BeNumerically("==", 10))
		})

		It("Deve somar 10 e 5 para retornar 15", func() {
			result, err := dice.RollDices("/r 10 + 5")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Total).To(BeNumerically("==", 15))
		})

		It("Deve subtrair 5 de 10 para retornar 5", func() {
			result, err := dice.RollDices("/r 10 - 5")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Total).To(BeNumerically("==", 5))
		})

		It("Deve multiplicar 10 por 5 para retornar 50", func() {
			result, err := dice.RollDices("/r 10 * 5")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Total).To(BeNumerically("==", 50))
		})

		It("Deve dividir 10 por 5 para retornar 2", func() {
			result, err := dice.RollDices("/r 10 / 5")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Total).To(BeNumerically("==", 2))
		})

		It("Deve multiplicar antes de subitrair e retornar 0", func() {
			result, err := dice.RollDices("/r 10 - 5 * 2")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Total).To(BeNumerically("==", 0))
		})

		It("Deve dividir antes de somar e retornar 7", func() {
			result, err := dice.RollDices("/r 10 / 5 + 5")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Total).To(BeNumerically("==", 7))
		})

		It("Deve rolar 1d2 e explodir o resultado", func() {
			var result models.DiceCommandResult
			var err error

			for i := 0; i < 50; i++ {
				result, err = dice.RollDices("/r 1d2ex")
				if len(result.Results) > 1 {
					break
				}
			}

			Expect(err).ToNot(HaveOccurred())
			Expect(len(result.Results)).To(BeNumerically(">", 1))
		})

		It("Ao rolar 1d1 deve ocorrer erro se explodido", func() {
			_, err := dice.RollDices("/r 1d1ex")
			Expect(err).To(HaveOccurred())
		})
	})
})
