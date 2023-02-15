package sublist

func Sublist(listA, listB []int) Relation {
	if len(listA) == len(listB) {
		return checkEquality(listA, listB)
	} else if len(listA) == 0 {
		return RelationSublist
	} else if len(listB) == 0 {
		return RelationSuperlist
	} else if len(listA) > len(listB) {
		return sublist(listA, listB, RelationSuperlist)
	} else {
		return sublist(listB, listA, RelationSublist)
	}
}

func sublist(long, short []int, onMatch Relation) Relation {
	for i := 0; i+len(short)-1 < len(long); i++ {
		if areEqual(long[i:i+len(short)], short) {
			return onMatch
		}
	}

	return RelationUnequal
}

func checkEquality(listA, listB []int) Relation {
	if areEqual(listA, listB) {
		return RelationEqual
	}
	return RelationUnequal
}

func areEqual(listA, listB []int) bool {
	if len(listA) != len(listB) {
		return false
	}

	for i := 0; i < len(listA); i++ {
		if listA[i] != listB[i] {
			return false
		}
	}

	return true
}
