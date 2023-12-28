package parts

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Rule struct {
	condition string
	action    string
}

func (r Rule) IsDefault() bool {
	// If the rule has no condition, it's a default rule
	return r.condition == ""
}

/**
 * Evaluate a part against a rule
 */
func (r Rule) evaluate(part Part) bool {
	if r.IsDefault() {
		return true
	}
	operator, operatorIdx := r.Operator()

	conditionValue, err := strconv.Atoi(r.condition[operatorIdx+1:])
	if err != nil {
		panic(err)
	}

	partValue := 0
	switch r.condition[:operatorIdx] {
	case "x":
		partValue = part.x
	case "m":
		partValue = part.m
	case "a":
		partValue = part.a
	case "s":
		partValue = part.s
	}

	if operator == ">" {
		return partValue > conditionValue
	}

	return partValue < conditionValue
}

/**
 * Get the rule's action
 */
func (r Rule) Action() string {
	return r.action
}

/**
 * Get the rule's operator
 */
func (r Rule) Operator() (string, int) {
	operator := ">"
	operatorIdx := strings.Index(r.condition, operator)
	if operatorIdx == -1 {
		operator = "<"
		operatorIdx = strings.Index(r.condition, operator)
	}

	return operator, operatorIdx
}

/**
 * Get the rule's key being evaluated
 */
func (r Rule) Key() string {
	_, operatorIdx := r.Operator()
	return r.condition[:operatorIdx]
}

/**
 * Get the max value that satisfies the rule
 */
func (r Rule) maxValue() (int, error) {
	operator, operatorIdx := r.Operator()

	if operator == ">" {
		return math.MaxInt64, nil
	}

	conditionValue, err := strconv.Atoi(r.condition[operatorIdx+1:])
	if err != nil {
		return 0, fmt.Errorf("error converting condition value to int: %s", err)
	}

	return conditionValue - 1, nil
}

/**
 * Get the min value that satisfies the rule
 */
func (r Rule) minValue() (int, error) {
	operator, operatorIdx := r.Operator()

	if operator == "<" {
		return 0, nil
	}

	conditionValue, err := strconv.Atoi(r.condition[operatorIdx+1:])
	if err != nil {
		return 0, fmt.Errorf("error converting condition value to int: %s", err)
	}

	return conditionValue + 1, nil
}

/**
 * Get the range of values that satisfy the rule
 */
func (r Rule) Range() (Range, error) {
	minValue, err := r.minValue()
	if err != nil {
		return Range{}, fmt.Errorf("error getting min value: %s", err)
	}

	maxValue, err := r.maxValue()
	if err != nil {
		return Range{}, fmt.Errorf("error getting max value: %s", err)
	}

	return Range{Min: minValue, Max: maxValue}, nil
}
