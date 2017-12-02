CREATE DATABASE authenticator_development;
CREATE TABLE credentials (id SERIAL, username TEXT, password_hash TEXT, created_at TIMESTAMP, updated_at TIMESTAMP);

-- PASSWORD: 123456
INSERT INTO credentials (username, password_hash, created_at, updated_at) VALUES ('user', '$2a$04$eXMv4jini2kxwtnC1Xkdc.0SrIpBTQSLuKGneejVu5giRdpzhS/wy', NOW(), NOW());
