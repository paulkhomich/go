package tempconv

// KToC преобразует температуру в Кельвинах в температуру в Цельсиях
func KToC (k Kelvin) Celcius {
	return Celcius(k - 273.15);
}

// CToK преобразует температуру в Цельсиях в температуру в Кельвинах
func CToK (c Celcius) Kelvin {
	return Kelvin(c + 273.15);
}