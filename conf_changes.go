package gconf

import "fmt"

type ChangeMode int

const (
	Created ChangeMode = iota
	Removed
	Modified
)

type confChanges struct {
	changes []Change
}

func (c *confChanges) GetNumOfChanges() int {
	return len(c.changes)
}

func (c *confChanges) GetChanges() []Change {
	return c.changes
}

func EmptyConfChanges() ConfChanges {
	var changes []Change

	return &confChanges{
		changes: changes,
	}
}

func copyMap(in map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{})

	for k, v := range in {
		out[k] = v
	}

	return out
}

func CalcConfChanges(current map[string]interface{}, prev map[string]interface{}) ConfChanges {
	var changes []Change

	current = copyMap(current)
	prev = copyMap(prev)

	if prev == nil && current == nil {
		return newConfChanges(changes)
	}

	if prev == nil {
		for k, v := range current {
			change := newChanges(k, Created, nil, v)
			changes = append(changes, change)
		}

		return newConfChanges(changes)
	}

	for k, v := range current {

		prevValue, ok := prev[k]
		delete(prev, k)

		if !ok {
			change := newChanges(k, Created, nil, v)
			changes = append(changes, change)
			continue
		}

		if prevValue != v {
			change := newChanges(k, Modified, prevValue, v)
			changes = append(changes, change)
			continue
		}
	}

	for k, v := range prev {
		change := newChanges(k, Removed, v, nil)
		changes = append(changes, change)
	}

	return newConfChanges(changes)
}

func newConfChanges(changes []Change) *confChanges {
	return &confChanges{
		changes: changes,
	}
}

type Change struct {
	KeyName string
	Mode    ChangeMode
	Prev    interface{}
	Current interface{}
}

func newChanges(keyName string, mode ChangeMode, prev interface{}, current interface{}) Change {
	return Change{
		KeyName: keyName,
		Mode:    mode,
		Prev:    prev,
		Current: current,
	}
}

func (c *Change) String() string {
	return fmt.Sprintf("[%s] key: %s, prev: %v, current: %v", changeModeToString(c.Mode), c.KeyName, c.Prev, c.Current)
}

func changeModeToString(mode ChangeMode) string {
	switch mode {
	case Created:
		return "created"
	case Removed:
		return "removed"
	case Modified:
		return "modified"
	default:
		return "unknown"
	}
}
