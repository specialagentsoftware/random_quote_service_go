package quote_model

type Quote struct {
	author   string
	category string
	quote    string
}

func NewQuote(author string, category string, quote string) Quote {
	/*
		this is another poor mans constructor that takes three string arguments and returns a Quote object
	*/
	q := Quote{author: author, category: category, quote: quote}
	return q
}

func GetQuoteInfo(q Quote) (string, string, string) {
	/*
		GetQuoteInfo takes in a quote and returns data related to that quote.
	*/
	return q.author, q.category, q.quote
}
