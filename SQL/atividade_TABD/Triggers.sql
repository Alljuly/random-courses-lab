-- Triggers Aluno
DELIMITER $$

-- Trigger para impedir CPF duplicado no INSERT
CREATE TRIGGER verificar_cpf_antes_inserir
BEFORE INSERT ON Alunos
FOR EACH ROW
BEGIN
    IF EXISTS (
        SELECT 1 FROM Alunos WHERE cpf = NEW.cpf
    ) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Erro: CPF já está cadastrado!';
    END IF;
END$$

-- Trigger para impedir CPF duplicado no UPDATE
CREATE TRIGGER verificar_cpf_antes_atualizar
BEFORE UPDATE ON Alunos
FOR EACH ROW
BEGIN
    IF EXISTS (
        SELECT 1 FROM Alunos WHERE cpf = NEW.cpf AND alunoid != NEW.alunoid
    ) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Erro: CPF já está cadastrado!';
    END IF;
END$$

DELIMITER ;

-- Triggers Professor

DELIMITER $$

-- Trigger para impedir CPF duplicado no INSERT em Professores
CREATE TRIGGER verificar_cpf_professor_inserir
BEFORE INSERT ON Professores
FOR EACH ROW
BEGIN
    IF EXISTS (
        SELECT 1 FROM Professores WHERE cpf = NEW.cpf
    ) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Erro: CPF já está cadastrado para outro professor!';
    END IF;
END$$

-- Trigger para impedir CPF duplicado no UPDATE em Professores
CREATE TRIGGER verificar_cpf_professor_atualizar
BEFORE UPDATE ON Professores
FOR EACH ROW
BEGIN
    IF EXISTS (
        SELECT 1 FROM Professores WHERE cpf = NEW.cpf AND professorid != NEW.professorid
    ) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Erro: CPF já está cadastrado para outro professor!';
    END IF;
END$$

DELIMITER ;


-- Triggers Curso 

DELIMITER $$

-- Trigger para impedir nome duplicado no INSERT em Cursos
CREATE TRIGGER verificar_nome_curso_inserir
BEFORE INSERT ON Cursos
FOR EACH ROW
BEGIN
    IF EXISTS (
        SELECT 1 FROM Cursos WHERE nome = NEW.nome
    ) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Erro: Já existe um curso com este nome!';
    END IF;
END$$

-- Trigger para impedir nome duplicado no UPDATE em Cursos
CREATE TRIGGER verificar_nome_curso_atualizar
BEFORE UPDATE ON Cursos
FOR EACH ROW
BEGIN
    IF EXISTS (
        SELECT 1 FROM Cursos WHERE nome = NEW.nome AND cursoid != NEW.cursoid
    ) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Erro: Já existe um curso com este nome!';
    END IF;
END$$

DELIMITER ;

-- Triggers Matricula
-- a. Não permitir que o aluno seja matriculado 
-- em uma turma que não tenha vaga
DELIMITER $$

CREATE TRIGGER verificar_vagas_antes_matricular
BEFORE INSERT ON Matriculas
FOR EACH ROW
BEGIN
    DECLARE vagas_disponiveis INT;
    SELECT vagas INTO vagas_disponiveis FROM Turmas WHERE turmaid = NEW.turmaid;

    IF vagas_disponiveis <= 0 THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Erro: A turma não possui vagas disponíveis!';
    END IF;
END$$

DELIMITER ;

-- b. Não permitir que o aluno seja matriculado em um curso/nível
-- que ele já tenha sido aprovado

DELIMITER $$

CREATE TRIGGER verificar_aprovacao_anterior
BEFORE INSERT ON Matriculas
FOR EACH ROW
BEGIN
    IF EXISTS (
        SELECT 1 
        FROM Matriculas AS m
        INNER JOIN Turmas AS t ON m.turmaid = t.turmaid
        WHERE m.alunoid = NEW.alunoid 
          AND t.cursoid = (SELECT cursoid FROM Turmas WHERE turmaid = NEW.turmaid)
          AND t.nivel = (SELECT nivel FROM Turmas WHERE turmaid = NEW.turmaid)
          AND m.statusmatricula = 'Aprovado'
    ) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Erro: O aluno já foi aprovado neste curso e nível!';
    END IF;
END$$

DELIMITER ;


-- c. Atualizar a quantidade de vagas 
-- disponíveis ao matricular
DELIMITER $$

CREATE TRIGGER atualizar_vagas_ao_matricular
AFTER INSERT ON Matriculas
FOR EACH ROW
BEGIN
    UPDATE Turmas
    SET vagas = vagas - 1
    WHERE turmaid = NEW.turmaid;
END$$

DELIMITER ;
-- d. Não permitir matricular aluno em turma fechada
DELIMITER $$

