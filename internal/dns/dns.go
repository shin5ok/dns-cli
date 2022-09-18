package dns

type Record struct {
	RType  string
	RData  string
	RKey   string
	TTL    int
	Status string
}

type Recorder interface {
	Get(string) (*Record, error)
	Set(*Record) error
}

type Rr struct {
	Domain string
}

func (r *Rr) Get(key string) (*Record, error) {
	return &Record{}, nil
}

func (r *Rr) Set(*Record) error {
	return nil
}
