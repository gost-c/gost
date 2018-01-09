// Copyright 2014-present Codehack. All rights reserved.
// For mobile and web development visit http://codehack.com
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package scrypto

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"math"
	"strconv"
	"strings"

	"golang.org/x/crypto/scrypt"
)

const (
	// SaltLen is the default length of the random salt.
	SaltLen = 32
	// DefaultN is the default value for Scrypt N.
	DefaultN = 16384
	// DefaultR is the default value for Scrypt r.
	DefaultR = 8
	// DefaultP is the default value for Scrypt p.
	DefaultP = 1
	// KeyLen is the length for the generated key.
	KeyLen = 32
)

// NewSalt returns a randomly generated salt of 'size' length.
// On failure, this returns nil and the error.
func NewSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}

// Hash hashes a password with optional scrypt values N, r, and p.
// password is a plaintext password.
// args corresponds to N, r, p values.
// This returns the hashed password using the format: $s0$VALUES$SALT$HASH
// On failure, empty string and error are returned.
func Hash(password string, args ...int) (string, error) {
	salt, err := NewSalt(SaltLen)
	if err != nil {
		return "", err
	}

	// set scrypt defaults
	N, r, p := DefaultN, DefaultR, DefaultP

	if args != nil {
		switch len(args) {
		case 3:
			p = args[2]
			fallthrough
		case 2:
			r = args[1]
			fallthrough
		case 1:
			N = args[0]
		}
	}

	hashed, err := scrypt.Key([]byte(password), salt, N, r, p, KeyLen)
	if err != nil {
		return "", err
	}

	values := int(math.Log2(float64(N)))<<16 | r<<8 | p

	hpass := "$s0$" +
		strconv.Itoa(values) + "$" +
		base64.RawStdEncoding.EncodeToString(salt) + "$" +
		base64.RawStdEncoding.EncodeToString(hashed)

	return hpass, nil
}

// Compare compares a plaintext password to a hashed password.
// password is a plaintext password.
// hpass is hashed password string in $s0$ format.
// returns true if passwords match, false otherwise.
func Compare(password, hpass string) bool {
	seg := strings.Split(hpass, "$")
	if len(seg) != 5 || seg[1] != "s0" {
		return false
	}

	values, _ := strconv.ParseInt(seg[2], 10, 64)
	salt, _ := base64.RawStdEncoding.DecodeString(seg[3])
	hashed1, _ := base64.RawStdEncoding.DecodeString(seg[4])

	N := int(math.Pow(2, float64(values>>16&0xffff)))
	r := int(values) >> 8 & 0xff
	p := int(values) & 0xff

	hashed2, _ := scrypt.Key([]byte(password), salt, N, r, p, KeyLen)

	len1 := len(hashed1)
	if len1 != len(hashed2) {
		return false
	}

	var comp int
	for i := 0; i < len1; i++ {
		comp |= int(hashed1[i] ^ hashed2[i])
	}

	return comp == 0
}
