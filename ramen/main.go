package main

import (
	"fmt"
	"strings"
)

// RamenOrder はラーメンの注文を表す構造体
type RamenOrder struct {
	garlic     Amount // にんにく
	vegetables Amount // 野菜
	oilLevel   Amount // 油の量
}

// RamenOption はラーメンのオプションを設定する関数の型
type RamenOption func(*RamenOrder)

// デフォルトのラーメン設定
func defaultRamenSetting() *RamenOrder {
	return &RamenOrder{
		garlic:     Regular,
		vegetables: Regular,
		oilLevel:   Regular,
	}
}

type Amount int

const (
	Light Amount = iota
	Regular
	Extra
	Ultimate
)

// String メソッドを実装して、各値を文字列として表現できるようにする
func (a Amount) String() string {
	switch a {
	case Light:
		return "少なめ"
	case Regular:
		return "普通"
	case Extra:
		return "マシ"
	case Ultimate:
		return "マシマシ"
	default:
		return "不明な量"
	}
}

// WithGarlic にんにく量を指定するオプション
func WithGarlic(a Amount) RamenOption {
	return func(r *RamenOrder) {
		r.garlic = a
	}
}

// WithVegetables 野菜の量を指定するオプション
func WithVegetables(a Amount) RamenOption {
	return func(r *RamenOrder) {
		r.vegetables = a
	}
}

// WithOilLevel 油の量を設定するオプション
func WithOilLevel(a Amount) RamenOption {
	return func(r *RamenOrder) {
		r.oilLevel = a
	}
}

// NewRamenOrder 新しいラーメンの注文を作成する
func NewRamenOrder(opts ...RamenOption) *RamenOrder {
	r := defaultRamenSetting()
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// Call はRamenのオプション情報を文字列として返す
func (r *RamenOrder) Call() string {
	options := []string{}

	if r.garlic != Regular {
		options = append(options, fmt.Sprintf("にんにく%s", r.garlic))
	}

	if r.vegetables != Regular {
		options = append(options, fmt.Sprintf("野菜%s", r.vegetables))
	}

	if r.oilLevel != Regular {
		options = append(options, fmt.Sprintf("油%s", r.oilLevel))
	}

	optionStr := ""
	if len(options) > 0 {
		optionStr = fmt.Sprintf("、%sで!!", strings.Join(options, "、"))
	}

	return fmt.Sprintf("注文入りました！ラーメン一丁%s", optionStr)
}

func main() {
	// はじめて
	ramenOrder1 := NewRamenOrder()
	fmt.Println(ramenOrder1.Call())

	// 健康志向
	ramenOrder3 := NewRamenOrder(
		WithVegetables(Ultimate),
	)
	fmt.Println(ramenOrder3.Call())

	// 深夜の背徳感
	ramenOrder2 := NewRamenOrder(
		WithGarlic(Extra),
		WithOilLevel(Extra),
	)
	fmt.Println(ramenOrder2.Call())

	// フードファイター
	ramenOrder4 := NewRamenOrder(
		WithGarlic(Ultimate),
		WithVegetables(Ultimate),
		WithOilLevel(Ultimate),
	)
	fmt.Println(ramenOrder4.Call())
}
