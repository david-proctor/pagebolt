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

var _ = Describe("AssemblePage", func() {

	Context("when calling AssemblePage()", func() {
        It("sets provided name", func() {
            Expect(AssemblePage("Name", "Arglebargle").Name()).To(Equal("Name"))
        })
        It("sets name to EMPTY when template is empty", func() {
            Expect(AssemblePage("Name", "").Name()).To(Equal("EMPTY"))
        })
		It("does not panic when source is empty", func() {
			Expect(func() { AssemblePage("Name", "") }).NotTo(Panic())
		})
        It("is empty when source is empty", func() {
            Expect(AssemblePage("Name", "").Contents()).To(BeEmpty())
        })
        It("returns one-level string when no sub-templates are included", func() {
            template := AssemblePage("Name", "Arglebargle")
            Expect(template.String()).To(Equal("Arglebargle"))
        })
        It("returns placeholder when template token is included", func() {
            template := AssemblePage("Name", "Arglebargle <# templatename #> faffernaff")
            Expect(template.Contents()[0].String()).To(Equal("Arglebargle "))
            Expect(template.Contents()[1]).To(BeAssignableToTypeOf(TemplatePlaceholder{}))
            Expect(template.Contents()[1].Name()).To(Equal("templatename"))
            Expect(template.Contents()[2].String()).To(Equal(" faffernaff"))
        })
        It("returns complete ordered types when many template tokens are included", func() {
            template := AssemblePage("Name", "Arglebargle <# template1 #> faffernaff <# template2 #> morblewoosh")
            Expect(template.Contents()[0].String()).To(Equal("Arglebargle "))
            Expect(template.Contents()[1]).To(BeAssignableToTypeOf(TemplatePlaceholder{}))
            Expect(template.Contents()[2].String()).To(Equal(" faffernaff "))
            Expect(template.Contents()[3]).To(BeAssignableToTypeOf(TemplatePlaceholder{}))
            Expect(template.Contents()[4].String()).To(Equal(" morblewoosh"))
        })
        It("panics on invalid template tokens", func() {
            Expect(func() { AssemblePage("Name", "Arglebargle <# template1 ># faffernaff") }).To(Panic())
        })
	})

    Context("When calling AssembleTemplateCache()", func() {
        literal1 := AssemblePage("literal1", "Literal 1")
        templateWithLiteral1 := AssemblePage("Template", "TemplateWithLiteral1 [<# literal1 #>]")
        directoryScanner := MockDirectoryScanner { }

        It("panics when directory scanner has no results", func() {
            directoryScanner := MockDirectoryScanner{}

            Expect(func(){AssembleTemplateCache(directoryScanner)}).To(Panic())
        })
        It("Collects correct templates in cache", func() {
            directoryScanner.Setup(literal1, templateWithLiteral1)

            cache := AssembleTemplateCache(directoryScanner)

            literalCheck := func() bool { return cache["literal1"].String() == "Literal 1" }
            Expect(literalCheck()).To(BeTrue())

            Expect(cache["Template"].Contents()[1].Name()).To(Equal(cache["literal1"].Name()))
        })
        It("Correctly substitutes placeholder values when calling ProcessedString", func() {
            directoryScanner.Setup(literal1, templateWithLiteral1)

            cache := AssembleTemplateCache(directoryScanner)
            expected := "TemplateWithLiteral1 [Literal 1]"
            actual := cache["Template"].ProcessedString(cache)

            Expect(actual).To(Equal(expected))
        })
    })
})

type MockDirectoryScanner struct {
    templates []Template
}

func (s *MockDirectoryScanner) Setup (templates ...Template) {
    s.templates = make([]Template, len(templates))
    for i,t := range templates {
        s.templates[i] = t
    }
}

func (s MockDirectoryScanner) Templates () []Template {
    return s.templates
}