// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package secure

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomBytes(t *testing.T) {
	for _, size := range []int{5, 7, 10, 15, 25, 32, 64, 96, 128, 256, 512, 1024} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			for i := 0; i < 100; i++ {
				b, err := GenerateRandomBytes(size)
				assert.NoError(t, err)
				assert.Equal(t, size, len(b))
			}
		})
	}
}

func BenchmarkGenerateRandomBytes64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomBytes(64)
	}
}

func BenchmarkGenerateRandomBytes128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomBytes(128)
	}
}

func TestGenerateSecureString(t *testing.T) {
	for _, size := range []int{5, 7, 10, 15, 25, 32, 64, 96, 128, 256, 512, 1024} {
		t.Run(fmt.Sprint(size), func(t *testing.T) {
			for i := 0; i < 100; i++ {
				s, err := GenerateRandomString(size)
				assert.NoError(t, err)
				assert.Equal(t, size, len(s))
			}
		})
	}
}

func BenchmarkGenerateRandomString64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomString(64)
	}
}

func BenchmarkGenerateRandomString128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomString(128)
	}
}
