package slopeone

// Algorithm implements the Slope One collaborative filtering algorithm.
type Algorithm struct {
	// differentials stores the differential rating matrix
	differentials map[string]map[string]float64
	// frequencies stores the rating count matrix
	frequencies map[string]map[string]int
}

func New() *Algorithm {
	return &Algorithm{
		differentials: make(map[string]map[string]float64),
		frequencies:   make(map[string]map[string]int),
	}
}

// Clear resets both the differential rating matrix and frequency matrix.
func (a *Algorithm) Clear() *Algorithm {
	a.differentials = make(map[string]map[string]float64)
	a.frequencies = make(map[string]map[string]int)
	return a
}

func (a *Algorithm) Update(data []map[string]float64) *Algorithm {
	// Process each user's ratings
	for _, ratings := range data {
		for item1, rating1 := range ratings {
			// Initialize maps if they don't exist
			if a.frequencies[item1] == nil {
				a.frequencies[item1] = make(map[string]int)
			}
			if a.differentials[item1] == nil {
				a.differentials[item1] = make(map[string]float64)
			}

			// Compare with all other items in the user's ratings
			for item2, rating2 := range ratings {
				a.frequencies[item1][item2]++
				a.differentials[item1][item2] += rating1 - rating2
			}
		}
	}

	return a
}

// Predict generates rating predictions for items not yet rated by a user.
func (a *Algorithm) Predict(preferences map[string]float64) map[string]float64 {
	predictions := make(map[string]float64)
	frequencies := make(map[string]int)

	// Calculate weighted predictions based on differential matrix
	for item, rating := range preferences {
		for diffItem, diffRatings := range a.differentials {
			if a.frequencies[diffItem] != nil && a.frequencies[diffItem][item] > 0 {
				frequency := a.frequencies[diffItem][item]
				avgDiff := diffRatings[item] / float64(frequency)
				predictions[diffItem] += float64(frequency) * (avgDiff + rating)
				frequencies[diffItem] += frequency
			}
		}
	}

	// Calculate final predictions, excluding items already rated
	results := make(map[string]float64)
	for item, value := range predictions {
		if _, exists := preferences[item]; !exists && frequencies[item] > 0 {
			results[item] = value / float64(frequencies[item])
		}
	}

	return results
}
