// Code generated by github.com/karalabe/ssz. DO NOT EDIT.

package consensus_spec_tests

import "github.com/karalabe/ssz"

// Cached static size computed on package init.
var staticSizeCacheAttestationDataVariation1 = ssz.PrecomputeStaticSizeCache((*AttestationDataVariation1)(nil))

// SizeSSZ returns the total size of the static ssz object.
func (obj *AttestationDataVariation1) SizeSSZ(sizer *ssz.Sizer) (size uint32) {
	if fork := int(sizer.Fork()); fork < len(staticSizeCacheAttestationDataVariation1) {
		return staticSizeCacheAttestationDataVariation1[fork]
	}
	if sizer.Fork() >= ssz.ForkFuture {
		size += 8
	}
	size += 8 + 8 + 32 + (*Checkpoint)(nil).SizeSSZ(sizer) + (*Checkpoint)(nil).SizeSSZ(sizer)
	return size
}

// DefineSSZ defines how an object is encoded/decoded.
func (obj *AttestationDataVariation1) DefineSSZ(codec *ssz.Codec) {
	ssz.DefineUint64PointerOnFork(codec, &obj.Future, ssz.ForkFilter{Added: ssz.ForkFuture}) // Field  (0) -          Future -  8 bytes
	ssz.DefineUint64(codec, &obj.Slot)                                                       // Field  (1) -            Slot -  8 bytes
	ssz.DefineUint64(codec, &obj.Index)                                                      // Field  (2) -           Index -  8 bytes
	ssz.DefineStaticBytes(codec, &obj.BeaconBlockHash)                                       // Field  (3) - BeaconBlockHash - 32 bytes
	ssz.DefineStaticObject(codec, &obj.Source)                                               // Field  (4) -          Source -  ? bytes (Checkpoint)
	ssz.DefineStaticObject(codec, &obj.Target)                                               // Field  (5) -          Target -  ? bytes (Checkpoint)
}