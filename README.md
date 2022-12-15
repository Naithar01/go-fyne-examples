# go-fyne-examples

## ex01

1. Container
2. Window Size
3. Label, Button, Entry
4. Change Label Text

## ex02

- Container Split - NewHSplit
- Center App - CenterOnScreen
- fyne.URI - file path
- NewExtensionFileFilter - save file option ex) md MD txt

- -

- Entry - NewMultiLineEntry ( line Input Entry )
- Label - NewRichTextFromMarkdown (line Label Markdown Option )
- edit.OnChanged = preview.ParseMarkdown // when entry changed, label value is change

- -

- NewMenuItem - New Menu widget
- NewMenu - Menu Item Wrapper (menu := fyne.NewMainMenu(fileMenu)) (win.SetMainMenu(menu))

- storage Writer - storage.Writer(path)
- writer.Write([]byte(c.EditWidget.Text))

- -

- data, err := ioutil.ReadAll(read), Read File string
- c.EditWidget.SetText(string(data))
- c.CurrenFile = read.URI() - Set path
- win.SetTitle(win.Title() + " - " + read.URI().Name()) - Set file Name
- openDialog.SetFilter(openFilter) - top filter

- -

- c.SaveMenuItem.Disabled = false - menu click disable
- saveDialog.SetFileName("untitled.md") - file default name
