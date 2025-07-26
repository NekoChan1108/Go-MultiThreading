package main

import "math"

// Vector2D 向量类型
/**
 *(x,y)表示向量
 */
type Vector2D struct {
	x, y float64
}

// Add 向量相加
func (v Vector2D) Add(v1 Vector2D) Vector2D {
	return Vector2D{v.x + v1.x, v.y + v1.y}
}

// Subtract Sub 向量相减
func (v Vector2D) Subtract(v1 Vector2D) Vector2D {
	return Vector2D{v.x - v1.x, v.y - v1.y}
}

// Multiply 向量相乘
func (v Vector2D) Multiply(v1 Vector2D) Vector2D {
	return Vector2D{v.x * v1.x, v.y * v1.y}
}

// AddVal 增加指定值
func (v Vector2D) AddVal(val float64) Vector2D {
	return Vector2D{v.x + val, v.y + val}
}

// MultiplyVal 扩大指定倍
func (v Vector2D) MultiplyVal(val float64) Vector2D {
	return Vector2D{v.x * val, v.y * val}
}

// DivideVal 缩小指定倍
func (v Vector2D) DivideVal(val float64) Vector2D {
	return Vector2D{v.x / val, v.y / val}
}

// Limit 限制范围在lower到upper之间
func (v Vector2D) Limit(lower, upper float64) Vector2D {
	return Vector2D{math.Min(math.Max(v.x, lower), upper), math.Min(math.Max(v.y, lower), upper)}
}

// Distance 计算两个向量的距离
func (v Vector2D) Distance(v2 Vector2D) float64 {
	return math.Sqrt(math.Pow(v.x-v2.x, 2) + math.Pow(v.y-v2.y, 2))
}
