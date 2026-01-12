package payment

type Status string

const (
	StatusInitiated Status = "INITIATED"
	StatusSuccess   Status = "SUCCESS"
	StatusFailed    Status = "FAILED"
)

type Payment struct {
	ID      uint `gorm:"primaryKey"`
	OrderID uint `gorm:"uniqueIndex"`
	Amount  int64
	Method  string
	Status  string
}
