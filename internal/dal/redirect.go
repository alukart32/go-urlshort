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

func (r *Redirect) Create(path *models.Redirect) (*models.RedirectEntity, error) {
	if path == nil {
		return nil, errors.New("redirect can not be nil")
	}

	e := &models.RedirectEntity{
		Redirect: models.Redirect{
			Path: path.Path,
			Url:  path.Url,
		},
		Entity: models.Entity{
			ID: r.GenerateUUID(),
		},
	}

	const stmt = `insert into public.redirects (id, path, url) values($1, $2, $3)`
	if _, err := r.DB.Exec(stmt, e.ID, e.Path, e.Url); err != nil {
		return nil, fmt.Errorf("redirect create insert %w", err)
	}
	return e, nil
}

func (r *Redirect) FetchByID(id string) (*models.RedirectEntity, error) {
	if len(id) != uuidLength {
		return nil, fmt.Errorf("redirect fetch by id length %d", len(id))
	}

	const stmt = `select id, path, url from public.redirects where id = $1`
	var e *models.RedirectEntity

	err := r.DB.QueryRow(stmt, id).Scan(&e.ID, &e.Path, &e.Url)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, ErrNoRedirect
	case err != nil:
		return nil, err
	default:
		return e, nil
	}
}

func (r *Redirect) FetchAll() ([]*models.RedirectEntity, error) {
	const stmt = `select id, path, url from public.redirects`
	rows, err := r.DB.Query(stmt)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, ErrNoRedirect
	case err != nil:
		return nil, fmt.Errorf("redirect fetch all query: %w", err)
	default:
	}
	defer rows.Close()

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

	return entities, nil
}

func (r *Redirect) Update(id string, redirect *models.Redirect) (*models.RedirectEntity, error) {
	switch {
	case len(id) != uuidLength:
		return nil, fmt.Errorf("redirect fetch by id length %d", len(id))
	case redirect == nil:
		return nil, errors.New("redirect can not be nil")
	default:
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

	const stmt = `update public.redirects SET path = $1, url = $2 where id = $3`
	if _, err := r.DB.Exec(stmt, e.Path, e.Url, id); err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Redirect) Delete(id string) error {
	if len(id) != uuidLength {
		return fmt.Errorf("redirect fetch by id length %d", len(id))
	}

	_, err := r.FetchByID(id)
	if err != nil {
		return err
	}

	const stmt = `delete public.redirects where id = $1`
	if _, err := r.DB.Exec(stmt, id); err != nil {
		return err
	}
	return nil
}
