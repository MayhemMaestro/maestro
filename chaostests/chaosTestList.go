package chaostests

import (
	cpu "github.com/MayhemMaestro/maestro/chaostests/cpuTests"
	mem "github.com/MayhemMaestro/maestro/chaostests/memTests"
)

// Takes user input to run a test
func CheckList(selectedComponent string, selectedChaosTest string, threshold string, duration string) string {

	switch selectedComponent {
	case "cpu":
		if selectedChaosTest == "saturation" {
			return cpu.CPUSaturation(threshold, duration)

		}
	// 	return "No tests found for:" + selectedChaosTest

	case "mem":
		return mem.MemSaturation()
	default:
		return "No tests found for:" + selectedChaosTest
	}
	return ""

}
