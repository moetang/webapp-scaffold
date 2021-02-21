package frmgin

import (
	"html"
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
)

func InitBuiltinFunc(c *gin.Engine) {
	c.FuncMap["mx_simple_rich_text"] = func(s string) template.HTML {
		return template.HTML(SimpleFormatRichText(s))
	}
}

func SimpleFormatRichText(s string) string {
	s = html.EscapeString(s)
	ss := strings.Split(s, "\n")
	var sb strings.Builder
	for _, v := range ss {
		if len(strings.TrimSpace(v)) > 0 {
			sb.WriteString("<p>")
			sb.WriteString(v)
			sb.WriteString("</p>")
		} else {
			sb.WriteString("<br>")
		}
	}
	return sb.String()
}
