package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
