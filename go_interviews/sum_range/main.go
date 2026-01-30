package main

import "errors"

var (
	ErrEmptyData = errors.New("empty")
	ErrIncorrectIdx = errors.New("incorrect idx")
)

type (
	Data struct {
		prefixs []int
	}
)

func New(raw []int) (*Data, error) {
	if (len(raw) == 0) {
		return nil, ErrEmptyData
	} 
	out := make([]int, len(raw) + 1)
	out[0] = 0
	for i, s := range raw {
		out[i+1] = out[i] + s
	}
	return &Data{prefixs: out}, nil
}

func (d *Data) SumByRange(left, right int) (int, error) {
	if left > right || left < 0 || right > len(d.prefixs) - 2 {
		return -1, ErrIncorrectIdx
	}

	return d.prefixs[right + 1] - d.prefixs[left], nil
}

func main() {

}