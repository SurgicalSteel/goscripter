package goscripter

import (
	"fmt"
)

// BuildJavascriptBundle wraps a javascript file using standard HTML tag for scripts
func BuildJavascriptBundle(body string) string {
	prefix := "<script type='text/javascript'>"
	suffix := "</script>"
	return fmt.Sprintf("%s%s%s", prefix, body, suffix)
}

// BuildCSSBundle wraps a css file using standard HTML tag for styles
func BuildCSSBundle(body string) string {
	prefix := "<style>"
	suffix := "</style>"
	return fmt.Sprintf("%s%s%s", prefix, body, suffix)
}
