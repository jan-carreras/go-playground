package twobucket

import (
	"errors"
	"fmt"
)

type bucket struct {
	value int
	cap   int
}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (goalBucket string, numSteps, otherBucketLevel int, err error) {
	// Sanity checks
	if goalAmount <= 0 || (goalAmount > sizeBucketOne && goalAmount > sizeBucketTwo) || sizeBucketOne < 0 || sizeBucketTwo < 0 {
		return "", 0, 0, errors.New("invalid argument")
	}
	if goalAmount%gcd(sizeBucketOne, sizeBucketTwo) != 0 {
		return "", 0, 0, errors.New("no solution")
	}

	buckets := [2]*bucket{
		{cap: sizeBucketOne},
		{cap: sizeBucketTwo},
	}

	start, other := 0, 1
	switch startBucket {
	case "one":
	case "two":
		buckets[start], buckets[other] = buckets[other], buckets[start]
	default:
		err = errors.New("unknown starting bucket")
		return
	}

	for buckets[start].value != goalAmount && buckets[other].value != goalAmount {
		numSteps++
		switch {
		case buckets[start].value == 0: // Empty bucket
			buckets[start].value = buckets[start].cap // Fill the bucket
		case buckets[other].cap == goalAmount: // If the other bucket is exactly like the capacity we're looking for, just fill it
			buckets[other].value = goalAmount
		case buckets[other].value == buckets[other].cap: // If it's full,
			buckets[other].value = 0 // empty it
		default:
			pour := min(buckets[other].cap-buckets[other].value, buckets[start].value)
			buckets[start].value -= pour
			buckets[other].value += pour
		}
		fmt.Println(*buckets[0], *buckets[1])
	}

	if buckets[start].value == goalAmount {
		goalBucket = "one"
		otherBucketLevel = buckets[1].value
	} else {
		goalBucket = "two"
		otherBucketLevel = buckets[0].value
	}
	return goalBucket, numSteps, otherBucketLevel, nil
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
