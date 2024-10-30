package types

type FieldComparator int

const (
	Year FieldComparator = iota
	Size
	Rate
)

type BookComparator struct {
	field FieldComparator
}

func (c BookComparator) Compare(left Book, right Book) bool {
	switch c.field {
	case Year:
		return left.Year() > right.Year()
	case Size:
		return left.Size() > right.Size()
	case Rate:
		return left.Rate() > right.Rate()
	default:
		return false
	}
}

func NewBookComparator(field FieldComparator) *BookComparator {
	return &BookComparator{field: field}

}
