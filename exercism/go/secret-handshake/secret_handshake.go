package secret

/**
1:    "wink",
	10:   "double blink",
	100:  "close your eyes",
	1000: "jump",
*/

// reverse = 10000

var greetings = []string{"wink", "double blink", "close your eyes", "jump"}

const reverse = 10000

func Handshake(in uint) []string {
	var resp []string
	for i := 0; i < len(greetings); i = i + 1 {
		mask := uint(1 << i)
		if (in & mask) > 0 {
			resp = append(resp, greetings[i])
		}
	}

	if (in & reverse) > 0 {
		for i := 0; i < len(resp)/2; i++ {
			j := len(resp) - i - 1
			resp[i], resp[j] = resp[j], resp[i]
		}
	}

	return resp
}
