package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// 創建新的Fyne應用
	myApp := app.New()
	// 創建應用窗口
	myWindow := myApp.NewWindow("NKNU_411034018_Golang_File Copier")
	// 初始化標籤和進度條
	sourceFileLabel := widget.NewLabel("No source file selected")
	targetDirLabel := widget.NewLabel("No target directory selected")
	progressBar := widget.NewProgressBar()

	var sourceFilePath string
	var targetDirPath string

	// 選擇源文件按鈕，打開文件選擇對話框
	selectSourceButton := widget.NewButton("Select Source File", func() {
		dialog.ShowFileOpen(func(reader fyne.URIReadCloser, err error) {
			if err == nil && reader != nil {
				sourceFilePath = reader.URI().Path()
				sourceFileLabel.SetText("Source: " + sourceFilePath)
			}
		}, myWindow)
	})
	// 選擇目標目錄按鈕，打開目錄選擇對話框
	selectTargetButton := widget.NewButton("Select Target Directory", func() {
		dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
			if err == nil && list != nil {
				targetDirPath = list.Path()
				targetDirLabel.SetText("Target: " + targetDirPath)
			}
		}, myWindow)
	})
	// 開始複製按鈕，觸發文件複製過程
	copyButton := widget.NewButton("Start Copy", func() {
		if sourceFilePath == "" || targetDirPath == "" {
			dialog.ShowInformation("Error", "Please select both source file and target directory.", myWindow)
			return
		}

		go func() {
			copyFileWithProgress(sourceFilePath, targetDirPath, progressBar, myWindow)
		}()
	})
	// 將所有組件放入垂直容器中
	content := container.NewVBox(
		selectSourceButton,
		sourceFileLabel,
		selectTargetButton,
		targetDirLabel,
		copyButton,
		progressBar,
	)
	// 設置窗口內容並顯示窗口
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}

// copyFileWithProgress 複製文件並顯示複製進度
func copyFileWithProgress(sourceFile, targetDir string, progressBar *widget.ProgressBar, window fyne.Window) {
	sourceFileStat, err := os.Stat(sourceFile)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	// 構造目標文件路徑
	targetFile := filepath.Join(targetDir, filepath.Base(sourceFile))
	source, err := os.Open(sourceFile)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer source.Close()

	target, err := os.Create(targetFile)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer target.Close()

	bufferSize := 1024 * 1024 // 1 MB
	buffer := make([]byte, bufferSize)
	totalBytesCopied := 0

	for {
		bytesRead, err := source.Read(buffer)
		if err != nil && err != io.EOF {
			fmt.Println("Error: ", err)
			return
		}
		if bytesRead == 0 {
			break
		}

		bytesWritten, err := target.Write(buffer[:bytesRead])
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}

		totalBytesCopied += bytesWritten
		progress := float64(totalBytesCopied) / float64(sourceFileStat.Size())
		progressBar.SetValue(progress)

		// 模擬慢速複製，方便查看進度條變化
		time.Sleep(10 * time.Millisecond)
	}
	// 複製完成後，將進度條設置為100%
	progressBar.SetValue(1)
	// 顯示複製成功的信息對話框
	dialog.ShowInformation("Success", "File copied successfully!", window)
}
