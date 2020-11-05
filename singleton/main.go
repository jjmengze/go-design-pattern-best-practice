package main

import (
	"fmt"
	"sync"
)

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

type fruit struct {
	name string
}

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

func main() {
	cm := NewConnectionManager("apple")
	fmt.Println(cm.GetAddress())

	//  can not new different object, cm is apple forever
	cm = NewConnectionManager("banana")
	fmt.Println(cm.GetAddress())

	// when import a package the fruit object will be set up , everyone can not revise it anymore in runtime.
	f := GetFruit()
	fmt.Println(f.GetName())
}
