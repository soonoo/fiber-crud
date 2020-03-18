CREATE TABLE IF NOT EXISTS users (
  id            SERIAL NOT NULL,
  github_login  TEXT NOT NULL,
  email         TEXT NOT NULL,
  avatar_url    TEXT NOT NULL
);
ALTER TABLE users ADD CONSTRAINT users_pk PRIMARY KEY (id);

CREATE TABLE IF NOT EXISTS repos (
  id            SERIAL NOT NULL,
  name          TEXT NOT NULL,
  owner         TEXT NOT NULL
);
ALTER TABLE repos ADD CONSTRAINT repos_pk PRIMARY KEY (id);

CREATE TABLE IF NOT EXISTS user_repos (
  user_id       INTEGER NOT NULL,
  repo_id       INTEGER NOT NULL
);
ALTER TABLE user_repos ADD CONSTRAINT user_repos_pk PRIMARY KEY (user_id, repo_id);
ALTER TABLE user_repos ADD CONSTRAINT user_repos_user_fk FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE user_repos ADD CONSTRAINT user_repos_repo_fk FOREIGN KEY (repo_id) REFERENCES repos(id);

CREATE TABLE IF NOT EXISTS commits (
  id            SERIAL NOT NULL,
  user_id       INTEGER NOT NULL,
  repo_id       INTEGER NOT NULL
);
ALTER TABLE commits ADD CONSTRAINT commits_pk PRIMARY KEY (id);
ALTER TABLE commits ADD CONSTRAINT commits_user_fk FOREIGN KEY (user_id) REFERENCES users(id);
ALTER TABLE commits ADD CONSTRAINT commits_repo_fk FOREIGN KEY (repo_id) REFERENCES repos(id);
