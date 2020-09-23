package dict

type Dict struct {
	d map[string]string
}

func NewDict() *Dict {
	return &Dict{
		d: make(map[string]string),
	}
}

func (d *Dict) Add(key, value string) {
	d.d[key] = value
}

func (d *Dict) Get(key string) string {
	return d.d[key]
}
