package templates

import (
    "reflect"
)

func AssembleTemplateCache(scanner DirectoryScanner) map[string]Template {
    templates := scanner.Templates()

    if len(templates) == 0 {
        panic("Failed to find any valid templates")
    }

    cache := buildCache(templates)

    return cache
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