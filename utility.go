package goscripter

import (
	"fmt"
)

func BuildJavascriptBundle(body string) string {
	prefix := "<script type='text/javascript'>"
	suffix := "</script>"
	return fmt.Sprintf("%s%s%s", prefix, body, suffix)
}

func BuildCSSBundle(body string) string {
	prefix := "<style>"
	suffix := "</style>"
	return fmt.Sprintf("%s%s%s", prefix, body, suffix)
}
