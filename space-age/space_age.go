package space

type Planet string

const EarthOrbitalPeriod = 31557600 // sec of year

var OrbitalPeriods = map[Planet]float64{
	"Earth": 1,
	"Mercury": 0.2408467,
	"Venus": 0.61519726,
	"Mars": 1.8808158,
	"Jupiter": 11.862615,
	"Saturn": 29.447498,
	"Uranus": 84.016846,
	"Neptune": 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	r, ok := OrbitalPeriods[planet]
	if !ok {
		return 0
	}

	return seconds / (r * EarthOrbitalPeriod)
}
