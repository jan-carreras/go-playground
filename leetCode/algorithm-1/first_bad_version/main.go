package first_bad_version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
var expected = 0

func isBadVersion(version int) bool {
	return version >= expected
}

func FirstBadVersion(n int) int {
	low, high := 1, n
	lastBadVersion := n
	middle := 0
	for low <= high {
		middle = ((high - low) / 2) + low
		switch {
		case isBadVersion(middle):
			lastBadVersion = middle
			high = middle - 1
		default:
			low = middle + 1
		}
	}

	return lastBadVersion
}
