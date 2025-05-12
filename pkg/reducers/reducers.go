package reducers

type Reducer interface {
	// Read reads the intermediate file
	Read() error
	// Write writes the output file
	Write() error
	// Close notify the masters that the reducer is done
	Close() error
}
