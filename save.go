package ck2save

type property string
type SaveData struct {
	property map[string]property
}

func NewSaveData() *SaveData {
	sd := &SaveData{}
	sd.property = make(map[string]property)
	return sd
}

func (sd *SaveData) AddProperty(name string, prop property) {
	sd.property[name] = prop
}

func (sd *SaveData) Property(name string) property {
	return sd.property[name]
}
