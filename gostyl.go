package styles

import (
	"os"
	"strings"
)

type Gostyl struct {
	codes   []string
	noColor bool
}

// NewStyle creates a new style with auto-detect colour support
// Color is disabled if NO_COLOR env var is set or if TERM "dumb" or empty
func NewStyle() Gostyl {
	return Gostyl{
		codes: []string{},
		noColor: os.Getenv("NO_COLOR") != "" ||
			os.Getenv("TERM") == "dumb" ||
			os.Getenv("TERM") == "",
	}
}

// NewStyleWithNoColor creates a style with explicit color control.
// This is primarily useful for testing
func NewStyleWithNoColor(noColor bool) Gostyl {
	return Gostyl{
		codes:   []string{},
		noColor: noColor,
	}
}

func (g Gostyl) NoColor() bool {
	return g.noColor
}

// addCode return a new Gostyl with an additional ANSI code
func (g Gostyl) addCode(code string) Gostyl {
	codes := append([]string(nil), g.codes...)
	codes = append(codes, code)
	return Gostyl{
		codes:   codes,
		noColor: g.noColor,
	}
}

// Text styles

// Bold returns a style with bold text
func (g Gostyl) Bold() Gostyl {
	return g.addCode("1")
}

// Dim returns a style with dim text
func (g Gostyl) Dim() Gostyl {
	return g.addCode("2")
}

// Italic returns a style with italic text
func (g Gostyl) Italic() Gostyl {
	return g.addCode("3")
}

// Underline returns a style with underline text
func (g Gostyl) Underline() Gostyl {
	return g.addCode("4")
}

// Foreground colors

// Red returns a style with red text
func (g Gostyl) Red() Gostyl {
	return g.addCode("31")
}

// Green returns a style with green text
func (g Gostyl) Green() Gostyl {
	return g.addCode("32")
}

// Yellow returns a style with yellow text
func (g Gostyl) Yellow() Gostyl {
	return g.addCode("33")
}

// Cyan returns a style with cyan text
func (g Gostyl) Cyan() Gostyl {
	return g.addCode("36")
}

// Sprint applies the style to given text
func (g Gostyl) Sprint(text string) string {
	if len(g.codes) == 0 || g.noColor {
		return text
	}

	var sb strings.Builder
	sb.WriteString("\033[")
	sb.WriteString(g.codes[0])

	for _, code := range g.codes[1:] {
		sb.WriteString(";")
		sb.WriteString(code)
	}

	sb.WriteString("m")
	sb.WriteString(text)
	sb.WriteString("\033[0m")

	return sb.String()
}
