package variables

import "fmt"

func Run() {
	var count int
	count = 5

	fmt.Println(count)

	limit := 2
	fmt.Println(limit)

	numbers := []int{3, 5, 7}
	fmt.Println(numbers)

	nums := make([]int, 2)
	nums[0] = 1

	students := map[string]string{"John": "1", "Doe": "2"}
	fmt.Println(students)
}
