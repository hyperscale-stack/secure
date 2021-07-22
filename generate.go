// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package secure

import (
	"crypto/rand"
	"fmt"
)

// GenerateRandomBytes returns securely generated random bytes.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)

	if _, err := rand.Read(b); err != nil {
		return nil, fmt.Errorf("read rand failed: %w", err)
	}

	return b, nil
}

// GenerateRandomString returns a securely generated random string.
func GenerateRandomString(length int) (string, error) {
	result := make([]byte, length)

	if _, err := rand.Read(result); err != nil {
		return "", fmt.Errorf("read rand failed: %w", err)
	}

	for i := 0; i < length; i++ {
		result[i] &= 0x7F

		//  48 : 0
		//  57 : 9
		//  65 : A
		//  90 : Z
		//  97 : a
		// 122 : z
		// [0-9A-Za-z]
		for (result[i] < 48 || result[i] > 57) && (result[i] < 65 || result[i] > 90) && (result[i] < 97 || result[i] > 122) {
			if _, err := rand.Read(result[i : i+1]); err != nil {
				return "", fmt.Errorf("read rand failed: %w", err)
			}

			result[i] &= 0x7F
		}
	}

	return string(result), nil
}
