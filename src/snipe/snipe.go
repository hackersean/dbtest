package snipe
import(
//    . "pool"
    . "dblib"
    "fmt"
//    . "parameter"
) 
var Pool_MAX_Count uint=5;

type Snipe struct{
    target string      //链接字符串
    dbtype string      //数据库类型
    dblink *DBlink
    
    arms_count int
    arms [5]*Armory
    
}

func (this *Snipe) Init(dbtype string,target string){
    this.target=target
    
    this.dblink=new(DBlink)
    this.dblink.Construct(&dbtype,&target)  //创建链接（还没有真的链接）
}

//增加军队
func (this *Snipe) Add_Arms() int{
    //注意，这里需要加锁
    if this.arms_count >= 5 {
        return 1
    }
    ptr:=this.arms_count
    this.arms_count+=1
    //初始化arms
    this.arms[ptr]=new(Armory)
    this.arms[ptr].Init(this.dblink)
    return 0
}

//设置线程池大小
func (this *Snipe) Set_Thread_Pool(id uint,thread_count_want uint){
    this.arms[id].ammo_prod.Run(thread_count_want)
    this.arms[id].fire_point.Run(thread_count_want)
    return
}

//读取线程池大小
func (this *Snipe) Show_Thread_Count(id uint) uint{
    return this.arms[id].fire_point.Show()
}


func init(){
    fmt.Println("snipe")
}
