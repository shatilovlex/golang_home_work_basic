package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookComparator_Compare(t *testing.T) {
	type args struct {
		left  Book
		right Book
	}
	tests := []struct {
		name  string
		field FieldComparator
		args  args
		want  bool
	}{
		{
			field: Year,
			args: args{
				left: Book{
					id:     1,
					title:  "The Go Programming Language. - Addison-Wesley Professional",
					size:   272,
					rate:   10,
					year:   2015,
					author: "Donovan A., Kernighan В.",
				},
				right: Book{
					id:     2,
					title:  "Test Driven Development: By Example",
					size:   172,
					rate:   9,
					year:   2002,
					author: "Kent Beck",
				},
			},
			want: true,
		},
		{
			field: Rate,
			args: args{
				left: Book{
					id:     1,
					title:  "The Go Programming Language. - Addison-Wesley Professional",
					size:   272,
					rate:   10,
					year:   2015,
					author: "Donovan A., Kernighan В.",
				},
				right: Book{
					id:     2,
					title:  "Test Driven Development: By Example",
					size:   172,
					rate:   9,
					year:   2002,
					author: "Kent Beck",
				},
			},
			want: true,
		},
		{
			field: Size,
			args: args{
				left: Book{
					id:     1,
					title:  "The Go Programming Language. - Addison-Wesley Professional",
					size:   272,
					rate:   10,
					year:   2015,
					author: "Donovan A., Kernighan В.",
				},
				right: Book{
					id:     2,
					title:  "Test Driven Development: By Example",
					size:   172,
					rate:   9,
					year:   2002,
					author: "Kent Beck",
				},
			},
			want: true,
		},
		{
			field: Rate,
			args: args{
				left: Book{
					id:     1,
					title:  "The Go Programming Language. - Addison-Wesley Professional",
					size:   272,
					rate:   10,
					year:   2015,
					author: "Donovan A., Kernighan В.",
				},
				right: Book{
					id:     2,
					title:  "Test Driven Development: By Example",
					size:   172,
					rate:   10,
					year:   2002,
					author: "Kent Beck",
				},
			},
			want: false,
		},
		{
			field: Rate,
			args: args{
				left: Book{
					id:     1,
					title:  "The Go Programming Language. - Addison-Wesley Professional",
					size:   272,
					rate:   9,
					year:   2015,
					author: "Donovan A., Kernighan В.",
				},
				right: Book{
					id:     2,
					title:  "Test Driven Development: By Example",
					size:   172,
					rate:   10,
					year:   2002,
					author: "Kent Beck",
				},
			},
			want: false,
		},
		{
			field: 100500,
			args: args{
				left: Book{
					id:     1,
					title:  "The Go Programming Language. - Addison-Wesley Professional",
					size:   272,
					rate:   10,
					year:   2015,
					author: "Donovan A., Kernighan В.",
				},
				right: Book{
					id:     2,
					title:  "Test Driven Development: By Example",
					size:   172,
					rate:   9,
					year:   2002,
					author: "Kent Beck",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewBookComparator(tt.field)
			got := c.Compare(tt.args.left, tt.args.right)
			assert.Equal(t, tt.want, got)
		})
	}
}
