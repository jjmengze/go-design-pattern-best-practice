# Decorator
使用者像是吃糖果，糖果紙裡面包什麼糖果（只要糖果符合某個規範就好）可以自行決定。

包裝是動態的，使用者可以自由決定包裝什麼（只要符合包裝的某個規範）。

以個範例是我們要在圖畫紙上畫一個垂直的拉霸，這個垂直的拉霸會復用`simplewindow`的實作draw畫面的部分（簡易的畫風），再加上自己draw拉霸的部分。
使用起來就像
```go
var window = NewVerticalScrollBars(
		NewVerticalScrollBars(new(SimpleWindow)))
    window.Draw()
	fmt.Println(window.GetDescription())
```

如果現在範例變成我們在圖畫紙上畫一格垂直的拉霸，這個垂直的拉霸要付用`complexwindow`的實作draw換面的部分（複雜的畫風），再加上自己draw拉霸的部分。
使用起來就像
```go
var window = NewVerticalScrollBars(
		NewVerticalScrollBars(new(ComplexWindow)))
    window.Draw()
	fmt.Println(window.GetDescription())
```