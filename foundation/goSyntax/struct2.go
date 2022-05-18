package main 

import "fmt"

type vehicle struct{
	doors int
	color string
}
type truck struct{
	vehicle
	fourWheel bool
}
type sedan struct{
	vehicle
	luxury bool
}

func main(){
	Sedan := sedan {
		vehicle : vehicle{
			doors : 4,
			color : "blue",
		},
		luxury : true,
	}
	Truck := truck {
		vehicle : vehicle{
			doors : 2,
			color : "red",
		},
		fourWheel :  true,
	}
	bakkie := struct {
		doors int
		color string
		fourWheel bool
		cupholder []string
	}{
		doors : 4,
		color : "white",
		fourWheel : true,
		cupholder : []string{"one"},
	}

	p := make([]int , 5 , 5)
	p = append(p , 5)
	fmt.Printf("Car = %T\n Doors = %v\n",Sedan , Sedan.vehicle.doors)
	fmt.Printf("Car = %T\n Doors = %v\n",Truck , Truck.vehicle.doors)
	fmt.Printf("Car = %T\n Doors = %v\n",bakkie , bakkie.doors)
	fmt.Println(cap(p) , len(p))
}