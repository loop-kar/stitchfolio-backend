package entities

type DressType struct {
	*Model `mapstructure:",squash"`

	Name         string `json:"name"`
	Measurements string `json:"measurements"` //CSV of mesurement types Hip, Waist, Chest
}

func (DressType) TableName() string {
	return "stitch.DressTypes"
}

func (DressType) TableNameForQuery() string {
	return "\"stitch\".\"DressTypes\" E"
}
