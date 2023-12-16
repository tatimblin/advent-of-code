package modules

import (
	"strconv"
	"strings"
)

type Hash struct {
	Value string
	Slot  int
	Box   int
}

type HashMap struct {
	Map           [256][9]Hash
	FocusingPower int
}

type HashMapInterface interface {
	Get(box int, slot int) Hash
	Set(Hash)
	Remove(Hash)
	Encode(string) Hash
}

func (hm *HashMap) Get(box int, slot int) Hash {
	return (*hm).Map[box][slot]
}

func (hm *HashMap) Set(hash Hash) {
	for i, slot := range (*hm).Map[hash.Box] {
		if slot.Value == hash.Value || slot == (Hash{}) {
			(*hm).Map[hash.Box][i] = hash
			break
		}
	}
}

func (hm *HashMap) Remove(hash Hash) {
	for i, slot := range (*hm).Map[hash.Box] {
		if slot.Value == hash.Value {
			(*hm).Map[hash.Box][i] = Hash{}
		}
	}

	for i := 0; i < len((*hm).Map[hash.Box])-1; i++ {
		if (*hm).Map[hash.Box][i] == (Hash{}) && (*hm).Map[hash.Box][i+1] != (Hash{}) {
			(*hm).Map[hash.Box][i], (*hm).Map[hash.Box][i+1] = (*hm).Map[hash.Box][i+1], (*hm).Map[hash.Box][i]
		}
	}
}

func Encode(str string) Hash {
	var slot int
	parts := strings.Split(str, "=")
	if len(parts) == 2 {
		slot, _ = strconv.Atoi(parts[1])
	} else {
		parts = strings.Split(str, "-")
	}

	var box int
	chars := []rune(parts[0])
	for i := 0; i < len(chars); i++ {
		box += int(chars[i])
		box *= 17
		box %= 256
	}
	return Hash{
		Value: parts[0],
		Box:   box,
		Slot:  slot,
	}
}
