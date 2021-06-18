package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

	timeFromatDemo()
}

func timeFromatDemo() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)

	//时间计算
	now := time.Now()
	fmt.Println("当前时间:", now)
	//两小时之后的时间
	h, _ := time.ParseDuration("2h")
	nowAfter2Hour := now.Add(h)
	fmt.Println("两小时之后的时间:", nowAfter2Hour)

	//两小时之前的时间
	negativeH, _ := time.ParseDuration("-2h")
	nowBefore2Hour := now.Add(negativeH)
	fmt.Println("两小时之前的时间:", nowBefore2Hour)

	//十分钟之后的时间
	m, _ := time.ParseDuration("10m")
	nowAfter10Minute := now.Add(m)
	fmt.Println("十分钟之后的时间:", nowAfter10Minute)

	//十分钟之前的时间
	negativeM, _ := time.ParseDuration("-10m")
	nowBefore10Minute := now.Add(negativeM)
	fmt.Println("十分钟之前的时间:", nowBefore10Minute)

	//30s之后的时间
	s, _ := time.ParseDuration("30s")
	nowAfter30Second := now.Add(s)
	fmt.Println("30s之后的时间:", nowAfter30Second)

	//30s之前的时间
	negativeS, _ := time.ParseDuration("-30s")
	nowBefore30Second := now.Add(negativeS)
	fmt.Println("30s之前的时间:", nowBefore30Second)

	//1年1个月1天 之后的时间
	fmt.Println("1年2个月3天之后的时间:", now.AddDate(1, 2, 3))

	//2年2个月2天之前的时间
	fmt.Println("2年3个月4天之前的时间:", now.AddDate(-2, -3, -4))

	// 时间戳转时间
	fmt.Println(time.Unix(1623161791954/1000, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Unix(1623161790949/1000, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Unix(1623161183803/1000, 0).Format("2006-01-02 15:04:05"))
	fmt.Println(time.Unix(1623310157520/1000, 0).Format("2006-01-02 15:04:05"))
}
