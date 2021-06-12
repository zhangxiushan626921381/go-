package main

import "fmt"

/*
type skills struct{
	名称
	耗蓝
	冷却
	范围
	伤害
}
定义结构体切片保存技能信息
type role struct{
	名称
	等级
	经验
	钻石
	金币
	生命值
	攻击力
	暴击
	防御
	血量
	蓝量
}
*/
/*
type 信用卡 struct{
	卡号
	名字
	余额
	有效期
	密码
	银行信息
	消费记录
}
type 消费记录 struct{
	卡号
	时间
	消费ID
	流水号
	消费金额
	备注
}
*/
type role struct {
	skills   string
	shanghai int
}

func main() {
	//结构体
	var s role
	s.shanghai = 569
	s.skills = "杀"
	fmt.Println(s)
	//数组
	s1 := [2]role{{"ss", 102}, {"sss", 104}}
	fmt.Println(s1)
	//切片
	s2 := []role{{"aaa", 56}, {"asa", 89}}
	s2 = append(s2, role{skills: "sss", shanghai: 56})
	fmt.Println(s2)

}
