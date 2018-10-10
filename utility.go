package goscripter

import (
	"fmt"
)

func BuildJavascript(body string) string {
	prefix := "<script type='text/javascript'>"
	suffix := "</script>"
	return fmt.Sprintf("%s%s%s", prefix, body, suffix)
}

func BuildCSS(body string) string {
	prefix := "<style>"
	suffix := "</style>"
	return fmt.Sprintf("%s%s%s", prefix, body, suffix)
}