CREATE TRIGGER verificar_status_turma_antes_matricular
BEFORE INSERT ON Matriculas
FOR EACH ROW
BEGIN
    DECLARE status_turma ENUM('Aberta', 'Fechada');
    SELECT status INTO status_turma FROM Turmas WHERE turmaid = NEW.turmaid;

    IF status_turma = 'Fechada' THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Erro: Não é possível matricular o aluno em uma turma fechada!';
    END IF;
END$$

DELIMITER ;

-- e. Não permitir matricular o aluno se ele não foi aprovado em nível anterior

DELIMITER $$

CREATE TRIGGER verificar_nivel_anterior
BEFORE INSERT ON Matriculas
FOR EACH ROW
BEGIN
    DECLARE nivel_atual ENUM('Básico', 'Intermediário', 'Avançado');
    DECLARE cursoid_atual INT;
    DECLARE aprovado_anterior BOOLEAN;

    SELECT nivel, cursoid INTO nivel_atual, cursoid_atual 
    FROM Turmas 
    WHERE turmaid = NEW.turmaid;

    IF nivel_atual != 'Básico' THEN
        SELECT EXISTS (
            SELECT 1 
            FROM Matriculas AS m
            INNER JOIN Turmas AS t ON m.turmaid = t.turmaid
            WHERE m.alunoid = NEW.alunoid 
              AND t.cursoid = cursoid_atual
              AND (t.nivel = 
                  CASE nivel_atual
                      WHEN 'Intermediário' THEN 'Básico'
                      WHEN 'Avançado' THEN 'Intermediário'
                  END)
              AND m.statusmatricula = 'Aprovado'
        ) INTO aprovado_anterior;

        IF aprovado_anterior = FALSE THEN
            SIGNAL SQLSTATE '45000'
            SET MESSAGE_TEXT = 'Erro: O aluno não foi aprovado no nível anterior!';
        END IF;
    END IF;
END$$

DELIMITER ;

-- Calcular Media do aluno

-- a. Calcular a média das notas ao alterar uma das notas

DELIMITER $$

CREATE TRIGGER calcular_media_ao_alterar_notas
AFTER UPDATE ON Matriculas
FOR EACH ROW
BEGIN
    -- Calcula a média das notas
    IF NEW.nota1 IS NOT NULL AND NEW.nota2 IS NOT NULL THEN
        UPDATE Matriculas
        SET notafinal = (NEW.nota1 + NEW.nota2) / 2
        WHERE matriculaid = NEW.matriculaid;
    END IF;
END$$

DELIMITER ;

-- b. Não permitir alterar notas com valores maiores que 10 ou menores que 0

DELIMITER $$

CREATE TRIGGER validar_notas_ao_atualizar
BEFORE UPDATE ON Matriculas
FOR EACH ROW
BEGIN
    -- Valida a nota1
    IF NEW.nota1 IS NOT NULL AND (NEW.nota1 < 0 OR NEW.nota1 > 10) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Erro: A nota1 deve estar entre 0 e 10!';
    END IF;

    -- Valida a nota2
    IF NEW.nota2 IS NOT NULL AND (NEW.nota2 < 0 OR NEW.nota2 > 10) THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Erro: A nota2 deve estar entre 0 e 10!';
    END IF;
END$$

DELIMITER ;

-- Fechar turma
-- a. Atualizar o status de cada aluno baseado na média

DELIMITER $$

CREATE TRIGGER atualizar_status_alunos_ao_fechar_turma
AFTER UPDATE ON Turmas
FOR EACH ROW
BEGIN
    -- Verifica se o status da turma foi alterado para 'Fechada'
    IF NEW.status = 'Fechada' AND OLD.status != 'Fechada' THEN
        -- Atualiza o status de matrícula dos alunos com base na nota final
        UPDATE Matriculas
        SET statusmatricula = CASE
            WHEN notafinal >= 7 THEN 'Aprovado'
            ELSE 'Reprovado'
        END
        WHERE turmaid = NEW.turmaid;
    END IF;
END$$

DELIMITER ;

-- Auditoria
-- a. Registrar mudanças baseadas em ações no banco de dados
DELIMITER $$

DELIMITER $$

CREATE TRIGGER registrar_auditoria
AFTER INSERT ON Matriculas
FOR EACH ROW
BEGIN
    INSERT INTO Auditoria (tabelaalterada, acao, dataalteracao)
    VALUES ('Matriculas', 'INSERIR', NOW());
END$$

CREATE TRIGGER registrar_auditoria_update
AFTER UPDATE ON Matriculas
FOR EACH ROW
BEGIN
    INSERT INTO Auditoria (tabelaalterada, acao, dataalteracao)
    VALUES ('Matriculas', 'ATUALIZAR', NOW());
END$$

CREATE TRIGGER registrar_auditoria_delete
AFTER DELETE ON Matriculas
FOR EACH ROW
BEGIN
    INSERT INTO Auditoria (tabelaalterada, acao, dataalteracao)
    VALUES ('Matriculas', 'EXCLUIR', NOW());
END$$

DELIMITER ;







