package snipe
import(
    . "pool"
//    . "dblib"
 //   "fmt"
)   


type Fire_Pool struct{
    Pool
    sql string
    para []string   //参数列表,用户输入的。
}


