package templates

import (
    "strings"
)

func Assemble(source string) Template {
    if len(source) == 0 {
        return makeEmptyTemplate()
    }

    sections := splitSource(source)

    if len(sections) == 1 {
        contents := make([]Template, 1)
        result := Container {
            contents: contents,
        }
        insertLiteral(source, result, 0)
        return result
    }

    contents := make([]Template, len(sections) * 2 - 1)
    template := Container {
        contents: contents,
    }

    for index,element := range sections {
        if strings.HasPrefix(element, "<#") {
            insertPlaceholder(element, template, index)
        } else {
            insertLiteral(element, template, index)
        }
    }

    return template
}

func splitSource(source string) []string {
    splitOnOpenTags := strings.Split(source, "<#")
    sections := make([]string, len(splitOnOpenTags) * 2 - 1)
    for index,element := range splitOnOpenTags {
        if index == 0 {
            sections[0] = element
            continue
        }
        firstIndex := index * 2 - 1
        splitOnCloseTags := strings.Split(element, "#>")
        sections[firstIndex] = "<# " + splitOnCloseTags[0] + "#>"
        sections[firstIndex + 1] = splitOnCloseTags[1]
    }

    return sections
}

func makeEmptyTemplate() Template {
    contents := make([]Template, 0)
    template := Container {
        contents: contents,
    }
    return template
}

func insertLiteral(content string, container Container, index int) {
    literal := Literal {
        textContent: content,
    }
    container.contents[index] = literal
}

func insertPlaceholder(name string, container Container, index int) {
    name = name[3 : len(name) - 3]
    placeholder := TemplatePlaceholder {
        name: name,
    }
    container.contents[index] = placeholder
}