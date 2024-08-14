package simplify_path

import "testing"

func TestExample1(t *testing.T) {
	output := SimplifyPath("/home/")
	expected := "/home"
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample2(t *testing.T) {
	output := SimplifyPath("/home//foo/")
	expected := "/home/foo"
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample3(t *testing.T) {
	output := SimplifyPath("/home/user/Documents/../Pictures")
	expected := "/home/user/Pictures"
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample4(t *testing.T) {
	output := SimplifyPath("/../")
	expected := "/"
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}

func TestExample5(t *testing.T) {
	output := SimplifyPath("/.../a/../b/c/../d/./")
	expected := "/.../b/d"
	if output != expected {
		t.Errorf("Test failed, expected: '%v', got: '%v'", expected, output)
	}
}
