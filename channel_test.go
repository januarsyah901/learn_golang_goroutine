package learngolanggoroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	// channel <- "januarsyah"

	// data := <- channel

	// fmt.Println(data)

	// kalo pengen langsung tinggal fmt.Println(<-channel)

	defer close(channel) // lebih afdol kalo di close

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "januarsyah"
		fmt.Println("Data has been sent to channel")
	}()
	data := <-channel
	fmt.Println(data)
	time.Sleep(5  * time.Second)


}
func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "januarsyah akbar"
}
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	defer close(channel) 

	go GiveMeResponse(channel)
	data := <-channel
	fmt.Println(data)
	time.Sleep(5  * time.Second)
}
func OnlyIn(channel chan <- string) {
	time.Sleep(2 * time.Second)
	channel <- "Januarsyah"
}
func OnlyOut(channel <- chan string) {
	data := <-channel
	fmt.Println(data)
}
func TestInOutChannel(t *testing.T) {
	channel := make(chan string)

	defer close(channel) 

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5  * time.Second)
}
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func ()  {
		channel <- "januarsyah"
		channel <- "akbar"
	}()
	go func ()  {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()
	time.Sleep(2 * time.Second)
}
func TestRangeChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func ()  {
		for i := 0; i < 10; i++ {
			channel <- "Data " + strconv.Itoa(i)
		}
		close(channel)
	}()
	for data := range channel {
		fmt.Println(data)
	}
}
func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0

	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2", data)
			counter++
		default :
			fmt.Println("Waiting for data...")
		}
		if counter == 2 {
			break
		}

	}

}

