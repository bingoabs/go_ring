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

func TestSearch(t *testing.T) {
	tree := &BalanceTree{}
	_ = tree.Add(3)
	_ = tree.Add(10)
	_ = tree.Add(15)
	_ = tree.Add(20)

	val := tree.Find(1)
	if val != 3 {
		t.Errorf("BalanceTree can't find the match val")
	}
	val = tree.Find(10)
	if val != 10 {
		t.Errorf("BalanceTree can't find the match val")
	}
	val = tree.Find(11)
	if val != 15 {
		t.Errorf("BalanceTree can't find the match val")
	}
	val = tree.Find(19)
	if val != 20 {
		t.Errorf("BalanceTree can't find the match val")
	}
	val = tree.Find(25)
	if val != 3 {
		t.Errorf("BalanceTree can't find the match val")
	}
}

func TestRemove(t *testing.T) {
	tree := &BalanceTree{}
	_ = tree.Add(11)
	_ = tree.Add(12)
	_ = tree.Add(13)
	_ = tree.Add(14)

	log.Println("TestRemove nodes before removing: ", tree.Layout())
	tree.Remove(14)
	log.Println("TestRemove nodes after removing: ", tree.Layout())
	is_exists := tree.Exists(14)
	if is_exists != false {
		t.Errorf("TestRemove nodes still exists")
	}
}
