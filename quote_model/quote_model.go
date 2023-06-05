package quote_model

type Quote struct {
	author   string
	category string
	quote    string
}

func NewQuote(author string, category string, quote string) Quote {
	q := Quote{author: author, category: category, quote: quote}
	return q
}

func GetQuoteInfo(q Quote) (string, string, string) {
	return q.author, q.category, q.quote
}
