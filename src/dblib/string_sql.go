package dblib
import(
    "fmt"
 //   "runtime"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


//负责数据库sql拼凑
type DBsql struct{
    db *sql.DB
}

//--------------------------------
//SQL执行函数
//--------------------------------
func (this *DBsql) Construct() int{
    fmt.Println("sql Construct")
    return 0
}

func (this *DBsql) Run() int{

    rows, err := this.db.Query("select name from mytest;")
    if err != nil {
        fmt.Println("查询数据库失败", err.Error())
        return 1
    }
    for rows.Next() { //开始循环
        var lvs string
        rerr := rows.Scan(&lvs)  //数据指针，会把得到的数据，往刚才id 和 lvs引入
        if rerr == nil {
 //             fmt.Println(lvs)      
        }
    }

 //   rows.Close()
    return 0
}
func (this *DBsql) Destory() int{
    fmt.Println("Destory")
    return 0
}