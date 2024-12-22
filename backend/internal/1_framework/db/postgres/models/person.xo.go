package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
)

// Person represents a row from 'public.persons'.
type Person struct {
	ID          int            `json:"id"`           // id
	Name        sql.NullString `json:"name"`         // name
	MailAddress sql.NullString `json:"mail_address"` // mail_address
	CreatedAt   sql.NullTime   `json:"created_at"`   // created_at
	UpdatedAt   sql.NullTime   `json:"updated_at"`   // updated_at
	// xo fields
	_exists, _deleted bool
}

// Exists returns true when the [Person] exists in the database.
func (p *Person) Exists() bool {
	return p._exists
}

// Deleted returns true when the [Person] has been marked for deletion
// from the database.
func (p *Person) Deleted() bool {
	return p._deleted
}

// Insert inserts the [Person] to the database.
func (p *Person) Insert(ctx context.Context, db DB) error {
	switch {
	case p._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case p._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (manual)
	const sqlstr = `INSERT INTO public.persons (` +
		`id, name, mail_address, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)`
	// run
	logf(sqlstr, p.ID, p.Name, p.MailAddress, p.CreatedAt, p.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID, p.Name, p.MailAddress, p.CreatedAt, p.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Update updates a [Person] in the database.
func (p *Person) Update(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case p._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.persons SET ` +
		`name = $1, mail_address = $2, created_at = $3, updated_at = $4 ` +
		`WHERE id = $5`
	// run
	logf(sqlstr, p.Name, p.MailAddress, p.CreatedAt, p.UpdatedAt, p.ID)
	if _, err := db.ExecContext(ctx, sqlstr, p.Name, p.MailAddress, p.CreatedAt, p.UpdatedAt, p.ID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the [Person] to the database.
func (p *Person) Save(ctx context.Context, db DB) error {
	if p.Exists() {
		return p.Update(ctx, db)
	}
	return p.Insert(ctx, db)
}

// Upsert performs an upsert for [Person].
func (p *Person) Upsert(ctx context.Context, db DB) error {
	switch {
	case p._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	const sqlstr = `INSERT INTO public.persons (` +
		`id, name, mail_address, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`)` +
		` ON CONFLICT (id) DO ` +
		`UPDATE SET ` +
		`name = EXCLUDED.name, mail_address = EXCLUDED.mail_address, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, p.ID, p.Name, p.MailAddress, p.CreatedAt, p.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID, p.Name, p.MailAddress, p.CreatedAt, p.UpdatedAt); err != nil {
		return logerror(err)
	}
	// set exists
	p._exists = true
	return nil
}

// Delete deletes the [Person] from the database.
func (p *Person) Delete(ctx context.Context, db DB) error {
	switch {
	case !p._exists: // doesn't exist
		return nil
	case p._deleted: // deleted
		return nil
	}
	// delete with single primary key
	const sqlstr = `DELETE FROM public.persons ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, p.ID)
	if _, err := db.ExecContext(ctx, sqlstr, p.ID); err != nil {
		return logerror(err)
	}
	// set deleted
	p._deleted = true
	return nil
}

// PersonByID retrieves a row from 'public.persons' as a [Person].
//
// Generated from index 'persons_pkey'.
func PersonByID(ctx context.Context, db DB, id int) (*Person, error) {
	// query
	const sqlstr = `SELECT ` +
		`id, name, mail_address, created_at, updated_at ` +
		`FROM public.persons ` +
		`WHERE id = $1`
	// run
	logf(sqlstr, id)
	p := Person{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, id).Scan(&p.ID, &p.Name, &p.MailAddress, &p.CreatedAt, &p.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &p, nil
}
