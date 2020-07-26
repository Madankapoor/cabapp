package models

// Passenger ...
type Passenger struct {
	ID        int64  `db:"id, primarykey, autoincrement" json:"id"`
	Email     string `db:"email" json:"email"`
	Name      string `db:"name" json:"name"`
	MobileNo  string `db:"mobileno" json:"mobileno"`
	UpdatedAt int64  `db:"updated_at" json:"-"`
	CreatedAt int64  `db:"created_at" json:"-"`
}

// PassengerModel ...
type PassengerModel struct{}
