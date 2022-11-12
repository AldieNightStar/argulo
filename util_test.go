package argulo

import "testing"

func TestToMap(t *testing.T) {
	mp := toMap([]string{
		"a", "b", "-age", "18", "-33", "-done", "-name", "Ihor", "-applied",
	})
	if mp["default"][0] != "a" {
		t.Fatal()
	}
	if mp["default"][1] != "b" {
		t.Fatal()
	}
	if mp["age"][0] != "18" {
		t.Fatal()
	}
	if mp["age"][1] != "-33" {
		t.Fatal()
	}
	if mp["name"][0] != "Ihor" {
		t.Fatal()
	}
	if mp["applied"][0] != "true" {
		t.Fatal()
	}
	if mp["done"][0] != "true" {
		t.Fatal()
	}
}
