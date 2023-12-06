package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const arab = 1
const rom = 2

var validArabNubmers = map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10}
var validRomNumbers = map[string]int{"M": 1000, "CM": 900, "D": 500, "CD": 400, "C": 100, "XC": 90, "L": 50, "XL": 40, "X": 10, "IX": 9, "V": 5, "IV": 4, "I": 1}
var validOperators = []string{"+", "-", "*", "/"}

func main() {
	reader := bufio.NewReader(os.Stdin)

	task_text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Input error")
	}

	task_text = strings.TrimSpace(task_text)
	task_text = strings.ToUpper(task_text)
	task := strings.Split(strings.Replace(task_text, " ", "", -1), "") // строка с заданием преобразована в массив строк

	if len(task) < 3 {
		fmt.Println("Input error") // если в массиве меньше трёх элементов - ошибка
		return
	}

	args, fault := get_arguments(task)
	if fault != "" {
		fmt.Println(fault)
		return
	}

	result, fault := calculate(args)

	if fault == "" {
		fmt.Println(result)
	} else {
		fmt.Println(fault)
		return
	}
}

// Функция осуществляет заданные операции над числами
func calculate(args []int) (result string, fault string) {
	var res int

	if args[0] == arab {
		res, fault = calculate_arab(args[1], args[2], args[3])
		return strconv.Itoa(res), fault
	} else if args[0] == rom {
		res, fault = calculate_rom(args[1], args[2], args[3])
		return int_to_roman(res), fault
	}

	return "", ""
}

func calculate_rom(operator int, first_num int, second_num int) (result int, fault string) {
	switch operator {
	case 0:
		result = first_num + second_num
		if result <= 0 {
			return 0, "The result of the operation is less than or equal to 0"
		} else {
			return result, ""
		}
	case 1:
		result = first_num - second_num
		if result <= 0 {
			return 0, "The result of the operation is less than or equal to 0"
		} else {
			return result, ""
		}
	case 2:
		result = first_num * second_num
		if result <= 0 {
			return 0, "The result of the operation is less than or equal to 0"
		} else {
			return result, ""
		}
	case 3:
		if second_num == 0 {
			return 0, "Division by zero"
		} else {
			result = first_num / second_num
			if result <= 0 {
				return 0, "The result of the operation is less than or equal to 0"
			} else {
				return result, ""
			}
		}
	default:
		return -1, ""
	}
}

func calculate_arab(operator int, first_num int, second_num int) (result int, fault string) {
	switch operator {
	case 0:
		return first_num + second_num, ""
	case 1:
		return first_num - second_num, ""
	case 2:
		return first_num * second_num, ""
	case 3:
		if second_num == 0 {
			return 0, "Division by zero"
		} else {
			return first_num / second_num, ""
		}
	default:
		return -1, ""
	}
}

// Функция преобразует число в строку с римским числом
func int_to_roman(arg int) string {
	var (
		romanFigure = []int{1000, 100, 10, 1}
		romanDigitA = []string{1: "I", 10: "X", 100: "C", 1000: "M"}
		romanDigitB = []string{1: "V", 10: "L", 100: "D", 1000: "MMMMM"}
	)

	if arg < 1 || arg > 4000 {
		return "Out of range"
	}

	var roman strings.Builder
	x := ""

	for _, f := range romanFigure {
		digit, i, v := int(arg/f), romanDigitA[f], romanDigitB[f]
		switch digit {
		case 1:
			roman.WriteString(i)
		case 2:
			roman.WriteString(i + i)
		case 3:
			roman.WriteString(i + i + i)
		case 4:
			roman.WriteString(i + v)
		case 5:
			roman.WriteString(v)
		case 6:
			roman.WriteString(v + i)
		case 7:
			roman.WriteString(v + i + i)
		case 8:
			roman.WriteString(v + i + i + i)
		case 9:
			roman.WriteString(i + x)
		}

		arg -= digit * f
		x = i
	}

	return roman.String()
}

/*
   Функция преобразовывает массив строк в массив чисел,
   где первый элемент - тип чисел (арабские или римские),
   второй элемент - оператор,
   третий элемент - первое число
   четвёртый элемент - второе число
*/
func get_arguments(task []string) (arguments []int, fault string) {
	var first_arg []int
	var second_arg []int
	var operator_index int = -1

	arguments = append(arguments, 0)  // тип чисел по умолчанию
	arguments = append(arguments, -1) // оператор по умолчанию

	for index, value := range task {
		val, ok := contains(validOperators, value) // поиск оператора в строке с заданием
		if ok {
			if arguments[1] == -1 {
				arguments[1] = val // найден оператор
				operator_index = index
			} else {
				operator_index = -1
				return arguments, "Operator repeats!"
			}
		}
	}

	if operator_index < 1 {
		return arguments, "An operator was found instead of a digit!"
	}

	first_arg, typeOfNum := get_numbers(task[0:operator_index]) // поиск первого числа, в случае успеха возвращается срез чисел

	if typeOfNum == 0 {
		return arguments, "The first argument is incorrect!"
	}

	arguments[0] = typeOfNum

	second_arg, typeOfNum = get_numbers(task[operator_index:]) // поиск второго числа, в случае успеха возвращается срез чисел

	if typeOfNum != arguments[0] {
		return arguments, "The second argument is incorrect!"
	}
	// в срез arguments добавляется первое и второе число
	arguments = append(arguments, array_to_int(first_arg, arguments[0]))
	arguments = append(arguments, array_to_int(second_arg, arguments[0]))

	return arguments, ""
}

//Функция преобразует срез чисе в число в соотвествии с типом чисел в срезе
func array_to_int(arr []int, typeOfNumber int) int {
	var p int
	result := 0

	if typeOfNumber == arab {
		p = len(arr) - 1
		for index, val := range arr {
			pow := math.Pow(10, float64((p - index)))
			result += val * int(pow)
		}
	} else if typeOfNumber == rom {
		for _, val := range arr {
			if val > result {
				result = val - result
			} else {
				result += val
			}
		}
	}

	return result
}

// Функция преобразует срез строк в срез чисел, а также, определяет тип чисел
func get_numbers(arg []string) (numbers []int, typeOfNumber int) {
	var val int
	var ok bool
	typeOfNumber = 0

	for _, value := range arg {
		val, ok = validArabNubmers[value]
		if ok {
			if typeOfNumber == 0 {
				typeOfNumber = arab
			}
			if typeOfNumber == arab {
				numbers = append(numbers, val)
				continue
			}
		} else {
			val, ok = validRomNumbers[value]
			if ok {
				if typeOfNumber == 0 {
					typeOfNumber = rom
				}
				if typeOfNumber == rom {
					numbers = append(numbers, val)
					continue
				}
			}
		}
	}

	return numbers, typeOfNumber
}

// Функция ищет строку в срезе строк и возвращает индекс, в случае успеха.
func contains(s []string, e string) (int, bool) {
	for index, a := range s {
		if a == e {
			return index, true
		}
	}

	return 0, false
}
