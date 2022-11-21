package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
)

func main() {
	group := errgroup.Group{}

	n := 10
	group.SetLimit(n)
	for i := 0; i < n; i++ {
		id := i
		/*group.Go(func() error {
			fmt.Printf("[%d] something\n", id)
			return errors.New("random error Go")
		})*/

		ok := group.TryGo(func() error {
			fmt.Printf("[%d] TryGo\n", id)
			return fmt.Errorf("random error Go %d", id)
		})
		if !ok {
			fmt.Printf("could not initiate TryGo")
		}
	}

	err := group.Wait()
	fmt.Printf("error=====", err)

}
