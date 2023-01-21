package main

import (
	"fmt"
)

func findGroups(relationships []string) [][]string {
	// Create a new HashMap to store the relationships
	network := make(map[string][]string)

	// Iterate through the list of relationships and update the network HashMap
	for _, relationship := range relationships {
		var personA, personB string
		fmt.Sscanf(relationship, "%s is friends with %s", &personA, &personB)
		network[personA] = append(network[personA], personB)
		network[personB] = append(network[personB], personA)
	}

	// Create a new HashMap to keep track of visited people
	visited := make(map[string]bool)

	// Create a variable to store the friend groups
	var groups [][]string

	// Iterate through the network HashMap
	for person := range network {
		if !visited[person] {
			// Create a new variable to store the current friend group
			var group []string
			// Use DFS to find all the friends in the current group
			dfs(person, network, &visited, &group)
			groups = append(groups, group)
		}
	}

	return groups
}

func dfs(person string, network map[string][]string, visited *map[string]bool, group *[]string) {
	// Mark the current person as visited
	(*visited)[person] = true
	// Add the current person to the current friend group
	*group = append(*group, person)

	// Iterate through the current person's friends
	for _, friend := range network[person] {
		if !(*visited)[friend] {
			// Use DFS to find all the friends of the current person's friend
			dfs(friend, network, visited, group)
		}
	}
}

func main() {
	relationships := []string{
		"Alice is friends with Bob",
		"Bob is friends with Charlie",
		"Charlie is friends with Alice",
		"David is friends with Eve",
		"Eve is friends with Frank",
	}

	groups := findGroups(relationships)

	for i, group := range groups {
		fmt.Printf("Group %d: %v\n", i+1, group)
	}
}
