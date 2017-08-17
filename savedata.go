package ck2save

type SaveData struct {
	property    []*Property
	propertyMap map[string]*Property
}

func NewSaveData() *SaveData {
	return &SaveData{
		propertyMap: make(map[string]*Property),
	}
}

func (sd *SaveData) AddProperty(prop *Property) {
	sd.property = append(sd.property, prop)
	sd.propertyMap[prop.Name] = prop
}

func (sd *SaveData) Property(name string) *Property {
	return sd.propertyMap[name]
}
