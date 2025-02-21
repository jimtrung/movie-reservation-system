INSERT INTO movies (title, description, image, genre) VALUES
('Inception', 'A thief who enters the dreams of others to steal secrets.', 'inception.jpg', 'Sci-Fi'),
('The Dark Knight', 'Batman faces Joker, a criminal mastermind.', 'dark_knight.jpg', 'Action'),
('Interstellar', 'Astronauts travel through a wormhole to save humanity.', 'interstellar.jpg', 'Sci-Fi'),
('Titanic', 'A romance set on the ill-fated Titanic.', 'titanic.jpg', 'Drama'),
('The Matrix', 'A hacker discovers the truth about his reality.', 'matrix.jpg', 'Sci-Fi'),
('Avengers: Endgame', 'Superheroes fight to undo Thanos'' snap.', 'endgame.jpg', 'Action'),
('Gladiator', 'A betrayed general seeks revenge in ancient Rome.', 'gladiator.jpg', 'Action'),
('Jurassic Park', 'Scientists clone dinosaurs, chaos ensues.', 'jurassic_park.jpg', 'Adventure'),
('The Godfather', 'The rise of a mafia family.', 'godfather.jpg', 'Crime'),
('Pulp Fiction', 'Interwoven crime stories in Los Angeles.', 'pulp_fiction.jpg', 'Crime'),
('The Lion King', 'A young lion prince flees his kingdom.', 'lion_king.jpg', 'Animation'),
('Forrest Gump', 'A slow-witted man experiences historical events.', 'forrest_gump.jpg', 'Drama'),
('The Shawshank Redemption', 'A banker is imprisoned for a crime he didn’t commit.', 'shawshank.jpg', 'Drama'),
('Fight Club', 'A disillusioned man forms an underground fight club.', 'fight_club.jpg', 'Thriller'),
('The Lord of the Rings: The Fellowship of the Ring', 'A hobbit embarks on a quest to destroy a powerful ring.', 'lotr_fellowship.jpg', 'Fantasy'),
('The Lord of the Rings: The Two Towers', 'The fellowship faces new challenges.', 'lotr_two_towers.jpg', 'Fantasy'),
('The Lord of the Rings: The Return of the King', 'The final battle for Middle-earth begins.', 'lotr_return_king.jpg', 'Fantasy'),
('Harry Potter and the Sorcerer''s Stone', 'A boy discovers he is a wizard.', 'harry_potter_1.jpg', 'Fantasy'),
('Harry Potter and the Chamber of Secrets', 'A hidden chamber in Hogwarts is opened.', 'harry_potter_2.jpg', 'Fantasy'),
('Harry Potter and the Prisoner of Azkaban', 'A prisoner escapes from Azkaban.', 'harry_potter_3.jpg', 'Fantasy'),
('Star Wars: A New Hope', 'A young hero joins the Rebellion against the Empire.', 'star_wars_4.jpg', 'Sci-Fi'),
('Star Wars: The Empire Strikes Back', 'The Empire strikes back against the Rebellion.', 'star_wars_5.jpg', 'Sci-Fi'),
('Star Wars: Return of the Jedi', 'The final battle against the Empire.', 'star_wars_6.jpg', 'Sci-Fi'),
('Spider-Man: No Way Home', 'Peter Parker deals with multiverse chaos.', 'spiderman_nwh.jpg', 'Superhero'),
('Deadpool', 'A wisecracking mercenary seeks revenge.', 'deadpool.jpg', 'Action');

