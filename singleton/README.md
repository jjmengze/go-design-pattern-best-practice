# 單例模式(Singleton Pattern)
一個類只允許建立唯一一個物件（或者實例），那這個類就是一個單例類，這種設計模式就叫作單例設計模式，簡稱單例模式Singleton。

## 使用sync.Once package 實作單例模式

```go
var connectionManagerOnce sync.Once
var connectionManager *ConnectionManager

type ConnectionManager struct {
	address string
}

func (cm *ConnectionManager) GetAddress() string {
	return cm.address
}

func NewConnectionManager(address string) *ConnectionManager {
	connectionManagerOnce.Do(func() {
		connectionManager = &ConnectionManager{
			address: address,
		}
	})
	return connectionManager
}
```

## 使用init function 實作模式

缺點很明顯我們無法在import package 時帶入參數，例如fruit name是什麼。
```go
var singletonFruit *fruit

func init() {
	//initialize static instance on load
	singletonFruit = &fruit{name: "fruit"}
}

//GetFruit - get singleton instance pre-initialized
func GetFruit() *fruit {
	return singletonFruit
}

func (f *fruit) GetName() string {
	return f.name
}
```