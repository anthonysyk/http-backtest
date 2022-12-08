package example

type LogElement struct {
	URL string `json:"url"`
}

type Backtest struct{}

func (b Backtest) GetURLs() []string {

	return nil
}
