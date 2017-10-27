-- Defines the initial lists table
CREATE TABLE `lists` (
	`id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	`name`	TEXT NOT NULL,
	`slug`	TEXT NOT NULL UNIQUE,
	`list`	BLOB,
	`description`	TEXT,
	`created_at`	datetime NOT NULL,
	`updated_at`	datetime NOT NULL
);
