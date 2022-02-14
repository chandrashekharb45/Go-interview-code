package main

import (
	"fmt"
	"sync"
)

func main() {
	// roles := []string{"1", "2", "3"}
	// fmt.Println(stringifyUserIDs(roles))
	var wg sync.WaitGroup
	chErr := make(chan error)

	fmt.Println(Run(task))
	wg.Wait()
}

func task(wg *sync.WaitGroup, chErr chan<- error) {
	defer wg.Done()

}

// func stringifyUserIDs(IDs []string) string {
// 	return "[\"" + strings.Join(IDs, "\",\"") + "\"]"
// }

// func stringifyUserIDs(IDs []string) string {
// 	j, _ := json.Marshal(IDs)
// 	return string(j)
// }

// Run :
func Run(tasks ...func(*sync.WaitGroup, chan<- error)) error {
	n := len(tasks)
	errCh := make(chan error, n)

	w := &sync.WaitGroup{}
	w.Add(n)

	for _, t := range tasks {
		go t(w, errCh)
	}

	w.Wait()
	close(errCh)
	for err := range errCh {
		if err != nil {
			return err
		}
	}
	return nil
}
