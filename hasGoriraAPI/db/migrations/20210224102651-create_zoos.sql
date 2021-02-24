
-- +migrate Up
CREATE TABLE IF NOT EXISTS zoos (
    id int(15) AUTO_INCREMENT,
    pref_id int(15),
    name varchar(255),
    hasGorira tinyint(1),
    PRIMARY KEY (id),
    FOREIGN KEY (pref_id) REFERENCES prefectures(id)
);

-- +migrate Down
DROP TABLE IF EXISTS zoos;