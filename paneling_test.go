package paneling

import (
	"strings"
	"testing"
)

func TestNewGrid(t *testing.T) {
	type args struct {
		width     int64
		height    int64
		direction direction
	}
	tests := []struct {
		name string
		args args
		want *Grid
	}{
		{
			name: "Create a vertical grid",
			args: args{
				width:     10,
				height:    5,
				direction: VERTICAL,
			},
			want: &Grid{
				Width:     10,
				Height:    5,
				Direction: VERTICAL,
			},
		},
		{
			name: "Create a horizontal grid",
			args: args{
				width:     20,
				height:    10,
				direction: HORIZONTAL,
			},
			want: &Grid{
				Width:     20,
				Height:    10,
				Direction: HORIZONTAL,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGrid(tt.args.width, tt.args.height, tt.args.direction)
			if got.Width != tt.want.Width || got.Height != tt.want.Height || got.Direction != tt.want.Direction {
				t.Errorf("NewGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddChild(t *testing.T) {
	parent := NewGrid(50, 50, VERTICAL)
	child := NewChild(25, 25, "Child content")

	type args struct {
		child *Grid
	}
	tests := []struct {
		name string
		grid *Grid
		args args
		want int // want represents the expected number of children after adding
	}{
		{
			name: "Add child to grid",
			grid: parent,
			args: args{
				child: child,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.grid.AddChild(tt.args.child)
			if got := len(tt.grid.children); got != tt.want {
				t.Errorf("AddChild() resulted in %d children, want %d", got, tt.want)
			}
		})
	}
}

func TestRender(t *testing.T) {
	// Setup for a basic parent-child structure
	parentV := NewGrid(100, 20, VERTICAL)
	child1 := NewChild(100, 10, "First child content")
	child2 := NewChild(100, 10, "Second child content")

	parentV.AddChild(child1).AddChild(child2)

	type wants struct {
		output string
	}
	tests := []struct {
		name string
		grid *Grid
		want wants
	}{
		{
			name: "Render VERTICAL grid with children",
			grid: parentV,
			want: wants{
				output: "First child content" + strings.Repeat("\n", 10) + "Second child content" +  strings.Repeat("\n", 9),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.grid.Render(); got != tt.want.output {
				t.Errorf("Render() got = %v, want %v", got, tt.want.output)
			}
		})
	}
}
