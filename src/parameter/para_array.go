package parameter
import(
    "fmt"
    "pool"
)
var Default_PARA_ARRAY_LANG uint=10
var MAX_PARA_COUNT uint=10
var MAX_PARA_CHANNEL_DEEP uint=100
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
    pool.Pool_Thread_Base
    desc_array_store []Paramer      //指向Paramer的指针数组
    paras_ptr_arrays_channel chan *[]*[]Paramer      //指向（指向Paramer的指针数组的指针）的数组的指针
}

func (this *ParaProd)Init(para_desc_array []string) chan *[]*[]Paramer{
    this.paras_ptr_arrays_channel=make(chan *[]*[]Paramer,MAX_PARA_CHANNEL_DEEP)
    this.desc_array_store=make([]Paramer,0,MAX_PARA_COUNT)
    for _,para_str:=range(para_desc_array){   
        if len(para_str)>0 && para_str[0]=='#'{
        }else{
            tmp_prt:=new(StaPara)
            tmp_prt.init(&para_str)
            this.desc_array_store=append(this.desc_array_store,tmp_prt)
        }
    }
    return this.paras_ptr_arrays_channel
}

func (this *ParaProd) Construct() int{
    return 0
}

func (this *ParaProd) Run() int{
    paramers_ptr_array:=make([]*[]Paramer,Default_PARA_ARRAY_LANG)
    this.paras_ptr_arrays_channel <- &paramers_ptr_array
    return 0
}

func (this *ParaProd) Destory() int{
    return 0
}