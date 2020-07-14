package views


// Copyright 2017 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func FromImage(rootDirPath, imageName string) []Widget {
	walk.Resources.SetRootDirPath(rootDirPath)
	return []Widget{
		ImageView{
			Image:      imageName,
			Margin:     1,
			Mode:       ImageViewModeZoom,
		},
	}
}

func ShowImage(widgets []Widget, title string) func() {
	mainWindow := MainWindow{
		Title:    title,
		Size:     Size{600, 400},
		Layout:   Grid{Columns: 1},
		Children: widgets,
	}
	return func() {
		mainWindow.Run()
	}
}