package main

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tabbed Application")
	w.Resize(fyne.NewSize(600, 500))

	// Вкладка 1: Таблица с аватарами
	tab1 := createTab1()

	// Вкладка 2: Калькулятор и график
	tab2 := createTab2()

	// Вкладка 3: Загрузка файла
	tab3 := createTab3(w)

	// Создаем вкладки
	tabs := container.NewAppTabs(
		container.NewTabItem("Сотрудники", tab1),
		container.NewTabItem("График", tab2),
		container.NewTabItem("Загрузка", tab3),
	)

	// Основной контейнер (часы + вкладки)
	mainContainer := container.New(
		layout.NewStackLayout(),
		tabs,
	)

	w.SetContent(mainContainer)
	w.ShowAndRun()
}
func createTab1() fyne.CanvasObject {
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

	// Теперь table будет РАСТЯГИВАТЬСЯ на всю высоту контейнера!
	main := container.NewBorder(
		widget.NewLabelWithStyle(
			"Список сотрудников", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		nil, nil, nil,
		container.NewStack(table),
		// table,
	)
	return main
}
func createTab2() *fyne.Container {
	entryK := widget.NewEntry()
	entryB := widget.NewEntry()
	resultLabel := widget.NewLabel("")
	graphContainer := container.New(layout.NewCenterLayout())

	calculate := func() {
		k, err1 := strconv.ParseFloat(entryK.Text, 64)
		b, err2 := strconv.ParseFloat(entryB.Text, 64)

		if err1 != nil || err2 != nil {
			resultLabel.SetText("Ошибка ввода чисел")
			return
		}

		sum := k + b
		resultLabel.SetText(fmt.Sprintf("Сумма: %.2f", sum))
		drawGraph(graphContainer, k, b)
	}

	form := container.NewVBox(
		widget.NewLabel("y = kx + b"),
		container.NewHBox(
			widget.NewLabel("k:"),
			entryK,
			widget.NewLabel("b:"),
			entryB,
		),
		widget.NewButton("Сумма и график", calculate),
		resultLabel,
	)

	return container.NewBorder(
		form,
		nil, nil, nil,
		graphContainer,
	)
}

func drawGraph(container *fyne.Container, k, b float64) {
	container.RemoveAll()
	width := 400.0
	height := 300.0
	padding := 30.0

	// Создаем оси координат
	axes := canvas.NewRectangle(color.Transparent)
	axes.Resize(fyne.NewSize(float32(width), float32(height)))

	// Рисуем оси
	xAxis := canvas.NewLine(color.Black)
	xAxis.Position1 = fyne.NewPos(float32(padding), float32(height/2))
	xAxis.Position2 = fyne.NewPos(float32(width-padding), float32(height/2))

	yAxis := canvas.NewLine(color.Black)
	yAxis.Position1 = fyne.NewPos(float32(width/2), float32(padding))
	yAxis.Position2 = fyne.NewPos(float32(width/2), float32(height-padding))

	// Подписи осей
	xLabel := canvas.NewText("X", color.Black)
	xLabel.Move(fyne.NewPos(float32(width-padding+5), float32(height/2-15)))
	yLabel := canvas.NewText("Y", color.Black)
	yLabel.Move(fyne.NewPos(float32(width/2+5), float32(padding-20)))

	container.Add(axes)
	container.Add(xAxis)
	container.Add(yAxis)
	container.Add(xLabel)
	container.Add(yLabel)

	// Рассчитываем масштаб
	xMin, xMax := -10.0, 10.0
	yMin, yMax := -10.0, 10.0

	// Точки пересечения с осями
	if k != 0 {
		xRoot := -b / k
		if xRoot < xMin {
			xMin = xRoot - 2
		}
		if xRoot > xMax {
			xMax = xRoot + 2
		}
	}

	yRoot := b
	if yRoot < yMin {
		yMin = yRoot - 2
	}
	if yRoot > yMax {
		yMax = yRoot + 2
	}

	// Функция преобразования координат
	transform := func(x, y float64) fyne.Position {
		xPos := padding + (x-xMin)*(width-2*padding)/(xMax-xMin)
		yPos := height - padding - (y-yMin)*(height-2*padding)/(yMax-yMin)
		return fyne.NewPos(float32(xPos), float32(yPos))
	}

	// Рисуем график
	points := make([]fyne.Position, 0)
	for x := xMin; x <= xMax; x += 0.1 {
		y := k*x + b
		if y >= yMin && y <= yMax {
			points = append(points, transform(x, y))
		}
	}

	for i := 0; i < len(points)-1; i++ {
		line := canvas.NewLine(color.RGBA{R: 255, A: 255})
		line.Position1 = points[i]
		line.Position2 = points[i+1]
		container.Add(line)
	}

	// Добавляем метки на оси
	for i := int(xMin); i <= int(xMax); i += 10 {
		if i != 0 {
			pos := transform(float64(i), 0)
			label := canvas.NewText(fmt.Sprintf("%d", i), color.Gray{128})
			label.Move(fyne.NewPos(pos.X-5, float32(height/2+5)))
			container.Add(label)
		}
	}

	for i := int(yMin); i <= int(yMax); i += 10 {
		if i != 0 {
			pos := transform(0, float64(i))
			label := canvas.NewText(fmt.Sprintf("%d", i), color.Gray{128})
			label.Move(fyne.NewPos(float32(width/2+5), pos.Y-5))
			container.Add(label)
		}
	}
}

func createTab3(w fyne.Window) *fyne.Container {
	filePath := ""
	progress := widget.NewProgressBar()
	statusLabel := widget.NewLabel("Файл не выбран")

	selectButton := widget.NewButton("Выбрать файл", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err != nil || reader == nil {
				return
			}
			filePath = reader.URI().Path()
			statusLabel.SetText("Выбран: " + filePath)
		}, w)
	})

	uploadButton := widget.NewButton("Загрузить", func() {
		if filePath == "" {
			statusLabel.SetText("Сначала выберите файл!")
			return
		}

		statusLabel.SetText("Загрузка...")
		go func() {
			for i := 0.0; i <= 1.0; i += 0.05 {
				time.Sleep(100 * time.Millisecond)
				progress.SetValue(i)
			}
			statusLabel.SetText("Файл отправлен на сервер!")
		}()
	})

	return container.NewVBox(
		selectButton,
		statusLabel,
		uploadButton,
		progress,
	)
}
