# Paneling

`Paneling` is a super original Go package designed for creating a grids-of-grids layout. 

It does one thing, and does it well.

# Features

* Create grids with specified dimensions (width and height).
* Nest grids within grids to create complex layouts (vertical and horizontal).
* Add content to grids that automatically wraps to fit the grid size.
* That's it.

# Installation

To use `Paneling` in your project, install the package:

go get -u github.com/jacobkania/paneling

# Usage: Creating a Grid

### To create a new grid:

```go
import "github.com/jacobkania/paneling"

grid := paneling.NewGrid(screenWidth, screenHeight, paneling.VERTICAL)

child := paneling.NewChild(width, height, "Your content here")
grid.AddChild(child)
```

### To render the above grid and output its content:

```go
output := grid.Render()
fmt.Print(output)
```

### Example: Creating a Complex Layout

```go
// m is defined as a Model within the context of a bubble tea program

g := paneling.NewGrid(
	m.window.width,
	m.window.height,
	paneling.HORIZONTAL,
).AddChild(
	paneling.NewGrid(
		3,
		0,
		paneling.VERTICAL,
	).AddChild(
		&paneling.Grid{
			Height:  3,
			Content: chatHistoryText,
		},
	).AddChild(
		&paneling.Grid{
			Height:  1,
			Content: textboxText,
		},
	),
).AddChild(
	&paneling.Grid{
		Width:   1,
		Content: settingsText,
	},
)
```

# API documentation

## Struct: `Grid`

* `Width` (int): The width of the grid.
    * Default: 0
    * Required: No
    * For the top-level grid, this is the total width of the window that you would like to use.
    * For every child grid, this is a proportional value.
    * If there are two child grids, one with a width of 1 and the other with a width of 2, the first grid will take up 1/3 of the total width and the second grid will take up 2/3 of the total width.
* `Height` (int): The height of the grid.
    * Default: 0
    * Required: No
    * For the top-level grid, this is the total height of the window that you would like to use.
    * For every child grid, this is a proportional value.
    * If there are two child grids, one with a height of 1 and the other with a height of 2, the first grid will take up 1/3 of the total height and the second grid will take up 2/3 of the total height.
* `Content` (string): The content to be displayed within the grid.
    * Default: ""
    * Required: No
    * If the content is longer than the grid's width, it will be automatically wrapped to fit the grid.
    * Content will only be used if the grid has no children. Otherwise, this value is ignored.
* `Direction` (int): The direction in which the grid's children will be laid out.
    * Default: ""
    * Required: No
    * Options: `paneling.VERTICAL`, `paneling.HORIZONTAL`
    * If `paneling.VERTICAL`, the children will be laid out vertically.
    * If `paneling.HORIZONTAL`, the children will be laid out horizontally.
    * If the grid has no children, this value is ignored.
