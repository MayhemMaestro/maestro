package chaostests

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/Codehardt/go-cpulimit"
	"github.com/shirou/gopsutil/cpu"
)

func CPUSaturation(threshold string, duration string) string {
	done := make(chan int)
	timeDuration, err := time.ParseDuration(duration)
	if err != nil {
		fmt.Printf("Error parsing duration of test: %s", err)
	}

	thresholdFloat, err := strconv.ParseFloat(threshold, 64)
	if err != nil {
		fmt.Printf("Error parsing max CPU usage for test: %s", err)
	}

	// Initialize the CPU limiter
	limiter := &cpulimit.Limiter{
		MaxCPUUsage:        thresholdFloat,         // Set the maximum CPU usage
		MeasureInterval:    time.Millisecond * 333, // Measure CPU usage every 333 ms
		Measurements:       3,
		CurrentProcessOnly: true, // Use the average of the last 3 measurements
	}
	limiter.Start()
	defer limiter.Stop()

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
					limiter.Wait()
				}
			}
		}()
	}

	// Start a goroutine to print CPU utilization every second
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				percent, err := cpu.Percent(time.Second, false)
				if err != nil {
					fmt.Println("Error getting CPU utilization:", err)
					return
				}
				fmt.Printf("CPU Utilization: %.2f%%\n", percent[0])
				time.Sleep(time.Second)
			}
		}
	}()

	time.Sleep(time.Duration(timeDuration))

	close(done)

	return "Test complete"
}
