package main

import (
	"syscall"
	"unsafe"
)

var (
	user32          = syscall.NewLazyDLL("user32.dll")
	procMessageBoxW = user32.NewProc("MessageBoxW")
)

const (
	MB_OKCANCEL        = 0x00000001
	MB_ICONINFORMATION = 0x00000040

	IDOK     = 1
	IDCANCEL = 2
)

func messageBox(title, content string, style uintptr) int {
	ret, _, _ := procMessageBoxW.Call(
		0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(content))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		style,
	)
	return int(ret)
}

func showPersistentExecuteDialog() {
	for {
		result := messageBox(
			"すべてを実行",
			`以下のすべてを実行します:

1. データを消す
2. ストレージを買う
3. スタレを入れない

OKを押して続行してください`,
			MB_OKCANCEL|MB_ICONINFORMATION,
		)

		if result == IDOK {
			break
		}

		messageBox("警告", "逃げられると思いましたか？", MB_ICONINFORMATION)
	}
}

func main() {
	showPersistentExecuteDialog()
}
