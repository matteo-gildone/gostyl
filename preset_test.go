package gostyl

import "testing"

func TestStyles_Presets(t *testing.T) {
	t.Setenv("TERM", "xterm-256color")
	tests := []struct {
		name   string
		preset func(string) string
		input  string
		want   string
	}{
		{
			name:   "success style",
			preset: Success,
			input:  "task done",
			want:   "\033[92m✓ task done\033[0m",
		},
		{
			name:   "warning style",
			preset: Warning,
			input:  "check this",
			want:   "\033[93m! check this\033[0m",
		},
		{
			name:   "danger style",
			preset: Danger,
			input:  "something failed",
			want:   "\033[91mx something failed\033[0m",
		},
		{
			name:   "info style",
			preset: Info,
			input:  "something to notice",
			want:   "\033[96mi something to notice\033[0m",
		},
		{
			name:   "muted style",
			preset: Muted,
			input:  "something not really important",
			want:   "\033[90msomething not really important\033[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.preset(tt.input)

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}

}

func TestStyles_PresetsNoColor(t *testing.T) {
	t.Setenv("TERM", "")
	tests := []struct {
		name   string
		preset func(string) string
		input  string
		want   string
	}{
		{
			name:   "success style",
			preset: Success,
			input:  "task done",
			want:   "✓ task done",
		},
		{
			name:   "warning style",
			preset: Warning,
			input:  "check this",
			want:   "! check this",
		},
		{
			name:   "danger style",
			preset: Danger,
			input:  "something failed",
			want:   "x something failed",
		},
		{
			name:   "info style",
			preset: Info,
			input:  "something to notice",
			want:   "i something to notice",
		},
		{
			name:   "muted style",
			preset: Muted,
			input:  "something not really important",
			want:   "something not really important",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.preset(tt.input)

			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
