package main

import (
	"fmt"
	"encoding/json"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	var k user
	v , _ := json.Marshal(u1)
	_ = json.Unmarshal(v, &k)
	fmt.Println(k)
	t , _ := json.Marshal(u2)
	_ = json.Unmarshal(t, &k)
	fmt.Println(k)
	r , _ := json.Marshal(u3)
	_ = json.Unmarshal(r, &k)
	fmt.Println(k)

}


