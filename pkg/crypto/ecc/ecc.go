/*
	Sean Visser
	Implementation of Elliptic Curve

	http://www.infosecwriters.com/text_resources/pdf/Elliptic_Curve_AnnopMS.pdf
*/

package ecc

import (
	"math/big"
)

type EC interface {
	Params() *ECCParams //Return the parameters of the curve
}

type ECCParams struct {
	P      *big.Int // Prime number for finite field Fp
	A      *big.Int // Defining parameter of curve
	B      *big.Int // Defining parameter of curve
	Gx, Gy *big.Int // Generator point
	N      *big.Int // Order of elliptic curve
	H      *big.Int // Cofactor where H = #E(F2^M)/N
	Name   string   // Cannonical name of curve
}

func (curve *ECCParams) Params() *ECCParams {
	// Function on struct ECCParams to return the parameters of the curve
	return curve
}
