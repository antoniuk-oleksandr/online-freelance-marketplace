package model

type OrderDeliveryCancellation struct {
	CancellationReason string `json:"cancellationReason" db:"cancellationReason"`
	CancelledAt        string `json:"cancelledAt" db:"cancelledAt"`
}
