package channels
import(
    "fmt"
    "math/rand"
)
//后面考虑要不要单独分离成包

type Channel struct{
    chan_array []chan interface{}
    array_len int
}

func (this *Channel) get_subscript() int{
    //代码还没写，预留
    return rand.Intn(this.array_len)
}

func (this *Channel) Init(array_len int,chan_deep int){
    
    for i:=int(0);i<array_len;i++ {
        chan_tmp:=make(chan interface{},chan_deep)
        this.chan_array=append(this.chan_array,chan_tmp)
    }
    this.array_len=array_len
}

func (this *Channel) Push(ptr interface{}) int{
    fmt.Println("Push",len(this.chan_array[this.get_subscript()]))
    this.chan_array[this.get_subscript()] <- ptr
    return 0
}

func (this *Channel) Pull() interface{}{
    return <-this.chan_array[this.get_subscript()]
}

func init(){
    fmt.Println("chanlib")
}