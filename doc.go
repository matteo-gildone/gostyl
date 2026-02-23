// Package gostyl provide a tiny, zero dependency ANSI styling helper for Go CLI applications
//
// Color output is automatically disabled when NO_COLOR environment variable is set or when TERM is "dump" or empty.
//
// Basic usage:
//
//	s := gostyl.NewStyle()
//	fmt.Prinln(s.Bold().Red().Sprint("error"))
//	fmt.Prinln(gostyl.Danger().Sprint("something went wrong"))
package gostyl
