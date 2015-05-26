package snipe
import(
    . "pool"
    . "parameter"
    . "dblib"
//"fmt"
)   


type Armory struct{   
    sql string
    para []string   //参数列表,用户输入的。
    fire_point Pool
    ammo_prod  Pool
}

func (this *Armory) Init(dblink *DBlink) int{
    //弹药生产
    paradesc:=[]string{"teset","6"}
    var prod *ParaProd=new(ParaProd)
    ch:=prod.Init(paradesc)

    this.ammo_prod.Init(prod)
    //火力点
    tmp:="SELECT id FROM mytest WHERE name=? AND id=?"
    
    this.fire_point.Init(NewDBpresql(dblink.DBptr(),&tmp,ch))
    return 0
}





