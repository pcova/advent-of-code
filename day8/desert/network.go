package desert

import (
	"bufio"
	"strings"
)

/**
 * parseNode parses a string into a node and its children.
 *
 * The input is expected to be in the following format: "AAA = (BBB, CCC)"
 * The output will be "AAA" and ["BBB", "CCC"].
 */
func parseNode(input string) (string, []string) {
	parts := strings.Split(input, " = ")

	node := strings.TrimSpace(parts[0])

	values := strings.Trim(parts[1], "() ")
	instructions := strings.Split(values, ", ")

	return node, instructions
}

/**
 * ParseNetwork parses a network of nodes into a map of node to children.
 *
 * The input is expected to be in the following format:
 *   AAA = (BBB, CCC)
 *   BBB = (DDD, EEE)
 *   CCC = (FFF, GGG)
 *   ...
 */
func ParseNetwork(scanner *bufio.Scanner) map[string][]string {
	network := make(map[string][]string)

	for scanner.Scan() {
		node, instructions := parseNode(scanner.Text())

		network[node] = instructions
	}

	return network
}
