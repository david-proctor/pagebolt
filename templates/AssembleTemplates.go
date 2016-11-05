package templates

import (
    "reflect"
)

func AssembleTemplates(scanner DirectoryScanner) []Template {
    templates := scanner.Templates()
    if len(templates) == 0 {
        panic("Failed to find any valid templates")
    }

    cache := buildCache(templates)
    templates = replacePlaceholdersWithCached(templates, cache)

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

func replacePlaceholdersWithCached(templates []Template, cache map[string]Template) []Template {
    for _,template := range templates {
        switch t := template.(type) {
        case Container:
            for i,content := range t.contents {
                if reflect.TypeOf(content) == reflect.TypeOf(TemplatePlaceholder{}) {
                    t.contents[i] = cache[content.Name()]
                }
            }
        case TemplatePlaceholder:
        case Literal:
        default:
            continue
        }
    }
    return templates
}