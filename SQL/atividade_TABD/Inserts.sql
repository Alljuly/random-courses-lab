-- Insert de dados fake pro banco
use baiaowordsacademy;

-- Cursos
INSERT INTO Cursos (cursoid, nome, descricao, duracao) VALUES
(1, 'Sistemas de informação', 'Curso de sistemas e informações', 8),
(2, 'Administração', 'Curso de gestão', 6),
(3, 'Letras', 'Curso de comunicação, e letras', 4);

-- Professores
INSERT INTO Professores (professorid, nome, email, cpf, telefone) VALUES
(1, 'Cleberson', 'cbles@escola.com', '11122233344', '(11) 98765-4321'),
(2, 'Luiza', 'luluiza@escola.com', '22233344455', '(21) 99876-5432'),
(3, 'Reppi', 'Reppi@escola.com', '33344455566', '(31) 91234-5678');

-- Alunos
INSERT INTO Alunos (alunoid, nome, email, telefone, endereco, cpf) VALUES
(1, 'Carlos', 'carlos.eduardo@aluno.com', '(41) 98765-1234', 'Rua A, 123', '44455566677'),
(2, 'Fernanda', 'fernanda.lima@aluno.com', '(51) 91234-5678', 'Rua B, 456', '55566677788'),
(3, 'Lucas', 'lucas.pereira@aluno.com', '(61) 99876-5432', 'Rua C, 789', '66677788899');

-- Turmas
INSERT INTO Turmas (turmaid, nome, cursoid, professorid, nivel, status, vagas) VALUES
(1, 'Turma A - SI', 1, 1, 'Básico', 'Aberta', 30),
(2, 'Turma B - ADM', 2, 2, 'Intermediário', 'Aberta', 25),
(3, 'Turma C - Letras', 3, 3, 'Avançado', 'Fechada', 0);

-- Matriculas
INSERT INTO Matriculas (matriculaid, alunoid, turmaid, datamatricula, nota1, nota2, notafinal, statusmatricula) VALUES
(1, 1, 1, '2024-01-15', 8.0, 9.0, 'Aprovado'),
(2, 2, 2, '2024-01-20', 7.0, 6.5, 'Reprovado'),
(3, 3, 3, '2024-01-25', NULL, NULL, 'Pendente');

-- Auditoria
INSERT INTO Auditoria (auditid, tabelaalterada, acao, dataalteracao) VALUES
(1, 'Alunos', 'INSERIR', '2024-01-15 10:30:00'),
(2, 'Turmas', 'ATUALIZAR', '2024-01-20 14:45:00'),
(3, 'Matriculas', 'EXCLUIR', '2024-01-25 16:00:00');
