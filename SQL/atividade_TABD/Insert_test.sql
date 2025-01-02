-- Teste de inserção
-- 1 - novo curso 
-- 2 - verificar se foi adicionado
INSERT INTO Cursos (cursoid, nome, descricao, duracao) VALUES (4, 'Direito', 'Legislações, Judiciário e demais ações civis', 4);
SELECT * FROM Cursos;

-- 1 - novo professor
-- 2 - verificar se foi adicionado
INSERT INTO Professores (professorid, nome, email, cpf, telefone) VALUES (4, 'Pedro Silva', 'pedro.silva@escola.com', '77788899900', '(11) 96543-2100');
SELECT * FROM Professores;

-- 1 - novo aluno
-- 2 - e verificar se foi adicionado
INSERT INTO Alunos (alunoid, nome, email, telefone, endereco, cpf) VALUES (4, 'Mariana Dias', 'mariana.dias@aluno.com', '(71) 98765-4321', 'Rua D, 101', '88899900011');
SELECT * FROM Alunos;