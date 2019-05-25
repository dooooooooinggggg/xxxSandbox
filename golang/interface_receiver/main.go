package main

import "fmt"

type Taiyaki interface {
	GetNakami() string
	SetNakami(nakami string)
}

/* 実体 */
func (t *normalTaiyaki) GetNakami() string {
	return t.Nakami
}

func (t *normalTaiyaki) SetNakami(nakami string) {
	t.Nakami = nakami
}

type normalTaiyaki struct {
	Nakami string
}

// interfaceを返す．ここで結びつけるのか
// interface型のstructを返す
func NewNormalTaiyaki() Taiyaki {
	return &normalTaiyaki{}
}

func main() {
	t1 := NewNormalTaiyaki()
	t1.SetNakami("あんこ")
	fmt.Println(t1.GetNakami()) // あんこ
}
