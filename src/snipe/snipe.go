package snipe
import(
    . "pool"
    . "dblib"
    "fmt"
) 

type Snipe struct{
    target string      //链接字符串
    dbtype string
    dblink *DBlink     //和数据库的链接
    fire_pool *Pool
    
}

func (this *Snipe) Init(dbtype string,target string){
    this.target=target
    
    this.dblink=new(DBlink)
    this.dblink.Construct(&dbtype,&target)  //创建链接（还没有真的链接）
    

    this.fire_pool=new(Pool)
    this.fire_pool.Init(NewDBsql(this.dblink.DBptr()))
}

//设置线程池大小
func (this *Snipe) Set_Thread_Pool(thread_count_want uint){
    this.fire_pool.Run(thread_count_want)
    return
}


func init(){
    fmt.Println("snipe")
}
