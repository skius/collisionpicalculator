package main

type Block struct {
	m float64
	v float64
}

func Bounce(b1, b2 *Block) {
	sumM := b1.m + b2.m
	newV1 := ((b1.m - b2.m)/sumM) * b1.v + (2*b2.m/sumM) * b2.v
	newV2 := (2*b1.m/sumM) * b1.v + (b2.m - b1.m)/sumM * b2.v
	b1.v = newV1
	b2.v = newV2
}

func (b *Block) Reverse() {
	b.v *= -1
}
