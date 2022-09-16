package ffc

type InsightBuilder struct {
}

func NewInsightBuilder() *InsightBuilder {
	return &InsightBuilder{}
}

func (i *InsightBuilder) CreateInsightProcessor(context Context) InsightProcessor {
	return &Insight{}
}
