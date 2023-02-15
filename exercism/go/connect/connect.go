package connect

func ResultOf(lines []string) (string, error) {
	// Question: How do I know the initial positions/Sides of each player?
	//
	// Insight: If there is a path from any stone from the starting side, to the other side, that player wins.
	// It's about pathfinding (BFS). If we could encode the information of the board on a graph, the algorithms could be
	// to search any node from the "starting side" has a link to the "finish side". That would be a winner. Otherwise
	// that player has not won.
	//
	// Caveat: The work of writing such graph might be as complicated as searching for the solution directly
	//
	// Idea: We might have a function that returns us the "next positions".

	return "", nil
}

func resultOfPlayer(lines []string, player string) error {
	// Find one starting edge: check top edge, check left edge

	return nil
}

/*func topEdgePositions(lines []string, player string) []*/
