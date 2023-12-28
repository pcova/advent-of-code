package parts

import (
	"fmt"
	"strings"
)

type Workflow struct {
	rules []Rule
}

func (w Workflow) Evaluate(part Part) (string, error) {
	for _, rule := range w.rules {
		if rule.evaluate(part) {
			return rule.Action(), nil
		}
	}

	return "", fmt.Errorf("no rule matched")
}

func ParseWorkflow(line string) (string, *Workflow) {
	var workflow *Workflow = new(Workflow)

	rulesStart := strings.Index(line, "{")

	name := line[:rulesStart]
	rulesStr := strings.TrimSuffix(line[rulesStart+1:], "}")

	rules := make([]Rule, 0)
	rulesSplit := strings.Split(rulesStr, ",")
	for i, ruleStr := range rulesSplit {
		if i == len(rulesSplit)-1 {
			rules = append(rules, Rule{"", ruleStr})
			continue
		}

		ruleSplit := strings.Split(ruleStr, ":")
		rules = append(rules, Rule{ruleSplit[0], ruleSplit[1]})
	}

	workflow.rules = rules

	return name, workflow
}

type Workflows map[string]Workflow

/**
 * Given a list of parts, returns a list of approved parts
 */
func (w Workflows) ApprovedParts(parts []Part) []Part {
	approvedParts := make([]Part, 0)

	firstWorkflow := w["in"]
	for _, part := range parts {
		currentWorkflow := firstWorkflow
		for {
			result, err := currentWorkflow.Evaluate(part)
			if err != nil {
				panic(err)
			}

			if result == "A" {
				// If the result is accept, add the part to the list of approved parts and break
				approvedParts = append(approvedParts, part)
				break
			} else if result == "R" {
				// If the result is reject, do nothing (break)
				break
			}

			// If the result is not accept or reject, update the current workflow
			currentWorkflow = w[result]
		}
	}

	return approvedParts
}

/**
 * Given a range of parts, returns the number of possible combinations
 */
func (w Workflows) combinations(workflowName string, ranges PartRange) int {
	combinations := 0
	workflow := w[workflowName]
	for _, rule := range workflow.rules {
		// If the rule is a default rule
		if rule.IsDefault() {
			// If the action is accept, return the number of possibilities
			if rule.action == "A" {
				return combinations + ranges.Possibilities()
			}

			// If the action is other update the workflow name and break
			workflowName = rule.action
			break
		}

		operator, _ := rule.Operator()
		key := rule.Key()

		ruleRange, err := rule.Range()
		if err != nil {
			panic(err)
		}

		var newRanges PartRange
		if operator == ">" {
			// If the operator is >, split the range by the min value
			ranges, newRanges = ranges.splitRange(key, ruleRange.Min)
		} else {
			// If the operator is <, split the range by the max value + 1
			newRanges, ranges = ranges.splitRange(key, ruleRange.Max+1)
		}

		// If the action is accept, add the number of possibilities for the new part range
		if rule.action == "A" {
			combinations += newRanges.Possibilities()
			continue
		}

		// If the action is other than reject, calculate the number of possibilities for the new workflow
		if rule.action != "R" {
			combinations += w.combinations(rule.action, newRanges)
		}
	}

	// If the workflow name is not reject, add the number of possibilities for the new workflow
	if workflowName != "R" {
		return combinations + w.combinations(workflowName, ranges)
	}

	return combinations
}

/**
 * Calculates the total number of possible combinations
 */
func (w Workflows) PossibleCombinations() int {
	ranges := PartRange{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}

	return w.combinations("in", ranges)
}
