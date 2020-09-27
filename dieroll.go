package dieroll

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

/*******************************************************
*
* contains(valid, max)
*
* Search the valid array for the max.
*
* The valid array contains the valid die types as ints:  d4, d6, d8, d10, d12, d10, d100
*
*
*******************************************************/

func contains(valid [7]int, max int) bool {
	for _, element := range valid {
		if element == max {
			return true
		} // if
	} // for

	return false

} // contains()

/*******************************************************
*
* Init()
*
*
*******************************************************/

func Init() {
	rand.Seed(time.Now().UTC().UnixNano())

} // Init

/*******************************************************
*
* randomInt(max)
*
* Returns a random integer between 1 and max.
*
*******************************************************/

func randomInt(max int) int {
	min := 1
	return min + rand.Intn(max-min)

} // randomInt()

/*******************************************************
* Roll (max, num)
*
* Rolls a number of times equal to num.  Each result is
* between 1 and max.
*
* Returns either the number if only one die was rolled
* or the sum of all dice
*
* Max must be equal to 4, 6, 8, 10, 12, 20 or 100.  The
* function will handle the details for rand.Intn itself.
*
* TBD:
*
* Figure out better error handling
*
*******************************************************/

func Roll(max int, num int) int {
	sum := 0

	// max must be 5 for d4, 7 for d6, 9 for d8, 11 for d10, 13 for d12 and 21 for d20
	// max += 1
	// Limited to the following:
	var valid = [7]int{4, 6, 8, 10, 12, 20, 100}

	valid_check := contains(valid, max)

	if valid_check == false {
		fmt.Println("dieroll().Roll(): max value = ", max, "is not correct value.")
		os.Exit(3)
	} // if

	// Increment max by one for randomInt calculation.
	max += 1

	if num == 1 {
		sum = randomInt(max)

	} else if num > 1 {
		count := 0

		for count < num {
			new_num := randomInt(max)
			sum += new_num
			count += 1
		}

	} else {
		fmt.Println("dieroll.Roll: Incorrect number of dice passed to function: ", num)
		os.Exit(3)
	}

	return sum

} // Roll()

/*******************************************************
* Rolldice (max, num)
*
* Rolls a number of times equal to num.  Each result is
* between 1 and max.
*
* Returns a list of all results.
*
* Max must be equal to 4, 6, 8, 10, 12, 20 or 100.  The
* function will handle the details for rand.Intn itself.
*
*******************************************************/

func Rolldice(max int, num int) []int {
	results := make([]int, num)

	// max must be 5 for d4, 7 for d6, 9 for d8, 11 for d10, 13 for d12 and 21 for d20
	// max += 1
	// Limited to the following:
	var valid = [7]int{4, 6, 8, 10, 12, 20, 100}

	valid_check := contains(valid, max)

	if valid_check == false {
		fmt.Println("dieroll().Roll(): max value = ", max, "is not correct value.")
		os.Exit(3)
	} // if

	// Increment max by one for randomInt calculation.
	max += 1

	if num == 1 {
		results[0] = randomInt(max)

	} else if num > 1 {
		count := 0

		for count < num {
			roll := randomInt(max)
			results[count] = roll
			count += 1
		}

	} else {
		fmt.Println("dieroll.Roll: Incorrect number of dice passed to function: ", num)
		os.Exit(3)
	}

	return results

} // Rolldice()

/*******************************************************
*
* RollStat()
*
* Rolls 4d6 for a stat and takes the top 3
*
*******************************************************/

func RollStat() int {
	rolls := make([]int, 4, 4)

	// d6 uses 7 in randomInt
	max := 7
	i := 0

	for i < 4 {
		roll := randomInt(max)
		rolls[i] = roll
		i++
	}

	sort.Ints(rolls)
	sliced_rolls := rolls[1:]
	sum := 0

	for _, roll := range sliced_rolls {
		sum += roll
	} // for

	return sum

} // RollStat()

/*******************************************************
*
* RandomKey()
*
* Returns a random key in a map.
*
*******************************************************/

func RandomKey(m map[string]string) string {
	// TBD
	// Don't use until you understand this:
	//
	// keys := reflect.ValueOf(mapI).MapKeys()
	// return keys[rand.Intn(len(keys))].Interface()

	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	} // for

	return keys[rand.Intn(len(keys))]

} // RandomKey()
