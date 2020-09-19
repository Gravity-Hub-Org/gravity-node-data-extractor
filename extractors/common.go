package extractors

// swagger:model
type RawData = byte

type IExtractor interface {
	DataFeedTag() string
	Description() string
	// raw and formated data types
	// first arg should represent type model, second one primitive
	Data() (interface{}, interface{})
	Info() *ExtractorInfo
	// extractData(params interface{}) []RawData
	// mapData(extractedData []RawData) interface{}
}

// swagger:model
type ExtractorInfo struct {
	Description string `json:"description"`
	DataFeedTag string `json:"datafeedtag"`
}

type ExtractorEnumeration = string
type ExtractorEnumerator struct {
	IBportWavesEth, LUportWavesEth, IBportEthWaves, LUportEthWaves ExtractorEnumeration
}