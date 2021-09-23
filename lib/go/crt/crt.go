package crt

// START HERE
func crtsafemod(x, m int) int {
	x %= m
	if x < 0 {
		x += m
	}
	return x
}
func crtinvgcd(a, b int) (int, int) {
	a = crtsafemod(a, b)
	if a == 0 {
		return b, 0
	}
	s, t, m0, m1 := b, a, 0, 1
	for t != 0 {
		u := s / t
		s -= t * u
		m0 -= m1 * u
		s, t, m0, m1 = t, s, m1, m0
	}
	if m0 < 0 {
		m0 += b / s
	}
	return s, m0
}
func crt(r, m []int) (int, int) {
	if len(r) != len(m) {
		panic("Mismatched length in crt")
	}
	for _, mm := range m {
		if mm <= 0 {
			panic("CRT error -- non-positive modulus")
		}
	}
	n, r0, m0 := len(r), 0, 1
	for i := 0; i < n; i++ {
		r1, m1 := crtsafemod(r[i], m[i]), m[i]
		if m0 < m1 {
			r0, r1, m0, m1 = r1, r0, m1, m0
		}
		if m0%m1 == 0 {
			if r0%m1 != r1 {
				return 0, 0
			}
			continue
		}
		g, im := crtinvgcd(m0, m1)
		u1 := m1 / g
		if (r1-r0)%g != 0 {
			return 0, 0
		}
		x := (r1 - r0) / g % u1 * im % u1
		r0 += x * m0
		m0 *= u1
		if r0 < 0 {
			r0 += m0
		}
	}
	return r0, m0
}
