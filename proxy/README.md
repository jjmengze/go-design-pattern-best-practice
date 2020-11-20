# Proxy

使用者想使用某些功能例如展示出照片，一般來說可以直接使用imageShower.ShowImage()的方式。
使用起來就像
```go
var image imageShower
	image = NewRealImage("realImage")
	image.ShowImage()
```

但有些跟業務邏輯無關的事情，例如Audit 、Monitoring、Logging等放在業務邏輯似乎破壞了struct的內聚性。

透過proxy 的方式可以讓proxy裡面呼叫Audit 、Monitoring、Logging等行為，最後在幫我們呼叫原本要做的展示出照片就好。

使用起來就像
```go
	var imageAudit imageShower
	imageAudit = NewProxyImage("proxyImage") // loadFileFromDisk() deferred
	fmt.Println("I want to use image , but user have to be audit so let talk to imageAudit proxy !")
	imageAudit.ShowImage()
	fmt.Println("thanks you imageAudit proxy !")
```
imageAudit偷偷的幫我們做了audit的動作，既不破壞image的內聚性也讓其他附屬的功能與image解耦。
