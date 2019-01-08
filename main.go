package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	list := os.Args
	if len(list) == 1 {
		fmt.Println("missing args")
		os.Exit(1)
	}
	result := ""
	for index, value := range list {
		//filter script self
		if index == 0 {
			continue
		}
		//convert to int
		intValue, err := strconv.Atoi(value)
		//convert str to timestamp
		if err != nil {
			timeStamp, _ := time.Parse("2006-01-02", value)
			result = strconv.FormatInt(timeStamp.Unix(), 10)
		} else {
			//convert timestamp to str
			tm := time.Unix(int64(intValue), 0)
			result = tm.Format("2006-01-02 15:04:05")
		}
		output := value + " => " + result
		fmt.Println(output)
	}
}
