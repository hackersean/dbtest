package snipe
import(
//    . "pool"
    . "dblib"
    "fmt"
    . "parameter"
) 
var Pool_MAX_Count uint=5;

type Snipe struct{
    target string      //链接字符串
    dbtype string      //数据库类型
    dblink *DBlink

    pools [5]*Fire_Pool
    
}

func (this *Snipe) Init(dbtype string,target string){
    paradesc:=[]string{"1","helloworld"}
    var prod *ParaProd=new(ParaProd)
    prod.Init(paradesc)


    this.target=target
    
    this.dblink=new(DBlink)
    this.dblink.Construct(&dbtype,&target)  //创建链接（还没有真的链接）
    
    this.pools[0]=new(Fire_Pool)
    tmp:="SELECT name FROM mytest WHERE id=?"
    this.pools[0].Init(NewDBpresql(this.dblink.DBptr(),&tmp))
}

//设置线程池大小
func (this *Snipe) Set_Thread_Pool(id uint,thread_count_want uint){
    this.pools[id].Run(thread_count_want)
    return
}

//读取线程池大小
func (this *Snipe) Show_Thread_Count(id uint) uint{
    return this.pools[id].Show()
}


func init(){
    fmt.Println("snipe")
}
