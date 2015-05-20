package dblib
import(
    . "fmt"
    "runtime"
)


//负责数据库连接建立
type DBlink struct{
    target *string
    
}

//负责数据库sql拼凑，随机数生成
type DBsql struct{
    
}
//--------------------------------
//数据库建立函数
//--------------------------------

func (this *DBlink) Construct() int{
    Println("Construct")
    return 0
}

func (this *DBlink) Run() int{
    runtime.Gosched()
    return 0
}

func (this *DBlink) Destory() int{
    Println("Destory")
    return 0
}




//--------------------------------
//SQL拼凑函数
//--------------------------------


//-------------------------
//函数库
//----------------------------
func NewDBlink(target *string) *DBlink{
    return &DBlink{target}
}

func init(){
    Println("dblib")
}