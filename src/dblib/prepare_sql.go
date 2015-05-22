package dblib
import(
    "fmt"
 //   "runtime"
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
    this.stmt,err=this.db.Prepare("SELECT id FROM mytest WHERE name=?" ) // ? = 占位符  
    if err != nil {  
        panic(err.Error())  
    }  
    return 0
}
func (this *DBPresql) Run() int{
    //操作数据库   
    var tmp *sql.Rows
    tmp, err := this.stmt.Query("teset") // 执行插入  
    if err != nil {  
        panic(err.Error())  
    }  
    for tmp.Next(){
        var stmp int
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