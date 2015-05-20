package snipe
import(
    . "pool"
    . "dblib"
    . "fmt"
) 

type Snipe struct{
    target string
    fire_pool *Pool
    
}

func (this *Snipe) Init(target string){
    this.target=target

    this.fire_pool=new(Pool)
    this.fire_pool.Init(NewDBlink(&this.target))
}

//设置线程池大小
func (this *Snipe) Set_Thread_Pool(thread_count_want uint){
    this.fire_pool.Run(thread_count_want)
    return
}


func init(){
    Println("snipe")
}
