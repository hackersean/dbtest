package dblib
import(
    "fmt"
 //   "runtime"
//    . "channels"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

//派生自DBsql，使用预处理
type DBPresql struct{
    DBsql      //匿名字段
    stmt *sql.Stmt
}

//-------------------------
//预编译插件的函数
//-------------------------
//在construct函数中预编译。
func (this *DBPresql) Construct() int{
    fmt.Println("prepare Construct")
    var err error
    this.stmt,err=this.db.Prepare(*this.sql) // ? = 占位符  
    if err != nil {  
        panic(err.Error())  
    }  
    return 0
}
func (this *DBPresql) Run() int{
    //操作数据库   
    
    paras_ptr:=this.ch.Pull().(*[]*[]string)
    /*
    for index:=range(*paras_ptr){
        fmt.Println((*paras_ptr)[index])
    }
    */

    //这里需要加循环
    var tmp *sql.Rows
    
    tmp, err := this.stmt.Query("1") // 执行插入  
    if err != nil {  
        panic(err.Error())  
    }  
    for tmp.Next(){
        var stmp string
        tmp.Scan(&stmp)
        fmt.Println(stmp)
    }
   
    fmt.Println("")
    return 0
}

func (this *DBPresql) Destory() int{
    fmt.Println("prepare Destory")
    this.stmt.Close() // main结束是关闭  
    return 0
}
/*
func (this *DBPresql) init(db *sql.DB,sql *string) int{
    this.db=db
    this.sql=sql
    return 0
}
*/