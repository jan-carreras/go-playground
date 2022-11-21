package daily_temperatures

func DailyTemperaturesBruteforce(temperatures []int) []int {
	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		if i > 0 && temperatures[i-1] == temperatures[i] {
			result[i] = result[i-1] - 1
		}

		for x := i + 1; x < len(temperatures); x++ {
			if temperatures[x] > temperatures[i] {
				result[i] = x - i
				break
			}
		}
	}

	return result
}

func DailyTemperaturesBruteforceRepeat(temperatures []int) []int {
	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for x := i + 1; x < len(temperatures); x++ {
			if temperatures[x] > temperatures[i] {
				result[i] = x - i
				break
			}
		}
	}

	return result
}

type Node struct {
	Temp     int
	Position int
}

func DailyTemperatures(temperatures []int) []int {
	output := make([]int, len(temperatures))

	var stack []Node
	
	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 {
			n := len(stack) - 1 // Top element
			node := stack[n]    // Read top element
			if node.Temp >= temperatures[i] {
				break
			}
			output[node.Position] = i - node.Position
			stack = stack[:n] // Remove top element
		}
		stack = append(stack, Node{temperatures[i], i})
	}

	return output
}
