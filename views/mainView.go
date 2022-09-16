package views

import (
	"github.com/therecipe/qt/widgets"
	"os"
)

func MainView() *widgets.QApplication {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetMinimumSize2(250, 200)
	window.SetWindowTitle("Testwindow")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("Search...")
	widget.Layout().AddWidget(input)

	button := widgets.NewQPushButton2("click me", nil)
	button.ConnectClicked(func(bool) {
		widgets.QMessageBox_Information(nil, "Ok", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Abort)
	})
	widget.Layout().AddWidget(button)

	window.Show()
	return app
}
