package gostyl

func Success(s string) string {
	return NewStyle().Green().Sprint(s)
}

func Successf(format string, a ...any) string {
	return NewStyle().Green().Sprintf(format, a)
}

func SuccessWithIcon(s string) string {
	return NewStyle().Green().Sprint("✓ " + s)
}

func SuccessWithIconf(format string, a ...any) string {
	return NewStyle().Green().Sprintf("✓ "+format, a)
}
