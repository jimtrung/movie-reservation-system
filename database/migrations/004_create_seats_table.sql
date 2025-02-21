-- Seats table
CREATE TABLE seats (
    seat_id SERIAL PRIMARY KEY,
    showtime_id INT NOT NULL REFERENCES showtimes(showtime_id) ON DELETE CASCADE,
    seat_number VARCHAR(10) NOT NULL, 
    is_available BOOLEAN DEFAULT TRUE
);

CREATE INDEX seats_idx ON seats (showtime_id, seat_number);