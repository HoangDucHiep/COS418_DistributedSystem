package cos418_hw1_1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	//"github.com/go-delve/delve/pkg/dwarf/reader"
)

// Sum numbers from channel `nums` and output sum to `out`.
// You should only output to `out` once.
// Do NOT modify function signature.
func sumWorker(nums chan int, out chan int) {
	localSum := 0

	for num := range nums {
		localSum += num
	}

	out <- localSum
}

// Read integers from the file `fileName` and return sum of all values.
// This function must launch `num` go routines running
// `sumWorker` to find the sum of the values concurrently.
// You should use `checkError` to handle potential errors.
// Do NOT modify function signature.
func sum(num int, fileName string) int {
	// HINT: use `readInts` and `sumWorkers`
	// HINT: used buffered channels for splitting numbers between workers

	reader, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return 0
	}
	defer reader.Close()

	nums, err := readInts(reader)

	if err != nil {
		fmt.Println("Error reading integers:", err)
		return 0
	}

	// Create channels
	numsChan := make(chan int, len(nums))
	outChan := make(chan int, num)

	for i := 0; i < num; i++ {
		go sumWorker(numsChan, outChan)
	}

	for i := 0; i < len(nums); i++ {
		numsChan <- nums[i]
	}

	// Close numsChan after all numbers are sent
	close(numsChan)

	// Collect results from workers
	totalSum := 0
	for i := 0; i < num; i++ {
		totalSum += <-outChan
	}

	return totalSum
}

// Read a list of integers separated by whitespace from `r`.
// Return the integers successfully read with no error, or
// an empty slice of integers and the error that occurred.
// Do NOT modify this function.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var elems []int
	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return elems, err
		}
		elems = append(elems, val)
	}
	return elems, nil
}

func main() {
	res := sum(4, "q2_test1.txt")
	fmt.Println(res)
}
