package lens

import "fmt"

type Box struct {
	index      int
	lens       map[string]int
	sortedKeys []string
}

func NewBox(index int) *Box {
	return &Box{
		index: index,
		lens:  make(map[string]int),
	}
}

func (b Box) String() string {
	str := fmt.Sprintf("Box %d:\n", b.index)
	for _, key := range b.sortedKeys {
		str += fmt.Sprintf("  %s: %d\n", key, b.lens[key])
	}
	return str
}

func (b *Box) Index() int {
	return b.index
}

func (b *Box) AddLen(key string, value int) {
	b.lens[key] = value
	b.sortedKeys = append(b.sortedKeys, key)
}

func (b *Box) UpdateLen(key string, value int) {
	b.lens[key] = value
}

func (b *Box) GetLen(key string) (int, error) {
	if value, ok := b.lens[key]; ok {
		return value, nil
	}
	return 0, fmt.Errorf("no such key: %s", key)
}

func (b *Box) RemoveLen(key string) {
	delete(b.lens, key)
	for i, k := range b.sortedKeys {
		if k == key {
			b.sortedKeys = append(b.sortedKeys[:i], b.sortedKeys[i+1:]...)
			break
		}
	}
}

func (b *Box) FocusingPower() int {
	var power int
	for i, key := range b.sortedKeys {
		power += (b.index + 1) * (i + 1) * b.lens[key]
	}
	return power
}
