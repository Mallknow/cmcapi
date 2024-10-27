package cmcapi

type CmcRestApiFearAndGreedLatestRes struct {
	Value               int    `json:"value"`                //integer [ 0 .. 100 ] The value of CMC Fear and Greed.
	UpdateTime          string `json:"update_time"`          //string <date> Timestamp (ISO 8601) of the last time this record was updated.
	ValueClassification string `json:"value_classification"` //string The value classication of CMC Fear and Greed.
}
