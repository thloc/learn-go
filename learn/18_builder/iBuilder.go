package builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFoor()
	getHouse() House // Build House
}

func GetBuilder(builderType string) IBuilder {
	switch builderType {
	case "normal":
		return &normalBuilder{}
	case "igloo":
		return &iglooBuilder{}
	}

	return nil
}
