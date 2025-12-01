CREATE DATABASE IF NOT EXISTS devbook;
USE devbook;

DROP TABLE If EXISTS publicacoes;
DROP TABLE If EXISTS seguidores;
DROP TABLE If EXISTS usuarios;

CREATE TABLE usuarios(
    id INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(100) NOT NULL,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP()
)ENGINE=INNODB;

CREATE TABLE seguidores(
    usuario_id INT NOT NULL,
    FOREIGN KEY (usuario_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    seguidor_id INT NOT NULL,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    PRIMARY KEY(usuario_id, seguidor_id)
)ENGINE=INNODB;


CREATE TABLE publicacoes(
    id int AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(100) NOT NULL,
    conteudo VARCHAR(300) NOT NULL,
    curtidas int DEFAULT 0,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    autor_id int NOT NULL,
    FOREIGN KEY (autor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE
)ENGINE=INNODB;