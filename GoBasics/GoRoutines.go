package GoBasics

import (
	"awesomeProject/DataBases"
	"fmt"
	"sync"
	"time"
)

/*
*
GOROUTINES
Go routines are lightweight threads of execution. They are used to execute functions concurrently.
*/
var loopCount = 1000
var waitTime = 0

var wg = sync.WaitGroup{}

var roster = []string{}
var stopWatch = StopWatch{}
var timeSequencial time.Duration
var timeConcurrent time.Duration

func ExecuteConcurrentReadTest() {
	var err error

	//Sequential Reads
	Label("Sequential Reads")

	stopWatch.Start()
	executeRead(false)
	stopWatch.Stop()
	timeSequencial, err = stopWatch.Elapse()
	if err != nil {
		panic(err)
	}

	Line()

	//Concurrent Reads
	Label("Concurrent Reads")
	stopWatch.Reset()
	stopWatch.Start()
	executeRead(true)
	stopWatch.Stop()
	timeConcurrent, err = stopWatch.Elapse()
	if err != nil {
		panic(err)
	}

	Line()
	println("Sequential Time: " + timeSequencial.String())
	println("Concurrent Time: " + timeConcurrent.String())

}

func ExecuteConcurrentWriteTest() {
	var err error

	//No Mutex
	Label("No Mutex")

	roster = []string{}
	stopWatch.Start()
	executeWrite(roster, false)
	stopWatch.Stop()
	timeSequencial, err = stopWatch.Elapse()
	if err != nil {
		panic(err)
	}

	for index, value := range roster {
		fmt.Println(index, value)
	}

	Line()

	//Concurrent Reads
	Label("With Mutex")
	roster = []string{}
	stopWatch.Reset()
	stopWatch.Start()
	executeWrite(roster, false)
	stopWatch.Stop()
	timeConcurrent, err = stopWatch.Elapse()
	if err != nil {
		panic(err)
	}

	Line()

	for index, value := range roster {
		fmt.Println(index, value)
	}
	println("Sequential Time: " + timeSequencial.String())
	println("Concurrent Time: " + timeConcurrent.String())

}

func executeRead(isConcurrent bool) {
	//Read from the database non concurrently
	for i := 0; i < loopCount; i++ {
		//Concurrent Reading
		if isConcurrent {
			wg.Add(1)
			go func() {
				fighter := DataBases.ReadFighterDb(waitTime)
				fmt.Println(fighter)
				wg.Done()
			}()
		} else {
			//Sequential Reading
			fighter := DataBases.ReadFighterDb(waitTime)
			fmt.Println(fighter)
		}
	}

	if isConcurrent {
		wg.Wait()
	}
}

func executeWrite(roster []string, isMutex bool) {
	mutex := sync.Mutex{}
	//Read from the database non concurrently
	for i := 0; i < loopCount; i++ {
		//Concurrent Reading
		wg.Add(1)
		go func() {
			if isMutex {
				mutex.Lock() //This is like the Syncrhonize Keyword In Java.  It makes sure each thread waits for the other to finish before it starts.
			}
			fighter := DataBases.ReadFighterDb(waitTime)
			roster = append(roster, fighter)
			if isMutex {
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	if isMutex {
		wg.Wait()
	}
}
