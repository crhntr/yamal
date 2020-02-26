package yamal_test

import (
	"bufio"
	"os"
	"testing"

	Ω "github.com/onsi/gomega"

	"github.com/crhntr/yamal"
)

func TestLexicalScanner(t *testing.T) {
	Ω.RegisterTestingT(t)

	data, fileErr := os.Open("test_data/key_value.yml")
	Ω.Expect(fileErr).NotTo(Ω.HaveOccurred())

	var tokens []yamal.Token

	err := yamal.LexicalScanner(bufio.NewReader(data), &tokens)
	Ω.Expect(err).NotTo(Ω.HaveOccurred())

	Ω.Expect(len(tokens) > 0).To(Ω.BeTrue())
	Ω.Expect(tokens[0].Value).To(Ω.Equal([]byte("---")))
	Ω.Expect(tokens[1].Value).To(Ω.Equal([]byte("\n")))
	Ω.Expect(tokens[2].Value).To(Ω.Equal([]byte("key")))
}
