package domain

type Stock struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	ProductID uint `json:"product_id" gorm:"not null;index;unique"`
	Quantity  int  `json:"quantity" gorm:"default:0"`
	Reserved  int  `json:"reserved" gorm:"default:0"`

	Product *Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
}

func (s *Stock) Available() int {
	return s.Quantity - s.Reserved
}
