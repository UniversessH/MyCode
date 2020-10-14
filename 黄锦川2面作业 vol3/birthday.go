package main

import (
	"fmt"
	"time"
)
var year,mon,day,t1,t2,t3 int

func main() {
	var a int
	t1=int(time.Now().Month())
	t2=int(time.Now().Day())
	t3=int(time.Now().Year())	
    var eachmon[13] int
	eachmon[1]=31
	eachmon[3]=31
	eachmon[5]=31
	eachmon[7]=31
	eachmon[8]=31
	eachmon[10]=31
	eachmon[12]=31
	eachmon[4]=30
	eachmon[6]=30
	eachmon[9]=30
	eachmon[11]=30
	if(t3%400==0){
		eachmon[2]=29;
	}else{
        if(t3%4==0&&t3%100!=0){
			eachmon[2]=29;
	    }else{
			eachmon[2]=28;
		}
	}/*闰年判断*/
	fmt.Printf("What year were you born?")
	fmt.Scanln(&year)
	fmt.Printf("What month were you born?")
	fmt.Scanln(&mon)
	fmt.Printf("What day were you born?")
	fmt.Scanln(&day)
	fmt.Printf("It looks like you were born on %d / %d / %d\n",day,mon,year)
	if(t1==mon){
		if(t2==day){
			fmt.Printf("Aha,today is your birthday!\nHappy birthday to you!")
		}
		if(t2<day){
			a=futrue(eachmon,13,t1,t2)
			fmt.Printf("It looks like there are %d days from your birthday.\nHope you're looking forward to it!",a)
		}
		if(t2>day){
			a=past(eachmon,13,t1,t2)
			fmt.Printf("It looks like your birthday has passed %d days.\nHope you're looking forward to the next one!",a)
		}
	}
	if(t1>mon){
		a=past(eachmon,13,t1,t2)
		fmt.Printf("It looks like your birthday has passed %d days.\nHope you're looking forward to the next one!",a)
	}
	if(t1<mon){
		a=futrue(eachmon,13,t1,t2)
		fmt.Printf("It looks like there are %d days from your birthday.\nHope you're looking forward to it!",a)
	}
}
/*past指生日已经过去 futrue同理，输入均为现在日期*/
func past(eachmons [13]int,size int,mon1,day1 int) int{
	var result int
	if(mon1==mon){
		result=day1-day
	}
	if(mon1>mon){
		result=day1+(eachmons[mon]-day);
		for i:=mon; i<=mon1-2; i++{
			result=result+eachmons[i];
		}
	}
	return result
}
func futrue(eachmons [13]int,size int,mon2,day2 int) int{
	var result int
	if(mon2==mon){
		result=day-day2
	}
	if(mon2<mon){
		result=day+(eachmons[mon]-day2);
		for i:=mon2; i<=mon-2; i++{
			result=result+eachmons[i];
		}
	}
	return result
}