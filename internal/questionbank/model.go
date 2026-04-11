package questionbank

type Dimension string

const (
	DimensionEI Dimension = "EI"
	DimensionSN Dimension = "SN"
	DimensionTF Dimension = "TF"
	DimensionJP Dimension = "JP"
)

type Strength string

const (
	StrengthStrongA   Strength = "strong_a"
	StrengthModerateA Strength = "moderate_a"
	StrengthSlightA   Strength = "slight_a"
	StrengthSlightB   Strength = "slight_b"
	StrengthModerateB Strength = "moderate_b"
	StrengthStrongB   Strength = "strong_b"
)

type ThresholdRange [2]int

type Bank struct {
	Meta      Meta       `json:"meta"`
	Questions []Question `json:"questions"`
}

type Meta struct {
	Title      string                      `json:"title"`
	TitleZH    string                      `json:"title_zh"`
	Version    string                      `json:"version"`
	Total      int                         `json:"total"`
	Dimensions map[Dimension]DimensionMeta `json:"dimensions"`
	Scoring    ScoringMeta                 `json:"scoring"`
}

type DimensionMeta struct {
	NameEN        string `json:"name_en"`
	NameZH        string `json:"name_zh"`
	PoleA         string `json:"pole_a"`
	PoleB         string `json:"pole_b"`
	DescriptionEN string `json:"description_en"`
	DescriptionZH string `json:"description_zh"`
	Count         int    `json:"count"`
}

type ScoringMeta struct {
	Description string                      `json:"description"`
	Thresholds  map[Strength]ThresholdRange `json:"thresholds"`
}

type Question struct {
	ID        string        `json:"id"`
	Dimension Dimension     `json:"dimension"`
	Reverse   bool          `json:"reverse"`
	Scenario  LocalizedText `json:"scenario"`
	Options   []Option      `json:"options"`
}

type LocalizedText struct {
	ZH string `json:"zh"`
	EN string `json:"en"`
}

type Option struct {
	Code  string        `json:"code"`
	Label LocalizedText `json:"label"`
	Score int           `json:"score"`
}
