CREATE TABLE spells (
    id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid (),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    level INTEGER NOT NULL,
    school VARCHAR(255),
    duration VARCHAR(255) NOT NULL,
    casting_time VARCHAR(255) NOT NULL,
    range VARCHAR(255) NOT NULL,
    components VARCHAR(255)
);

INSERT INTO spells (name, description, level, school, duration, casting_time, range, components) VALUES ('Acid Arrow', 'A shimmering green arrow streaks toward a target within range and bursts in a spray of acid.', 3, 'evocation', '5 minutes', '1 action', '20 feet', 'VSM');