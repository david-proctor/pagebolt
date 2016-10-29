package templates

import (
    //"fmt"
)

type Template interface {
    Contents() []Template
    String() string
}

type Container struct {
    contents []Template
}

func (c Container) Contents() []Template {
    return c.contents
}

func (c Container) String() string {
    var output string
    for _,element := range c.contents {
        output += element.String()
    }
    return output
}

type Literal struct {
    textContent string
}

func (l Literal) Contents() []Template {
    return make([]Template, 0)
}

func (l Literal) String() string {
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