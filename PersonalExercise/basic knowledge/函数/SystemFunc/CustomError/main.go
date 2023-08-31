package main

import (
	"errors"
	"fmt"
)

func main(){
	err:=test()
	//为nil不报错，不为nil报错
	if err !=nil{
		fmt.Println("自定义错误；",err)
		panic(err)//如果报错则后面的代码就不执行
	}

	fmt.Println("上面的除法操作执行成功。")
	fmt.Println("正常执行下面的逻辑。")
}
func test()(err error){//error类型的函数为专门判断是否出错的函数类型
	num1:=10
	num2:=0
	if num2==0{
		return errors.New("除数不能为0")
	}else{
		result:=num1/num2
		fmt.Println(result)
		//没有错误，返回0值
		return nil
	}

}