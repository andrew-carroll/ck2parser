package ck2save

type SaveData struct {
	property    []*Property
	propertyMap map[string]*Property
}

func NewSaveData() *SaveData {
	sd := &SaveData{}
	sd.propertyMap = make(map[string]*Property)
	return sd
}

func (sd *SaveData) AddProperty(name string, prop *Property) {
	sd.property = append(sd.property, prop)
	sd.propertyMap[name] = prop
}

func (sd *SaveData) Property(name string) *Property {
	return sd.propertyMap[name]
}
