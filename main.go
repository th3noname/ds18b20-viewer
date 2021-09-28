package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gosuri/uilive"
	"github.com/yryz/ds18b20"
)

func main() {
	out := make(chan string)
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	go readSensors(sensors, out)

	writer := uilive.New()
	writer.Start()
	defer writer.Stop()

	for {
		select {
		case s := <-out:
			fmt.Fprint(writer, s)
		case <-sigs:
			os.Exit(0)
		}
	}
}

func readSensors(sensors []string, out chan<- string) {
	for {
		var s string
		for _, sensor := range sensors {
			t, err := ds18b20.Temperature(sensor)
			if err != nil {
				panic(fmt.Errorf("Error reading sensor %s: %w", sensor, err))
			}

			s = fmt.Sprintf("%ssensor: %s temperature: %.2f°C\n", s, sensor, t)
		}

		s = fmt.Sprintf("%s\n%s\n", s, time.Now().Format("15:04:05"))

		out <- s
	}
}
