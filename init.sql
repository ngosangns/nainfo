-- auth-service schema
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL
);

-- profile-service schema
CREATE TABLE profiles (
  id int AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  facebook VARCHAR(255) NOT NULL,
  linkedin VARCHAR(255) NOT NULL,
  github VARCHAR(255) NOT NULL
);