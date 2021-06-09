package refresh

import (
	"github.com/taurusgroup/cmp-ecdsa/pkg/hash"
	"github.com/taurusgroup/cmp-ecdsa/pkg/math/curve"
	"github.com/taurusgroup/cmp-ecdsa/pkg/math/polynomial"

	//"github.com/taurusgroup/cmp-ecdsa/pkg/pedersen"
	"github.com/taurusgroup/cmp-ecdsa/pkg/round"
)

// localParty is the state we store for a remote party.
// The messages are embedded to make access to the attributes easier.
type localParty struct {
	*round.Party

	// A are the schnorr commitments to the polynomial in the exponent
	A []*curve.Point // Aⱼ

	commitment hash.Commitment // H(msg2, decommitment)

	// rho = ρⱼ
	// if keygen, then this is RIDⱼ
	rho []byte

	// polyExp = Fⱼ(X) = fⱼ(X)•G
	polyExp *polynomial.Exponent

	// sharesReceived = xʲᵢ is the share received from party j
	shareReceived *curve.Scalar
}