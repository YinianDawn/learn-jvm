package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (entries CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range entries {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (entries CompositeEntry) String() string {
	str := make([]string, len(entries))
	for i, entry := range entries {
		str[i] = entry.String()
	}
	return strings.Join(str, pathListSeparator)
}
