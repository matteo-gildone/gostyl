package gostyl

import (
	"fmt"
	"os"
	"strings"
)

type Gostyl struct {
	codes   []string
	noColor bool
}

// NewStyle creates a new style with auto-detect color support
// Color is disabled if NO_COLOR env var is set or if TERM "dumb" or empty
func NewStyle() Gostyl {
	return Gostyl{
		codes: []string{},
		noColor: os.Getenv("NO_COLOR") != "" ||
			os.Getenv("TERM") == "dumb" ||
			os.Getenv("TERM") == "",
	}
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

// Strikethrough returns a style with strikethrough text
func (g Gostyl) Strikethrough() Gostyl {
	return g.addCode("9")
}

// Inverse returns a style with strikethrough text
func (g Gostyl) Inverse() Gostyl {
	return g.addCode("7")
}

// Foreground colors

// Black returns a style with black text
func (g Gostyl) Black() Gostyl {
	return g.addCode("30")
}

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

// Blue returns a style with blue text
func (g Gostyl) Blue() Gostyl {
	return g.addCode("34")
}

// Magenta returns a style with magenta text
func (g Gostyl) Magenta() Gostyl {
	return g.addCode("35")
}

// Cyan returns a style with cyan text
func (g Gostyl) Cyan() Gostyl {
	return g.addCode("36")
}

// White returns a style with white text
func (g Gostyl) White() Gostyl {
	return g.addCode("37")
}

// BrightBlack returns a style with bright black text
func (g Gostyl) BrightBlack() Gostyl {
	return g.addCode("90")
}

// BrightRed returns a style with bright red text
func (g Gostyl) BrightRed() Gostyl {
	return g.addCode("91")
}

// BrightGreen returns a style with bright green text
func (g Gostyl) BrightGreen() Gostyl {
	return g.addCode("92")
}

// BrightYellow returns a style with bright yellow text
func (g Gostyl) BrightYellow() Gostyl {
	return g.addCode("93")
}

// BrightBlue returns a style with bright blue text
func (g Gostyl) BrightBlue() Gostyl {
	return g.addCode("94")
}

// BrightMagenta returns a style with bright magenta text
func (g Gostyl) BrightMagenta() Gostyl {
	return g.addCode("95")
}

// BrightCyan returns a style with  bright cyan text
func (g Gostyl) BrightCyan() Gostyl {
	return g.addCode("96")
}

// BrightWhite returns a style with bright white text
func (g Gostyl) BrightWhite() Gostyl {
	return g.addCode("97")
}

// Background colors

// BgBlack returns a style with black background
func (g Gostyl) BgBlack() Gostyl {
	return g.addCode("40")
}

// BgRed returns a style with background red background
func (g Gostyl) BgRed() Gostyl {
	return g.addCode("41")
}

// BgGreen returns a style with background green background
func (g Gostyl) BgGreen() Gostyl {
	return g.addCode("42")
}

// BgYellow returns a style with yellow background
func (g Gostyl) BgYellow() Gostyl {
	return g.addCode("43")
}

// BgBlue returns a style with blue background
func (g Gostyl) BgBlue() Gostyl {
	return g.addCode("44")
}

// BgMagenta returns a style with magenta background
func (g Gostyl) BgMagenta() Gostyl {
	return g.addCode("45")
}

// BgCyan returns a style with cyan background
func (g Gostyl) BgCyan() Gostyl {
	return g.addCode("46")
}

// BgWhite returns a style with white background
func (g Gostyl) BgWhite() Gostyl {
	return g.addCode("47")
}

// BgBrightBlack returns a style with bright black background
func (g Gostyl) BgBrightBlack() Gostyl {
	return g.addCode("100")
}

// BgBrightRed returns a style with bright red background
func (g Gostyl) BgBrightRed() Gostyl {
	return g.addCode("101")
}

// BgBrightGreen returns a style with bright green background
func (g Gostyl) BgBrightGreen() Gostyl {
	return g.addCode("102")
}

// BgBrightYellow returns a style with bright yellow background
func (g Gostyl) BgBrightYellow() Gostyl {
	return g.addCode("103")
}

// BgBrightBlue returns a style with bright blue background
func (g Gostyl) BgBrightBlue() Gostyl {
	return g.addCode("104")
}

// BgBrightMagenta returns a style with bright magenta background
func (g Gostyl) BgBrightMagenta() Gostyl {
	return g.addCode("105")
}

// BgBrightCyan returns a style with  bright cyan background
func (g Gostyl) BgBrightCyan() Gostyl {
	return g.addCode("106")
}

// BgBrightWhite returns a style with bright white background
func (g Gostyl) BgBrightWhite() Gostyl {
	return g.addCode("107")
}

// apply applies the style to given text
func (g Gostyl) apply(text string) string {
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

func (g Gostyl) Sprint(a ...any) string {
	return g.apply(fmt.Sprint(a...))
}

func (g Gostyl) Sprintf(format string, a ...any) string {
	return g.apply(fmt.Sprintf(format, a...))
}

func (g Gostyl) Sprintln(a ...any) string {
	return g.apply(fmt.Sprintln(a...))
}
