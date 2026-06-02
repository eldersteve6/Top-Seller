package main

func topSeller(sellers []Seller) string {
	bestScore := -1
	bestName := ""

	for _, seller := range sellers {
		if !seller.Active {
			continue
		}

		score := seller.Sales * seller.Rating

		if score > bestScore {
			bestScore = score
			bestName = seller.Name
		}
	}

	return bestName
}
