//go:build generate

//go:generate -command mkwinsyscall go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go
//go:generate mkwinsyscall comctl32.go comdlg32.go gdi32.go ole32.go shell32.go user32.go

package win