// Package tempconv для конвертации температуры
package tempconv

import "fmt"

// Celcius temp
type Celcius float32

// Kelvin temp
type Kelvin float32

func (c Celcius) String() string {
	return fmt.Sprintf("%fºC", c)
}

func (k Kelvin) String() string {
	return fmt.Sprintf("%fK", k)
}
