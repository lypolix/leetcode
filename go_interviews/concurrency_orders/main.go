package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Warehouse struct {
	items map[string]int
	mu sync.RWMutex
}

func NewWarehouse() *Warehouse {
	return &Warehouse{
		items: map[string]int{
			"phone":  10,
			"laptop": 25,
			"tablet": 5,
		},
	}
}

func (w *Warehouse) ReserveItem(item string, qty int) bool {
	w.mu.Lock()

	defer w.mu.Unlock()
	count, ok := w.items[item] 

	if !ok {
		return false
	}

	if count >= qty {
		w.items[item] -= qty  
		return true
	}


	return false

}

type Order struct {
	ID       int
	Item     string
	Quantity int
}

func ProcessOrders(warehouse *Warehouse, order Order) {
	if warehouse.ReserveItem(order.Item, order.Quantity) {
		fmt.Println("Заказ обработан", order.ID)
	} else {
		fmt.Println("Заказ отменен", order.ID)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	warehouse := NewWarehouse()
	workerpool := make(chan struct{}, 3)
	defer close(workerpool)
	var wg sync.WaitGroup

	orders := make(chan Order, 10)

	go func() {
		for i := 1; i <= 10; i++ {
			item := []string{"phone", "laptop", "tablet"}[rand.Intn(3)]
			qty := rand.Intn(3) + 1
			orders <- Order{ID: i, Item: item, Quantity: qty}
			time.Sleep(time.Millisecond * 200)
		}

		close(orders)

	}()

	t := time.NewTicker(time.Second * 3)

	loop: 
	for order := range orders {

		select {
		case <- t.C:
			break loop
		default:
			workerpool <- struct{}{}
			wg.Add(1)

			go func() { 
				defer func() {
					wg.Done()
					<-workerpool
				}()

				ProcessOrders(warehouse, order)
			}()
		}
		
	}

	wg.Wait()

	fmt.Println("Все заказы обработаны")
}
