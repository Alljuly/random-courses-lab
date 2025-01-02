-- Teste de atualizações
-- Atualizar o telefone de um professor
UPDATE Professores SET telefone = '(11) 99999-9999' WHERE professorid = 1;
SELECT * FROM Professores WHERE professorid = 1;

-- Alterar a nota de um aluno e verificar se a média foi recalculada
UPDATE Matriculas SET nota1 = 9.0 WHERE matriculaid = 1;
SELECT * FROM Matriculas WHERE matriculaid = 1;

-- Excluir um aluno e verificar se ele é removido
DELETE FROM Alunos WHERE alunoid = 4;
SELECT * FROM Alunos;

-- Testar a matrícula de um aluno
CALL MatricularAluno(3, 1);

-- Verificar se as vagas foram atualizadas
SELECT * FROM Turmas WHERE turmaid = 1;

-- Consultar as alterações feitas no sistema
SELECT * FROM Auditoria;
