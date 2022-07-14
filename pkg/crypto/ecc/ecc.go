/*
	Sean Visser
	Implementation of Elliptic Curve

	http://www.infosecwriters.com/text_resources/pdf/Elliptic_Curve_AnnopMS.pdf

	Afine curves to Jacobian curves and back to Afine curves to perform calculations are faster than calculating on afine curve,
	because the use of / operator has been removed https://medium.com/computronium/gradient-based-optimizations-jacobians-jababians-hessians-b7cbe62d662d
*/
package ecc

import (
	"math/big"
)

type EC interface {
	Params() *ECCParams //Return the parameters of the curve
	Add(x1, y1, x2, y2s *big.Int) (x, y *big.Int)
	//Double (x1, y1 *big.Int) (x, y *big.Int)
	//ScalarMult(x1, y1 *big.Int, k []byte) (x, y *big.Int) // ScalarMult returns k*(Bx, By) where K is a number in big-endian form
	//ScalarBaseMult(k []byte) (x, y *big.Int) //ScalarBaseMult returns k*G where k is int in big-endian form and G is basepoint
	//IsOnCurve(x, y *big.Int) bool
	//
}

type ECCParams struct {
	P      *big.Int // Prime number for finite field Fp
	A      *big.Int // Defining parameter of curve
	B      *big.Int // Defining parameter of curve, constant
	Gx, Gy *big.Int // Generator point
	N      *big.Int // Order of elliptic curve base point
	Name   string   // Cannonical name of curve
}

func (curve *ECCParams) Params() *ECCParams {
	// Function on struct ECCParams to return the parameters of the curve
	return curve
}

func zForAffine(x, y *big.Int) *big.Int {
	var z *big.Int = new(big.Int)       //Creates new big.Int Z
	if x.Sign() != 0 || y.Sign() != 0 { //Checks if y OR x point is larger than 0
		z.SetInt64(1) // Set Z to one
	}
	return z
}

func (curve *ECCParams) affineFromJacobian(x, y, z *big.Int) (xl, yl *big.Int) {
	if z.Sign() == 0 {
		return new(big.Int), new(big.Int) // Als z 0 is, geef dan 0,0 terug. Point to infinity qq
	}

	xl, yl = new(big.Int), new(big.Int)

	// Z * X = 1 (mod P)
	var zinv *big.Int = new(big.Int).ModInverse(z, curve.P)

	// X*X
	var zinvsq *big.Int = new(big.Int).Mul(zinv, zinv)

	// XL = x * zinvsq (mod P)
	xl.Mul(x, zinvsq)
	xl = curve.wrapAroundP(xl)

	// YL = Y * ZINVSQ (mod P)
	yl.Mul(y, zinvsq)
	yl = curve.wrapAroundP(yl)

	return xl, yl
}

func (curve *ECCParams) wrapAroundP(n *big.Int) *big.Int {
	/*
		Function to wrap a number around prime number for finite field Fp (p)
		Every point is wrapped around P

		BAD OPTIMALIZATION FUNCTION SINCE WE CREATE NEW BIG INT EVERY TIME SOMETHING NEEDS WRAPPING
		Better solution: use .Mod method on *big.Int types
	*/

	var nw *big.Int = new(big.Int).Mod(n, curve.P) // nw = n % P
	return nw
}

/*
	ECC Method 1: Addition

	Point addition is the addition of two points (J & K) on an elliptic curve to obtain another point (L) on the same curve

	Point J = xj, yj
	Point K = xk, yk
	Point L = xl, yl

	L = J + K
*/

