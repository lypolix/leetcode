package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Data struct {
	Temperture map[string]int
	mu sync.RWMutex
}

func NewData(interval time.Duration) *Data {
	ticker := time.NewTicker(interval)

	newData := &Data{}

	go func() {
		defer ticker.Stop()
		for range ticker.C {
			newData.UpdateTemperature()
		}
	}()

	return newData
}

func (d *Data) UpdateTemperature() {
	wg := &sync.WaitGroup{}
	for city := range d.Temperture {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tmp := WeatherForecast(city)
			d.mu.Lock()
			d.Temperture[city] = tmp
			d.mu.Unlock()
		}()
	}
	wg.Wait()
	
	
}

func (d *Data) GetTemperature(city string) (int, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	t, ok := d.Temperture[city]
	if !ok {
		return 0, fmt.Errorf("city %s not found", city)
	}


	return t, nil
}

		
func WeatherForecast(city string) int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}																																															

func main() {
	data := NewData(1 * time.Minute)
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request){
		temp, err := data.GetTemperature("Moscow")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNoContent)
		}

		fmt.Fprintf(w, "{\"temperature\":%d}\n", temp)
	})
	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}