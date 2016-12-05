package templates

import (
    "reflect"
    "fmt"
)

func AssembleTemplates(scanner DirectoryScanner) []Template {
    templates := scanner.Templates()

    for _,t := range templates {
        fmt.Println("===============")
        fmt.Println("Template name: ", t.Name())
        fmt.Println("Template contents: ", t.String())
    }

    if len(templates) == 0 {
        panic("Failed to find any valid templates")
    }

    cache := buildCache(templates)

    for _,t := range templates {
        fmt.Println("======= !!! =======")
        fmt.Println("Template name: ", t.Name())
        fmt.Println("Template contents: ", t.ProcessedString(cache))
    }

    // At this point we have correct templates when we call ProcessedString.
    // Still need a way to return them as templates.

    templates = replacePlaceholdersWithCached(&templates, cache)

    return templates
}

func buildCache(templates []Template) map[string]Template {
    cache := make(map[string]Template)
    for _,template := range templates {
        if cachedTemplate,found := cache[template.Name()]; found {
            if reflect.TypeOf(cachedTemplate) == reflect.TypeOf(TemplatePlaceholder{}) &&
                reflect.TypeOf(template) != reflect.TypeOf(TemplatePlaceholder{}) {
                cache[template.Name()] = template
            }
            continue
        }
        cache[template.Name()] = template
    }
    return cache
}

func replacePlaceholdersWithCached(rawTemplates *[]Template, cache map[string]Template) []Template {
    output := make([]Template, 0)
    for _,template := range *rawTemplates {
        if IsPlaceholder(template) {
            output = append(output, cache[template.Name()])
        } else {
            output = append(output, template)
        }
    }
    return output
}