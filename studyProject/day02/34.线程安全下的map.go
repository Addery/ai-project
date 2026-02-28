//package main
//
//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//var mp = map[string]string{}
//var lock1 = sync.Mutex{}
//var wait2 sync.WaitGroup
//
//func reader() {
//	for {
//		lock1.Lock()
//		fmt.Println(mp["time"])
//		lock1.Unlock()
//	}
//	wait2.Done()
//}
//func writer() {
//	for {
//		lock1.Lock()
//		mp["time"] = time.Now().Format("15:04:05")
//		lock1.Unlock()
//	}
//	wait2.Done()
//}
//
//func main() {
//	wait2.Add(2)
//	go reader()
//	go writer()
//	wait2.Wait()
//}

package main

import (
	"fmt"
	"sync"
	"time"
)

var wait2 sync.WaitGroup
var mp = sync.Map{}

func reader() {
	for {
		fmt.Println(mp.Load("time"))
	}
	wait2.Done()
}
func writer() {
	for {
		mp.Store("time", time.Now().Format("15:04:05"))
	}
	wait2.Done()
}

func main() {
	wait2.Add(2)
	go reader()
	go writer()
	wait2.Wait()

}
