package views

// Copyright 2017 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"goexamer/store"
	"math"
	"strings"
)

type itemEditMainWindow struct {
	*walk.MainWindow
	txt   *walk.TextEdit
	lb    *walk.ListBox
	model *ItemInfoModel
}

type ItemInfo struct {
	Qus string
	Ans []string
}

type ItemInfoModel struct {
	walk.ListModelBase
	items []ItemInfo
}

func (m *ItemInfoModel) ItemCount() int {
	return len(m.items)
}

func (m *ItemInfoModel) Value(index int) interface{} {
	return m.items[index].Qus
}

func NewItemInfoModel(batchName string) *ItemInfoModel {
	m := new(ItemInfoModel)
	m.items = make([]ItemInfo, 0, 10)
	for qus, ans := range store.GetBatch(batchName).GetAllQus() {
		m.items = append(m.items, ItemInfo{qus, ans})
	}
	return m
}

func (mw *itemEditMainWindow) CurrentIndexChanged() {
	qus := mw.model.items[mw.lb.CurrentIndex()]
	qusStr, ansArr := qus.Qus, qus.Ans
	s := "#" + qusStr + ":"
	for _, line := range ansArr {
		s += line + "\r\n"
	}
	mw.txt.SetText(s)
}

func (mw *itemEditMainWindow) ItemActivated() {
	str := strings.ReplaceAll(mw.txt.Text(), "\r\n", "\n")
	communicator.Send(SelectItemSave, str)
	mw.Close()
}

func FromItem(batchName string) (mainWindow MainWindow) {
	var imw = &itemEditMainWindow{
		model: NewItemInfoModel(batchName),
	}
	mainWindow = MainWindow{
		AssignTo: &imw.MainWindow,
		Title:    "Item Editor",
		Size:     Size{Width: 400, Height: 300},
		Layout:   VBox{},
		MenuItems: []MenuItem{
			Action{
				Text: "Save Item",
				OnTriggered: func() {
					imw.ItemActivated()
				},
			},
		},
		Children: []Widget{
			TextEdit{
				AssignTo: &imw.txt,
				ReadOnly: false,
				HScroll: true,
				VScroll: true,
				Font: Font{PointSize: 12},
				MinSize:Size{Width: 300, Height: 200},
			},
			ListBox{
				AssignTo: &imw.lb,
				Model: imw.model,
				OnCurrentIndexChanged: imw.CurrentIndexChanged,
				OnItemActivated:       imw.ItemActivated,
				MinSize:Size{Width: 300, Height: int(math.Min(float64(25 * len(imw.model.items)), 125.0))},
			},
		},
	}
	return
}
