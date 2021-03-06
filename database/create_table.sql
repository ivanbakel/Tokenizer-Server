DROP TABLE IF EXISTS organisations CASCADE;

CREATE TABLE organisations (
  id	UUID		PRIMARY KEY DEFAULT uuid_generate_v4(),
  name	VARCHAR(50)	NOT NULL
);

DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE users (
  id		UUID		PRIMARY KEY DEFAULT uuid_generate_v4(),
  facebook_id	VARCHAR(128)	NOT NULL,
  org_id	UUID		NULL REFERENCES organisations(id)
);

DROP TABLE IF EXISTS tokens CASCADE;

CREATE TABLE tokens (
  id		UUID		PRIMARY KEY DEFAULT uuid_generate_v4(),
  name		VARCHAR(100)	NOT NULL,
  expires	TIMESTAMP	NOT NULL,
  org_id	UUID		NOT NULL REFERENCES organisations(id)
);

DROP TABLE IF EXISTS user_tokens;

CREATE TABLE user_tokens (
  user_id	UUID		NOT NULL REFERENCES users(id),
  token_id	UUID		NOT NULL REFERENCES tokens(id), 
  number	SMALLINT	DEFAULT 1,
  PRIMARY KEY(user_id, token_id)
);

