package main

// SQL queries for table creation
const (
	productsTable = `
		CREATE TABLE IF NOT EXISTS products (
			id smallint(6) unsigned NOT NULL auto_increment,
			pid smallint(6) unsigned NULL,
			user smallint(6) unsigned NULL,
			p_type varchar(20) NULL,
			label varchar(1500) NULL,
			details varchar(2500) NULL,
			image BLOB NULL,
			scid varchar(64) NULL,	
			inventory int(25),
			lastupdate timestamp DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		) ENGINE=InnoDB;
	`
	iAddressesTable = `
		CREATE TABLE IF NOT EXISTS i_addresses (
			id smallint(6) unsigned NOT NULL auto_increment,
			iaddr_id varchar(100) NULL,
			user smallint(6) unsigned NULL,
			product_id varchar(100) NULL,
			iaddr varchar(1500) NULL,
			ask_amount int(25) NULL,
			comment varchar(500) NULL,
			ia_scid varchar(64) NULL,	
			status tinyint(1) NULL,
			ia_inventory int(25) NULL,	
			lastupdate timestamp DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		) ENGINE=InnoDB;
	`
	usersTable = `
		CREATE TABLE IF NOT EXISTS users (
			userid smallint(6) unsigned NOT NULL auto_increment,
			username varchar(128) NULL,
			wallet varchar(128) NULL,
			uuid varchar(128) NULL,
			status tinyint(1) NULL,
			checkin timestamp NULL,
			lastupdate timestamp DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (userid)
		) ENGINE=InnoDB;
	`
	ordersTable = `
		CREATE TABLE IF NOT EXISTS orders (
			id smallint(6) unsigned NOT NULL auto_increment,
			uuid varchar(128) NULL,
			ia_id smallint(6) unsigned  NULL,
			user smallint(6) unsigned  NULL,
			lastupdate timestamp DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		) ENGINE=InnoDB;
	`
)
