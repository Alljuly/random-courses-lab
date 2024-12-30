CREATE VIEW TurmasPorCurso AS
SELECT 
    t.turmaid,
    t.nome AS turma_nome,
    c.nome AS curso_nome,
    p.nome AS professor_nome,
    t.nivel
FROM Turmas t
JOIN Cursos c ON t.cursoid = c.cursoid
JOIN Professores p ON t.professorid = p.professorid;

SELECT * 
FROM TurmasPorCurso 
WHERE curso_nome = 'Nome do Curso';

CREATE VIEW AlunosPorTurma AS
SELECT 
    m.turmaid,
    t.nome AS turma_nome,
    a.nome AS aluno_nome,
    m.nota1,
    m.nota2,
    m.notafinal
FROM Matriculas m
JOIN Alunos a ON m.alunoid = a.alunoid
JOIN Turmas t ON m.turmaid = t.turmaid;

SELECT * 
FROM AlunosPorTurma 
WHERE turma_nome = 'Nome da Turma';

CREATE VIEW RegistrosAuditoria AS
SELECT 
    auditid,
    tabelaalterada,
    acao,
    dataalteracao
FROM Auditoria;

SELECT * 
FROM RegistrosAuditoria;

CREATE VIEW AlunosCertificados AS
SELECT 
    a.nome AS aluno_nome,
    c.nome AS curso_nome
FROM Alunos a
JOIN Matriculas m ON a.alunoid = m.alunoid
JOIN Turmas t ON m.turmaid = t.turmaid
JOIN Cursos c ON t.cursoid = c.cursoid
WHERE m.statusmatricula = 'Aprovado'
GROUP BY a.alunoid, c.cursoid
HAVING COUNT(DISTINCT t.nivel) = 3;

SELECT * 
FROM AlunosCertificados;

CREATE VIEW CursosPorAluno AS
SELECT 
    a.nome AS aluno_nome,
    c.nome AS curso_nome
FROM Alunos a
JOIN Matriculas m ON a.alunoid = m.alunoid
JOIN Turmas t ON m.turmaid = t.turmaid
JOIN Cursos c ON t.cursoid = c.cursoid;

SELECT * 
FROM CursosPorAluno 
WHERE aluno_nome = 'Nome do Aluno';



