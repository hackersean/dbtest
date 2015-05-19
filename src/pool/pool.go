package pool
import(
    . "fmt"
)


struct Pool{
    var thread_amount uint
    var thread_max uint

}

func (this *Pool) go_thread(uint id,){
    //建立数据库连接
    
    for{
        //sql请求
        Println(id)
        //如果id大于thread_max，就退出线程。
        if id>thread_max{
            break
        }

    }

    //数据库链接断开
}

func init(){
    
    Println("init完毕")
}

