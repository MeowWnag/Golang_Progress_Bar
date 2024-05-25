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

	//創建或覆蓋檔案，並在操作完成後正確關閉檔案
	target, err := os.Create(targetFile)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer target.Close()
	// 設定緩衝區大小為 1 MB
	bufferSize := 1024 * 1024 
	// 建立一個長度為 bufferSize 的 byte slice 作為緩衝區
	buffer := make([]byte, bufferSize)
	// 初始化已複製的總位元組數為 0
	totalBytesCopied := 0

	for {
		// 從源中讀取數據到緩衝區
		bytesRead, err := source.Read(buffer)
		if err != nil && err != io.EOF {// 如果發生錯誤且不是檔案結束錯誤
			fmt.Println("Error: ", err)
			return
		}
		if bytesRead == 0 {// 如果讀取的位元組數為 0，表示檔案已讀取完畢
			break
		}
		// 將緩衝區中的數據寫入目標檔案
		bytesWritten, err := target.Write(buffer[:bytesRead])
		if err != nil {// 如果寫入發生錯誤
			fmt.Println("Error: ", err)
			return
		}
		// 更新已複製的總位元組數
		totalBytesCopied += bytesWritten
		// 計算複製進度
		progress := float64(totalBytesCopied) / float64(sourceFileStat.Size())
		// 模擬慢速複製，方便查看進度條變化
		progressBar.SetValue(progress)

		// 模擬慢速複製，方便查看進度條變化
		time.Sleep(10 * time.Millisecond)
	}
	// 複製完成後，將進度條設置為100%
	progressBar.SetValue(1)
	// 顯示複製成功的信息對話框
	dialog.ShowInformation("Success", "File copied successfully!", window)
}
