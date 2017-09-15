package dyn

import (
	"errors"
	"fmt"
	"testing"

	"github.com/nesv/go-dynect/dynect"
)

type mockMyDyn struct{}

func NewMockMyDyn(m myDynI) myDynI {
	return &mockMyDyn{}
}

func (m *mockMyDyn) CreateRecord(record *dynect.Record) error {
	return errors.New("FAKE ERROR")
}

func TestCreateDNS(t *testing.T) {
	mock := NewMyDyn(&mockMyDyn{})
	err := mock.CreateDNS()
	if err != nil {
		t.Logf("Sweet, got expected err: %v", err)
	}
}

func TestCreateDNS2(t *testing.T) {
	var test string
	var Test string
	var TEST string

	fmt.Println(test)
	fmt.Println(Test)
	fmt.Println(TEST)
}
