package main
import (
    "fmt"   
    "runtime"
    "sync"
)
var waitgroup sync.WaitGroup  
func play(){
    i:=18
    for{
       i=i*i
       runtime.Gosched()
    }
    waitgroup.Done()
}
func main(){
    num:=runtime.NumCPU()*2
    fmt.Println("线程数为：",num)
    runtime.GOMAXPROCS(num)
    for i:=0;i<128;i++{
        waitgroup.Add(1)
        go play()
    }
    waitgroup.Wait()
    return
}