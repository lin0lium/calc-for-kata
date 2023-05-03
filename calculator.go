// Данная программа написана самостоятельно, человеком, имеющим представление о Go, но не работающим с ним ранее.
// Прошу простить за 4-х уровневую вложенность, обработать косяки регулярными выражениями мне не удалось.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	operators = map[string]func() int{
		"+": func() int { return *a + *b },
		"-": func() int { return *a - *b },
		"/": func() int { return *a / *b },
		"*": func() int { return *a * *b },
	}
	roman = map[string]int{
		"C": 100, "L": 50, "XL": 40, "X": 10,
		"IX": 9, "VIII": 8, "VII": 7, "VI": 6, "V": 5, "IV": 4, "III": 3, "II": 2, "I": 1,
	}
	convIntToRoman = [14]int{100, 90, 50, 40, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	data           []string
	a, b           *int
)

// Обработка валидности значений и основные условия калькулятора
func base(s string) {
	var (
		operator     string
		stringsFound int
	)
	numbers := make([]int, 0)
	romansToInt := make([]int, 0)
	romans := make([]string, 0)
	for idx := range operators {
		for _, val := range s {
			if idx == string(val) {
				operator += idx
				data = strings.Split(s, operator)
			}
		}
	}
	if len(operator) != 1 {
		panic("Строка не является математической операцией. Требования: два операнда и один оператор (+, -, /, *).")
	}
	for _, elem := range data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch stringsFound {
	case 1:
		panic("Одновременно используются разные системы счисления.")
	case 0:
		errCheck := numbers[0] > 0 && numbers[0] < 11 && numbers[1] > 0 && numbers[1] < 11
		if val, ok := operators[operator]; ok && errCheck == (numbers[0] > 0 && numbers[0] < 11 && numbers[1] > 0 && numbers[1] < 11) {
			a, b = &numbers[0], &numbers[1]
			fmt.Println(val())
		} else {
			panic("Калькулятор умеет работать только с арабскими целыми числами или римскими цифрами от 1 до 10 включительно")
		}
	case 2:
		for _, elem := range romans {
			if val, ok := roman[elem]; ok && 0 < val && val < 11 {
				romansToInt = append(romansToInt, val)
			} else {
				panic("Калькулятор умеет работать только с арабскими целыми числами или римскими цифрами от 1 до 10 включительно")
			}
		}
		if val, ok := operators[operator]; ok {
			a, b = &romansToInt[0], &romansToInt[1]
			intToRoman(val())
		}
	}
}

// Обработка условий и задача вывода в римской сс
func intToRoman(romanResult int) {
	var romanNum string
	if romanResult < 1 {
		panic("В римской системе счисления нет отрицательных чисел/нуля.")
	} else {
		for _, elem := range convIntToRoman {
			for i := elem; i <= romanResult; {
				for index, value := range roman {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanNum)
}

// Главный движок
func main() {
	fmt.Println("CALCULATOR")
	reader := bufio.NewReader(os.Stdin)
	for {
		console, _ := reader.ReadString('\n')
		s := strings.ReplaceAll(console, " ", "")
		base(strings.ToUpper(strings.TrimSpace(s)))
	}
}
