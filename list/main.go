package main

import "fmt"


type Product struct{
	id int
	title string
	price float64
}


func new(id int,title string,price float64)Product{
	return Product{
		id: id,
		title: title,
		price: price,
	}
}

func main(){


	hobbies := [3]string{"Games","Sports","Workout"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	twoHobbies := hobbies[1:]
	fmt.Println(twoHobbies)
	slicedHobby1 := hobbies[0:2]
	slicedHobby2 := hobbies[:2]
	fmt.Println(slicedHobby1)

	favoriteHobbies := slicedHobby2[1:3]
	fmt.Println(favoriteHobbies)

	courses := []string{"CS","Biology"}
	courses[1] = "Math"
	courses = append(courses,"Biology")
	fmt.Println(courses)


	products := []Product{
		new(1,"A",100),
		new(2,"B",200),
	}
	fmt.Println(products)


	products = append(products,new(3,"C",300) )
	fmt.Println(products)
}
