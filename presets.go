package gostyl

func Success(s string) string {
	return NewStyle().BrightGreen().Sprint("✓ " + s)
}

func Successf(format string, a ...any) string {
	return NewStyle().BrightGreen().Sprintf("✓ "+format, a...)
}

func Successln(s string) string {
	return NewStyle().BrightGreen().Sprintln("✓ " + s)
}

func Warning(s string) string {
	return NewStyle().BrightYellow().Sprint("! " + s)
}

func Warningf(format string, a ...any) string {
	return NewStyle().BrightYellow().Sprintf("! "+format, a...)
}

func Warningln(s string) string {
	return NewStyle().BrightYellow().Sprintln("! " + s)
}

func Danger(s string) string {
	return NewStyle().BrightRed().Sprint("x " + s)
}

func Dangerf(format string, a ...any) string {
	return NewStyle().BrightRed().Sprintf("x "+format, a...)
}

func Dangerln(s string) string {
	return NewStyle().BrightRed().Sprintln("x " + s)
}

func Info(s string) string {
	return NewStyle().BrightCyan().Sprint("i " + s)
}

func Infof(format string, a ...any) string {
	return NewStyle().BrightCyan().Sprintf("i "+format, a...)
}

func Infoln(s string) string {
	return NewStyle().BrightCyan().Sprintln("i " + s)
}

func Muted(s string) string {
	return NewStyle().BrightBlack().Sprint(s)
}

func Mutedf(format string, a ...any) string {
	return NewStyle().BrightBlack().Sprintf(format, a...)
}

func Mutedln(s string) string {
	return NewStyle().BrightBlack().Sprintln(s)
}
