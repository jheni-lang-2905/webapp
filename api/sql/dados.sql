INSERT INTO usuarios(nome, nick, email, senha)
VALUES
('user1', 'user1', 'user1@gmail.com', '$2a$10$qYCsJx63gYciqoiFDyFJMudgQFTAarVLX8RGRFR7sTJ4hPigjIhyu'),
('user2', 'user2', 'user2@gmail.com', '$2a$10$qYCsJx63gYciqoiFDyFJMudgQFTAarVLX8RGRFR7sTJ4hPigjIhyu'),
('user3', 'user3', 'user3@gmail.com', '$2a$10$qYCsJx63gYciqoiFDyFJMudgQFTAarVLX8RGRFR7sTJ4hPigjIhyu'),
('user4', 'user4', 'user4@gmail.com', '$2a$10$qYCsJx63gYciqoiFDyFJMudgQFTAarVLX8RGRFR7sTJ4hPigjIhyu'),
('user5', 'user5', 'user5@gmail.com', '$2a$10$qYCsJx63gYciqoiFDyFJMudgQFTAarVLX8RGRFR7sTJ4hPigjIhyu');


INSERT INTO seguidores(usuario_id, seguidor_id)
VALUES
(5, 7);

INSERT INTO publicacoes(titulo, conteudo, autor_id)
VALUES
("publicacao1", "conteudo1", 1),
("publicacao2", "conteudo2", 2),
("publicacao3", "conteudo3", 3),
("publicacao4", "conteudo4", 4);