/*
	In order to efficiently add points, we need convert it to Jacobian curve first(get Z for point (xj, yj) and (xk, yk),
	then we can use jacobian method for adding.

	Optimalisation made: No longer requires the use of / operator to calculate the slope :D
*/
func (curve *ECCParams) Add(xj, yj, xk, yk *big.Int) (*big.Int, *big.Int) {
	/*
		Adding to points on curve

		Optimalization has been made by converting points (x,y) to jacobian-form (x,y,z)
		For actual addition algorithm, see addJacobianPoints
	*/
	var zj *big.Int = zForAffine(xj, yj)                                             //Get Z-value for Xj, Yj
	var zk *big.Int = zForAffine(xk, yk)                                             //Get Z-value for Xk, Yk
	return curve.affineFromJacobian(curve.addJacobianPoints(xj, yj, zj, xk, yj, zk)) //affine from jacobian
}

func (curve *ECCParams) addJacobianPoints(xj, yj, zj, xk, yk, zk *big.Int) (x, y, z *big.Int) {
	/*
		For more information about jacobian addition, see https://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian-3.html#addition-add-2007-bl
	*/

	xl, yl, zl := new(big.Int), new(big.Int), new(big.Int)

	var zjzj *big.Int = new(big.Int).Mul(zj, zj) //Z1Z1 = Z1**2
	zjzj = curve.wrapAroundP(zjzj)               // All points are wrapped around P-parameter from curve

	var zkzk *big.Int = new(big.Int).Mul(zk, zk) //Z2Z2 = Z2**2
	zkzk = curve.wrapAroundP(zkzk)

	var uj *big.Int = new(big.Int).Mul(xj, zkzk) //U1 = x1 * Z2Z2
	uj = curve.wrapAroundP(uj)

	var uk *big.Int = new(big.Int).Mul(xk, zjzj) //U2 = X2 * Z1Z1
	uk = curve.wrapAroundP(uk)

	var h *big.Int = new(big.Int).Sub(uk, uj) // H = U2 - U1. No need to wrap round P, h is not a point

	var xEqual bool = h.Sign() == 0
	if h.Sign() == -1 { //if h is negative, add P to h
		h.Add(h, curve.P)
	}

	// SJ = YJ * ZK * ZKZK
	var sj *big.Int = new(big.Int).Mul(yj, zk)
	sj.Mul(sj, zkzk)
	sj = curve.wrapAroundP(sj)

	// SK = YK * ZJ * ZJZJ
	var sk *big.Int = new(big.Int).Mul(yk, zj)
	sk.Mul(sk, zjzj)
	sk = curve.wrapAroundP(sk)

	// I = (2*H)^2
	var i *big.Int = new(big.Int).Lsh(h, 1) //Left shift to *2 because Mul takes two *big.Int
	i.Mul(i, i)

	//J = I * H
	var j *big.Int = new(big.Int).Mul(i, h)

	// R = 2*(S2-S1)
	var r *big.Int = new(big.Int).Sub(sk, sj)
	r.Lsh(r, 1) //Left shift to *2 because Mul takes two *big.Int
	if r.Sign() == -1 {
		r.Add(r, curve.P) // if r is negative, add P to r
	}

	var yEqual bool = r.Sign() == 0

	if xEqual && yEqual {
		//If r and h are both 0, use point doubling instead
		// return curve.doubleJacobianPoints(x, y, z) TODO

	}

	var v *big.Int = new(big.Int).Mul(uj, i) // V = U1 * I

	// X3 = r**2 - J - 2*v
	xl.Set(r)
	xl.Mul(r, r)
	xl.Sub(xl, j)
	xl.Sub(xl, v)
	xl.Sub(xl, v)
	xl = curve.wrapAroundP(xl)

	// Y3 = R * (V-X3) - 2 * SJ * J
	yl.Set(r)
	v.Sub(v, xl)
	yl.Mul(yl, v)
	sj.Mul(sj, j)
	sj.Lsh(sj, 1)
	yl.Sub(yl, sj)
	yl = curve.wrapAroundP(yl)

	// Z3 = ((Z1+Z2)^2-Z1Z1-Z2Z2)*H
	zl.Add(zj, zk)
	zl.Mul(zl, zl)
	zl.Sub(zl, zjzj)
	zl.Sub(zl, zkzk)
	zl.Mul(zl, h)

	return xl, yl, zl

}
