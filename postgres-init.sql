CREATE TABLE spells (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid (),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    level INTEGER NOT NULL,
    school VARCHAR(255),
    duration VARCHAR(255) NOT NULL,
    casting_time VARCHAR(255) NOT NULL,
    range VARCHAR(255) NOT NULL,
    components VARCHAR(255)
);
CREATE TABLE characters (
   id uuid PRIMARY KEY NOT NULL DEFAULT gen_random_uuid (),
   name VARCHAR(255),
   level INTEGER,
   class VARCHAR(255),
   subclass VARCHAR(255),
   ac INTEGER,
   str INTEGER,
   dex INTEGER,
   con INTEGER,
   int INTEGER,
   wis INTEGER,
   cha  INTEGER,
   initiative INTEGER,
   speed INTEGER,

);
CREATE TABLE characters_spells(
    character_id UUID REFERENCES characters(id),
    spell_id UUID REFERENCES spells(id),
    PRIMARY KEY (character_id, spell_id)
);
