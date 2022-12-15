package main

import (
	"io/ioutil"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

// * URI는 경로를 말함
// ex) file://C:/Users/사용자/Videos/untitled1.md

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrenFile    fyne.URI // 파일경로
	SaveMenuItem  *fyne.MenuItem
}

var (
	cfg        config             = config{}
	openFilter storage.FileFilter = storage.NewExtensionFileFilter([]string{".md", ".MD", ".txt"}) // 사용 가능 확장자 리스트
)

func main() {
	app := app.New()
	win := app.NewWindow("Markdown Editor")

	edit, preview := cfg.makeUI()
	cfg.createItemsMenu(win)

	win.SetContent(container.NewHSplit(preview, edit)) // 화면에 split하여 content를 보여줌

	win.Resize(fyne.Size{Width: 700, Height: 700})
	win.CenterOnScreen() // 실행시에 가운데로
	win.ShowAndRun()

}

func (c *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()            // 텍스트 input 위젯
	preview := widget.NewRichTextFromMarkdown("") // 텍스트 label 위젯

	c.EditWidget = edit
	c.PreviewWidget = preview

	// edit input의 값이 바뀌면 preview에 값이 전달
	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}

func (c *config) createItemsMenu(win fyne.Window) {
	openMenuItem := fyne.NewMenuItem("Open Menu", c.openFunc(win))

	saveMenuItem := fyne.NewMenuItem("Save", c.saveFunc(win))

	c.SaveMenuItem = saveMenuItem
	c.SaveMenuItem.Disabled = true // disable이 true인 이유는 파일을 열기 혹은 새로저장하기를 안 했을 때 저장을 하는 모순이 생기면 안 돼서

	saveAsMenuItem := fyne.NewMenuItem("Save as", c.saveAsFunc(win))

	// 옵션, 옵션 리스트는 위 3개
	fileMenu := fyne.NewMenu("File", openMenuItem, saveMenuItem, saveAsMenuItem)

	// 새로운 메뉴 New로 추가
	menu := fyne.NewMainMenu(fileMenu)

	// Set Menu
	win.SetMainMenu(menu)
}

func (c *config) saveFunc(win fyne.Window) func() {
	return func() {
		if c.CurrenFile != nil {
			// CurrenFile = URI (경로)인데
			// 파일을 오픈, 새로 저장을 하면 currenFile에 URI가 생겨서 Save가 disable이 false가 됨
			writer, err := storage.Writer(c.CurrenFile)

			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			writer.Write([]byte(c.EditWidget.Text))

			defer writer.Close()
		}
	}
}

func (c *config) openFunc(win fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if read == nil {
				return
			}

			defer read.Close()

			// 파일 열기 URI
			data, err := ioutil.ReadAll(read)

			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			// Edit Widget에 data를 넣음, 그러면 preview도 자동적으로 바뀜
			c.EditWidget.SetText(string(data))

			c.CurrenFile = read.URI()
			win.SetTitle(win.Title() + " - " + read.URI().Name()) // 파일의 이름 (경로 제외)
			c.SaveMenuItem.Disabled = false

		}, win)
		openDialog.SetFilter(openFilter)
		openDialog.Show()
	}
}

func (c *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		saveDialog := dialog.NewFileSave(func(writer fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if writer == nil {
				return
			}

			// 파일 확장자가 md 여야지 가능
			if !strings.HasSuffix(strings.ToLower(writer.URI().String()), ".md") {
				dialog.ShowInformation("Error", "Save Only md", win)
			}

			writer.Write([]byte(c.EditWidget.Text)) // Input 위젯의 Text를 바이트로 저장
			c.CurrenFile = writer.URI()

			defer writer.Close()

			win.SetTitle(win.Title() + " - " + writer.URI().Name()) // 파일의 이름만 고름 (경로 제외)
			c.SaveMenuItem.Disabled = false

		}, win)
		saveDialog.SetFileName("untitled.md") // 파일 저장 전의 초기이름
		saveDialog.SetFilter(openFilter)
		saveDialog.Show()
	}
}
