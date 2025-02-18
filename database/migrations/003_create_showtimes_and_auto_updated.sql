CREATE TABLE showtimes (
    showtime_id SERIAL PRIMARY KEY,
    movie_id INT NOT NULL REFERENCES movies(movie_id) ON DELETE CASCADE,
    showtimes TIMESTAMPTZ[] NOT NULL,

    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX showtimes_idx ON showtimes USING GIN (showtimes);

CREATE FUNCTION update_showtimes_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_showtimes_updated_at
BEFORE UPDATE ON showtimes
FOR EACH ROW EXECUTE FUNCTION update_showtimes_updated_at();
