DROP TABLE IF EXISTS organisations CASCADE;

CREATE TABLE organisations (
  id	char(36)		PRIMARY KEY,
  name	VARCHAR(50)	NOT NULL
);

DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE users (
  id		char(36)		PRIMARY KEY,
  facebook_id	VARCHAR(128)	NOT NULL,
  org_id	char(36)		NULL REFERENCES organisations(id)
);

DROP TABLE IF EXISTS tokens CASCADE;

CREATE TABLE tokens (
  id		char(36)		PRIMARY KEY,
  name		VARCHAR(100)	NOT NULL,
  expires	TIMESTAMP	NOT NULL,
  org_id	char(36)		NOT NULL REFERENCES organisations(id)
);

DROP TABLE IF EXISTS user_tokens;

CREATE TABLE user_tokens (
  user_id	char(36)		NOT NULL REFERENCES users(id),
  org_id	char(36)		NOT NULL REFERENCES organisations(id), 
  number	SMALLINT	DEFAULT 1
);

