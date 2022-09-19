package dal

import (
	"database/sql"
	"errors"
	"fmt"

	"alukart32.com/urlshort/internal/models"
)

const uuidLength = 36

// Redirect handles all of the database actions.
type Redirect struct {
	DB           *sql.DB
	GenerateUUID GenerateUUID
}

// DAO methods for models.RedirectEntity.
type RedirectDAO interface {
	Create(r *models.Redirect) (*models.RedirectEntity, error)
	FetchByID(id string) (*models.RedirectEntity, error)
	FetchAll() ([]*models.RedirectEntity, error)
	Update(id string, r *models.Redirect) (*models.RedirectEntity, error)
	Delete(id string) error
}

func (r *Redirect) Create(rd *models.Redirect) (*models.RedirectEntity, error) {
	fail := func(err error) (*models.RedirectEntity, error) {
		return nil, fmt.Errorf("Create redirect: %v", err)
	}

	if rd == nil {
		return fail(errors.New("can not be nil"))
	}

	e := &models.RedirectEntity{
		Redirect: models.Redirect{
			Path: rd.Path,
			Url:  rd.Url,
		},
		Entity: models.Entity{
			ID: r.GenerateUUID(),
		},
	}

	tx, err := r.DB.Begin()
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	const stmt = `insert into public.redirects (id, path, url) values($1, $2, $3)`
	if _, err := tx.Exec(stmt, e.ID, e.Path, e.Url); err != nil {
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return e, nil
}

func (r *Redirect) FetchByID(id string) (*models.RedirectEntity, error) {
	if len(id) != uuidLength {
		return nil, fmt.Errorf("redirect fetch by id length %d", len(id))
	}

	fail := func(err error) (*models.RedirectEntity, error) {
		return nil, fmt.Errorf("FetchByID: %v", err)
	}

	tx, err := r.DB.Begin()
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	const stmt = `select id, path, url from public.redirects where id = $1`
	var e *models.RedirectEntity

	err = tx.QueryRow(stmt, id).Scan(&e.ID, &e.Path, &e.Url)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fail(ErrNoRedirect)
		}
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}
	return e, nil
}

func (r *Redirect) FetchAll() ([]*models.RedirectEntity, error) {
	fail := func(err error) (*models.RedirectEntity, error) {
		return nil, fmt.Errorf("FetchAll: %v", err)
	}

	tx, err := r.DB.Begin()
	if err != nil {
		fail(err)
	}
	defer tx.Rollback()

	const stmt = `select id, path, url from public.redirects`
	rows, err := r.DB.Query(stmt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fail(ErrNoRedirect)
		}
		fail(err)
	}

	entities := []*models.RedirectEntity{}
	for rows.Next() {
		e := &models.RedirectEntity{}
		if err := rows.Scan(&e.ID, &e.Path, &e.Url); err != nil {
			return nil, fmt.Errorf("redirect row scan error %w", err)
		}
		entities = append(entities, e)
	}
	if rows.Err() != nil {
		return nil, err
	}
	if err = rows.Close(); err != nil {
		fail(err)
	}

	if err = tx.Commit(); err != nil {
		fail(err)
	}

	return entities, nil
}

func (r *Redirect) Update(id string, redirect *models.Redirect) (*models.RedirectEntity, error) {
	fail := func(err error) (*models.RedirectEntity, error) {
		return nil, fmt.Errorf("Update redirect: %v", err)
	}

	if len(id) != uuidLength {
		fail(fmt.Errorf("redirect fetch by id length %d", len(id)))
	}
	if redirect == nil {
		fail(fmt.Errorf("redirect can not be nil"))
	}

	e, err := r.FetchByID(id)
	if err != nil {
		return nil, err
	}
	if len(redirect.Path) > 0 {
		e.Path = redirect.Path
	}
	if len(redirect.Url) > 0 {
		e.Url = redirect.Url
	}

	tx, err := r.DB.Begin()
	if err != nil {
		fail(err)
	}
	defer tx.Rollback()

	const stmt = `update public.redirects SET path = $1, url = $2 where id = $3`
	if _, err := r.DB.Exec(stmt, e.Path, e.Url, id); err != nil {
		fail(err)
	}
	if err = tx.Commit(); err != nil {
		fail(err)
	}

	return e, nil
}

func (r *Redirect) Delete(id string) error {
	fail := func(err error) error {
		return fmt.Errorf("Delete redirect: %v", err)
	}

	if len(id) != uuidLength {
		fail(fmt.Errorf("redirect fetch by id length %d", len(id)))
	}

	_, err := r.FetchByID(id)
	if err != nil {
		return err
	}

	tx, err := r.DB.Begin()
	if err != nil {
		fail(err)
	}
	defer tx.Rollback()

	const stmt = `delete public.redirects where id = $1`
	if _, err := r.DB.Exec(stmt, id); err != nil {
		fail(err)
	}

	if err = tx.Commit(); err != nil {
		fail(err)
	}
	return nil
}
