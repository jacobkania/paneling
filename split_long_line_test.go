package paneling

import "testing"

func TestSplitLongLine(t *testing.T) {
	type args struct {
		line  string
		width int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Split a long line into smaller lines",
			args: args{
				line:  "This is a very long line that needs to be split into smaller lines.",
				width: 10,
			},
			want: []string{"This is a", "very long", "line that", "needs to", "be split", "into", "smaller", "lines."},
		},
		{
			name: "Handle a really long word",
			args: args{
				line:  "Thisisaverylongwordthatislongerthantheallowedwidth",
				width: 10,
			},
			want: []string{"Thisisaver"},
		},
		{
			name: "Handle a single word",
			args: args{
				line:  "Hello",
				width: 10,
			},
			want: []string{"Hello"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitLongLine(tt.args.line, tt.args.width); !compareStringSlices(got, tt.want) {
				t.Errorf("SplitLongLine() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// compareStringSlices is a helper function to compare two slices of strings for equality.
func compareStringSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
