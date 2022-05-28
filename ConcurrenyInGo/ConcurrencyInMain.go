package main

import (
	"fmt"
	"sync"
	"time"
)

type TimeObj struct
{
	timeValue string
	hour int
	min int
	sec int
}


func GetTimeWG(wg *sync.WaitGroup){
	defer wg.Done()

	for i := 0; i < 5; i++{
		var timeObj TimeObj
	
		currentTimeStamp := time.Now()
	
	
		timeObj.hour = currentTimeStamp.Hour()
		timeObj.min = currentTimeStamp.Minute()
		timeObj.sec = currentTimeStamp.Second()
	
		timeObj.timeValue = fmt.Sprintf("%02d*%02d*%02d...\n", timeObj.sec, timeObj.min, timeObj.hour)
	
		fmt.Print(timeObj.timeValue)

		time.Sleep(time.Second)

	}
}



func GetTimeChannel(channel chan TimeObj){
	
	var timeObj TimeObj

	currentTimeStamp := time.Now()


	timeObj.hour = currentTimeStamp.Hour()
	timeObj.min = currentTimeStamp.Minute()
	timeObj.sec = currentTimeStamp.Second()

	timeObj.timeValue = fmt.Sprintf("%02d*%02d*%02d...\n", timeObj.sec, timeObj.min, timeObj.hour)

	channel <- timeObj

	
}



func main(){
	
	fmt.Println("Wellcome to the clock")

	var wg sync.WaitGroup

	wg.Add(1)

	go GetTimeWG(&wg)

	wg.Wait()

	fmt.Println("New channel clock")
	
	var x TimeObj
	channel := make(chan TimeObj)

	for count := 0; count < 10; count++{

		go GetTimeChannel(channel)

		x = <- channel

		fmt.Print(x.timeValue)


		time.Sleep(time.Second)

	}

	wg.Add(1)

	timeWGFunc := func(){
	fmt.Println("New in line func clock")
		for i := 0; i < 5; i++{
			var timeObj TimeObj
		
			currentTimeStamp := time.Now()
		
		
			timeObj.hour = currentTimeStamp.Hour()
			timeObj.min = currentTimeStamp.Minute()
			timeObj.sec = currentTimeStamp.Second()
		
			timeObj.timeValue = fmt.Sprintf("%02d*%02d*%02d...\n", timeObj.sec, timeObj.min, timeObj.hour)
		
			fmt.Print(timeObj.timeValue)

			time.Sleep(time.Second)
		}
		wg.Done()
	}

	timeWGFunc()

	
}