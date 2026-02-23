# Gostyl

A tiny, focused styling helper for Go.

## Features

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

## Presets

```go

gostyl.Success("task done")
gostyl.Successf("%d: %s", 1, "task done")
gostyl.Successln("task done")

gostyl.Warning("check this")
gostyl.Warningf("%s", "check this")
gostyl.Warningln("check this")

gostyl.Danger("something failed")
gostyl.Dangerf("%s", "something failed")
gostyl.Dangerln("something failed")

gostyl.Info("something to notice")
gostyl.Infof("%s", "something to notice")
gostyl.Infoln("something to notice")

gostyl.Muted("something not really important")
gostyl.Mutedf("%s", "something not really important")
gostyl.Mutedln("something not really important")
```
