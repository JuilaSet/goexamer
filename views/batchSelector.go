package views


// Copyright 2017 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"goexamer/store"
)

type batchSelectorMainWindow struct {
	*walk.MainWindow
}

func FromBatch(batchGroup map[string]*store.Batch) (mainWindow MainWindow) {
	var ws []Widget
	var bmw = new(batchSelectorMainWindow)
	for _, batch := range batchGroup {
		name := batch.Name
		ws = append(ws, PushButton{
			MinSize: Size{Width: 100, Height: 50},
			Text: batch.Name,
			OnClicked: func() {
				communicator.Send(SelectPost, name)
				bmw.Close()
			},
		})
	}
	mainWindow = MainWindow{
		AssignTo: &bmw.MainWindow,
		Title:    "Batch Selector",
		Size:     Size{Width: 600, Height: 400},
		Layout:  VBox{},
		Children: ws,
	}
	return
}
