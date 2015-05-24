package parameter
import(
    "fmt"
)
var MAX_PARA_COUNT uint=10
//参数获取的接口
type Paramer interface{
    Get() *string
}

func init(){
    fmt.Println("Parameter")
}
//----------------------------------
//para量产基地
//--------------------------------
type ParaProd struct{
    desc_array_store []Paramer      //指向Paramer的指针数组

}

func (this *ParaProd)Init(para_desc_array []string){
    this.desc_array_store=make([]Paramer,0,MAX_PARA_COUNT)
    for _,para_str:=range(para_desc_array){
        
        if len(para_str)>0 && para_str[0]=='#'{
        }else{
            tmp_prt:=new(StaPara)
            tmp_prt.init(&para_str)
            this.desc_array_store=append(this.desc_array_store,tmp_prt)
        }
    }
}