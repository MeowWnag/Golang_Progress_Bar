# Golang_Progress_Bar
使用Golang實作複製檔案與進度條  
這是我第一個Golang程式，所以除了如何使用這個程式與程式的結構外，還會簡單說明GO語言的語法跟學習心得  
***
##  Golang(GO)語言簡介:  
Go 是一種於 2007 年在 Google 開發的跨平台、開源程式語言，可用於創建高效能應用程式，以其簡單性和高效性而聞名，語法類似C++
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/81ee49a7-380e-4491-a491-f34c721d028a)
***
##  如何使用這個程式?:
1.  方法1:下載完程式後，在終端執行以下代碼  
```
go run ProgressBar.go
```
2.  方法2:下載完程式後，在終端執行以下代碼，之後會在同一個目錄生成一個.exe執行檔  
```
go build run ProgressBar.go
```
---
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/d317e269-780e-4488-aba1-da5e1d306ba7)   
3.  點擊"Select Source File"和"Select Target Directory"選擇要複製的檔案與要複製的位置( **注意**!!這裡只有英文能正常顯示，中文會變成?圖示)  
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/aaff98fd-526f-49e7-8a73-c0d0ffb4f99e)  
4.  點擊"Start Copy"程式就會開始執行複製，同時下面的進度條也會開始跑  
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/3636e2a6-94da-4c91-b497-985c7247c8c1)  
5.  複製完成後會跳出"File copied successfully!"的訊息  
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/ef8bb815-d764-4b25-8593-26a08d916dab)  
***
## 程式碼簡介
第 1 行：在 Go 中，每個程式都是套件的一部分。我們使用關鍵字來定義它 package。在此範例中，程式屬於 main套件。
第 2 行： 讓我們匯入fmt套件中包含的檔案，
"fmt" 和 "io" 是Go語言標準庫中的包,用於格式化輸入/輸出和實現基本的I/O操作。
"os" 是Go語言標準庫中的包,提供了一些與操作系統相關的功能,比如文件操作、進程控制等。
"path/filepath" 是Go語言標準庫中的包,提供了一些路徑相關的實用函數。
"time" 是Go語言標準庫中的包,用於處理時間和日期相關的操作。
"fyne.io/fyne/v2" 是Fyne跨平台GUI工具包的主包,提供了基本的GUI組件和事件處理機制。
"fyne.io/fyne/v2/app" 是Fyne工具包中的一個包,用於創建和管理應用程序實例。
"fyne.io/fyne/v2/container" 是Fyne工具包中的一個包,提供了一些常用的容器組件,如水平或垂直佈局等。
"fyne.io/fyne/v2/dialog" 是Fyne工具包中的一個包,提供了一些常用的對話框,如文件選擇對話框、消息對話框等。
"fyne.io/fyne/v2/widget" 是Fyne工具包中的一個包,提供了一些常用的GUI組件,如按鈕、標籤、輸入框等。
總而言之,這段代碼導入了Go語言標準庫中的一些常用包,以及Fyne跨平台GUI工具包中的一些核心包,為構建GUI應用程序做好準備
![image](https://github.com/MeowWnag/Golang_Progress_Bar/assets/119922838/56a12293-193e-4495-bf82-e7a8b2148ab6)






