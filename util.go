package main

import (
	"errors"
	"log"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func isWindow(obj glib.IObject) (*gtk.ApplicationWindow, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.ApplicationWindow); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.ApplicationWindow")
}

func isAssistant(obj glib.IObject) (*gtk.Assistant, error) {
	// Make type assertion (as per gtk.go).
	if win, ok := obj.(*gtk.Assistant); ok {
		return win, nil
	}
	return nil, errors.New("not a *gtk.Assistant")
}

func isDialog(obj glib.IObject) (*gtk.Dialog, error) {
	// Make type assertion (as per gtk.go).
	if dialog, ok := obj.(*gtk.Dialog); ok {
		return dialog, nil
	}
	return nil, errors.New("not a *gtk.Dialog")
}

func isEntry(obj glib.IObject) (*gtk.Entry, error) {
	// Make type assertion (as per gtk.go).
	if entry, ok := obj.(*gtk.Entry); ok {
		return entry, nil
	}
	return nil, errors.New("not a *gtk.Entry")
}

func isComboBoxText(obj glib.IObject) (*gtk.ComboBoxText, error) {
	// Make type assertion (as per gtk.go).
	if comboBoxText, ok := obj.(*gtk.ComboBoxText); ok {
		return comboBoxText, nil
	}
	return nil, errors.New("not a *gtk.Entry")
}

func errorCheck(e error) {
	if e != nil {
		// panic for any errors.
		log.Panic(e)
	}
}

func isLabel(obj glib.IObject) (*gtk.Label, error) {
	if label, ok := obj.(*gtk.Label); ok {
		return label, nil
	}
	return nil, errors.New("not a *gtk.Label")
}

func isSpinButton(obj glib.IObject) (*gtk.SpinButton, error) {
	if spinBtn, ok := obj.(*gtk.SpinButton); ok {
		return spinBtn, nil
	}
	return nil, errors.New("not a *gtk.SpinButton")
}

func isPopover(obj glib.IObject) (*gtk.Popover, error) {
	if popover, ok := obj.(*gtk.Popover); ok {
		return popover, nil
	}
	return nil, errors.New("not a *gtk.Popover")
}

func isButton(obj glib.IObject) (*gtk.Button, error) {
	if btn, ok := obj.(*gtk.Button); ok {
		return btn, nil
	}
	return nil, errors.New("not a *gtk.Button")
}

func isSearchEntry(obj glib.IObject) (*gtk.SearchEntry, error) {
	if searchEntry, ok := obj.(*gtk.SearchEntry); ok {
		return searchEntry, nil
	}
	return nil, errors.New("not a *gtk.Button")
}

func isListStore(obj glib.IObject) (*gtk.ListStore, error) {
	if listStore, ok := obj.(*gtk.ListStore); ok {
		return listStore, nil
	}
	return nil, errors.New("not a *gtk.Button")
}

func isEntryCompletion(obj glib.IObject) (*gtk.EntryCompletion, error) {
	if entryCompletion, ok := obj.(*gtk.EntryCompletion); ok {
		return entryCompletion, nil
	}
	return nil, errors.New("not a *gtk.Button")
}

func isSpinner(obj glib.IObject) (*gtk.Spinner, error) {
	if spinner, ok := obj.(*gtk.Spinner); ok {
		return spinner, nil
	}
	return nil, errors.New("not a *gtk.Spinner")
}

func isImage(obj glib.IObject) (*gtk.Image, error) {
	if image, ok := obj.(*gtk.Image); ok {
		return image, nil
	}
	return nil, errors.New("not a *gtk.Image")
}

func isGrid(obj glib.IObject) (*gtk.Grid, error) {
	if grid, ok := obj.(*gtk.Grid); ok {
		return grid, nil
	}
	return nil, errors.New("not a *gtk.Grid")
}

func isTreeView(obj glib.IObject) (*gtk.TreeView, error) {
	if treeView, ok := obj.(*gtk.TreeView); ok {
		return treeView, nil
	}
	return nil, errors.New("not a *gtk.TreeView")
}

func isBox(obj glib.IObject) (*gtk.Box, error) {
	if box, ok := obj.(*gtk.Box); ok {
		return box, nil
	}
	return nil, errors.New("not a *gtk.Box")
}

func isPopoverMenu(obj glib.IObject) (*gtk.PopoverMenu, error) {
	if popoverMenu, ok := obj.(*gtk.PopoverMenu); ok {
		return popoverMenu, nil
	}
	return nil, errors.New("not a *gtk.PopoverMenu")
}

func isCellRenderer(obj glib.IObject) (*gtk.CellRenderer, error) {
	if cellRenderer, ok := obj.(*gtk.CellRenderer); ok {
		return cellRenderer, nil
	}
	return nil, errors.New("not a *gtk.CellRenderer")
}

func isCellRendererText(obj glib.IObject) (*gtk.CellRendererText, error) {
	if cellRendererText, ok := obj.(*gtk.CellRendererText); ok {
		return cellRendererText, nil
	}
	return nil, errors.New("not a *gtk.CellRendererText")
}

func isFileChooserDialog(obj glib.IObject) (*gtk.FileChooserDialog, error) {
	if fileChooserDialog, ok := obj.(*gtk.FileChooserDialog); ok {
		return fileChooserDialog, nil
	}
	return nil, errors.New("not a *gtk.FileChooserDialog")
}

func isTextBuffer(obj glib.IObject) (*gtk.TextBuffer, error) {
	if textBuffer, ok := obj.(*gtk.TextBuffer); ok {
		return textBuffer, nil
	}
	return nil, errors.New("not a *gtk.TextBuffer")
}

func isTextView(obj glib.IObject) (*gtk.TextView, error) {
	if textView, ok := obj.(*gtk.TextView); ok {
		return textView, nil
	}
	return nil, errors.New("not a *gtk.TextView")
}
