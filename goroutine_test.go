package learngolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)


func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")

	time.Sleep( 1 * time.Second)
}

func DisplayNumber(number int){
	fmt.Println("Display", number)
}

func TestManyGoroutine(t *testing.T) {
	for  i := 0; i<= 10000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep( 5 * time.Second)
}
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}
func(account *BankAccount) AddBalance(amount int){
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}
func(account *BankAccount) GetBalance() int{
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()	
	return balance
}
func TestRWMutex(t *testing.T) {
	account := BankAccount{}
	for i := 0; i<100; i++ {
		go func() {
			for j := 0; j<100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance", account.GetBalance())
}
type UserBalance struct {
	sync.Mutex
	Name string
	Balance int
}
func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}
func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}
func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}
func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}
func TestDeadlock(t *testing.T)  {
	user1 := UserBalance{
		Name: "janu",
		Balance: 1000000,
	}
	user2 := UserBalance{
		Name: "andi",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)


	time.Sleep(3 * time.Second)

	fmt.Println("User ", user1.Name, "Balance", user1.Balance)
	fmt.Println("User ", user2.Name, "Balance", user2.Balance)

}