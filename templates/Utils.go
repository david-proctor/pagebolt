package templates

import "reflect"

func IsPlaceholder(template Template) bool {
	return reflect.TypeOf(template) == reflect.TypeOf(TemplatePlaceholder{})
}
