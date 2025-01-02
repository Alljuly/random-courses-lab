# Decisões
 _Justifique as suas decisões relacionadas à criação de índices, enums, views, stored
procedures, functions e triggers durante o processo de construção do banco de dados.
Explique como cada um desses elementos contribui para a eficiência, organização e
manutenção do sistema. Além disso, analise como essas escolhas impactam na segurança,
na clareza do modelo e na facilidade de futuras modificações ou expansões do banco de
dados._
---

#### **Justificativa:**
1. **Eficiência**: Índices reduzem o tempo de execução das consultas, especialmente ao lidar com grandes volumes de dados.
2. **Manutenção**: Facilitam a identificação rápida de registros sem a necessidade de varrer a tabela inteira.
3. **Segurança**: Melhoram a integridade referencial ao garantir que as consultas que dependem de chaves estrangeiras sejam mais rápidas e precisas.

---

### ENUMs
Foram utilizados  **ENUMs** para colunas com valores restritos e estáveis, a fim de reduzir erros de entrada e facilitar a validação e a consistência dos dados.
ex.:
- `nivel` (`Básico`, `Intermediário`, `Avançado`);
- `status` (`Aberta`, `Fechada`).

#### **Justificativa:**
1. **Clareza**: Torna o modelo mais compreensível, especificando valores permitidos diretamente no esquema.
2. **Segurança**: Previne valores inválidos ao limitar as entradas possíveis.
3. **Eficiência**: Ocupa menos espaço no banco de dados e oferece melhor desempenho em comparação com tabelas de referência separadas para pequenas listas de valores.

---

### Views
As **views** foram criadas para consultas recorrentes e complexas, como listar alunos por turma ou turmas por curso.

#### **Justificativa:**
1. **Organização**: Encapsulam consultas complexas, tornando o código SQL mais limpo e reutilizável.
2. **Eficiência**: Melhora a manutenção, já que as atualizações na lógica de consulta podem ser feitas diretamente na **view**.
3. **Segurança**: Restringem o acesso direto às tabelas, fornecendo apenas os dados necessários.

---

### Stored Procedures e Functions
As **stored procedures** usadas para operações complexas, como cálculos de médias, e **functions** para validações específicas (e.g., verificar se o aluno já foi aprovado em um nível anterior).

#### **Justificativa:**
1. **Eficiência**: Executam no servidor, reduzindo o tráfego entre cliente e servidor.
2. **Manutenção**: Centralizam a lógica de negócios, facilitando ajustes futuros.
3. **Segurança**: Restringem a manipulação de dados diretamente pelo cliente, reduzindo a chance de erros ou ataques.

---

### Triggers
Os **triggers** são usados para automatizar ações, como calcular médias, atualizar vagas de turmas, e registrar alterações no sistema.

#### **Justificativa:**
1. **Eficiência**: Automatizam processos que seriam repetitivos ou propensos a erro se feitos manualmente.
2. **Segurança**: Garantem que regras de negócio sejam aplicadas consistentemente, independentemente da origem da operação.
3. **Organização**: Mantêm a lógica de negócios próxima aos dados.

---


