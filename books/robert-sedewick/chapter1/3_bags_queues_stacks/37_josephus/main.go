package main

// Josephus problem. In the Josephus problem from antiquity, N people are in dire
// straits and agree to the following strategy to reduce the population. They
// arrange themselves in a circle (at positions numbered from 0 to N–1) and
// proceed around the circle, eliminating every Mth person until only one person
// is left. Legend has it that Josephus figured out where to sit to avoid being
// eliminated. Write a Queue client Josephus that takes N and M from the command
// line and prints out the order in which people are eliminated (and thus would
// show Josephus where to sit in the circle).
//
// Input: java Josephus 7 2
//        1 3 5 0 4 2 6

// Josephus returns the order into which the people is going to die. The last one
// in the list gets saved.
func Josephus(n, m int) (deaths []int, err error) {
	q := newQueue(m)

	for counter := 1; q.length > 0; counter++ {
		if person, err := q.Pop(); err != nil {
			return nil, err
		} else if counter%n == 0 {
			deaths = append(deaths, person)
		} else {
			q.Push(person)
		}
	}

	return deaths, nil
}

func newQueue(m int) *Queue[int] {
	q := new(Queue[int])
	for i := 0; i < m; i++ {
		q.Push(i)
	}
	return q
}
