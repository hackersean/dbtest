package dblib
import(
    "fmt"
 //   "runtime"
    . "channels"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)


//负责数据库连接建立
type DBlink struct{
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

//返回DB指针
func (this *DBlink) DBptr() *sql.DB{
    return this.db
}

//销毁链接
func (this *DBlink) Destory() int{
    fmt.Println("Destory")
    err:=this.db.Close()
    fmt.Println(err)
    return 0
}


//-------------------------
//函数库
//----------------------------
func NewDBpresql(dbptr *sql.DB,sql *string,ch *Channel) *DBPresql{
    dbsql:= new(DBPresql)
    dbsql.init(dbptr,sql,ch)
    return dbsql
}
func NewDBsql(dbptr *sql.DB,sql *string,ch *Channel) *DBsql{
    dbsql:=new(DBsql)
    dbsql.init(dbptr,sql,ch)
    return dbsql
}
func init(){
    fmt.Println("dblib")
}