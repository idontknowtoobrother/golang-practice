package calculator

func Add(a, b int) int {
	return a + b
}

func PublicMultiply(a, b int) int { // This is a public function that call a private function
	return multiply(a, b) // This is a private function
}
