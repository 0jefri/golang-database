package dto

type MonthlyOrder struct {
	Month      string `db:"month"`
	TotalOrder int    `db:"total_order"`
}
