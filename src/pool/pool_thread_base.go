package pool
import(
    "runtime"
)

type Pool_Thread_Base struct{
    
}
//构造函数
func (this *Pool_Thread_Base) Construct() int{
    return 0
}
//运行的线程

func (this *Pool_Thread_Base) Run() int{
    runtime.Gosched()
    return 0
}

//线程退出时的析构函数
func (this *Pool_Thread_Base) Destory() int{
    return 0
}
