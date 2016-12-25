package templates

type Template interface {
    Contents() []Template
    String() string
    ProcessedString(cache TemplateCache) string
    Name() string
}

type Container struct {
    contents []Template
    name string
}

func (c Container) Contents() []Template {
    return c.contents
}

func (c Container) String() string {
    output := ""
    for _,element := range c.contents {
        output += element.String()
    }
    return output
}

func (c Container) ProcessedString(cache TemplateCache) string {
    output := ""
    for _,element := range c.contents {
        if IsPlaceholder(element) {
            output += cache.Get(element.Name()).ProcessedString(cache)
        } else {
            output += element.String()
        }
    }
    return output
}

func (c Container) Name() string {
    return c.name
}

func (c* Container) AppendLiteral(name string, text string) {
    literal := Literal {
        name: name,
        textContent: text,
    }
    c.contents = append(c.contents, literal)
}

func (c* Container) AppendPlaceholder(name string) {
    placeholder := TemplatePlaceholder {
        name: name,
    }
    c.contents = append(c.contents, placeholder)
}

type Literal struct {
    textContent string
    name string
}

func (l Literal) Contents() []Template {
    return make([]Template, 0)
}

func (l Literal) String() string {
    return l.textContent
}

func (l Literal) ProcessedString (TemplateCache) string {
    return l.String()
}

func (l Literal) Name() string {
    return l.textContent
}

type TemplatePlaceholder struct {
    name string
}

func (p TemplatePlaceholder) Contents() []Template {
    return make([]Template, 0)
}

func (p TemplatePlaceholder) String() string {
    return "PLACEHOLDER<" + p.name + ">"
}

func (p TemplatePlaceholder) ProcessedString (TemplateCache) string {
    panic("Cannot process string on a TemplatePlaceholder")
}

func (p TemplatePlaceholder) Name() string {
    return p.name
}