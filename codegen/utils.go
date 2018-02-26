package codegen

func bool2int8(b bool) int8 {
	var bits int8 = 0
	if b {
		bits = 1
	}
	return bits
}

