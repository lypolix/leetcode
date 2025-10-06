package main

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type Sayer interface {
	Say()
}

type Base struct {
	name string
}

type Child struct {
	lastName string
	Base
}

func (b Base) Say() { 
	fmt.Printf("Hello, %s!\n", b.name)
}

func (ch Child) Say() {
	fmt.Printf("Hello, %s %s!\n", ch.lastName, ch.name) 
}

func NewObject(str string) Sayer {
	switch str {
	case "Base":
		return Base{name: "Parent"}
	case "Child":
		return Child{lastName: "Inherited", Base: Base{name: "Child"}}
	default:
		return nil
	}

}

func ObjectGenerator(ctx context.Context, kind string, out chan<- Sayer, d time.Duration) {
	ticker := time.NewTicker(d) 
	defer ticker.Stop()
	for {
		select {
		case <- ctx.Done():
			return
		case <- ticker.C :
			out <- NewObject(kind)
		}
		
	}
}



func main() {
	b1 := Base{
		name: "Parent",
	}

	c1 := Child{
		lastName: "Inherited",
		Base: Base{
			name: "Child",
		},
	}

	b1.Say()

	c1.Base.Say() 

	arr := []Sayer{b1, c1} 

	for _, el := range arr {
		el.Say()
	}
	
	ch := make(chan Sayer, 8) 
	ctx, cancel := context.WithTimeout(context.Background(), 11 * time.Second)
	defer cancel()

	go ObjectGenerator(ctx, "Base", ch, 1 * time.Second)
	go ObjectGenerator(ctx, "Child", ch, 2 * time.Second)

	for {
		select {
		case obj := <-ch: 
			obj.Say()
		case <- ctx.Done():
			return
		}
	}

}


func TestNewObject(t *testing.T) { 
	tests := []struct { 
		name string
		str string
		res Sayer
	}{
		{
		name: "Test 1",
		str : "Base",
		res : Base{name: "Parent"},
		}, 
		{
		name: "Test 2",
		str : "Child",
		res : Child{lastName: "Inherited", Base: Base{name: "Child"}},
		},
		{
		name: "Test 3",
		str : "qwwywu",
		res : nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			got := NewObject(tt.str)

			switch w := tt.res.(type) { 
			case nil: 
				if got != nil {
					t.Fatalf("want nil, got %T=%v", got, got) 
				}
			case Base:
				b, ok := got.(Base)
				
				if !ok{
					t.Fatalf("want Base, got %T", got)
				}

				if b.name != w.name { 
					t.Fatalf("Base.name: want %q, got %q", w.name, b.name) 
				}
			case Child:
				ch, ok := got.(Child) 

				if !ok {
					t.Fatalf("want Child, got %T", got)
				}

				if ch.name != w.name {
					t.Fatalf("Child.name: want %q, got %q", w.name, ch.name)
				}

				if ch.lastName != w.lastName {
					t.Fatalf("Child.lastName: want %q, got %q", w.lastName, ch.lastName)
				}
			default:
				t.Fatalf("unexpected want type %T", tt.res)
			}
		})

	} 
}

