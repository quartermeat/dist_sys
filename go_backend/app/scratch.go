package app

import (
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func RunChat() {
	var inputTextEdit *walk.TextEdit
	var chatHistoryEdit *walk.TextEdit

	// Create a new tool tip for the "Send" button
	toolTip, err := walk.NewToolTip()
	if err != nil {
		log.Fatal(err)
	}

	mainWindow := MainWindow{
		Title:   "Simple Chat",
		MinSize: Size{Width: 600, Height: 400},
		Layout:  VBox{},
		Children: []Widget{
			TextEdit{
				AssignTo: &chatHistoryEdit,
				ReadOnly: true,
				VScroll:  true,
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					TextEdit{AssignTo: &inputTextEdit},
					PushButton{
						Text:        "Send",
						ToolTipText: "Click to send the message",
						OnClicked: func() {
							message := inputTextEdit.Text()
							if message != "" {
								chatHistoryEdit.SetText(chatHistoryEdit.Text() + message + "\r\n")
								inputTextEdit.SetText("")
							}
						},
						// Set the tool tip for the "Send" button
						AssignTo: &PushButton{},
						OnMouseEnter: func(x, y int, dragState walk.MouseDragState) {
							toolTip.SetText("Click to send the message")
							toolTip.SetWindow(PushButton.Parent())
							toolTip.SetEnabled(true)
							toolTip.TriggerPopUp(x, y)
						},
						OnMouseLeave: func() {
							toolTip.SetEnabled(false)
						},
					},
				},
			},
		},
	}

	if _, err := mainWindow.Run(); err != nil {
		log.Fatal(err)
	}
}
