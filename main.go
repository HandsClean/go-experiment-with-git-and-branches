// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// type Mystring string

// func main() {
// 	fmt.Println("Hello, playground")

// 	var str string = "abc"
// 	var my_str Mystring = Mystring(str) //gives a compile error
// 	//Provide seed
// 	rand.Seed(time.Now().Unix())

// 	//Generate a random array of length n
// 	fmt.Println(rand.Perm(10))
// 	var slice = rand.Perm(10)
// 	fmt.Println(my_str)
// 	for i, val := range slice {
// 		fmt.Println(i, val)
// 	}
// }
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Result struct {
	Time  time.Time
	Value int
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	//Provide seed
	rand.Seed(time.Now().Unix())

	//Generate a random array of length n
	fmt.Println(rand.Perm(25))
	fmt.Println(rand.Perm(10))

	//	var wg2 sync.WaitGroup
	// Perform 3 concurrent transactions against the database.

	//var slice = rand.Perm(25)
	var slice = []int{12, 32, 56, 1, 2, 6, 8, 112, 7, 8, 13, 14, 21, 63}
	fmt.Println(slice)
	chmin := make(chan int, len(slice))
	chmax := make(chan int, len(slice))
	chanexit1 := make(chan Result, len(slice))
	chanexit2 := make(chan Result, len(slice))
	done := make(chan struct{}, 2)
	pivot := 17

	go sender(pivot, slice, chmin, chmax, &wg)
	go receiverchmin(chmin, chanexit1, done, &wg)
	go receinverchmax(chmax, chanexit2, done, &wg)

	//go Print(chmin, &wg, " min")
	//	go Print(chmax, &wg, " max")

	//go PrintStruct(chanexit1, " minstruct", &wg)
	//	go PrintStruct(chanexit2, " maxstruct", &wg)

	//var wg2 sync.WaitGroup
	//	wg2.Add(2)
	go PrintStruct(chanexit1, " min", " minstruct", &wg)
	go PrintStruct(chanexit2, " max", " maxstruct", &wg)
	//	wg2.Wait()
	wg.Wait()
	//go receiverchmin(chmin, chanexit1, done, &waitGroup)
	//go receiverchmin(chmax, chanexit2, done, &waitGroup)

	// go func() {
	// 	waitGroup.Wait()
	// 	close(chanexit1)
	// 	close(chanexit2)
	// }()
	// for {
	// 	s, more := <-chanexit1
	// 	if more == true {
	// 		fmt.Println(s, " min")
	// 		//	fmt.Println("received job", j)
	// 	} else {
	// 		fmt.Println("received all jobs min")

	// 		// _, ok := <-done
	// 		// if !ok {
	// 		// 	done <- struct{}{}
	// 		// }

	// 		break
	// 	}
	// }

	//<-done
	fmt.Println("ended")
}

func sender(pivot int, sl []int, chmin chan int, chmax chan int, wg *sync.WaitGroup) {
	for i, _ := range sl {
		if sl[i] <= pivot {
			chmin <- sl[i]
		} else {
			chmax <- sl[i]
		}

	}

	defer wg.Done()
	defer close(chmax)
	defer close(chmin)
	fmt.Println(len(chmin), " sender min")
	fmt.Println(len(chmax), " sender max")
	// defer func() {
	// 	done <- struct{}{}
	// 	done <- struct{}{}
	// }()
	return
}
func receiverchmin(chmin chan int, chanexit1 chan Result, done chan struct{}, wg *sync.WaitGroup) {
	//<-done
	for {
		select {
		case msg1, ok := <-chmin:
			if ok != true {
				for ms := range chmin {
					res := Result{
						time.Now(),
						ms,
					}
					chanexit1 <- res
				}

				fmt.Println("received job min after close chmin ", msg1)

				close(chanexit1)
				wg.Done()
				fmt.Println(len(chmin), " what left min")
				fmt.Println(len(chanexit1), " what have min")
				return
			} else {
				res := Result{
					time.Now(),
					msg1,
				}
				fmt.Println("received job min ", msg1)
				chanexit1 <- res
			}

		}
	}

	// fmt.Println(len(chmin), " begin min")
	// for n := range chmin {

	// 	res := Result{
	// 		time.Now(),
	// 		n,
	// 	}
	// 	fmt.Println("received job min ", n)
	// 	chanexit1 <- res
	// }
	// defer close(chanexit1)
	// defer wg.Done()
	// defer fmt.Println(len(chmin), " what left min")
	// defer fmt.Println(len(chanexit1), " what have min")
	// return

	// for {
	// 	j, more := <-chmin
	// 	if more == true {
	// 		res := Result{
	// 			time.Now(),
	// 			j,
	// 		}
	// 		chanexit1 <- res
	// 		//	fmt.Println("received job", j)
	// 	} else {
	// 		fmt.Println("received all jobs min")

	// 		// _, ok := <-done
	// 		// if !ok {
	// 		// 	done <- struct{}{}
	// 		// }
	// 		defer close(chanexit1)
	// 		defer wg.Done()
	// 		return

	// 	}
	// }

}
func receinverchmax(chmax chan int, chanexit2 chan Result, done chan struct{}, wg *sync.WaitGroup) {

	for {
		select {
		case msg1, ok := <-chmax:
			if ok != true {
				for ms := range chmax {
					res := Result{
						time.Now(),
						ms,
					}
					chanexit2 <- res
				}

				fmt.Println("received job max after close chamax ", msg1)

				close(chanexit2)
				wg.Done()
				fmt.Println(len(chmax), " what left max")
				fmt.Println(len(chanexit2), " what have max")
				return
			} else {
				res := Result{
					time.Now(),
					msg1,
				}
				fmt.Println("received job max ", msg1)
				chanexit2 <- res
			}

		}
	}

	//	<-done
	// fmt.Println(len(chmax), " begin max")
	// for n := range chmax {

	// 	res := Result{
	// 		time.Now(),
	// 		n,
	// 	}
	// 	fmt.Println("received job max ", n)
	// 	chanexit2 <- res
	// }

	// defer close(chanexit2)
	// defer wg.Done()
	// defer fmt.Println(len(chmax), " what left max")
	// defer fmt.Println(len(chanexit2), " what have max")
	// return
	// for {
	// 	j, more := <-chmax
	// 	if more == true {
	// 		res := Result{
	// 			time.Now(),
	// 			j,
	// 		}
	// 		chanexit2 <- res
	// 		//	fmt.Println("received job", j)
	// 	} else {
	// 		fmt.Println("received all jobs max")
	// 		// _, ok := <-done
	// 		// if !ok {
	// 		// 	done <- struct{}{}
	// 		// }
	// 		defer close(chanexit2)
	// 		defer wg.Done()
	// 		return
	// 	}
	// }

}
func Print(ch <-chan int, wg *sync.WaitGroup, str string) {
	for n := range ch {
		fmt.Println(n, str)
	}

	defer wg.Done()
}
func PrintStruct(ch <-chan Result, str2, str string, wg *sync.WaitGroup) {

	for {
		select {
		case msg1, ok := <-ch:
			if ok != true {
				for n := range ch {
					fmt.Println("PRINT job min after close receiver of results ", n, str)
				}

				wg.Done()

				return
			} else {
				fmt.Println(msg1, str)
			}

		}
	}

	// 	fmt.Println(len(ch), " begin to print", str2)
	// 	for n := range ch {
	// 		fmt.Println(n, str)
	// 	}
	// 	defer wg.Done()
}
