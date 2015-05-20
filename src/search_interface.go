package main

import "fmt"

type PType struct {
     x int
}
type Inter interface {
     post()
     change()
}


// 接收者是指针
func (t *PType) post() {
    fmt.Println(t.x)
}

func (t *PType) change() {
    t.x=5
}

func main(){
    var it Inter
    //var it *Inter //接口不能定义为指针
    pty := new(PType)
    
    it = pty // 将变量赋值给接口，OK
    it.change() // 接口调用方法，OK
    it.post() // 接口调用方法，OK
    pty.post()

}