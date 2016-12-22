package main

import (
	"fmt"
	"time"
	"math"
)
const (
	Y_M = "2006-01"
	Y_M_D = "2006-01-02"
	Y_M_D_2 = "2006年01月02日"
	Y_M_D_H_I_S = "2006-01-02 15:04:05"
	Y_M_D_H_I_S_2 = "2006年01月02日 15:04:05"
	H_I_S = "15:04:05"
)
func main() {
	/**
	先考察一下，女神个性爱好等等都要了解清楚
	字符型
	 */
	var str string = "你喜欢我什么？"
	fmt.Println(str);
	/**
	 另一种方式
	 */
	str1 := "你喜欢我什么？"
	fmt.Println(str1);
	/**
	表白的 日期时间还记得么，要记得哦
	 */
	now := time.Now()
	//Format 里面是格式，这个时间格式比较奇葩，是GO语言的诞生时间
	fmt.Println("这个是现在时间", now.Format("2006-01-02 15:04:05"));
	//      2016 1 2 3 4 5 6 7
	//      年   月 日 时 分 秒 年 时 区
	day2 := "2015-07-24"
	fmt.Println("这个才是表白日期", day2);
	/**
	女神虽然不告诉你年龄，但你要千方百计的打听得到，以备不时之需
	 */
	birthday := "1993-05-05"
	// Sub 计算两个日期差  得到 年龄
	date1,_ := time.Parse("2006-01-02", birthday)
	subM := now.Sub(date1)
	age :=math.Floor(subM.Hours()/24/365)
	fmt.Println("女神的年龄是", age);
	age1 := 23
	fmt.Println("女神的年龄是", age1);
	/**
	生日你更要知道哦
	 */
	//从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串

	fmt.Println("女神的生日是", birthday);

	/**
	女神有哪些数据是不变的
	常量
	 */
	const sex string = "女"
	fmt.Println("女神的性别当然是：", sex);
	const (
		province = "湖北"
		city = "武汉"
		district = "隐藏信息"
	)
	fmt.Println("省：", province);
	fmt.Println("市：", city);
	fmt.Println("区：", district);
	/**
	多行字符串
	 */
	content := `小家碧玉 天生丽质 完美无暇，
	娇羞可爱 温文尔雅  闭月羞花 端庄优雅 大家闺秀
	楚楚动人 壁月羞花，国色天香 冰清玉洁 聪明伶俐， 眉目如画 `
	fmt.Println("评价：", content);

	/**
	字符串中的字符快速修改
	 */
	content2 := "she is my love"
	// 将字符串 content2 转换为rune数组
	c := []rune(content2)

	// 修改数组第一个元素,这个时候要用单引号
	c[0] = 'S'

	// 创建新的字符串contents保存修改
	contents := string(c)
	fmt.Printf("%s\n", contents)

	fmt.Println("这是我的女神，不是你的，所以你知道得太多了");

	//字符串判断 是否有值 是否为空
	//实际查看的是字符串个数是否为 0
	n := len(str)
	if n == 0 {
		//return 0
	}
	//now:=time.Now()
	t,_ := time.Parse(Y_M_D_H_I_S, now.Format(Y_M_D_H_I_S))
	//tmp:=time.Time(t).Format(Y_M_D_H_I_S)
	fmt.Println(t)

	n1:=Format(time.Now(),Y_M_D_H_I_S)
	fmt.Println(n1)
}
func Format(str interface{}, layout string) string {
	var date time.Time
	var err error
	//判断变量类型
	switch str.(type) {
	case time.Time:
		date = str.(time.Time)
	case string:
		//如果是字符串则转换成 标准日期时间格式
		fmt.Println(str)
		date, err = time.Parse(layout, str.(string))
		if err != nil {
			return ""
		}
	}

	return date.Format(layout)
}
//当前日期时间
func Now() string {
	return Format(time.Now(), Y_M_D_H_I_S)
}
//当前日期
func Date() string {
	return Format(time.Now(), Y_M_D)
}
//当前时间
func Time() string {
	return Format(time.Now(), H_I_S)
}