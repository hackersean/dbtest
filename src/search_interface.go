package main
import "fmt"
type PType struct {
     x int
}
// 接收者是指针
func (t *PType) post() {
    fmt.Println(t.x)
}
func (t *PType) change() {
    t.x=5
}
type Inter interface {
     post()
     change()
}

func main(){
    var it Inter
    //var it *Inter //接口不能定义为指针
    pty := PType{4}
    fmt.Println("通过指针输出原值：")
    pty.post()
    it = pty           // 将变量赋值给接口，OK
    fmt.Println("通过接口调用change")
    it.change()        // 接口调用方法，OK
    fmt.Println("通过接口输出值：")
    it.post()          // 接口调用方法，OK
    fmt.Println("通过指针输出值：")
    pty.post()
}