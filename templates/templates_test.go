package templates_test

import (
	. "github.com/pagebolt/templates"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestTemplates(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pagebolt Suite")
}

var _ = Describe("TemplateAssembler", func() {

	Context("when calling Assemble()", func() {
		It("does not panic when source is empty", func() {
			Expect(func() { Assemble("") }).NotTo(Panic())
		})
        It("is empty when source is empty", func() {
            Expect(Assemble("").Contents()).To(BeEmpty())
        })
        It("returns one-level string when no sub-templates are included", func() {
            template := Assemble("Arglebargle")
            Expect(template.String()).To(Equal("Arglebargle"))
        })
        It("returns placeholder when template token is included", func() {
            template := Assemble("Arglebargle <# templatename #> faffernaff")
            Expect(template.Contents()[0].String()).To(Equal("Arglebargle "))
            Expect(template.Contents()[1]).To(BeAssignableToTypeOf(TemplatePlaceholder{}))
            Expect(template.Contents()[2].String()).To(Equal(" faffernaff"))
        })
        It("returns complete ordered types when many template tokens are included", func() {
            template := Assemble("Arglebargle <# template1 #> faffernaff <# template2 #> morblewoosh")
            Expect(template.Contents()[0].String()).To(Equal("Arglebargle "))
            Expect(template.Contents()[1]).To(BeAssignableToTypeOf(TemplatePlaceholder{}))
            Expect(template.Contents()[2].String()).To(Equal(" faffernaff "))
            Expect(template.Contents()[3]).To(BeAssignableToTypeOf(TemplatePlaceholder{}))
            Expect(template.Contents()[4].String()).To(Equal(" morblewoosh"))
        })
        It("argle", func() {
            Expect("Aaag"[0:2]).To(Equal("Aa"))
        })
	})
})
