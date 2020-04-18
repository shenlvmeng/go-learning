package tempconv

import "fmt"

type Celsius float64
type Fahreheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
func (f Fahreheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
func (k Kelvin) String() string {
	return fmt.Sprintf("%g°K", k)
}

func CTOF(c Celsius) Fahreheit {
	return Fahreheit(c*9/5 + 32)
}
func FTOC(f Fahreheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
func CTOK(c Celsius) Kelvin {
	return Kelvin(c + BoilingC)
}
func KTOC(k Kelvin) Celsius {
	return Celsius(k - BoilingC)
}
func KTOF(k Kelvin) Fahreheit {
	return Fahreheit((k-BoilingC)*9/5 + 32)
}
