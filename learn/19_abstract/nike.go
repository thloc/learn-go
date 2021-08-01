package abstract

type Nike struct {
}

func (a *Nike) MakeShoe() IShoe {
	return &NikeShoe {
		shoe: shoe {
			logo: "nike",
			size: 14,
		},
	}
}

func (a *Nike) MakeShort() IShort  {
	return &NikeShort {
		short: short {
			logo: "nike",
			size: 14,
		}
	}
}
