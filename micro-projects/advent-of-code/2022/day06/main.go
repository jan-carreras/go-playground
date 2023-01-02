package day6

const startOfPacketSize = 4
const startOfMessageSize = 14

func startOfTransmission(input string) int {
	return searchMarker(input, startOfPacketSize)
}

func startMessageMarker(input string) int {
	return searchMarker(input, startOfMessageSize)
}

func searchMarker(input string, markerSize int) int {
	window := make([]string, markerSize)

	for i := 0; i < len(input); i++ {
		window[i%markerSize] = string(input[i])
		if i >= 3 && !hasDuplicates(window) {
			return i + 1
		}
	}

	return -1

}

func hasDuplicates(input []string) bool {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return true
			}
		}
	}
	return false
}
