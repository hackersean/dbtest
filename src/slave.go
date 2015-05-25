package main
 
import(
    . "snipe"
    "fmt"
)

func main(){
    fmt.Println("go")
    ptr:=new(Snipe)
    ptr.Init("mysql","test:test@tcp(localhost:3306)/test?charset=utf8")
    if ptr==nil{
        fmt.Println("error")
    }
    ptr.Add_Arms()
    var x uint
    for{
        fmt.Scanf("%d",&x)
  //      fmt.Println(x)
        if(x==9){break}
        ptr.Set_Thread_Pool(0,x)
    }
    return 
}