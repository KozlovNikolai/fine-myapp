package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var data = [][]string{{"top left", "top right"},
	{"bottom left", "bottom right"}}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")

	/* 	list := widget.NewTable(
	func() (int, int) {
		return len(data), len(data[0])
	},
	func() fyne.CanvasObject {
		return widget.NewLabel("wide content")
	},
	func(i widget.TableCellID, o fyne.CanvasObject) {
		o.(*widget.Label).SetText(data[i.Row][i.Col])
	}) */
	headers := []string{"Имя", "Аватар"}
	employees := []struct {
		Name   string
		Avatar string
	}{
		{"Иван Петров", "👨‍💼"},
		{"Мария Сидорова", "👩‍🔬"},
		{"Алексей Иванов", "👨‍🎓"},
		{"Елена Смирнова", "👩‍💻"},
	}
	rowCount := len(employees) + 1

	table := widget.NewTable(
		func() (int, int) { return rowCount, 2 },
		func() fyne.CanvasObject { return widget.NewLabel("") },
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			label := co.(*widget.Label)
			if tci.Row == 0 {
				label.SetText(headers[tci.Col])
				label.TextStyle = fyne.TextStyle{Bold: true}
			} else {
				emp := employees[tci.Row-1]
				if tci.Col == 0 {
					label.SetText(emp.Name)
				} else {
					label.SetText(emp.Avatar)
				}
				label.TextStyle = fyne.TextStyle{}
			}
		},
	)
	table.SetColumnWidth(0, 200)
	table.SetColumnWidth(1, 100)

	mainContainer := container.New(
		layout.NewStackLayout(),
		table,
	)

	myWindow.SetContent(mainContainer)
	myWindow.ShowAndRun()
}
