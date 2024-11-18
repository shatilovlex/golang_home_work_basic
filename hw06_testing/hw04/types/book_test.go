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

const (
	id     = 1
	year   = 2024
	size   = 100
	rate   = 10.0
	title  = "title"
	author = "author"
)

func TestBook(t *testing.T) {
	b := &Book{
		id:     id,
		year:   year,
		size:   size,
		rate:   rate,
		title:  title,
		author: author,
	}

	assert.Equal(t, id, b.ID())
	assert.Equal(t, year, b.Year())
	assert.Equal(t, size, b.Size())
	assert.Equal(t, rate, b.Rate())
	assert.Equal(t, title, b.Title())
	assert.Equal(t, author, b.Author())
}

func TestBook_Setters(t *testing.T) {
	b := &Book{}

	b.SetID(id)
	b.SetYear(year)
	b.SetSize(size)
	b.SetRate(rate)
	b.SetTitle(title)
	b.SetAuthor(author)

	assert.Equal(t, id, b.ID())
	assert.Equal(t, year, b.Year())
	assert.Equal(t, size, b.Size())
	assert.Equal(t, rate, b.Rate())
	assert.Equal(t, title, b.Title())
	assert.Equal(t, author, b.Author())
}
