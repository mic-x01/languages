package main

type DiscountedProduct struct {
	*Product `json:",omitempty"`
	Discount float64 `json:",string"`
}
