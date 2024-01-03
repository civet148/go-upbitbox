package upbitbox

import (
	"fmt"
	"strconv"
)

func Version() string {
	return "v1.0.0"
}

func GetTick(price float64) float64 {
	if price < 1000 {
		return 1
	} else if price < 5000 {
		return 5
	} else if price < 10000 {
		return 10
	} else if price < 50000 {
		return 50
	} else if price < 100000 {
		return 100
	} else if price < 500000 {
		return 500
	}

	return 1000
}

func ParseFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func Find[T any](datas []T, f func(x T) bool) T {
	var ret T
	for _, data := range datas {
		if f(data) {
			ret = data
			break
		}
	}
	return ret
}

func Map[T, M any](datas []T, f func(x T) M) []M {
	result := []M{}
	for _, data := range datas {
		result = append(result, f(data))
	}
	return result
}

func Reverse[T any](datas []T) []T {
	newArr := make([]T, 0, len(datas))
	for i := len(datas) - 1; i >= 0; i-- {
		newArr = append(newArr, datas[i])
	}
	return newArr
}

func GetSymbol(acc Account) string {
	return fmt.Sprintf("%s-%s", acc.UnitCurrency, acc.Currency)
}
