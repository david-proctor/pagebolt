package main_test

import (
	. "github.com/pagebolt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
    "testing"
)

func TestPagebolt(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Pagebolt Suite")
}

var _ = Describe("Main", func() {
	Context("When running main", func(){
		It("Does not throw", func() {
            main := func() { PrintLogo() }
			Expect(main).NotTo(Panic())
		})
	})
})