package main

/*

## 1) валидатор

**Цель:** научиться возвращать ошибки и не усложнять.

**Задание:** реализуй функции:

- `ParsePositiveInt(s string) (int, error)`
    - парсит `s` в int
    - если `s` не число → верни ошибку от `strconv.Atoi` (как есть)
    - если число `<= 0` → верни `errors.New("must be positive")`
- `SumTwo(a, b string) (int, error)`
    - вызывает `ParsePositiveInt` для `a` и `b`
    - если ошибка → сразу возвращает её
    - иначе возвращает сумму

**Готово, если:**

- `SumTwo("10","20") == 30`
- `SumTwo("x","20")` возвращает ошибку от `Atoi`
- `SumTwo("-1","20")` возвращает `"must be positive"`

*/

import (
	"fmt"
	"strconv"
)

func ParsePositiveInt(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	if n <= 0 {
		return 0, fmt.Errorf("must be positive")
	}

	return n, nil
}

func SumTwo(a, b string) (int, error) {
	n1, err := ParsePositiveInt(a)
	if err != nil {
		return 0, err
	}

	n2, err := ParsePositiveInt(b)
	if err != nil {
		return 0, err
	}

	return n1 + n2, nil
}

func main() {
	sum1, err1 := SumTwo("10", "20")
	fmt.Println(sum1, err1)

	sum2, err2 := SumTwo("x", "20")
	fmt.Println(sum2, err2)

	sum3, err3 := SumTwo("-1", "20")
	fmt.Println(sum3, err3)
}

/* Вывод из консоли:
30 <nil>
0 strconv.Atoi: parsing "x": invalid syntax
0 must be positive
*/
