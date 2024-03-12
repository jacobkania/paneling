package paneling

import (
	"fmt"
	"strings"
)

type direction string

const (
	VERTICAL   direction = "vertical"
	HORIZONTAL direction = "horizontal"
)

type Grid struct {
	Width     int64 // absolute on parent, relative on children
	Height    int64 // absolute on parent, relative on children
	Direction direction

	Content string

	children []*Grid
}

func (g *Grid) safeWidth() int64 {
	if g.Width <= 0 {
		return 1
	}

	return g.Width
}

func (g *Grid) safeHeight() int64 {
	if g.Height <= 0 {
		return 1
	}

	return g.Height
}

// NewGrid creates a new grid with the given width, height and direction
// This type of grid is used for containing children
func NewGrid(width, height int64, direction direction) *Grid {
	return &Grid{
		Width:     width,
		Height:    height,
		Direction: direction,
	}
}

// NewChild creates a new grid with the given width, height and content
// This type of grid is used for containing content
func NewChild(width, height int64, content string) *Grid {
	return &Grid{
		Width:   width,
		Height:  height,
		Content: content,
	}
}

func (g *Grid) AddChild(child *Grid) *Grid {
	g.children = append(g.children, child)

	return g
}

func (g *Grid) Render() string {
	return strings.Join(g.render(0, 0), "\n")
}

func (g *Grid) render(widthConstraint, heightConstraint int64) []string {
	if len(g.children) == 0 {
		return g.renderContent(widthConstraint, heightConstraint)
	}

	return g.renderChildren(widthConstraint, heightConstraint)
}

func (g *Grid) renderContent(widthConstraint, heightConstraint int64) []string {
	availableHeight := fallbackOnZero(heightConstraint, g.safeHeight())
	availableWidth := fallbackOnZero(widthConstraint, g.safeWidth())

	lines := []string{}

	pendingContent := strings.Split(g.Content, "\n")
	for _, line := range pendingContent {
		if int64(len(line)) > availableWidth {
			lines = append(lines, SplitLongLine(line, int(availableWidth))...)
		} else {
			lines = append(lines, line)
		}
	}

	// fill the remaining space with empty lines
	for i := int64(len(lines)); i < availableHeight; i++ {
		lines = append(lines, "")
	}

	// return only the lines that fit the actual height
	return lines[:availableHeight]
}

func (g *Grid) renderChildren(widthConstraint, heightConstraint int64) []string {
	// TODO: I can remove this fallbackOnZero and just use the constraints
	//		THEN, I can move the g.Height and g.Width checks up to the Render()
	//		since they're only used the first time
	availableHeight := fallbackOnZero(heightConstraint, g.safeHeight())
	availableWidth := fallbackOnZero(widthConstraint, g.safeWidth())

	lines := []string{}

	if g.Direction == VERTICAL {
		totalUsedHeight := int64(0)
		for i, child := range g.children {
			proportionalHeight := int64((float64(child.safeHeight()) / float64(g.totalChildrenHeight())) * float64(availableHeight))

			// ensure the last child gets the remaining space if division is not exact
			totalUsedHeight += proportionalHeight
			if i == len(g.children)-1 {
				proportionalHeight += availableHeight - totalUsedHeight
			}

			childLines := child.render(availableWidth, proportionalHeight)

			lines = append(lines, childLines...)
		}
	}

	if g.Direction == HORIZONTAL {
		childLines := []childLine{}

		// lines = append(lines

		for _, child := range g.children {
			proportionalWidth := int64((float64(child.safeWidth()) / float64(g.totalChildrenWidth())) * float64(availableWidth))

			childLine := childLine{
				Content: child.render(proportionalWidth, availableHeight),
				Width:   proportionalWidth,
			}

			childLines = append(childLines, childLine)
		}

		for i := int64(0); i < availableHeight; i++ {
			line := ""
			for _, child := range childLines {
				line += fmt.Sprintf("%-*v", child.Width, child.Content[i])
			}
			lines = append(lines, line)
		}
	}

	// return only the lines that fit the actual height
	return lines[:availableHeight]
}

type childLine struct {
	Content []string
	Width   int64
}

func (g *Grid) totalChildrenHeight() int64 {
	var total int64
	for _, child := range g.children {
		total += child.safeHeight()
	}

	return total
}

func (g *Grid) totalChildrenWidth() int64 {
	var total int64
	for _, child := range g.children {
		total += child.safeWidth()
	}

	return total
}

func fallbackOnZero(a, b int64) int64 {
	if a == 0 {
		return b
	}
	return a
}
