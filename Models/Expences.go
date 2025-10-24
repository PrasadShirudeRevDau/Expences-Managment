package models

import "time"

type User struct {
	UserID uint `json:"user_id" gorm:"primaryKey"`
	UserName string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"created_at"`

	Expenses []Expense `gorm:"foreignKey:UserID"`
}

type Expense struct {
	Id         uint		`json:"id" gorm:"primaryKey"`
	UserID    uint		`json:"user_id" gorm:"not null"`
	Amount     float64 `json:"amount"`
	Category   ExpenseCategory `gorm:"type:enum('Housing','Utilities','Transportation','Food_Groceries','Health_Fitness','Entertainment','Education','Clothing','Personal_Care','Gifts_Donation','Travel_Vacation','Insurance','Miscellaneous');"`
	Date       time.Time `json:"date"`
	Note       string	`json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ExpenseCategory string

const (
	Housing ExpenseCategory = "Housing"
	Utilities ExpenseCategory = "Utilities"
	Transportation ExpenseCategory = "Transportation"
	Food_Groceries ExpenseCategory = "Food_Groceries"
	Health_Fitness ExpenseCategory = "Health_Fitness"
	Entertainment ExpenseCategory = "Entertainment" 
	Education ExpenseCategory = "Education"
	Clothing ExpenseCategory = "Clothing"
	Personal_Care ExpenseCategory = "Personal_Care"
	Gifts_Donation ExpenseCategory = "Gifts_Donation"
	Travel_Vacation ExpenseCategory = "Travel_Vacation"
	Insurance ExpenseCategory = "Insurance"
	Miscellaneous ExpenseCategory = "Miscellaneous"
)

func GetAllCategories () []ExpenseCategory {
	return []ExpenseCategory{Housing,Utilities,Transportation,Food_Groceries,Health_Fitness,Entertainment,Education,Clothing,Personal_Care,Gifts_Donation,Travel_Vacation,Insurance,Miscellaneous}
}