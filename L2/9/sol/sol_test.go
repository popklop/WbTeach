package sol

import (
	"testing"
)

func TestDefault(t *testing.T) {
	res, err := Stringdecoder("a4bc2d5e")
	if err != nil {
		t.Error(err)
	}
	exp := "aaaabccddddde"
	if res != exp {
		t.Error("Результат не совпадает с ожидаемым!")
	}
}
func TestAllchar(t *testing.T) {
	res, err := Stringdecoder("abcd")
	if err != nil {
		t.Error(err)
	}
	exp := "abcd"
	if res != exp {
		t.Error("Результат не совпадает с ожидаемым!")
	}
}
func TestAlldigit(t *testing.T) {
	_, err := Stringdecoder("45")
	if err == nil {
		t.Error(err)
	}
	expectError := "строка не может начинатся с цифры"
	if err.Error() != expectError {
		t.Error("Expected another error")
	}
}
func TestEmptystring(t *testing.T) {
	res, err := Stringdecoder("")
	if err != nil {
		t.Error(err)
	}
	exp := ""
	if res != exp {
		t.Error("Результат не совпадает с ожидаемым!")
	}
}
func TestBigData(t *testing.T) {
	res, err := Stringdecoder("a10bc11")
	if err != nil {
		t.Error(err)
	}
	exp := "aaaaaaaaaabccccccccccc"
	if res != exp {
		t.Error("Результат не совпадает с ожидаемым!")
	}
}
func TestEscape1(t *testing.T) {
	res, err := Stringdecoder("qwe\\4\\5")
	if err != nil {
		t.Error(err)
	}
	exp := "qwe45"
	if res != exp {
		t.Error("Результат не совпадает с ожидаемым!")
	}
}
func TestEscape2(t *testing.T) {
	res, err := Stringdecoder("qwe\\45")
	if err != nil {
		t.Error(err)
	}
	exp := "qwe44444"
	if res != exp {
		t.Error("Результат не совпадает с ожидаемым!")
	}
}
