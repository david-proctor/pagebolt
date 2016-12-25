package templates

import "reflect"

func IsPlaceholder(template Template) bool {
	return reflect.TypeOf(template) == reflect.TypeOf(TemplatePlaceholder{})
}

func Contains(collection []string, element string) bool {
    for _,s := range collection {
        if s == element {
            return true
        }
    }
    return false
}