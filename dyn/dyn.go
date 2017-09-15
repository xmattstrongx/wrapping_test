package dyn

import (
	"github.com/nesv/go-dynect/dynect"
)

type myDyn struct {
	M myDynI
}

func NewMyDyn(m myDynI) *myDyn {
	return &myDyn{m}
}

type myDynI interface {
	CreateRecord(record *dynect.Record) error
}

func (m *myDyn) CreateRecord(record *dynect.Record) error {
	return m.M.CreateRecord(record)
}

func (m *myDyn) CreateDNS() error {
	err := m.M.CreateRecord(&dynect.Record{})
	if err != nil {
		return err
	}
	return nil
}
