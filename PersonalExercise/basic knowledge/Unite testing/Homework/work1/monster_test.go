package work1

import "testing"

// 反序列化测试
func TestReStore(t *testing.T) {
	monster := Monster{}     //创建一个Monster的实例monster
	res := monster.ReStore() //res是一个bool类型的
	if !res {                //如果为真
		t.Fatalf("monster.ReStore()错误，希望=%v 实际=%v", true, res)
	}
	//进一步判断
	if monster.Name != "刘岱" {
		t.Fatalf("monster.ReStore()错误，希望=%v 实际=%v", "刘岱", monster.Name)
	}
	t.Logf("monster.Store()测试成功")
}

// 测试用例
/*想要先测试TestStore 应为要创建对应的路径和写入操作*/
func TestStore(t *testing.T) {
	monster := Monster{
		Name:  "刘岱",
		Age:   23,
		Skill: "唱跳",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误，希望=%v 实际=%v", true, res)
	}
	t.Logf("monster.Store()测试成功")
}
