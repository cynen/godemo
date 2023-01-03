package syncdemo

import (
	"demo_2/syncdemo"
	"testing"
	"time"
)

func TestDemo(t *testing.T) {
	//TestMap()
	//SyncMap()

	//go syncdemo.Lock()
	//go syncdemo.Rlock()
	//go syncdemo.WLock()
	//time.Sleep(time.Second * 3)

	go syncdemo.PersonDemo("Rancher_W1")
	go syncdemo.PersonDemo("Rancher_W2")
	syncdemo.PersonDemo("Rancher_W3")
	time.Sleep(time.Second)
}
