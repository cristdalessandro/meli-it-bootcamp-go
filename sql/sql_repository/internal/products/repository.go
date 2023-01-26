package products

import (
	"database/sql"
	"errors"

	"github.com/cristdalessandro/meli-it-bootcamp-go/sql/internal/domain"
	"github.com/go-sql-driver/mysql"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	GetOne(id int) (domain.Product, error)
	Create(p domain.Product) (domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

var (
	ErrNotFound   = errors.New("product not found")
	ErrInternal   = errors.New("an internal error")
	ErrDuplicated = errors.New("duplicated product")
)

func (repo *repository) GetAll() (res []domain.Product, err error) {

	var rows *sql.Rows
	query := `SELECT id, name, quantity, code_value, is_published, expiration, price 
				FROM products`
	rows, err = repo.db.Query(query)
	if err != nil {
		err = ErrInternal
		return
	}

	for rows.Next() {
		b := domain.Product{}
		err = rows.Scan(
			&b.Id,
			&b.Name,
			&b.Quantity,
			&b.Code_value,
			&b.Is_published,
			&b.Expiration,
			&b.Price,
		)
		if err != nil {
			err = ErrInternal
			return
		}
		res = append(res, b)
	}
	return
}

func (repo *repository) GetOne(id int) (res domain.Product, err error) {

	query := `SELECT id, name, quantity, code_value, is_published, expiration, price
				FROM products
				WHERE id = ?`
	row := repo.db.QueryRow(query, id)
	if row.Err() != nil {
		switch row.Err() {
		case sql.ErrNoRows:
			err = ErrNotFound
		default:
			err = ErrInternal
		}
		return
	}

	err = row.Scan(
		&res.Id,
		&res.Name, &res.Quantity,
		&res.Code_value,
		&res.Is_published,
		&res.Expiration,
		&res.Price,
	)
	if err != nil {
		err = ErrInternal
	}
	return
}

func (repo *repository) Create(p domain.Product) (res domain.Product, err error) {
	var (
		statement    *sql.Stmt
		result       sql.Result
		id           int
		rowsAffected int64
	)

	queryNextId := `SELECT MAX(t.id) 
					FROM products t`
	row := repo.db.QueryRow(queryNextId)
	if row.Err() != nil {
		err = ErrInternal
		return
	}
	row.Scan(&id)
	id++

	statement, err = repo.db.Prepare(
		`INSERT INTO products (id, name, quantity, code_value, is_published, expiration, price)
			VALUES (?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		err = ErrInternal
		return
	}
	defer statement.Close()

	result, err = statement.Exec(
		id,
		p.Name,
		p.Quantity,
		p.Code_value,
		p.Is_published,
		p.Expiration,
		p.Price,
	)
	if err != nil {
		driverErr, ok := err.(*mysql.MySQLError)
		if !ok {
			err = ErrInternal
			return
		}

		switch driverErr.Number {
		case 1062:
			err = ErrDuplicated
		default:
			err = ErrInternal
		}
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil || rowsAffected != 1 {
		err = ErrInternal
		return
	}

	res = p
	res.Id = id
	return
}
