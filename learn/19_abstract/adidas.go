package abstract

type Adidas struct {
}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe {
		shoe: shoe {
			logo: "adidas",
			size: 14,
		},
	}
}

func (a *Adidas) MakeShort() IShort  {
	return &AdidasShort {
		short: short {
			logo: "adidas",
			size: 14,
		}
	}
}
