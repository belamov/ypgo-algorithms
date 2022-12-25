package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// https://contest.yandex.ru/contest/22450/problems/A/

//###############################################################################################
//
// айди успешной посылки 80042465: https://contest.yandex.ru/contest/22450/run-report/80042465/
//
//###############################################################################################

// Тимофей ищет место, чтобы построить себе дом.
//Улица, на которой он хочет жить, имеет длину n, то есть состоит из n одинаковых идущих подряд участков.
//Каждый участок либо пустой, либо на нём уже построен дом.
//
//Общительный Тимофей не хочет жить далеко от других людей на этой улице.
//Поэтому ему важно для каждого участка знать расстояние до ближайшего пустого участка.
//Если участок пустой, эта величина будет равна нулю — расстояние до самого себя.
//
//Помогите Тимофею посчитать искомые расстояния. Для этого у вас есть карта улицы.
//Дома в городе Тимофея нумеровались в том порядке, в котором
//строились, поэтому их номера на карте никак не упорядочены.
//Пустые участки обозначены нулями.
func getDistances(street []bool) []int {
	distances := make([]int, len(street))

	// заполним расстояния максимальными значениями
	for i := 0; i < len(street); i++ {
		distances[i] = 1000001 //по условию задачи длина улицы не более 10^6, а значит и расстояние не может быть длиннее этого значения
	}

	latest0position := -1

	for i := 0; i < len(street); i++ {
		//если находим пустой участок
		if street[i] {
			//записываем растояние 0 для этого участка
			distances[i] = 0

			// пройдем назад по улице от пустого участка, заполняя дистации до этого пустого участка
			// пока не дойдем до растояний меньше или не дойдем до начала улицы
			// расстояние до пустого участка будет разница их индексов
			for j := i - 1; j >= 0 && distances[j] > i-j; j-- {
				distances[j] = i - j
			}

			//запомним индекс предыдущего пустого участка, чтобы заполнять дистанции
			//для последующих участков
			latest0position = i

			continue
		}

		//если мы уже встречали пустой участок, то его индекс будет в latest0position и эта переменная будет
		//не меньше 0
		//для участка заполняем его расстояние до последнего пустого участка
		//это будет расстояние между индексами участков
		if latest0position >= 0 {
			distances[i] = i - latest0position
		}
	}
	return distances
}

func main() {
	scanner := makeScanner()
	readLine(scanner)
	street := readStreet(scanner)
	printArray(getDistances(street))
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func readStreet(scanner *bufio.Scanner) []bool {
	scanner.Scan()
	listString := strings.Split(scanner.Text(), " ")
	arr := make([]bool, len(listString))
	for i := 0; i < len(listString); i++ {
		//нам не важны номера домов, важны только занятость/незанятость участка
		//поэтому можем вставлять true для незанятого участка и false для занятого
		arr[i] = listString[i] == "0"
	}
	return arr
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 100 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func printArray(arr []int) {
	writer := bufio.NewWriter(os.Stdout)
	for i := 0; i < len(arr); i++ {
		writer.WriteString(strconv.Itoa(arr[i]))
		writer.WriteString(" ")
	}
	writer.Flush()
}
