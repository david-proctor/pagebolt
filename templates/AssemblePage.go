package templates

import (
    "strings"
)

func AssemblePage(name string, source string) Template {
    if len(source) == 0 {
        return makeEmptyTemplate()
    }

    sections := splitSource(source)

    if len(sections) == 1 {
        result := Container {
            name: name,
            contents: make([]Template, 0),
        }
        result.AppendLiteral(name, source)
    }

    template := Container {
        name: name,
        contents: make([]Template, 0),
    }

    for index,element := range sections {
        if strings.HasPrefix(element, "<#") {
            placeholderName := getNameFromPlaceholderTag(element)
            if(placeholderName == name) {
                panic("Cannot assemble page from self-referencing template: " + name)
            }
            template.AppendPlaceholder(placeholderName)
        } else {
            template.AppendLiteral(name + string(index), element)
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
        name: "EMPTY",
        contents: contents,
    }
    return template
}

func getNameFromPlaceholderTag(tag string) string {
    name := tag[3 : len(tag) - 3]
    return strings.TrimSpace(name)
}