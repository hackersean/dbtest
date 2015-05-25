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
    Get(*string) int
}

func init(){
    fmt.Println("Parameter")
}
//----------------------------------
//para量产基地
//--------------------------------
type ParaProd struct{
    paras_count uint
    pool.Pool_Thread_Base
    desc_array_store []Paramer      //指向Paramer的数组
    paras_ptr_arrays_channel chan *[]*[]string      //指向（指向string(变量值)的指针数组的指针）的数组的指针
}

func (this *ParaProd)Init(para_desc_array []string) chan *[]*[]string{
    this.desc_array_store=make([]Paramer,0,MAX_PARA_COUNT)
    this.paras_ptr_arrays_channel=make(chan *[]*[]string,MAX_PARA_CHANNEL_DEEP)
    for index:=range(para_desc_array){   
        if len(para_desc_array[index])>0 && para_desc_array[index][0]=='#'{

        }else{
            tmp_prt:=new(StaPara)
            tmp_prt.init(&para_desc_array[index])
            this.desc_array_store=append(this.desc_array_store,tmp_prt)
        }
        this.paras_count++
    }
    return this.paras_ptr_arrays_channel
}

func (this *ParaProd) Construct() int{
    return 0
}

//获取一套parasd的string
func (this *ParaProd) Get_Para_array() *[]string{
    var para_ptr_array []string=make([]string,this.paras_count)
    for index:=range(this.desc_array_store){
        this.desc_array_store[index].Get(&para_ptr_array[index])
    }
    
    return &para_ptr_array
}

//获取多套paras，组合成paras数组，再把地址放入channel。
func (this *ParaProd) Run() int{
    paramers_ptr_array:=make([]*[]string,0,Default_PARA_ARRAY_LANG)
    for i:=uint(0);i<Default_PARA_ARRAY_LANG;i++ {
        paramers_ptr_array=append(paramers_ptr_array,this.Get_Para_array())
    }
    
    this.paras_ptr_arrays_channel <- &paramers_ptr_array
    return 0
}

func (this *ParaProd) Destory() int{
    return 0
}