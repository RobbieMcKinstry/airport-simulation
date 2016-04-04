package simulation

const (
	CommuterArrivalRate      = (40.0 / 60.0)
	InternationalArrivalMean = 75.0
	InternationalArrivalVar  = 50.0

	QueueingForCheckIn = iota
	PrintingBoardingPass
	CheckingBags
	MiscDelays
	QueueingForSecurity
	AtSecurity
	AtGate
	EmptyQueue
)

var (
	CommuterArrivalGen      = NewExpGenerator(CommuterArrivalRate)
	InternationalArrivalGen = NewNormalGenerator(InternationalArrivalMean, InternationalArrivalVar)
	BagCheckGen             = NewExpGenerator(1.0)
	BoardingPassGen         = NewExpGenerator(1.0 / 2.0)
	MiscGen                 = NewExpGenerator(1.0 / 3.0)
	CommuterBagGen          = NewGeoGenerator(0.80)
	InternationalBagGen     = NewGeoGenerator(0.60)
	BernoulliFirstGen       = NewBernGenerator(0.80)
	BernoulliCoachGen       = NewBernGenerator(0.85)
)
