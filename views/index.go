package views

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/pkg/errors"
	"goexamer/store"
	"os"
	"strings"
)

type MyMainWindow struct {
	*walk.MainWindow
}

var (
	mw           *MyMainWindow
	consoleTxt   *walk.TextEdit
	slider		 *walk.Slider
	aFontSlider	 *walk.Action
	communicator *Communicator
)

// 通讯协议
const (
	SelectPost = 2
	SelectYes     = 1
	SelectNo      = 0
	TitlePrefix   =  "记忆小工具-"
)

// 通讯通道
type Communicator struct {
	flag chan int
	ctx chan string
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


func init(){
	communicator = NweCommunicator()
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
						Text: "Batch Selector",
						OnTriggered: func() {
							FromBatch(store.GetAllBatch()).Run()
						},
					},
					Action{
						AssignTo: &aFontSlider,
						Text: "Font Size Slider",
						OnTriggered: func() {
							slider.SetVisible(!slider.Visible())
							if slider.Visible() {
								aFontSlider.SetText("Font Size Slider  √")
							} else {
								aFontSlider.SetText("Font Size Slider")
							}
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
		},
		Children: []Widget{
			TextEdit{AssignTo: &consoleTxt, ReadOnly:true, HScroll: true, VScroll: true, Font: Font{PointSize:12}},
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
						MinSize: Size{100, 50},
						Text: "Yes!",
						OnClicked: func() {
							communicator.Send(SelectYes, "")
						},
					},
					PushButton{
						MinSize: Size{100, 50},
						Text: "No!",
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

func Wait(){
	for {
		if consoleTxt != nil && mw != nil && slider != nil {
			return
		}
	}
}

func Clear(){
	consoleTxt.SetText("")
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
