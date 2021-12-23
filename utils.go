package main

import (
	"fmt"
        "strings"
	"github.com/ssimunic/gosensors"
)

func sensors()  string {
	sensors, err := gosensors.NewFromSystem()

	if err != nil {
		panic(err)
	}
	fmt.Println(sensors)
        var sens_data []string
	for chip := range sensors.Chips {
		for key, value := range sensors.Chips[chip] {
                         sens_data = append(sens_data,fmt.Sprintf("%s -> %s",key,value))
		}
	}
        if len(sens_data) > 0 {
        return strings.Join(sens_data, "\n")
        } else {
        return fmt.Sprintf("No sensors detected on the system\n")
        }
}
