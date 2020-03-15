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
* init()
*
*
*******************************************************/

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

} // init

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
* roll (max, num)
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

	// yeah, i know
	debug := 0

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

		if debug > 0 {
			fmt.Println("dieroll.Roll() number rolled is ", sum)
		} // if

	} else if num > 1 {
		count := 0

		for count < num {
			new_num := randomInt(max)

			if debug > 0 {
				fmt.Println("dieroll.Roll() number rolled is ", new_num)
			} // if

			sum += new_num
			count += 1

			if debug > 0 {
				fmt.Println("dieroll.Roll() sum is ", sum)
			} // debug

		} // for

		if debug > 0 {
			fmt.Println("dieroll.Roll() final sum is ", sum)
		} // debug

	} else {
		fmt.Println("dieroll.Roll: Incorrect number of dice passed to function: ", num)
		os.Exit(3)
	} // if

	return sum

} // roll()

/*******************************************************
*
* RollStat()
*
* Rolls 4d6 for a stat and takes the top 3
*
*******************************************************/

func RollStat() int {
	rolls := make([]int, 4, 4)

	// yeah, i know
	debug := 0

	// d6 uses 7 in randomInt
	max := 7
	i := 0

	for i < 4 {
		roll := randomInt(max)
		rolls[i] = roll
		i++
	}

	if debug > 0 {
		fmt.Println("Content of rolls prior to sorting:  ", rolls)
	} // if

	sort.Ints(rolls)

	if debug > 0 {
		fmt.Println("Content of rolls after sorting:  ", rolls)
	} // if

	sliced_rolls := rolls[1:]

	if debug > 0 {
		fmt.Println("Content of sliced_rolls after pruning:  ", sliced_rolls)
	} // if

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
