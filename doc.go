// Package gostyl provides a tiny, zero dependency ANSI styling helper for Go CLI applications
//
// Color output is automatically disabled when NO_COLOR environment variable is set or when TERM is "dumb" or empty.
//
// Basic usage:
//
//	s := gostyl.NewStyle()
//	fmt.Println(s.Bold().Red().Sprint("error"))
//	fmt.Println(gostyl.Danger().Sprint("something went wrong"))
package gostyl
