package enum

type Status int

const  (
	Incomplete = iota + 1
	InProgress
	Completed
	Cancelled
	Pending
	Failed
	AwaitingAcceptance
)