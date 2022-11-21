package main

/*
Urlify

* URLify: Write a method to replace all spaces in a string with '%20: You may
assume that the string has sufficient space at the end to hold the additional
characters, and that you are given the "true" length of the string. (Note: If
implementing in Java, please use a character array so that you can perform this
operation in place.) EXAMPLE Input: "Mr John Smith " J 13 Output: "Mr%20J
ohn%20Smith"
*/
func Urlify(s []byte, size int) []byte {
	endIdx := len(s) - 1
	for i := size - 1; i >= 0; i-- {
		if s[i] == ' ' {
			s[endIdx] = '0'
			s[endIdx-1] = '2'
			s[endIdx-2] = '%'
			endIdx -= 3
		} else {
			s[endIdx] = s[i]
			endIdx--
		}
	}

	return s
}
