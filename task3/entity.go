package task3

type Student struct {
	ID    uint   `gorm:"primaryKey;autoIncrement"`
	Name  string `gorm:"type:varchar(100)"`
	Age   int    `gorm:"type:int"`
	Grade string `gorm:"type:varchar(50)"`
}

type Account struct {
	ID      uint    `gorm:"primaryKey;autoIncrement"`
	Balance float64 `gorm:"type:decimal(10,2)"`
}

type Transaction struct {
	ID            uint    `gorm:"primaryKey;autoIncrement"`
	FromAccountID uint    `gorm:"type:int"`
	ToAccountID   uint    `gorm:"type:int"`
	Amount        float64 `gorm:"type:decimal(10,2)"`
}
