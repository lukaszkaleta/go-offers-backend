package postgres

func (pgDb *PgDb) createUserTable() error {
	var query = `drop table if exists business_user`
	_, err := pgDb.Database.Exec(query)
	if err != nil {
		return err
	}

	query = `CREATE TABLE if not exists business_user (
		id serial primary key,
		person_first_name text not null default '',
		person_last_name text not null default '',
		person_email text not null default '',
		person_phone text not null default '',
		address_line1 text not null default '',
		address_line2 text not null default '',
		address_city       text not null default '',
		address_postal_code text not null default '',
		address_district   text not null default '')`
	_, err = pgDb.Database.Exec(query)

	return err
}
