CREATE TABLE IF NOT EXISTS measurement(
    id SERIAL PRIMARY KEY,
    value REAL NOT NULL,
    unit VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS disc(
    name VARCHAR(200),
    distributor VARCHAR(400),
    max_weight_id INT REFERENCES measurement,
    diameter_id INT REFERENCES measurement,
    height_id INT REFERENCES measurement,
    rim_depth_id INT REFERENCES measurement,
    speed SMALLINT,
    glide SMALLINT,
    turn SMALLINT,
    fade SMALLINT,
    stability SMALLINT,
    primary_use VARCHAR(50),
    plastic_grades VARCHAR(50)[],
    link VARCHAR(400),
    PRIMARY KEY (name, distributor)
);