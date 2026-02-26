package gostyl

import (
	"fmt"
	"os"
	"strings"
)

type Style struct {
	codes []string
}

// NewStyle creates a new style with auto-detect color support
// Color is disabled if NO_COLOR env var is set or if TERM "dumb" or empty
func NewStyle() Style {
	return Style{}
}

func colorDisabled() bool {
	return os.Getenv("NO_COLOR") != "" ||
		os.Getenv("TERM") == "dumb" ||
		os.Getenv("TERM") == ""
}

// addCode return a new Style with an additional ANSI code
func (s Style) addCode(code string) Style {
	// Return unchanged if code already applied to avoid duplicate ANSI codes
	for _, c := range s.codes {
		if c == code {
			return s
		}
	}

	codes := append([]string(nil), s.codes...)
	codes = append(codes, code)
	return Style{
		codes: codes,
	}
}

// Text styles

// Bold returns a style with bold text
func (s Style) Bold() Style {
	return s.addCode("1")
}

// Dim returns a style with dim text
func (s Style) Dim() Style {
	return s.addCode("2")
}

// Italic returns a style with italic text
func (s Style) Italic() Style {
	return s.addCode("3")
}

// Underline returns a style with underline text
func (s Style) Underline() Style {
	return s.addCode("4")
}

// Inverse returns a style with inverse text
func (s Style) Inverse() Style {
	return s.addCode("7")
}

// Strikethrough returns a style with strikethrough text
func (s Style) Strikethrough() Style {
	return s.addCode("9")
}

// Foreground colors

// Black returns a style with black text
func (s Style) Black() Style {
	return s.addCode("30")
}

// Red returns a style with red text
func (s Style) Red() Style {
	return s.addCode("31")
}

// Green returns a style with green text
func (s Style) Green() Style {
	return s.addCode("32")
}

// Yellow returns a style with yellow text
func (s Style) Yellow() Style {
	return s.addCode("33")
}

// Blue returns a style with blue text
func (s Style) Blue() Style {
	return s.addCode("34")
}

// Magenta returns a style with magenta text
func (s Style) Magenta() Style {
	return s.addCode("35")
}

// Cyan returns a style with cyan text
func (s Style) Cyan() Style {
	return s.addCode("36")
}

// White returns a style with white text
func (s Style) White() Style {
	return s.addCode("37")
}

// BrightBlack returns a style with bright black text
func (s Style) BrightBlack() Style {
	return s.addCode("90")
}

// BrightRed returns a style with bright red text
func (s Style) BrightRed() Style {
	return s.addCode("91")
}

// BrightGreen returns a style with bright green text
func (s Style) BrightGreen() Style {
	return s.addCode("92")
}

// BrightYellow returns a style with bright yellow text
func (s Style) BrightYellow() Style {
	return s.addCode("93")
}

// BrightBlue returns a style with bright blue text
func (s Style) BrightBlue() Style {
	return s.addCode("94")
}

// BrightMagenta returns a style with bright magenta text
func (s Style) BrightMagenta() Style {
	return s.addCode("95")
}

// BrightCyan returns a style with bright cyan text
func (s Style) BrightCyan() Style {
	return s.addCode("96")
}

// BrightWhite returns a style with bright white text
func (s Style) BrightWhite() Style {
	return s.addCode("97")
}

// Background colors

// BgBlack returns a style with black background
func (s Style) BgBlack() Style {
	return s.addCode("40")
}

// BgRed returns a style with red background
func (s Style) BgRed() Style {
	return s.addCode("41")
}

// BgGreen returns a style with green background
func (s Style) BgGreen() Style {
	return s.addCode("42")
}

// BgYellow returns a style with yellow background
func (s Style) BgYellow() Style {
	return s.addCode("43")
}

// BgBlue returns a style with blue background
func (s Style) BgBlue() Style {
	return s.addCode("44")
}

// BgMagenta returns a style with magenta background
func (s Style) BgMagenta() Style {
	return s.addCode("45")
}

// BgCyan returns a style with cyan background
func (s Style) BgCyan() Style {
	return s.addCode("46")
}

// BgWhite returns a style with white background
func (s Style) BgWhite() Style {
	return s.addCode("47")
}

// BgBrightBlack returns a style with bright black background
func (s Style) BgBrightBlack() Style {
	return s.addCode("100")
}

// BgBrightRed returns a style with bright red background
func (s Style) BgBrightRed() Style {
	return s.addCode("101")
}

// BgBrightGreen returns a style with bright green background
func (s Style) BgBrightGreen() Style {
	return s.addCode("102")
}

// BgBrightYellow returns a style with bright yellow background
func (s Style) BgBrightYellow() Style {
	return s.addCode("103")
}

// BgBrightBlue returns a style with bright blue background
func (s Style) BgBrightBlue() Style {
	return s.addCode("104")
}

// BgBrightMagenta returns a style with bright magenta background
func (s Style) BgBrightMagenta() Style {
	return s.addCode("105")
}

// BgBrightCyan returns a style with bright cyan background
func (s Style) BgBrightCyan() Style {
	return s.addCode("106")
}

// BgBrightWhite returns a style with bright white background
func (s Style) BgBrightWhite() Style {
	return s.addCode("107")
}

// apply applies the style to given text
func (s Style) apply(text string) string {
	if len(s.codes) == 0 || colorDisabled() {
		return text
	}

	var sb strings.Builder
	sb.WriteString("\033[")
	sb.WriteString(s.codes[0])

	for _, code := range s.codes[1:] {
		sb.WriteString(";")
		sb.WriteString(code)
	}

	sb.WriteString("m")
	sb.WriteString(text)
	sb.WriteString("\033[0m")

	return sb.String()
}

func (s Style) Sprint(a ...any) string {
	return s.apply(fmt.Sprint(a...))
}

func (s Style) Sprintf(format string, a ...any) string {
	return s.apply(fmt.Sprintf(format, a...))
}

func (s Style) Sprintln(a ...any) string {
	return s.apply(strings.TrimSuffix(fmt.Sprintln(a...), "\n")) + "\n"
}
