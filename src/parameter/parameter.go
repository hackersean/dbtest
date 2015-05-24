package parameter

//参数结构体
//静态参数，用户的指定参数，无需随机什么的。
type StaPara struct{
    desc *string   //描述
}

func (this *StaPara) init(desc *string) int{
    this.desc=desc
    return 0
}

func (this *StaPara) Get() *string{
    return this.desc
}

//随机数字参数
type IntRandPara struct{
    max int64    //最大值
    min int64    //最小值

}

func (this *IntRandPara) Get() *string{
     return nil
}