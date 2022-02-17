package ring

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	tree := &BalanceTree{}
	ok := tree.Add(11)
	if !ok {
		t.Errorf("BalanceTree add error")
	}
	ok = tree.Add(12)
	if !ok {
		t.Errorf("BalanceTree add error")
	}
	ok = tree.Add(13)
	if !ok {
		t.Errorf("BalanceTree add error")
	}
	ok = tree.Add(14)
	if !ok {
		t.Errorf("BalanceTree add error")
	}
	log.Println("TestAdd nodes: ", tree.Layout())
}

func TestRemove(t *testing.T) {
	tree := &BalanceTree{}
	ok := tree.Add(11)
	if !ok {
		t.Errorf("BalanceTree add error")
	}
	ok = tree.Add(12)
	if !ok {
		t.Errorf("BalanceTree add error")
	}
	ok = tree.Add(13)
	if !ok {
		t.Errorf("BalanceTree add error")
	}
	ok = tree.Add(14)
	if !ok {
		t.Errorf("BalanceTree add error")
	}
	log.Println("TestAdd nodes before removing: ", tree.Layout())
	tree.Remove(14)
	log.Println("TestAdd nodes after removing: ", tree.Layout())
}
