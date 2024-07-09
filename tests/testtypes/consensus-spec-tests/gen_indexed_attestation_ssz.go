// Code generated by github.com/karalabe/ssz. DO NOT EDIT.

package consensus_spec_tests

import "github.com/karalabe/ssz"

// Cached static size computed on package init.
var staticSizeCacheIndexedAttestation = 4 + (*AttestationData)(nil).SizeSSZ() + 96

// SizeSSZ returns either the static size of the object if fixed == true, or
// the total size otherwise.
func (obj *IndexedAttestation) SizeSSZ(fixed bool) uint32 {
	var size = uint32(staticSizeCacheIndexedAttestation)
	if fixed {
		return size
	}
	size += ssz.SizeSliceOfUint64s(obj.AttestationIndices)

	return size
}

// DefineSSZ defines how an object is encoded/decoded.
func (obj *IndexedAttestation) DefineSSZ(codec *ssz.Codec) {
	// Define the static data (fields and dynamic offsets)
	ssz.DefineSliceOfUint64sOffset(codec, &obj.AttestationIndices) // Offset (0) - AttestationIndices -  4 bytes
	ssz.DefineStaticObject(codec, &obj.Data)                       // Field  (1) -               Data -  ? bytes (AttestationData)
	ssz.DefineStaticBytes(codec, obj.Signature[:])                 // Field  (2) -          Signature - 96 bytes

	// Define the dynamic data (fields)
	ssz.DefineSliceOfUint64sContent(codec, &obj.AttestationIndices, 2048) // Field  (0) - AttestationIndices - ? bytes
}
