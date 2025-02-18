CREATE TABLE movies (
    movie_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    image TEXT NOT NULL,
    genre TEXT,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX genre_idx ON movies USING GIN (genre);

CREATE FUNCTION update_updated_at_movie()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_updated_at
BEFORE UPDATE ON movies
FOR EACH ROW EXECUTE FUNCTION update_updated_at_movie();

SELECT * FROM movies;
