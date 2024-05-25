# Golang_Progress_Bar
使用Golang實作複製檔案與進度條  
這是我第一個Golang程式，所以除了如何使用這個程式與程式的結構外，還會簡單說明GO語言的語法跟學習心得  
---
##  目錄
- [如何使用這個程式?](#如何使用這個程式)
- [程式碼簡介](#程式碼簡介)
***
##  Golang(GO)語言簡介:  
### Go 是一種於 2007 年在 Google 開發的跨平台、開源程式語言，可用於創建高效能應用程式，以其簡單性和高效性而聞名，語法類似C++
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/81ee49a7-380e-4491-a491-f34c721d028a)
***
##  如何使用這個程式?:
### 1.  方法1:下載完程式後，在終端執行以下代碼  
```
go run ProgressBar.go
```
### 2.  方法2:下載完程式後，在終端執行以下代碼，之後會在同一個目錄生成一個.exe執行檔  
```
go build run ProgressBar.go
```
---
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/d317e269-780e-4488-aba1-da5e1d306ba7)   
### 3.  點擊"Select Source File"和"Select Target Directory"選擇要複製的檔案與要複製的位置( **注意**!!這裡只有英文能正常顯示，中文會變成?圖示)  
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/aaff98fd-526f-49e7-8a73-c0d0ffb4f99e)  
### 4.  點擊"Start Copy"程式就會開始執行複製，同時下面的進度條也會開始跑  
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/3636e2a6-94da-4c91-b497-985c7247c8c1)  
### 5.  複製完成後會跳出"File copied successfully!"的訊息  
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/ef8bb815-d764-4b25-8593-26a08d916dab)  
***
## 程式碼簡介
### 第 1 行：在 Go 中，每個程式都是套件的一部分。我們使用關鍵字來定義它 package。在此範例中，程式屬於 main套件。  
### 第 2 行： 讓我們匯入fmt套件中包含的檔案:  
- "fmt" 和 "io" 是Go語言標準庫中的包,用於格式化輸入/輸出和實現基本的I/O操作。  
- "os" 是Go語言標準庫中的包,提供了一些與操作系統相關的功能,比如文件操作、進程控制等。  
- "path/filepath" 是Go語言標準庫中的包,提供了一些路徑相關的實用函數。  
- "time" 是Go語言標準庫中的包,用於處理時間和日期相關的操作。  
- "fyne.io/fyne/v2" 是Fyne跨平台GUI工具包的主包,提供了基本的GUI組件和事件處理機制。  
- "fyne.io/fyne/v2/app" 是Fyne工具包中的一個包,用於創建和管理應用程序實例。    
- "fyne.io/fyne/v2/container" 是Fyne工具包中的一個包,提供了一些常用的容器組件,如水平或垂直佈局等。  
- "fyne.io/fyne/v2/dialog" 是Fyne工具包中的一個包,提供了一些常用的對話框,如文件選擇對話框、消息對話框等。  
- "fyne.io/fyne/v2/widget" 是Fyne工具包中的一個包,提供了一些常用的GUI組件,如按鈕、標籤、輸入框等。  
### 總而言之,這段代碼導入了Go語言標準庫中的一些常用包,以及Fyne跨平台GUI工具包中的一些核心包,為構建GUI應用程序做好準備  
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/56a12293-193e-4495-bf82-e7a8b2148ab6)  
***    
### 在 Golang 中，宣告變數時，可以直接指定值，或者使用型別推導，更簡單地用 := 來宣告局部變數。例如:  
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/8abe91d1-5e30-4d2e-a32f-fe223c1286b5)  
### 也可以使用 var 關鍵字進行普通聲明，然後指定變數名稱和類型，再賦予初始值。例如:  
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/37685800-5b66-4993-9af2-dd971f6e7410)  
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/fce2856d-fc5e-4434-8d8f-ced0435f55f4)
***
### 在Go語言中,nil是一個預設值,代表指標、函數、介面、Maps、Slices 和 Channels 的"無值"或空白值。它的語義稍有不同,具體取決於它所應用的類型，以33行來說，這裡的nil是用來檢查err和是否為空值
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/19ec21b3-bd5d-4626-b894-2f73035d06ae)  
***  
### 在第49行中，widget.Button對象並將其賦值給變量copyButton。widget.NewButton函數接受兩個參數:第一個是按鈕上顯示的文字"Start Copy"，第二個是一個匿名函數,該函數將在按鈕被點擊時執行。(**注意**左大括號{不能出現在行首)  
### 這裏檢查了sourceFilePath和targetDirPath兩個變量是否為空字符串。如果有一個為空,就會彈出一個對話框,提示用戶選擇源文件和目標目錄。   
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/2cdf7bba-55d5-458b-9af2-91426bcec270)  
***
### 在第60行，我們創建一個垂直盒子佈局(VBox)。這個垂直盒子包含了以下控制元件:  
- ### selectSourceButton: 一個按鈕控制元件,可能用於選擇源文件或目錄。  
- ### sourceFileLabel: 一個標籤控制元件,可能用於顯示選擇的源文件或目錄路徑。  
- ### selectTargetButton: 一個按鈕控制元件,可能用於選擇目標目錄。  
- ### targetDirLabel: 一個標籤控制元件,可能用於顯示選擇的目標目錄路徑。  
- ### copyButton: 一個按鈕控制元件,可能用於執行拷貝操作。  
- ### progressBar: 一個進度條控制元件,可能用於顯示拷貝進度。  
### 這些控制元件被垂直排列在一個盒子佈局中,可能用於構建一個文件拷貝工具的用戶界面。當用戶點擊相關按鈕時,應用程序會執行相應的操作,如選擇源文件、選擇目標目錄、執行拷貝操作等,並在標籤和進度條中實時顯示相關信息。
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/3a1652f4-0b19-47ef-a9b1-f8cd8cf1a879)
***
### 在第76行，os.Stat(sourceFile) 這行代碼的作用是獲取指定文件或目錄的文件信息(metadata)，其中 **os** 是 Go 語言中一個內建的包(package),提供了操作系統相關的功能,比如文件系統的操作。  
**Stat** 是 os 包中的一個函數,它的作用是返回一個描述指定文件對象(文件或目錄)的 FileInfo 類型值。  
**sourceFile** 是一個字符串類型的變量,表示你要獲取文件信息的文件路徑。  
### 該函數的返回值有兩個,第一個返回值 sourceFileStat 是一個 FileInfo 類型的值,包含了該文件的元數據信息,比如文件大小、修改時間、權限等。第二個返回值 err 是一個 error 類型的值,用於判斷操作是否出現錯誤。  
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/3cab14f8-790e-4883-b0da-578ea032efd0)  
***
### 第88行，這行代碼使用了 Go 語言中的 defer 關鍵字。它的作用是在當前函數執行完畢後,延遲執行 source.Close() 語句。  
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/9c067492-b8d2-405f-a37a-58c1655707f7)  


















