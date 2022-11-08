package model

import (
	"reflect"
	"testing"
)

type Foo struct {
	Foo string
	Bar string
}

type fooModel struct{}

// Convert takes an arbitrary map and converts it via the given mapping
func (f *fooModel) Convert(d map[string]interface{}, m *Mapping) ModelWithIndexes[Foo] {
	fs := Foo{
		Foo: d[m.mappings["foo"]].(string),
		Bar: d[m.mappings["bar"]].(string),
	}

	return ModelWithIndexes[Foo]{
		model:   fs,
		indexes: m.indexes,
	}
}

func Test_Convert(t *testing.T) {
	m := &Mapping{
		predicateToMap: "https://foo",
		indexes:        []string{"foo"},
		mappings:       map[string]string{"foo": "baz", "bar": "bar"},
	}

	d := map[string]interface{}{
		"baz": "a",
		"bar": "bar",
	}

	t.Run("simple convert", func(t *testing.T) {
		want := ModelWithIndexes[Foo]{
			model: Foo{
				Foo: "a",
				Bar: "bar",
			},
			indexes: []string{"foo"},
		}
		fm := &fooModel{}
		got := fm.Convert(d, m)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("DataModel.Convert() = %v, expected %v", got, want)
		}
	})

}

func Test_ConvertAny(t *testing.T) {
	m := &Mapping{
		predicateToMap: "https://foo",
		indexes:        []string{"foo"},
		mappings:       map[string]string{"foo": "baz", "bar": "bar"},
	}

	d := map[string]interface{}{
		"baz": "a",
		"bar": "bar",
	}

	t.Run("simple convert", func(t *testing.T) {
		want := ModelWithIndexes[Foo]{
			model: Foo{
				Foo: "a",
				Bar: "bar",
			},
			indexes: []string{"foo"},
		}
		got, err := ConvertAny[Foo](d, m)
		if err != nil {
			t.Errorf("DataModel.ConverAny() had error: %v, expected nil", err)
		}
		if !reflect.DeepEqual(*got, want) {
			t.Errorf("DataModel.ConvertAny() = %v, expected %v", *got, want)
		}
	})

}
