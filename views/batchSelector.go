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
	for _, batch := range store.BatchArray() {
		name, text := batch.Name, batch.Name
		if text == "" {
			text = "<default batch>"
		}
		if !batch.IsEmpty() {
			ws = append(ws, PushButton{
				MinSize: Size{Width: 100, Height: 50},
				Text: text,
				OnClicked: func() {
					communicator.Send(SelectPost, name)
					bmw.Close()
				},
			})
		}
	}
	mainWindow = MainWindow{
		AssignTo: &bmw.MainWindow,
		Title:    "Batch Selector",
		Size:     Size{Width: 300, Height: len(ws) * 50 + 5},
		Layout:  VBox{},
		Children: ws,
	}
	return
}
