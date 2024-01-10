package basic_test

import (
	"testing"

	"github.com/go-test/deep"
)

type Subject struct {
	Name  string
	Marks int
}

type Student struct {
	Name     string
	Subjects []Subject
}

func TestForFailure(t *testing.T) {
	foo := Student{
		Name: "foo",
		Subjects: []Subject{
			{
				Name:  "Maths",
				Marks: 70,
			},
			{
				Name:  "History",
				Marks: 60,
			},
		},
	}

	bar := Student{
		Name: "bar",
		Subjects: []Subject{
			{
				Name:  "Maths",
				Marks: 70,
			},
			{
				Name:  "History",
				Marks: 61,
			},
		},
	}

	if diff := deep.Equal(foo, bar); diff != nil {
		t.Error(diff)
	}
}

func TestForSuccess(t *testing.T) {
	foo := Student{
		Name: "foo",
		Subjects: []Subject{
			{
				Name:  "Maths",
				Marks: 70,
			},
			{
				Name:  "History",
				Marks: 60,
			},
		},
	}

	foo2 := Student{
		Name: "foo",
		Subjects: []Subject{
			{
				Name:  "Maths",
				Marks: 70,
			},
			{
				Name:  "History",
				Marks: 60,
			},
		},
	}

	if diff := deep.Equal(foo, foo2); diff != nil {
		t.Error(diff)
	}
}
