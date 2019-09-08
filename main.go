package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	//read n
	if _, err := fmt.Scan(&n); err != nil {
		fmt.Println(err)
		return
	}
	textN, _ := reader.ReadString('\n')
	textN = strings.TrimSuffix(textN, "\n")
	if _, err := fmt.Scan(&m); err != nil {
		fmt.Println(err)
		return
	}

	textM, _ := reader.ReadString('\n')
	textM = strings.TrimSuffix(textM, "\n")
	sN := strings.Split(textN, " ")
	sM := strings.Split(textM, " ")
	mC := make(map[string]int)
	for i := 0; i < m; i++ {
		mC[sM[i]]++
	}
	for i := 0; i < n; i++ {
		mC[sN[i]]--
	}
	var keys []int
	for k, v := range mC {
		if v != 0 {
			i, _ := strconv.Atoi(k)
			keys = append(keys, i)
		}
	}
	sort.Ints(keys)
	for _, z := range keys {
		fmt.Printf("%d ", z)
	}
}
