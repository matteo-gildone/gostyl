package gostyl

import "testing"

func TestStyles_New(t *testing.T) {
	style := NewStyle()

	if len(style.codes) != 0 {
		t.Errorf("expected empty list, got length %d", len(style.codes))
	}
}

func TestStyle_NewNoColor(t *testing.T) {
	tests := []struct {
		name  string
		setup func(t *testing.T)
		want  bool
	}{
		{
			name:  "NO_COLOR set",
			setup: func(t *testing.T) { t.Setenv("NO_COLOR", "1") },
			want:  true,
		},
		{
			name:  "TERM is dumb",
			setup: func(t *testing.T) { t.Setenv("TERM", "dumb") },
			want:  true,
		},
		{
			name:  "TERM is empty",
			setup: func(t *testing.T) { t.Setenv("TERM", "") },
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(t)
			got := NewStyle()

			if got.noColor != tt.want {
				t.Errorf("got %v, want %v", got.noColor, tt.want)
			}
		})
	}
}

func TestStyles_Sprint(t *testing.T) {
	tests := []struct {
		name  string
		style Gostyl
		input string
		want  string
	}{
		{
			name:  "base style",
			style: NewStyle(),
			input: "base style",
			want:  "base style",
		},
		{
			name:  "bold style",
			input: "bold style",
			style: NewStyle().Bold(),
			want:  "\033[1mbold style\033[0m",
		},
		{
			name:  "dim style",
			input: "dim style",
			style: NewStyle().Dim(),
			want:  "\033[2mdim style\033[0m",
		},
		{
			name:  "italic style",
			input: "italic style",
			style: NewStyle().Italic(),
			want:  "\033[3mitalic style\033[0m",
		},
		{
			name:  "underline style",
			input: "underline style",
			style: NewStyle().Underline(),
			want:  "\033[4munderline style\033[0m",
		},
		{
			name:  "inverse style",
			input: "inverse style",
			style: NewStyle().Inverse(),
			want:  "\033[7minverse style\033[0m",
		},
		{
			name:  "strikethrough style",
			input: "strikethrough style",
			style: NewStyle().Strikethrough(),
			want:  "\033[9mstrikethrough style\033[0m",
		},
		{
			name:  "black style",
			input: "black style",
			style: NewStyle().Black(),
			want:  "\033[30mblack style\033[0m",
		},
		{
			name:  "red style",
			input: "red style",
			style: NewStyle().Red(),
			want:  "\033[31mred style\033[0m",
		},
		{
			name:  "green style",
			input: "green style",
			style: NewStyle().Green(),
			want:  "\033[32mgreen style\033[0m",
		},
		{
			name:  "yellow style",
			input: "yellow style",
			style: NewStyle().Yellow(),
			want:  "\033[33myellow style\033[0m",
		},
		{
			name:  "blue style",
			input: "blue style",
			style: NewStyle().Blue(),
			want:  "\033[34mblue style\033[0m",
		},
		{
			name:  "magenta style",
			input: "magenta style",
			style: NewStyle().Magenta(),
			want:  "\033[35mmagenta style\033[0m",
		},
		{
			name:  "cyan style",
			input: "cyan style",
			style: NewStyle().Cyan(),
			want:  "\033[36mcyan style\033[0m",
		},
		{
			name:  "white style",
			input: "white style",
			style: NewStyle().White(),
			want:  "\033[37mwhite style\033[0m",
		},
		{
			name:  "bright black style",
			input: "bright black style",
			style: NewStyle().BrightBlack(),
			want:  "\033[90mbright black style\033[0m",
		},
		{
			name:  "bright red style",
			input: "bright red style",
			style: NewStyle().BrightRed(),
			want:  "\033[91mbright red style\033[0m",
		},
		{
			name:  "bright green style",
			input: "bright green style",
			style: NewStyle().BrightGreen(),
			want:  "\033[92mbright green style\033[0m",
		},
		{
			name:  "bright yellow style",
			input: "bright yellow style",
			style: NewStyle().BrightYellow(),
			want:  "\033[93mbright yellow style\033[0m",
		},
		{
			name:  "bright blue style",
			input: "bright blue style",
			style: NewStyle().BrightBlue(),
			want:  "\033[94mbright blue style\033[0m",
		},
		{
			name:  "bright magenta style",
			input: "bright magenta style",
			style: NewStyle().BrightMagenta(),
			want:  "\033[95mbright magenta style\033[0m",
		},
		{
			name:  "bright cyan style",
			input: "bright cyan style",
			style: NewStyle().BrightCyan(),
			want:  "\033[96mbright cyan style\033[0m",
		},
		{
			name:  "bright white style",
			input: "bright white style",
			style: NewStyle().BrightWhite(),
			want:  "\033[97mbright white style\033[0m",
		},
		{
			name:  "black background style",
			input: "black background style",
			style: NewStyle().BgBlack(),
			want:  "\033[40mblack background style\033[0m",
		},
		{
			name:  "red background style",
			input: "red background style",
			style: NewStyle().BgRed(),
			want:  "\033[41mred background style\033[0m",
		},
		{
			name:  "green background style",
			input: "green background style",
			style: NewStyle().BgGreen(),
			want:  "\033[42mgreen background style\033[0m",
		},
		{
			name:  "yellow background style",
			input: "yellow background style",
			style: NewStyle().BgYellow(),
			want:  "\033[43myellow background style\033[0m",
		},
		{
			name:  "blue background style",
			input: "blue background style",
			style: NewStyle().BgBlue(),
			want:  "\033[44mblue background style\033[0m",
		},
		{
			name:  "magenta background style",
			input: "magenta background style",
			style: NewStyle().BgMagenta(),
			want:  "\033[45mmagenta background style\033[0m",
		},
		{
			name:  "cyan background style",
			input: "cyan background style",
			style: NewStyle().BgCyan(),
			want:  "\033[46mcyan background style\033[0m",
		},
		{
			name:  "white background style",
			input: "white background style",
			style: NewStyle().BgWhite(),
			want:  "\033[47mwhite background style\033[0m",
		},
		{
			name:  "bright black background style",
			input: "bright black background style",
			style: NewStyle().BgBrightBlack(),
			want:  "\033[100mbright black background style\033[0m",
		},
		{
			name:  "bright red background style",
			input: "bright red background style",
			style: NewStyle().BgBrightRed(),
			want:  "\033[101mbright red background style\033[0m",
		},
		{
			name:  "bright green background style",
			input: "bright green background style",
			style: NewStyle().BgBrightGreen(),
			want:  "\033[102mbright green background style\033[0m",
		},
		{
			name:  "bright yellow background style",
			input: "bright yellow background style",
			style: NewStyle().BgBrightYellow(),
			want:  "\033[103mbright yellow background style\033[0m",
		},
		{
			name:  "bright blue background style",
			input: "bright blue background style",
			style: NewStyle().BgBrightBlue(),
			want:  "\033[104mbright blue background style\033[0m",
		},
		{
			name:  "bright magenta background style",
			input: "bright magenta background style",
			style: NewStyle().BgBrightMagenta(),
			want:  "\033[105mbright magenta background style\033[0m",
		},
		{
			name:  "bright cyan background style",
			input: "bright cyan background style",
			style: NewStyle().BgBrightCyan(),
			want:  "\033[106mbright cyan background style\033[0m",
		},
		{
			name:  "bright white background style",
			input: "bright white background style",
			style: NewStyle().BgBrightWhite(),
			want:  "\033[107mbright white background style\033[0m",
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

func TestStyles_Sprintf(t *testing.T) {
	tests := []struct {
		name   string
		style  Gostyl
		format string
		input  string
		want   string
	}{
		{
			name:   "base style",
			style:  NewStyle(),
			input:  "base style with format",
			format: "%s",
			want:   "base style with format",
		},
		{
			name:   "bold style",
			input:  "bold style with format",
			format: "%s",
			style:  NewStyle().Bold(),
			want:   "\033[1mbold style with format\033[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.style.Sprintf(tt.format, tt.input)

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}

}

func TestStyles_Sprintln(t *testing.T) {
	tests := []struct {
		name  string
		style Gostyl
		input string
		want  string
	}{
		{
			name:  "base style",
			style: NewStyle(),
			input: "base style",
			want:  "base style\n",
		},
		{
			name:  "bold style",
			input: "bold style",
			style: NewStyle().Bold(),
			want:  "\033[1mbold style\033[0m\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.style.Sprintln(tt.input)

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}

}

func TestStyleChaining(t *testing.T) {
	tests := []struct {
		name  string
		style Gostyl
		input string
		want  string
	}{
		{
			name:  "base style unchanged",
			style: NewStyle(),
			input: "text",
			want:  "text",
		},
		{
			name:  "red style independent",
			style: NewStyle().Red(),
			input: "text",
			want:  "\033[31mtext\033[0m",
		},
		{
			name:  "multiple styles independent",
			style: NewStyle().Red().Bold(),
			input: "text",
			want:  "\033[31;1mtext\033[0m",
		},
		{
			name:  "multiple styles duplicated",
			style: NewStyle().Red().Bold().Red().Bold(),
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
