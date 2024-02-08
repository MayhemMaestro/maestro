package cmd

import (
	cpu "maestro/cmd/chaostests/cpuTests"
	mem "maestro/cmd/chaostests/memTests"
)

// Takes user input to run a test
func CheckList(selectedComponent string, selectedChaosTest string) string {

	switch selectedComponent {
	case "cpu":
		if selectedChaosTest == "saturation" {
			return cpu.CPUSaturation()
		}
		return "No tests found for:" + selectedChaosTest

	case "memory":
		return mem.MemSaturation()
	default:
		return "No tests found for:" + selectedChaosTest
	}

}
