# Gostyl

A tiny, focused styling helper for Go.

## Feature

- 🎨Chain styles: `.Red().Bold().Underline()`
- 🔇Auto-respects `NO_COLOR` environment variable
- 📦Zero dependencies (stdlib only)

## Installation

```bash
go get github.com/matteo-gildone/gostyl
```

## Usage

```go
import "github.com/matteo-gildone/gostyl"

func main() {
	style := gostyl.NewStyle()
	
	// Simple colors
	fmt.Println(style.Red().Sprint("Error!"))
	
	// Chain styles
    fmt.Println(style.Green().Bold().Sprint("Success!"))
	
	// Multiple styles
	warning := style.Yellow().Bold()
    fmt.Println(warning.Sprint("Warning!"))
}
```
