// Code generated by github.com/karalabe/ssz. DO NOT EDIT.

package consensus_spec_tests

import "github.com/karalabe/ssz"

// SizeSSZ returns the total size of the static ssz object.
func (obj *SigningRoot) SizeSSZ() uint32 {
	return 32 + 8
}

// DefineSSZ defines how an object is encoded/decoded.
func (obj *SigningRoot) DefineSSZ(codec *ssz.Codec) {
	ssz.DefineStaticBytes(codec, obj.ObjectRoot[:]) // Field  (0) - ObjectRoot - 32 bytes
	ssz.DefineStaticBytes(codec, obj.Domain[:])     // Field  (1) -     Domain -  8 bytes
}