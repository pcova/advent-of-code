package hotsprings

import (
	"fmt"
	"strconv"
	"strings"
)

type Row struct {
	Springs []Spring
	// blocks is a list of the number of contiguous damaged springs in the row.
	blocks []int
}

func NewRow(springs []Spring, blocks []int) *Row {
	return &Row{
		Springs: springs,
		blocks:  blocks,
	}
}

func (r *Row) Copy() *Row {
	springs := append([]Spring{}, r.Springs...)
	contiguousDamagedSprings := append([]int{}, r.blocks...)

	return &Row{
		Springs: springs,
		blocks:  contiguousDamagedSprings,
	}
}

func (r Row) String() string {
	var sb strings.Builder
	for _, spring := range r.Springs {
		sb.WriteRune(rune(spring))
	}
	sb.WriteString(" ")
	for i, groupSize := range r.blocks {
		sb.WriteString(strconv.Itoa(groupSize))
		if i < len(r.blocks)-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

/**
 * IsValid returns true if the row is valid, i.e. if the number of damaged springs
 * in each group of damaged springs is the same as the number of damaged springs
 * in the row.
 */
func (r Row) IsValid() bool {
	contiguousDamagedSprings := []int{}

	current := 0
	for _, spring := range r.Springs {
		if spring.IsDamaged() {
			current++
		} else {
			if current > 0 {
				contiguousDamagedSprings = append(contiguousDamagedSprings, current)
				current = 0
			}
		}
	}

	if current > 0 {
		contiguousDamagedSprings = append(contiguousDamagedSprings, current)
	}

	if len(contiguousDamagedSprings) != len(r.blocks) {
		return false
	}

	for i, groupSize := range r.blocks {
		if groupSize != contiguousDamagedSprings[i] {
			return false
		}
	}

	return true
}

/**
 * This is a recursive function that tries to find all the possible combinations of
 * operational/damaged springs that can be used to satisfy the damaged springs groups
 * of the row.
 *
 * It starts by checking the first spring of the row. If it is unknown, it will replace
 * it with an operational spring and call itself recursively. Then it will replace it
 * with a damaged spring and call itself recursively. If the spring is already known,
 * it will just call itself recursively.
 *
 * The function will return the number of valid combinations of operational/damaged
 * springs for the rest of the row.
 */
func (r *Row) bruteForce(index int) int {
	if index == len(r.Springs) {
		if r.IsValid() {
			return 1
		}
		return 0
	}

	if r.Springs[index].IsUnknown() {
		operationalRow := r.Copy()
		damagedlRow := r.Copy()

		operationalRow.Springs[index] = Operational
		damagedlRow.Springs[index] = Damaged

		return operationalRow.bruteForce(index+1) + damagedlRow.bruteForce(index+1)
	}

	return r.bruteForce(index + 1)
}

var cache = map[string]int{}

func CacheClear() {
	cache = map[string]int{}
}

/**
 * i - index of the spring we are currently checking.
 * bi - index of the block we are currently checking.
 * current - number of contiguous damaged springs we have found so far (or size of the current block).
 */
func (r *Row) optimizedSolution(i, bi, current int) int {
	key := fmt.Sprintf("%d-%d-%d", i, bi, current)
	if value, ok := cache[key]; ok {
		return value
	}

	if i == len(r.Springs) {
		if bi == len(r.blocks) && current == 0 {
			return 1
		} else if bi == len(r.blocks)-1 && current == r.blocks[bi] {
			return 1
		}
		return 0
	}

	ans := 0
	for _, c := range []rune{'.', '#'} {
		if r.Springs[i] == Spring(c) || r.Springs[i].IsUnknown() {
			if c == '.' && current == 0 {
				// operationalRow := r.Copy()
				// operationalRow.Springs[springIndex].Type = Operational
				ans += r.optimizedSolution(i+1, bi, 0)
			} else if c == '.' && current > 0 && bi < len(r.blocks) && current == r.blocks[bi] {
				// operationalRow := r.Copy()
				// operationalRow.Springs[springIndex].Type = Operational
				ans += r.optimizedSolution(i+1, bi+1, 0)
			} else if c == '#' {
				// damagedRow := r.Copy()
				// damagedRow.Springs[springIndex].Type = Damaged
				ans += r.optimizedSolution(i+1, bi, current+1)
			}
		}
	}
	cache[key] = ans

	return ans
}

/**
 * ValidAlternativeRows returns the number of valid combinations of operational/damaged
 * springs for the row.
 */
func (r *Row) ValidAlternativeRows(bruteForce bool) int {
	if bruteForce {
		return r.bruteForce(0)
	}
	return r.optimizedSolution(0, 0, 0)
}
