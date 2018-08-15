package maths

func Mix(x, y, a float64) float64 {
	return x * (1.0 - a) + y * a
}

func MixFloat32(x, y, a float32) float32 {
	return float32(Mix(float64(x), float64(y), float64(a)))
}
