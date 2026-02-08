package domain

type TopProduct struct {
	Name    string
	QtySold int
}

type Report struct {
	TotalRevenue      int
	TotalTransactions int
	TopProduct        TopProduct
}
