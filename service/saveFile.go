package service

import (
	"goexamer/io"
	"goexamer/store"
)

func SaveFile(){
	output := io.NewFileOutPutter(store.CurFileName())
	defer output.Close()

	output.Clear()
	output.Print(store.GetBeforeTitle())

	// title
	output.SetTitle(store.GetTitle()...)

	// batch
	for _, b := range store.GetAllBatch() {
		output.Print("\n")
		output.Print(b.ToString())
	}
}
