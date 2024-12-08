package types

type Book struct {
	Id     int     `json:"id,omitempty" xml:"id,omitempty" yaml:"id,omitempty"`
	Year   int     `json:"year,omitempty" xml:"year,omitempty" yaml:"year,omitempty"`
	Size   int     `json:"size,omitempty" xml:"size,omitempty" yaml:"size,omitempty"`
	Rate   float64 `json:"rate,omitempty" xml:"rate,omitempty" yaml:"rate,omitempty"`
	Title  string  `json:"title,omitempty" xml:"title,omitempty" yaml:"title,omitempty"`
	Author string  `json:"author,omitempty" xml:"author,omitempty" yaml:"author,omitempty"`
	Sample []byte  `json:"sample,omitempty" xml:"sample,omitempty" yaml:"sample,omitempty"`
}
