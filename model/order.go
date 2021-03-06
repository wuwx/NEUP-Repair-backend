package model

import (
	"database/sql"
	"time"

	"github.com/pkg/errors"
)

// Order is the main struct (now the only) to request for
type Order struct {
	ID          int64  `json:"id,omitempty" db:"id,omitempty"`
	Name        string `json:"name, omitempty" db:"name,omitempty"`
	StuID       string `json:"stu_id,omitempty" db:"stu_id,omitempty"`
	Date        string `json:"date,omitempty" db:"date,omitempty"`
	Comment     string `json:"comment" db:"comment"`
	Agreement   string `json:"agreement" db:"agreement"`
	ServiceType string `json:"service_type" db:"service_type"`
	// Below are info to modify after order created
	Rating     int64     `json:"rating,omitempty" db:"rating"`
	OperaterID int64     `json:"operater_id,omitempty" db:"operater_id"`
	SecretID   string    `json:"secret_id,omitempty" db:"secret_id"`
	DoneFlag   bool      `json:"done_flag,omitempty" db:"done_flag"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (o *Order) Insert(db Storager) (err error) {
	sqlStr := "INSERT INTO orders (name, stu_id, date, comment, service_type, secret_id)VALUES(?,?,?,?,?,?)"
	_, err = db.Queryx(sqlStr, o.Name, o.StuID, o.Date, o.Comment, o.ServiceType, o.SecretID)
	if err != nil {
		err = errors.Wrap(err, "insert error")
		return
	}
	return
}

func OrderByStuID(db Storager, stuID string) (o Order, err error) {
	sqlStr := "SELECT * FROM orders where stu_id = ? AND done_flag = FALSE ORDER BY CREATE_TIME DESC LIMIT 1"
	err = db.QueryRowx(sqlStr, stuID).StructScan(&o)
	if err != nil && err != sql.ErrNoRows {
		err = errors.Wrap(err, "order by stu_id error")
		return
	}
	if err == sql.ErrNoRows {
		err = nil
		o.ID = -1
	}
	return
}

func OrderByID(db Storager, ID int) (err error) {
	return
}

func UpdateOrderDoneFlagBySecret(db Storager, secretID string) (err error) {
	sqlStr := "UPDATE orders SET done_flag = TRUE WHERE secret_id = ?"
	_, err = db.Exec(sqlStr, secretID)

	if err != nil && err != sql.ErrNoRows {
		err = errors.Wrap(err, "order by stu_id error")
		return
	}
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func OrderBySecret(db Storager, secretID string) (o Order, err error) {
	sqlStr := "SELECT * FROM orders where secret_id = ?"
	err = db.QueryRowx(sqlStr, secretID).StructScan(&o)
	if err != nil && err != sql.ErrNoRows {
		err = errors.Wrap(err, "order by stu_id error")
		return
	}
	if err == sql.ErrNoRows {
		err = nil
		o.ID = -1
	}
	return
}

func OrderPager(db Storager) (err error) {
	return
}
