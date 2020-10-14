package main

import (
	"fmt"
	"time"
)
var year,mon,day int
func main() {
	var a int
	fmt.Printf("What year were you born?")
	fmt.Scanln(&year)
	fmt.Printf("What month were you born?")
	fmt.Scanln(&mon)
	fmt.Printf("What day were you born?")
	fmt.Scanln(&day)
	fmt.Printf("It looks like you were born on %d / %d / %d\n",day,mon,year)
	t1:=int(time.Now().Month())
	t2:=int(time.Now().Day())
	if(t1==mon){
		if(t2==day){
			fmt.Printf("Aha,today is your birthday!\nHappy birthday to you!")
		}
		if(t2<day){
			a=futrue(t1,t2)
			fmt.Printf("It looks like there are %d days from your birthday.\nHope you're looking forward to it!",a)
		}
		if(t2>day){
			a=past(t1,t2)
			fmt.Printf("It looks like your birthday has passed %d days.\nHope you're looking forward to the next one!",a)
		}
	}
	if(t1>mon){
		a=past(t1,t2)
		fmt.Printf("It looks like your birthday has passed %d days.\nHope you're looking forward to the next one!",a)
	}
	if(t1<mon){
		a=futrue(t1,t2)
		fmt.Printf("It looks like there are %d days from your birthday.\nHope you're looking forward to it!",a)
	}
}
/*past指生日已经过去 futrue同理，输入均为现在日期*/
func past(mon1,day1 int) int{
	var result int
	if(mon1==mon){
		result=day1-day
	}
	if(mon1>mon){
		result=day1+(30-day)+(mon1-mon-1)*30
	}
	return result
}
func futrue(mon2,day2 int) int{
	var result int
	if(mon2==mon){
		result=day-day2
	}
	if(mon2<mon){
		result=day+(30-day2)+(mon-mon2-1)*30
	}
	return result
}