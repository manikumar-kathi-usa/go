package main

import (
	"fmt"
	"sync"
	"time"
)

// normal program with out paralal processing
/*func main() {
	foodsCenter := []string{"chicfil-a", "dominos-pitza", "startbuks", "bucees"}
	fmt.Println(foodsCenter)

	for _, f := range foodsCenter {

		pickupFood(f)

	}

}*/
// using normal go routines without waitedGroup
/* func main() {
	foodCenter := []string{"chickfil-a", "dominos-pitza", "startbuks", "bucees"}
	for _, f := range foodCenter {
		func(pickup string) {
			pickupFood(pickup)
		}(f)
	}
}
*/

func main() {
	started := time.Now()
	fc := []string{"chickfil-a", "dominos-pitza", "startbuks", "bucees"}

	var wg sync.WaitGroup

	wg.Add(len(fc)) // this 

	for _, p := range fc {
		go func(p string) {
			pickupFood(p)
			wg.Done()
		}(p)
	}
	wg.Wait()
	fmt.Printf("done in %v \n", time.Since(started))
}

func pickupFood(procCenter string) {
	fmt.Printf("picking up..... %s  \n", procCenter)
	time.Sleep(2 * time.Second)
	fmt.Printf("picked up  %s \n", procCenter)
}
