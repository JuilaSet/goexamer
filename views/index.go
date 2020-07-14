package views

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/pkg/errors"
	"strings"
)

var (
	mainWindow *MainWindow
	consoleTxt *walk.TextEdit
)

var (
	yesOrNo chan bool
)

func init(){
	yesOrNo = make(chan bool)
	mainWindow = &MainWindow{
		Title:   "记忆小工具",
		MinSize: Size{400, 300},
		Layout:  VBox{},
		Children: []Widget{
			TextEdit{AssignTo: &consoleTxt, ReadOnly:true},
			HSplitter{
				Children: []Widget{
					PushButton{
						MinSize: Size{100, 50},
						Text: "Yes!",
						OnClicked: func() {
							yesOrNo <- true
						},
					},
					PushButton{
						MinSize: Size{100, 50},
						Text: "No!",
						OnClicked: func() {
							yesOrNo <- false
						},
					},
				},
			},
		},
	}
}

func Wait(){
	for {
		if consoleTxt != nil {
			return
		}
	}
}

func Clear(){
	consoleTxt.SetText("")
}

func GetYesOrNo() <-chan bool {
	return yesOrNo
}

func SetText(str string) {
	if consoleTxt != nil {
		str = consoleTxt.Text() + str
		str = strings.ReplaceAll(str,"\n", "\r\n")
		consoleTxt.SetText(str)
		return
	}
	panic(errors.New("consoleTxt is nil"))
}

func Index() {
	mainWindow.Run()
}
