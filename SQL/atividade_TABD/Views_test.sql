-- Teste com as views
-- Fechar uma turma e verificar os status dos alunos
CALL FecharTurma(1);
SELECT * FROM Matriculas WHERE turmaid = 1;

-- Listar todas as turmas de um curso
SELECT * FROM ViewTurmasCursos WHERE cursoid = 1;

-- Listar alunos com direito a certificado
SELECT * FROM ViewCertificados;
