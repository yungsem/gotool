package maths

import "math"

// Round 对 v 进行四舍五入，scale 是要保留的小数位数
func Round(v float64, scale int) float64 {
	p := math.Pow10(scale)
	return math.Floor(v*p+0.5) / p
}
