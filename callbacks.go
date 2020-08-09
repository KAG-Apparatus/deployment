package main

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/gotk3/gotk3/gtk"
)

func onDialogDelete(dialog *gtk.Dialog) bool {
	dialog.Hide()
	return true
}

func onFileChooserDialogDelete(fileChooserDialog *gtk.FileChooserDialog) bool {
	fileChooserDialog.Hide()
	return true
}

// onMainWindowDestory is the callback that is linked to the
// on_main_window_destroy handler. It is not required to map this,
// and is here to simply demo how to hook-up custom callbacks.
func onMainWindowDestroy() {
	log.Println("Leaving application. Good bye!")
}

func onButtonDevFolderClicked(button *gtk.Button) {
	obj, err := builder.GetObject("dev_folder_chooser")
	errorCheck(err)
	fileChooserDialog, err := isFileChooserDialog(obj)
	errorCheck(err)
	fileChooserDialog.SetTitle("Select development folder")
	if params.sourceDir != "" {
		parent := filepath.Dir(params.sourceDir)
		fileChooserDialog.SetCurrentFolder(parent)
	}
	fileChooserDialog.Run()
}

func onButtonOkFolderClicked(button *gtk.Button) {
	obj, err := builder.GetObject("dev_folder_chooser")
	errorCheck(err)

	fileChooserDialog, err := isFileChooserDialog(obj)
	errorCheck(err)

	folder := strings.TrimPrefix(fileChooserDialog.GetURI(), "file://")

	params.sourceDir = folder

	obj, err = builder.GetObject("entry_dev_folder")
	errorCheck(err)

	entryDevFolder, err := isEntry(obj)
	errorCheck(err)

	entryDevFolder.SetText(folder)

	fileChooserDialog.Hide()
}

func onButtonCancelFolderClicked(button *gtk.Button) {
	obj, err := builder.GetObject("dev_folder_chooser")
	errorCheck(err)
	fileChooserDialog, err := isFileChooserDialog(obj)
	errorCheck(err)
	fileChooserDialog.Hide()
}
