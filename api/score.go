package api

type ScoredOffer struct {
	Offer Offer
	Score float64
}

func ScoreOffers(offers []Offer) []ScoredOffer {
	var scoredOffers = make([]ScoredOffer, 0, len(offers))
	for _, offer := range offers {
		score := 100.0
		scoredOffers = append(scoredOffers, ScoredOffer{Offer: offer, Score: score})
	}
	return scoredOffers
}
