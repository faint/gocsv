package gocsv

// Key the key of the csv content
type Key struct {
	Value []string
}

// GetIndex input the fieldName, output the index.
func (key *Key) GetIndex(fieldName string) (int, bool) {
	for k, v := range key.Value {
		if v == fieldName {
			return k, true
		}
	}

	return 0, false
}
