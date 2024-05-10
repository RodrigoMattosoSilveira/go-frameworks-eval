DROP TABLE IF EXISTS measurements;
CREATE TABLE measurements (
  id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  user_id INTEGER NOT NULL,
  weight FLOAT NOT NULL,
  height FLOAT NOT NULL,
  body_fat FLOAT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  CONSTRAINT fk_category FOREIGN KEY (user_id) REFERENCES user(id)
);
