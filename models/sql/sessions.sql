DROP TABLE sessions;

CREATE TABLE sessions(
    id SERIAL PRIMARY KEY,
    user_id INT UNIQUE REFERENCES users(id) ON DELETE CASCADE,
    token_hash TEXT UNIQUE NOT NULL
);


select users.id, users.email, sessions.id as session_id
from users
join sessions on users.id = sessions.user_id;



