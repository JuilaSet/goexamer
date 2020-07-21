package views

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/pkg/errors"
	"os"
	"strings"
)

type MyMainWindow struct {
	*walk.MainWindow
	dangerMode   bool
	saveAction  *walk.Action
	aFontAction  *walk.Action
	dangerAction *walk.Action
}

var (
	mw           *MyMainWindow
	consoleTxt   *walk.TextEdit
	slider       *walk.Slider
	communicator *Communicator
)

// 通讯协议
const (
	SelectItemSave = 6
	SelectItemEdit = 5
	SelectSave     = 4
	SelectFile     = 3
	SelectBatch    = 2
	SelectYes      = 1
	SelectNo       = 0
	TitlePrefix    = "记忆小工具 "
)

// 通讯通道
type Communicator struct {
	flag chan int
	ctx  chan string
}

func NweCommunicator() (communicator *Communicator) {
	communicator = new(Communicator)
	communicator.flag = make(chan int, 5)
	communicator.ctx = make(chan string, 5)
	return
}

func (communicator *Communicator) Send(flag int, ctx string) {
	communicator.flag <- flag
	communicator.ctx <- ctx
}

func (communicator *Communicator) Receive() (flag int, ctx string) {
	return <-communicator.flag, <-communicator.ctx
}

func GetCommunicator() *Communicator {
	return communicator
}

func init() {
	communicator = NweCommunicator()
	mw = &MyMainWindow{dangerMode: false}
	(MainWindow{
		Title:    TitlePrefix,
		MinSize:  Size{400, 300},
		Layout:   VBox{},
		AssignTo: &mw.MainWindow,
		MenuItems: []MenuItem{
			Menu{
				Text: "File",
				Items: []MenuItem{
					Action{
						Text: "Open",
						OnTriggered: func() {
							FromDir().Run()
						},
					},
					Action{
						AssignTo: &mw.saveAction,
						Text: "Save",
						Enabled: mw.dangerMode,
						OnTriggered: func() {
							communicator.Send(SelectSave, "")
						},
					},
					Action{
						Text: "Restart",
						OnTriggered: func() {
							communicator.Send(SelectFile, "")
						},
					},
					Action{
						Text: "Exit",
						OnTriggered: func() {
							os.Exit(0)
						},
					},
				},
			},
			Menu{
				Text: "Batch",
				Items: []MenuItem{
					Action{
						Text: "Select Batch",
						OnTriggered: func() {
							FromBatchGroup().Run()
						},
					},
					Action{
						Text: "Edit Item",
						OnTriggered: func() {
							communicator.Send(SelectItemEdit, "")
						},
					},
				},
			},
			Menu{
				Text: "Setting",
				Items: []MenuItem{
					Action{
						AssignTo: &mw.dangerAction,
						Text:     "Danger Mode",
						OnTriggered: func() {
							mw.dangerMode = !mw.dangerMode
							mw.saveAction.SetEnabled(mw.dangerMode)
							if mw.dangerMode {
								mw.dangerAction.SetText("Danger Mode  √")
							} else {
								mw.dangerAction.SetText("Danger Mode")
							}
						},
					},
					Action{
						AssignTo: &mw.aFontAction,
						Text:     "Font Size Slider",
						OnTriggered: func() {
							slider.SetVisible(!slider.Visible())
							if slider.Visible() {
								mw.aFontAction.SetText("Font Size Slider  √")
							} else {
								mw.aFontAction.SetText("Font Size Slider")
							}
						},
					},
				},
			},
		},
		Children: []Widget{
			TextEdit{AssignTo: &consoleTxt, ReadOnly: true, HScroll: true, VScroll: true, Font: Font{PointSize: 12}},
			Slider{
				ColumnSpan: 1,
				AssignTo:   &slider,
				MinValue:   5,
				MaxValue:   50,
				Value:      12,
				OnValueChanged: func() {
					font, _ := walk.NewFont("", slider.Value(), 0)
					consoleTxt.SetFont(font)
				},
			},
			HSplitter{
				Children: []Widget{
					PushButton{
						MinSize: Size{80, 50},
						Text:    "Yes!",
						OnClicked: func() {
							communicator.Send(SelectYes, "")
						},
					},
					PushButton{
						MinSize: Size{80, 50},
						Text:    "No!",
						OnClicked: func() {
							communicator.Send(SelectNo, "")
						},
					},
				},
			},
		},
	}).Create()
	slider.SetVisible(false)
}

func Wait() {
	for {
		if consoleTxt != nil && mw != nil && slider != nil {
			return
		}
	}
}

func ReadOnly(b bool){
	consoleTxt.SetReadOnly(b)
}

func Clear() {
	consoleTxt.SetText("")
}

func SetText(str string) {
	if consoleTxt != nil {
		str = consoleTxt.Text() + str
		str = strings.ReplaceAll(str, "\n", "\r\n")
		consoleTxt.SetText(str)
		return
	}
	panic(errors.New("consoleTxt is nil"))
}

func SetTitle(str ...string) {
	if mw != nil {
		mw.SetTitle(TitlePrefix + strings.Join(str, ""))
		return
	}
	panic(errors.New("wm is nil"))
}

//func CreateAction(menu *walk.Menu) (action *walk.Action) {
//	action = walk.NewMenuAction(menu)
//	action.SetText("dynamic action")
//	action.Triggered().Attach(func() {
//		fmt.Println("aaa")
//	})
//	return
//}

func Index() {
	//menu, _ := walk.NewMenu()
	//mw.Menu().Actions().Add(CreateAction(menu))
	mw.Run()
}
