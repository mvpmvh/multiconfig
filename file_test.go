package multiconfig

import (
	"os"
	"testing"
)

func TestFileLoad(t *testing.T) {
	t.Run("file loader should not return an error for a map source", loadFileMapSource)
	t.Run("file loader should return nonPointerError for non-pointer source", testFileNonPointerError)
}

func loadFileMapSource(t *testing.T) {
	loaders := []Loader{
		&JSONLoader{},
		&TOMLLoader{},
		&YAMLLoader{},
	}

	for _, l := range loaders {
		s := map[string]interface{}{
			"foo": "foo",
		}

		err := l.Load(&s)
		if err != nil {
			t.Error(err)
		}
	}
}

func testFileNonPointerError(t *testing.T) {
	loaders := []Loader{
		&JSONLoader{},
		&TOMLLoader{},
		&YAMLLoader{},
	}

	for _, l := range loaders {
		s := map[string]interface{}{
			"foo": "foo",
		}

		err := l.Load(s)
		if err == nil {
			t.Error("expected nonPointerError")
		}
	}

	//s2 := struct {
	//	Foo string
	//}{"foo"}
	//
	//err = l.Load(s2)
	//if err == nil {
	//	t.Error("expected nonPointerError")
	//}
}

func TestYAML(t *testing.T) {
	m := NewWithPath(testYAML)

	s := &Server{}
	if err := m.Load(s); err != nil {
		t.Error(err)
	}

	testStruct(t, s, getDefaultServer())
}

func TestYAML_Reader(t *testing.T) {
	f, err := os.Open(testYAML)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	l := MultiLoader(&TagLoader{}, &YAMLLoader{Reader: f})
	s := &Server{}
	if err := l.Load(s); err != nil {
		t.Error(err)
	}

	testStruct(t, s, getDefaultServer())
}
func TestToml(t *testing.T) {
	m := NewWithPath(testTOML)

	s := &Server{}
	if err := m.Load(s); err != nil {
		t.Error(err)
	}

	testStruct(t, s, getDefaultServer())
}

func TestToml_Reader(t *testing.T) {
	f, err := os.Open(testTOML)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	l := MultiLoader(&TagLoader{}, &TOMLLoader{Reader: f})
	s := &Server{}
	if err := l.Load(s); err != nil {
		t.Error(err)
	}

	testStruct(t, s, getDefaultServer())
}

func TestJSON(t *testing.T) {
	m := NewWithPath(testJSON)

	s := &Server{}
	if err := m.Load(s); err != nil {
		t.Error(err)
	}

	testStruct(t, s, getDefaultServer())
}

func TestJSON_Reader(t *testing.T) {
	f, err := os.Open(testJSON)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	l := MultiLoader(&TagLoader{}, &JSONLoader{Reader: f})
	s := &Server{}
	if err := l.Load(s); err != nil {
		t.Error(err)
	}

	testStruct(t, s, getDefaultServer())
}

// func TestJSON2(t *testing.T) {
// 	ExampleEnvironmentLoader()
// 	ExampleTOMLLoader()
// }
