package dns

type Record struct {
	RType string
	RData string
	RKey  string
}

type Recorder interface {
	Get(string) (Record, error)
	Set(Record) error
}

func (r *Record) Get(key string) (Record, error) {
	return Record{}, nil
}

func (r *Record) Set(Record) error {
	return nil
}
