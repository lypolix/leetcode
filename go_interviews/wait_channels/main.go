package main

func waitChannels(a, b chan struct{}) {
	for a != nil || b != nil{
		select {
		case _, ok := <- a:
			if !ok {
				a = nil 
			}
		
		case _, ok := <- b:
			if !ok {
				b= nil 
			}
		}
	
	}
}
