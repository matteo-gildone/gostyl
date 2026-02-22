package gostyl

import "testing"

func TestStyles_New(t *testing.T) {
	style := NewStyle()

	if len(style.codes) != 0 {
		t.Errorf("expected empty list, got length %d", len(style.codes))
	}
}

func TestStyles(t *testing.T) {
	tests := []struct {
		name         string
		style        Gostyl
		input        string
		wantSprint   string
		wantSprintf  string
		wantSprintln string
	}{
		{
			name:       "base style",
			style:      NewStyle(),
			input:      "base style",
			wantSprint: "base style",
		},
		{
			name:       "bold style",
			input:      "bold style",
			style:      NewStyle().Bold(),
			wantSprint: "\033[1mbold style\033[0m",
		},
		{
			name:       "dim style",
			input:      "dim style",
			style:      NewStyle().Dim(),
			wantSprint: "\033[2mdim style\033[0m",
		},
		{
			name:       "italic style",
			input:      "italic style",
			style:      NewStyle().Italic(),
			wantSprint: "\033[3mitalic style\033[0m",
		},
		{
			name:       "underline style",
			input:      "underline style",
			style:      NewStyle().Underline(),
			wantSprint: "\033[4munderline style\033[0m",
		},
		{
			name:       "inverse style",
			input:      "inverse style",
			style:      NewStyle().Inverse(),
			wantSprint: "\033[7minverse style\033[0m",
		},
		{
			name:       "strikethrough style",
			input:      "strikethrough style",
			style:      NewStyle().Strikethrough(),
			wantSprint: "\033[9mstrikethrough style\033[0m",
		},
		{
			name:       "red style",
			input:      "red style",
			style:      NewStyle().Red(),
			wantSprint: "\033[31mred style\033[0m",
		},
		{
			name:       "green style",
			input:      "green style",
			style:      NewStyle().Green(),
			wantSprint: "\033[32mgreen style\033[0m",
		},
		{
			name:       "yellow style",
			input:      "yellow style",
			style:      NewStyle().Yellow(),
			wantSprint: "\033[33myellow style\033[0m",
		},
		{
			name:       "cyan style",
			input:      "cyan style",
			style:      NewStyle().Cyan(),
			wantSprint: "\033[36mcyan style\033[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.style.Sprint(tt.input)

			if got != tt.wantSprint {
				t.Errorf("got %q, want %q", got, tt.wantSprint)
			}
		})
	}

}

func TestStyles_Sprintf(t *testing.T) {
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
			want:  "base style with format",
		},
		{
			name:  "bold style",
			input: "bold style",
			style: NewStyle().Bold(),
			want:  "\033[1mbold style with format\033[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.style.Sprintf("%s %s", tt.input, "with format")

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
		name    string
		style   Gostyl
		input   string
		noColor bool
		want    string
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
