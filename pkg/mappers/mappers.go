package mappers

type Mapper interface {
	// Write writes the intermediate file
	Write() error
	// Read reads the input file
	Read() error
	// Close notify the masters that the mapper is done
	Close() error
}