INSERT INTO showtimes (movie_id, showtimes) VALUES
(1, ARRAY['2025-03-01 14:00:00+00', '2025-03-01 18:00:00+00']::TIMESTAMPTZ[]),
(2, ARRAY['2025-03-02 15:00:00+00', '2025-03-02 19:00:00+00']::TIMESTAMPTZ[]),
(3, ARRAY['2025-03-03 16:00:00+00', '2025-03-03 20:00:00+00']::TIMESTAMPTZ[]),
(4, ARRAY['2025-03-04 17:00:00+00', '2025-03-04 21:00:00+00']::TIMESTAMPTZ[]),
(5, ARRAY['2025-03-05 18:00:00+00', '2025-03-05 22:00:00+00']::TIMESTAMPTZ[]),
(6, ARRAY['2025-03-06 19:00:00+00', '2025-03-06 23:00:00+00']::TIMESTAMPTZ[]),
(7, ARRAY['2025-03-07 14:00:00+00', '2025-03-07 18:00:00+00']::TIMESTAMPTZ[]),
(8, ARRAY['2025-03-08 15:00:00+00', '2025-03-08 19:00:00+00']::TIMESTAMPTZ[]),
(9, ARRAY['2025-03-09 16:00:00+00', '2025-03-09 20:00:00+00']::TIMESTAMPTZ[]),
(10, ARRAY['2025-03-10 17:00:00+00', '2025-03-10 21:00:00+00']::TIMESTAMPTZ[]),
(11, ARRAY['2025-03-11 18:00:00+00', '2025-03-11 22:00:00+00']::TIMESTAMPTZ[]),
(12, ARRAY['2025-03-12 19:00:00+00', '2025-03-12 23:00:00+00']::TIMESTAMPTZ[]),
(13, ARRAY['2025-03-13 14:00:00+00', '2025-03-13 18:00:00+00']::TIMESTAMPTZ[]),
(14, ARRAY['2025-03-14 15:00:00+00', '2025-03-14 19:00:00+00']::TIMESTAMPTZ[]),
(15, ARRAY['2025-03-15 16:00:00+00', '2025-03-15 20:00:00+00']::TIMESTAMPTZ[]),
(16, ARRAY['2025-03-16 17:00:00+00', '2025-03-16 21:00:00+00']::TIMESTAMPTZ[]),
(17, ARRAY['2025-03-17 18:00:00+00', '2025-03-17 22:00:00+00']::TIMESTAMPTZ[]),
(18, ARRAY['2025-03-18 19:00:00+00', '2025-03-18 23:00:00+00']::TIMESTAMPTZ[]),
(19, ARRAY['2025-03-19 14:00:00+00', '2025-03-19 18:00:00+00']::TIMESTAMPTZ[]),
(20, ARRAY['2025-03-20 15:00:00+00', '2025-03-20 19:00:00+00']::TIMESTAMPTZ[]),
(21, ARRAY['2025-03-21 16:00:00+00', '2025-03-21 20:00:00+00']::TIMESTAMPTZ[]),
(22, ARRAY['2025-03-22 17:00:00+00', '2025-03-22 21:00:00+00']::TIMESTAMPTZ[]),
(23, ARRAY['2025-03-23 18:00:00+00', '2025-03-23 22:00:00+00']::TIMESTAMPTZ[]),
(24, ARRAY['2025-03-24 19:00:00+00', '2025-03-24 23:00:00+00']::TIMESTAMPTZ[]),
(25, ARRAY['2025-03-25 14:00:00+00', '2025-03-25 18:00:00+00']::TIMESTAMPTZ[]);

DO $$ 
DECLARE 
    showtime_id INT;
    seat_row CHAR(1);
    seat_number INT;
BEGIN
    FOR showtime_id IN 1..25 LOOP  -- 25 showtimes
        FOR seat_row IN ASCII('A')..ASCII('J') LOOP  -- Convert 'A'-'J' to ASCII range
            FOR seat_number IN 1..10 LOOP  -- Seats 1 to 10
                INSERT INTO seats (showtime_id, seat_number, is_available)
                VALUES (showtime_id, CHR(seat_row) || seat_number, RANDOM() > 0.3); -- 70% availability
            END LOOP;
        END LOOP;
    END LOOP;
END $$;
