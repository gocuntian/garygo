package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Widget struct {
	ID    int64
	Name  string
	Score float32
	Added string
}

func main() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("ID", "Name", "Score", "Added")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	p := &Widget{}
	//&{0  0 }
	fmt.Println(p)
	p2 := new(Widget)
	fmt.Println(p2.Added)
	fmt.Println(p2)
	//&{0  0 }

	widget := Widget{
		ID:    100,
		Name:  "xingcuntian",
		Score: 90.09,
		Added: "2017-06-28",
	}
	fmt.Println(widget)
	//for _, widget := range getWidgets() {
	tbl.AddRow(widget.ID, widget.Name, widget.Score, widget.Added)
	//}

	tbl.Print()
}
