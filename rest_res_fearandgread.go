package cmcapi

// value
// integer [ 0 .. 100 ]
// The value of CMC Fear and Greed.
//
// When the value is closer to 0, the market is in Extreme Fear, and investors have over-sold irrationally.
//
// When the value is closer to 100, the market is in Extreme Greed, indicating a likely market correction.
//
// value_classification
// string
// The value classication of CMC Fear and Greed.
//
// 1 ≤ x < 20: Extreme Fear
// 20 ≤ x < 40: Fear
// 40 ≤ x < 60: Neutral
// 60 ≤ x < 80: Greed
// 80 ≤ x ≤ 100: Extreme Greed
//
// update_time
// string <date>
// Timestamp (ISO 8601) of the last time this record was updated.
type CmcRestApiFearAndGreedLatestRes struct {
	Value               int    `json:"value"`                //integer [ 0 .. 100 ] The value of CMC Fear and Greed.
	UpdateTime          string `json:"update_time"`          //string <date> Timestamp (ISO 8601) of the last time this record was updated.
	ValueClassification string `json:"value_classification"` //string The value classication of CMC Fear and Greed.
}
