// Package dog for understanding Dog's Loyalty
package dog

import "testing"

// Years of dog package converting human years into dog years
func Years(year int) int {
	yearsdog := year * 7
	return yearsdog
}

// BenchmarkYears in Dog Package
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(i + 10)
	}
}
