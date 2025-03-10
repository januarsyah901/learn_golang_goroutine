package learngolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}
	pool.Put("januarsyah")
	pool.Put("akbar")
	pool.Put("ganteng")

	for i := 0; i < 10; i++ {
		go func ()  {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}
	time.Sleep(4*time.Second)
	fmt.Println("Selesai")
}