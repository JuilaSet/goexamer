package views

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/pkg/errors"
	"strings"
)

type MyMainWindow struct {
	*walk.MainWindow
}

var (
	mw           *MyMainWindow
	consoleTxt   *walk.TextEdit
	communicator chan int
)

// 通讯协议
const (
	SelectYes = 1
	SelectNo = 0
	TitlePrefix =  "记忆小工具-"
)

func init(){
	communicator = make(chan int)
	mw = new(MyMainWindow)
	(MainWindow{
		Title:  TitlePrefix,
		MinSize: Size{400, 300},
		Layout:  VBox{},
		AssignTo: &mw.MainWindow,
		MenuItems: []MenuItem{
			Menu{
				Text: "Setting",
				Items: []MenuItem{
					Action{
						Text: "Exit",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
		},
		Children: []Widget{
			TextEdit{AssignTo: &consoleTxt, ReadOnly:true, HScroll: true, VScroll: true, Font: Font{PointSize:12}},
			HSplitter{
				Children: []Widget{
					PushButton{
						MinSize: Size{100, 50},
						Text: "Yes!",
						OnClicked: func() {
							communicator <- SelectYes
						},
					},
					PushButton{
						MinSize: Size{100, 50},
						Text: "No!",
						OnClicked: func() {
							communicator <- SelectNo
						},
					},
				},
			},
		},
	}).Create()
}

func Wait(){
	for {
		if consoleTxt != nil && mw != nil {
			return
		}
	}
}

func Clear(){
	consoleTxt.SetText("")
}

func GetYesOrNo() <-chan int {
	return communicator
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

func SetTitle(str... string) {
	if mw != nil {
		mw.SetTitle(TitlePrefix + strings.Join(str, ""))
		return
	}
	panic(errors.New("wm is nil"))
}

func Index() {
	mw.Run()
}
