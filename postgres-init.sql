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