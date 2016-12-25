package templates

import (
    "reflect"
)

type TemplateCache struct {
    cache map[string]Template
}

func (cache TemplateCache) Get (name string) Template {
    return cache.cache[name]
}

func (cache TemplateCache) GetAll () []Template {
    templates := make([]Template, 0)
    for _,t := range cache.cache {
        templates = append(templates, t)
    }
    return templates
}

func (cache TemplateCache) panicIfCircularReferenceDetected() {
    for _,template := range cache.cache {
        cache.panicIfTemplateRefersToParents(template, make([]string, 0))
    }
}

func MakeCache (scanner DirectoryScanner) TemplateCache {
    rawTemplates := scanner.Templates()

    if (len(rawTemplates) == 0) {
        panic("Failed to find any valid templates")
    }

    templateMap := make(map[string]Template)
    for _,template := range rawTemplates {
        if cachedTemplate,found := templateMap[template.Name()]; found {
            if reflect.TypeOf(cachedTemplate) == reflect.TypeOf(TemplatePlaceholder{}) &&
                reflect.TypeOf(template) != reflect.TypeOf(TemplatePlaceholder{}) {
                templateMap[template.Name()] = template
            }
            continue
        }
        templateMap[template.Name()] = template
    }
    cache := TemplateCache {
        cache: templateMap,
    }

    cache.panicIfCircularReferenceDetected()
    return cache
}

func (cache TemplateCache) panicIfTemplateRefersToParents (template Template, parentNames []string) {
    forbiddenNames := append(parentNames, template.Name())

    for _,child := range template.Contents() {
        if(IsPlaceholder(child)) {
            if (Contains(forbiddenNames, child.Name())) {
                panic("Detected circular reference in template " + child.Name())
            }
            cached := cache.Get(child.Name())
            cache.panicIfTemplateRefersToParents(cached, forbiddenNames)
        } else {
            cache.panicIfTemplateRefersToParents(child, forbiddenNames)
        }
    }
}