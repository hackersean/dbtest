package dblib
import(
    "fmt"
 //   "runtime"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


//负责数据库连接建立
type DBlink struct{
    db *sql.DB
}

//负责数据库sql拼凑，随机数生成
type DBsql struct{
    db *sql.DB
}
//--------------------------------
//数据库建立函数
//--------------------------------
//内部调用函数
func (this *DBlink) init() int{
    
    return 0
}
//建立连接
func (this *DBlink) Construct(dbtype *string,target *string) int{
    fmt.Println("Construct")
    db, err := sql.Open(*dbtype, *target)
    if err != nil {
        fmt.Println("准备失败", err.Error())
        return 1
    }
    db.SetMaxIdleConns(1)
 //   db.SetMaxOpenConns(5)
    
    err = db.Ping()
    if err != nil {
        fmt.Println("连接数据库失败", err.Error())
        return 1
    }
    this.db=db
    return 0
}

func (this *DBlink) DBptr() *sql.DB{
    return this.db
}

func (this *DBlink) Destory() int{
    fmt.Println("Destory")
    err:=this.db.Close()
    fmt.Println(err)
    return 0
}




//--------------------------------
//SQL执行函数
//--------------------------------
func (this *DBsql) Construct() int{
    fmt.Println("sql Construct")
    return 0
}

func (this *DBsql) Run() int{
    //查询数据库

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
    rows.Close()
    return 0
}
func (this *DBsql) Destory() int{
    fmt.Println("Destory")
    return 0
}

//-------------------------
//函数库
//----------------------------
func NewDBsql(dbptr *sql.DB) *DBsql{
    return &DBsql{dbptr}
}
func init(){
    fmt.Println("dblib")
}