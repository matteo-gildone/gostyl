package styles

import "testing"

func TestStyles_New(t *testing.T) {
	style := NewStyle()

	if len(style.codes) != 0 {
		t.Errorf("expected empty list, got length %d", len(style.codes))
	}
}

func TestStyles_NewWithNoColor(t *testing.T) {
	tests := []struct {
		name    string
		noColor bool
	}{
		{
			name:    "with color",
			noColor: false,
		},
		{
			name:    "without color",
			noColor: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			style := NewStyleWithNoColor(tt.noColor)

			if style.NoColor() != tt.noColor {
				t.Errorf("got %v, want %v", style.NoColor(), tt.noColor)
			}
		})
	}

}

func TestStyle_NoColor(t *testing.T) {
	tests := []struct {
		name        string
		style       Gostyl
		wantNoColor bool
	}{
		{
			name:        "explicit color enabled",
			style:       NewStyleWithNoColor(false),
			wantNoColor: false,
		},
		{
			name:        "explicit color disabled",
			style:       NewStyleWithNoColor(true),
			wantNoColor: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.style.NoColor(); got != tt.wantNoColor {
				t.Errorf("got %v, want %v", got, tt.wantNoColor)
			}
		})
	}
}

func TestStylesWithColor_Sprint(t *testing.T) {
	tests := []struct {
		name  string
		style Gostyl
		input string
		want  string
	}{
		{
			name:  "base style",
			style: NewStyleWithNoColor(false),
			input: "base style",
			want:  "base style",
		},
		{
			name:  "bold style",
			input: "bold style",
			style: NewStyleWithNoColor(false).Bold(),
			want:  "\033[1mbold style\033[0m",
		},
		{
			name:  "dim style",
			input: "dim style",
			style: NewStyleWithNoColor(false).Dim(),
			want:  "\033[2mdim style\033[0m",
		},
		{
			name:  "underline style",
			input: "underline style",
			style: NewStyleWithNoColor(false).Underline(),
			want:  "\033[4munderline style\033[0m",
		},
		{
			name:  "italic style",
			input: "italic style",
			style: NewStyleWithNoColor(false).Italic(),
			want:  "\033[3mitalic style\033[0m",
		},
		{
			name:  "red style",
			input: "red style",
			style: NewStyleWithNoColor(false).Red(),
			want:  "\033[31mred style\033[0m",
		},
		{
			name:  "green style",
			input: "green style",
			style: NewStyleWithNoColor(false).Green(),
			want:  "\033[32mgreen style\033[0m",
		},
		{
			name:  "yellow style",
			input: "yellow style",
			style: NewStyleWithNoColor(false).Yellow(),
			want:  "\033[33myellow style\033[0m",
		},
		{
			name:  "cyan style",
			input: "cyan style",
			style: NewStyleWithNoColor(false).Cyan(),
			want:  "\033[36mcyan style\033[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.style.Sprint(tt.input)

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}

}

func TestStylesWithNoColor_Sprint(t *testing.T) {
	tests := []struct {
		name  string
		style Gostyl
		input string
		want  string
	}{
		{
			name:  "base style",
			style: NewStyleWithNoColor(true),
			input: "base style",
			want:  "base style",
		},
		{
			name:  "bold style",
			input: "bold style",
			style: NewStyleWithNoColor(true).Bold(),
			want:  "bold style",
		},
		{
			name:  "dim style",
			input: "dim style",
			style: NewStyleWithNoColor(true).Dim(),
			want:  "dim style",
		},
		{
			name:  "underline style",
			input: "underline style",
			style: NewStyleWithNoColor(true).Underline(),
			want:  "underline style",
		},
		{
			name:  "italic style",
			input: "italic style",
			style: NewStyleWithNoColor(true).Italic(),
			want:  "italic style",
		},
		{
			name:  "red style",
			input: "red style",
			style: NewStyleWithNoColor(true).Red(),
			want:  "red style",
		},
		{
			name:  "green style",
			input: "green style",
			style: NewStyleWithNoColor(true).Green(),
			want:  "green style",
		},
		{
			name:  "yellow style",
			input: "yellow style",
			style: NewStyleWithNoColor(true).Yellow(),
			want:  "yellow style",
		},
		{
			name:  "cyan style",
			input: "cyan style",
			style: NewStyleWithNoColor(true).Cyan(),
			want:  "cyan style",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.style.Sprint(tt.input)

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}

}

func TestStyleChaining(t *testing.T) {
	tests := []struct {
		name    string
		style   Gostyl
		input   string
		noColor bool
		want    string
	}{
		{
			name:  "base style unchanged",
			style: NewStyleWithNoColor(false),
			input: "text",
			want:  "text",
		},
		{
			name:  "red style independent",
			style: NewStyleWithNoColor(false).Red(),
			input: "text",
			want:  "\033[31mtext\033[0m",
		},
		{
			name:  "multiple styles independent",
			style: NewStyleWithNoColor(false).Red().Bold(),
			input: "text",
			want:  "\033[31;1mtext\033[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.style.Sprint(tt.input)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
