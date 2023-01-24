package entities

type Sent struct {
	Lang string
	Sent string
}

type Pair struct {
	Sent_a Sent
	Sent_b Sent
	Type   string
}
