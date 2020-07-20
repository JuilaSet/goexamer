package views

// Copyright 2017 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"io/ioutil"
	"os"
	"time"
)

type fileSelectorMainWindow struct {
	*walk.MainWindow
	lb    *walk.ListBox
	model *FileInfoModel
}

type FileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
}

type FileInfoModel struct {
	walk.ListModelBase
	items []FileInfo
}

func (m *FileInfoModel) ItemCount() int {
	return len(m.items)
}

func (m *FileInfoModel) Value(index int) interface{} {
	return m.items[index].Name
}

func NewFooModel() *FileInfoModel {
	m := new(FileInfoModel)
	var fileInfoArr []FileInfo
	for _, info := range DirInfo() {
		name, isDir, size, modTime := info.Name(), info.IsDir(), info.Size(), info.ModTime()
		if !isDir {
			fileInfoArr = append(fileInfoArr, FileInfo{
				name,
				size,
				modTime,
			})
		}
	}
	m.items = fileInfoArr
	return m
}

func DirInfo() []os.FileInfo {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		panic(err)
	}
	return fileInfoList
}

func FromDir() (mainWindow MainWindow) {
	var mw = &fileSelectorMainWindow{
		model: NewFooModel(),
	}
	mainWindow = MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "File Selector",
		Size:     Size{Width: 500, Height: 300},
		Layout:   VBox{},
		Children: []Widget{
			ListBox{
				AssignTo: &mw.lb,
				Model: mw.model,
				OnCurrentIndexChanged: mw.CurrentIndexChanged,
				OnItemActivated:       mw.ItemActivated,
			},
		},
	}
	return
}

func (mw *fileSelectorMainWindow) CurrentIndexChanged() {}

func (mw *fileSelectorMainWindow) ItemActivated() {
	name := mw.model.items[mw.lb.CurrentIndex()].Name
	communicator.Send(SelectFile, name)
	mw.Close()
}