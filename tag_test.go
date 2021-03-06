package multiconfig

import "testing"

func TestTagLoad(t *testing.T) {
	t.Run("tag loader should not return an error for a map source", loadTagMapSource)
	t.Run("tag loader should not return an error for underlying pointer type map source", loadTagUnderlyingPointerMapSource)
	t.Run("tag loader should return an error for underlying type map source", loadTagUnderlyingMapSource)
	t.Run("tag loader should return nonPointerError for non-pointer source", testTagNonPointerError)
}

func loadTagUnderlyingPointerMapSource(t *testing.T) {
	type underlyingMap map[string]interface{}
	var u underlyingMap

	l := &TagLoader{}
	err := l.Load(&u)
	if err != nil {
		t.Error(err)
	}
}

func loadTagUnderlyingMapSource(t *testing.T) {
	type underlyingMap map[string]interface{}
	var u underlyingMap

	l := &TagLoader{}
	err := l.Load(u)
	if err != nonPointerError {
		t.Error("expected nonPointerError")
	}
}

func loadTagMapSource(t *testing.T) {
	l := &TagLoader{}
	s := map[string]interface{}{
		"foo": "foo",
	}

	err := l.Load(&s)
	if err != nil {
		t.Error(err)
	}
}

func testTagNonPointerError(t *testing.T) {
	l := &TagLoader{}
	s := map[string]interface{}{
		"foo": "foo",
	}

	err := l.Load(s)
	if err == nil {
		t.Error("expected nonPointerError")
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

func TestDefaultValues(t *testing.T) {
	m := &TagLoader{}
	s := new(Server)
	if err := m.Load(s); err != nil {
		t.Error(err)
	}

	if s.Port != getDefaultServer().Port {
		t.Errorf("Port value is wrong: %d, want: %d", s.Port, getDefaultServer().Port)
	}

	if s.Postgres.DBName != getDefaultServer().Postgres.DBName {
		t.Errorf("Postgres DBName value is wrong: %s, want: %s", s.Postgres.DBName, getDefaultServer().Postgres.DBName)
	}
}
