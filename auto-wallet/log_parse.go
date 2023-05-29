// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func main() {
// 	file, err := os.Open("auto_wallet_output.log")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	defer file.Close()
// 	tx_prepare_time := make(map[int]time.Time)
// 	tx_generate_time := make(map[int]time.Time)

// 	scanner := bufio.NewScanner(file)
// 	layout := "2006-01-02T15:04:05.999Z"
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		fields := strings.Split(line, " ")

// 		child_idx, err := strconv.ParseInt(fields[2], 10, 64)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		time, err := time.Parse(layout, fields[0][])

// 		if fields[4] == "Prepare" {

// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		fmt.Println(err)
// 	}
// }
