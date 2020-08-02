package commonModel

import (
	"context"
	"database/sql"
	"errors"
	"rating-service/db"
)

const (
	Driver    = "driver"
	Passenger = "passenger"
)

type User struct {
	UserID      int    `json:"userID"`
	RideID      int    `json:"rideID"`
	RatingValue int    `json:"rating"`
	RatingBase  int    `json:"rating_base"`
	UserType    string `json:"user_type,omitempty"`
}

func GetAggregateRating(ctx *context.Context, user User) (float64, error) {
	var row *sql.Row
	conn, err := db.RatingDB.Conn(*ctx)
	if err != nil {
		return 0, errors.New("something bad happened")
	}
	defer conn.Close()
	var averageRating float64
	row = conn.QueryRowContext(*ctx, `SELECT ROUND(avg(rating_value)/20, 2) as avgRating FROM user_rating where user_id=?`,
											user.UserID)
	if err = row.Scan(&averageRating); err != nil {
		return 0, errors.New("no rating(s) found")
	}
	return averageRating, nil
}

func SetRating(ctx *context.Context, user User) error {
	conn, err := db.RatingDB.Conn(*ctx)
	if err != nil {
		return errors.New("something bad happened")
	}
	defer conn.Close()
	_, err = conn.ExecContext(*ctx, `INSERT INTO user_rating (user_id,user_type,rating_value,rating_base) VALUES (?,?,?,?)`,
												user.UserID, user.UserType, user.RatingValue, user.RatingBase)
	if err != nil {
		return err
	}
	return nil
}
