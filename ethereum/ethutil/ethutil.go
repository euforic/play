// package ethutil is a general purpose package for ethereum helper utils
package ethutil

import (
	"math"
	"math/big"
	"strings"
)

type unit int

// Ethereum unit types
const (
	Wei  unit = 0
	Kwei unit = 3
	Ada
	Femtoether
	Mwei unit = 6
	Babbage
	Picoether
	Gwei unit = 9
	Shannon
	Nanoether
	Nano
	Microether unit = 12
	Micro
	Szabo
	Milliether unit = 15
	Milli
	Finney
	Ether  unit = 18
	Kether unit = 21
	Grand
	Einstein
	Mether unit = 24
	Gether unit = 27
	Tether unit = 30
)

func (u unit) String() string {
	v, ok := unitName[u]
	if !ok {
		return ""
	}
	return v
}

var unitName = map[unit]string{
	Wei:    "wei",
	Kwei:   "kwei",
	Mwei:   "mwei",
	Gwei:   "gwei",
	Szabo:  "szabo",
	Finney: "finney",
	Ether:  "ether",
}

// EthToWei converts eth value to wei
func EthToWei(v string) string {
	return Convert(v, Ether, Wei).FloatString(0)
}

// WeiToEth converts wei to eth
func WeiToEth(v string) string {
	eth := Convert(v, Wei, Ether).FloatString(18)
	return strings.TrimRight(eth, "0")
}

// Convert values from one type to another
func Convert(w string, from unit, to unit) *big.Rat {
	v, ok := new(big.Rat).SetString(w)
	if !ok {
		return nil
	}

	fromUnit := new(big.Int).SetInt64(int64(math.Pow10(int(from))))
	toUnit := new(big.Int).SetInt64(int64(math.Pow10(int(to))))

	return v.Mul(v, new(big.Rat).SetFrac(fromUnit, toUnit))
}
