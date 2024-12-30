create database BaiaoWordsAcademy;

use BaiaoWordsAcademy;

CREATE TABLE Cursos (
    cursoid INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    descricao TEXT NOT NULL,
    duracao INT NOT NULL
);

CREATE TABLE Professores (
    professorid INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    cpf CHAR(11) UNIQUE NOT NULL,
    telefone VARCHAR(15)
);

CREATE TABLE Alunos (
    alunoid INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    telefone VARCHAR(15),
    endereco TEXT,
    cpf CHAR(11) UNIQUE NOT NULL
);

CREATE TABLE Turmas (
    turmaid INT AUTO_INCREMENT PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    cursoid INT NOT NULL,
    professorid INT NOT NULL,
    nivel ENUM('Básico', 'Intermediário', 'Avançado') NOT NULL,
    status ENUM('Aberta', 'Fechada') NOT NULL,
    vagas INT DEFAULT 30,
    FOREIGN KEY (cursoid) REFERENCES Cursos(cursoid),
    FOREIGN KEY (professorid) REFERENCES Professores(professorid)
);

CREATE TABLE Matriculas (
    matriculaid INT AUTO_INCREMENT PRIMARY KEY,
    alunoid INT NOT NULL,
    turmaid INT NOT NULL,
    datamatricula DATE NOT NULL,
    nota1 DECIMAL(5, 2),
    nota2 DECIMAL(5, 2),
    notafinal DECIMAL(5, 2) GENERATED ALWAYS AS ((nota1 + nota2) / 2) STORED,
    statusmatricula ENUM('Pendente', 'Aprovado', 'Reprovado') AS (
        CASE 
            WHEN notafinal >= 6 THEN 'Aprovado'
            WHEN notafinal < 6 AND notafinal IS NOT NULL THEN 'Reprovado'
            ELSE 'Pendente'
        END
    ) STORED,
    FOREIGN KEY (alunoid) REFERENCES Alunos(alunoid),
    FOREIGN KEY (turmaid) REFERENCES Turmas(turmaid)
);

CREATE TABLE Auditoria (
    auditid INT AUTO_INCREMENT PRIMARY KEY,
    tabelaalterada VARCHAR(50) NOT NULL,
    acao ENUM('INSERIR', 'ATUALIZAR', 'EXCLUIR', 'FECHAR_TURMA') NOT NULL,
    dataalteracao DATETIME DEFAULT CURRENT_TIMESTAMP
);

