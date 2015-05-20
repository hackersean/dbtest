package main
 
import(
    . "snipe"
    . "fmt"
)

func main(){
    Println("go")
    ptr:=new(Snipe)
    ptr.Init("127.0.0.1:80808")
    if ptr==nil{
        Println("error")
    }
    var x uint
    for{
        Scanf("%d",&x)
  //      Println(x)
        if(x==9){break}
        ptr.Set_Thread_Pool(x)
    }
    return 
}