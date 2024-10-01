package main


import (
	"fmt"
	"strings"
	"strconv"
	"sort"
)

var numsR = make([]float64, 0, 100)

func main() {
	
	usCh := scanUserInp()


	getNums(&numsR)
	
	
	switch usCh {
	case "AVG":
		fmt.Printf("Среднее чисел равно %.2f", avg(numsR))
	case "MED":
		fmt.Printf("Медиана чисел равна %.2f", med(numsR))
	case "SUM":
		fmt.Printf("Сумма чисел равна: %.2f",sum(numsR))
	}
	
}


func scanUserInp() string {
	var usCh string
	for {
	fmt.Println("Выберите необходимую функцию (AVG, MED, SUM)")
	fmt.Scanln(&usCh)

	switch usCh {
	case "AVG", "MED", "SUM":
		return usCh
	default:
		fmt.Println("Введен некорректный параметр")
		continue
	}
	}
	return usCh
}

func getNums(numsR *[]float64) {
	var num string
	
	//var numsF = make([]float64, 0, 100)
	
	fmt.Println("Введите числа через запятую. (Пример: 10,20,30)")
	fmt.Scan(&num)
	

	nums := strings.Split(num, ",")
	for _, v := range nums {
		//v = strings.Replace(v, " ", "", -1)	
                n, err := strconv.ParseFloat(v, 64)
                if err != nil {
			fmt.Printf("Не удалось обработать строку. (Ошибка: %s)", err)
                } 
		*numsR = append(*numsR, n)
		
	}
	
	
}

func avg(nums []float64) float64 {
	sumA := sum(nums)
	avg := sumA/float64(len(nums))
	return avg
}

func med(nums []float64) float64 {
	sort.Float64s(nums)	
	ind := len(nums)/2
	if len(nums) % 2 == 0 {
		return (nums[ind]+nums[ind-1])/2
	} else {
		return nums[ind] 
}
}

func sum(nums []float64) float64 {
	var sum float64
	for _, n := range nums {
		sum += n
	}
	return sum
}
