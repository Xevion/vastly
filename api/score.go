package api

import (
	"fmt"
	"math"

	"go.uber.org/zap"
)

type ScoreReason struct {
	Reason       string
	Offset       float64
	Value        interface{}
	IsMultiplier bool
}

type ScoredOffer struct {
	Offer   Offer
	Score   float64
	Reasons []ScoreReason
}

var (
	sugar *zap.SugaredLogger
)

func init() {
	logger, err := zap.NewDevelopment()
	sugar = logger.Sugar()
	if err != nil {
		panic(err)
	}
}

func AddReason(reasonList []ScoreReason, reason string, offset float64, multiplier bool, value interface{}) []ScoreReason {
	return append(reasonList, ScoreReason{Reason: reason, Offset: offset, IsMultiplier: multiplier, Value: fmt.Sprintf("%v", value)})
}

func ScoreOffers(offers []Offer) []ScoredOffer {
	var scoredOffers = make([]ScoredOffer, 0, len(offers))

	for _, offer := range offers {
		reasons := make([]ScoreReason, 0, 16)

		// Judge cost
		targetCost := 0.42                                    // $/hour
		costMultiplier := targetCost / offer.Search.TotalHour // e.x $0.42 / $0.50 = x0.84
		// sugar.Debugw("Cost", "offer", offer.ID, "costMultiplier", costMultiplier, "targetCost", targetCost, "actualCost", offer.Search.TotalHour)
		if math.Abs(costMultiplier-1.0) > 0.05 {
			reasons = AddReason(reasons, "Cost Multiplier", costMultiplier, true, offer.Search.TotalHour)
		}

		// Judge DLPerf
		targetDLPerf := 85.0
		dlPerfMultiplier := offer.DLPerf / targetDLPerf // e.x 100 / 85 = x1.18
		if math.Abs(dlPerfMultiplier-1.0) > 0.03 {
			if dlPerfMultiplier > 1.0 {
				dlPerfMultiplier = math.Sqrt(dlPerfMultiplier)
			} else {
				dlPerfMultiplier = math.Pow(dlPerfMultiplier, 2.0)
			}
			reasons = AddReason(reasons, "DLPerf Multiplier", dlPerfMultiplier, true, offer.DLPerf)
		}

		// Judge Internet Download Speed
		if offer.InetDown < 150.0 {
			reasons = AddReason(reasons, "Very Poor Internet Download Speed", -7.0, false, offer.InetDown)
		} else if offer.InetDown < 300.0 {
			reasons = AddReason(reasons, "Poor Internet Download Speed", -3.0, false, offer.InetDown)
		} else if offer.InetDown < 500.0 {
			reasons = AddReason(reasons, "Decent Internet Download Speed", 1.0, false, offer.InetDown)
		} else if offer.InetDown < 1000.0 {
			reasons = AddReason(reasons, "Good Internet Download Speed", 3.0, false, offer.InetDown)
		} else if offer.InetDown >= 2000.0 {
			reasons = AddReason(reasons, "Great Internet Download Speed", 4.0, false, offer.InetDown)
		}

		// Judge Internet Upload Speed
		if offer.InetUp < 50.0 {
			reasons = AddReason(reasons, "Extremely Poor Internet Upload Speed", -9.0, false, offer.InetUp)
		} else if offer.InetUp < 100.0 {
			reasons = AddReason(reasons, "Very Poor Internet Upload Speed", -5.0, false, offer.InetUp)
		} else if offer.InetUp < 200.0 {
			reasons = AddReason(reasons, "Poor Internet Upload Speed", -3.0, false, offer.InetUp)
		} else if offer.InetUp < 400.0 {
			reasons = AddReason(reasons, "Decent Internet Upload Speed", 1.0, false, offer.InetUp)
		} else if offer.InetUp >= 800.0 {
			reasons = AddReason(reasons, "Great Internet Upload Speed", 3.0, false, offer.InetUp)
		} else if offer.InetUp >= 1000.0 {
			reasons = AddReason(reasons, "Amazing Internet Upload Speed", 4.0, false, offer.InetUp)
		}

		// Judge verification
		if offer.Verification == "verified" {
			reasons = AddReason(reasons, "Verified", 2.0, false, offer.Verification)
		} else {
			reasons = AddReason(reasons, "Not Verified", -5.0, false, offer.Verification)
		}

		// Judge direct port count
		if offer.DirectPortCount < 8 {
			reasons = AddReason(reasons, "Very low direct port count", -2.0, false, offer.DirectPortCount)
		} else if offer.DirectPortCount >= 32 {
			reasons = AddReason(reasons, "Decent port count", 0.5, false, offer.DirectPortCount)
		} else if offer.DirectPortCount >= 100 {
			reasons = AddReason(reasons, "High port count", 1.0, false, offer.DirectPortCount)
		}

		// Judge CPU memory
		if offer.CPURam < 16*1024 {
			reasons = AddReason(reasons, "Low CPU memory", -4.0, false, offer.CPURam)
		} else if offer.CPURam >= 31*1024 {
			reasons = AddReason(reasons, "High CPU memory", 2.0, false, offer.CPURam)
		} else if offer.CPURam >= 63*1024 {
			reasons = AddReason(reasons, "Very High CPU memory", 3.0, false, offer.CPURam)
		}

		// Judge GPU count
		if offer.NumGPUs < 1 {
			reasons = AddReason(reasons, "No GPUs", -20.0, false, offer.NumGPUs)
		} else if offer.NumGPUs == 2 {
			reasons = AddReason(reasons, "Dual GPU", 1.0, false, offer.NumGPUs)
		} else if offer.NumGPUs >= 3 {
			reasons = AddReason(reasons, "Multi GPU", 2.0, false, offer.NumGPUs)
		}

		// Judge CUDA version
		if offer.CudaMaxGood < 11.0 {
			reasons = AddReason(reasons, "CUDA version very outdated", -7.0, false, offer.CudaMaxGood)
		} else if offer.CudaMaxGood < 12.0 {
			reasons = AddReason(reasons, "CUDA version outdated", -5.0, false, offer.CudaMaxGood)
		} else if offer.CudaMaxGood >= 12.5 {
			reasons = AddReason(reasons, "CUDA version high", 5.0, false, offer.CudaMaxGood)
		} else if offer.CudaMaxGood >= 12.0 {
			reasons = AddReason(reasons, "CUDA version decent", 1.5, false, offer.CudaMaxGood)
		}

		// Judge effective core count
		if offer.CPUCoresEffective < 4.0 {
			reasons = AddReason(reasons, "Low core count", -5.0, false, offer.CPUCoresEffective)
		} else if offer.CPUCoresEffective >= 8.0 {
			reasons = AddReason(reasons, "High core count", 5.0, false, offer.CPUCoresEffective)
		} else if offer.CPUCoresEffective >= 6.0 {
			reasons = AddReason(reasons, "Decent core count", 2.0, false, offer.CPUCoresEffective)
		}

		// Judge disk space available
		if offer.DiskSpace < 100.0 {
			reasons = AddReason(reasons, "Low disk space", -5.0, false, offer.DiskSpace)
		} else if offer.DiskSpace >= 750.0 {
			reasons = AddReason(reasons, "Reasonable disk space", 1.0, false, offer.DiskSpace)
		} else if offer.DiskSpace >= 250.0 {
			reasons = AddReason(reasons, "Concerning disk space", -1.0, false, offer.DiskSpace)
		} else if offer.DiskSpace >= 2500.0 {
			reasons = AddReason(reasons, "High disk space", 2.0, false, offer.DiskSpace)
		}

		// Judge GPU architecture
		if offer.GPUArch == "nvidia" {
			reasons = AddReason(reasons, "Nvidia Preference", 1.0, false, offer.GPUArch)
		} else {
			reasons = AddReason(reasons, "Unknown/Incompatible GPU Architecture", -10.0, false, offer.GPUArch)
		}

		// Judge reliability
		if offer.Reliability2 < 0.98 {
			reasons = AddReason(reasons, "Low reliability", -5.0, false, offer.Reliability2)
		} else if offer.Reliability2 >= 0.999 {
			reasons = AddReason(reasons, "Very high reliability", 5.0, false, offer.Reliability2)
		} else if offer.Reliability2 >= 0.995 {
			reasons = AddReason(reasons, "High reliability", 2.0, false, offer.Reliability2)
		} else if offer.Reliability2 >= 0.99 {
			reasons = AddReason(reasons, "Decent reliability", 1.0, false, offer.Reliability2)
		}

		// Calculate base score
		score := 0.0
		for _, reason := range reasons {
			if !reason.IsMultiplier {
				score += reason.Offset
			}
		}

		// Apply multipliers
		multiplier := 1.0
		for _, reason := range reasons {
			if reason.IsMultiplier {
				multiplier *= reason.Offset
			}
		}
		newScore := score * multiplier
		// sugar.Infow("Multiplier Applied", "offer", offer.ID, "baseScore", score, "score", newScore, "multiplier", multiplier)
		score = newScore

		scoredOffers = append(scoredOffers, ScoredOffer{Offer: offer, Score: score, Reasons: reasons})
	}
	return scoredOffers
}
