package pool
import(
    "fmt"
)

//接口，可以用于传入不同类型的结构体。
type Tasker interface {
    Construct() int
    Run() int
    Destory() int
}

type Pool struct{
    thread_count_now uint
    task Tasker
}

//线程池初始化,不运行
func (this *Pool) Init(task Tasker){
    this.task=task
}

//指定协程数量，运行。
func (this *Pool) Run(thread_count_want uint){
    //此处需要加锁   
    for ;this.thread_count_now<thread_count_want;{
        this.thread_count_now+=1
        go this.thread(this.thread_count_now)
    }
    this.thread_count_now=thread_count_want
}

//设置协程数量
func (this *Pool) Set(thread_count_want uint){
    this.Run(thread_count_want)
}

func (this *Pool) thread(id uint){
    fmt.Println(id)
    //建立数据库连接
    this.task.Construct()
    for{
        //sql请求
        this.task.Run()
  //      Println(id)
        //如果id大于thread_max，就退出线程。
        if id>this.thread_count_now{
            break
        }

    }
    //数据库链接断开
    this.task.Destory()    
}

func NewPool(thread_count_want uint)(*Pool){
    ptr := new(Pool)
    ptr.Run(thread_count_want)
    return ptr
}


func init(){
    
    fmt.Println("pool")
}

