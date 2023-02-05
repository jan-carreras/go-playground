package space

const secondsOnEarth = 31557600

func secondsToEarthYears(seconds float64) float64 {
	return seconds / secondsOnEarth
}

type Planet string

func Age(seconds float64, planet Planet) float64 {
	ratio := planet.Ratio()
	if ratio == 0 {
		return -1
	}

	return secondsToEarthYears(seconds) / ratio
}

func (p Planet) Ratio() float64 {
	r := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	return r[p]
}
