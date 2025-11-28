package main

import "fmt"



func main(){

	numbers := []int{1,2,4,5,67,6,7}

	anonymous := func(value int)int{
		return value * 8
	}


	anonymousFn := transformNumbers(&numbers,anonymous)

	doubled := transformNumbers(&numbers,double)

	tripled := transformNumbers(&numbers,triple)

	fmt.Println(anonymousFn)
	fmt.Println(doubled)
	fmt.Println(tripled)


}


func transformNumbers(nums *[]int, transform func(int) int ) []int{

	transformed := []int{}


	for _, value := range *nums{
		transformed = append(transformed, transform(value) )
	}

	return transformed

}


func double(value int) int {
	return value * 2
}


func triple(value int) int {
	return value * 3
}