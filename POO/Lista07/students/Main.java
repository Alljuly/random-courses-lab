package Lista07.students;

import java.util.Scanner;

public class Main {
    public static void main(String[] args){
        StudentManager students = new StudentManager();
        Scanner e = new Scanner(System.in);

        int act;

        do {
                    System.out.println("\n----- Gerenciador de Estudantes -----");
                    System.out.println("1. Cadastrar um estudante");
                    System.out.println("2. Calcular média de um estudante");
                    System.out.println("3. Consultar matrícula de um estudante");
                    System.out.println("4. Consultar endereço de um estudante");
                    System.out.println("5. Calcular a média da turma");
                    System.out.println("6. Exibir os nomes dos estudantes reprovados");
                    System.out.println("0. Sair");
                    System.out.print("Escolha uma opção: ");
                    act = e.nextInt();
                    e.nextLine();

                    switch (act) {
                        case 1:
                            students.addNewStudent();
                            break;
                        case 2:
                            students.avgStudent();
                            break;
                        case 3:
                            students.getStudentByEntry();
                            break;
                        case 4:
                            students.getStudentsByAddres();
                            break;
                        case 5:
                            students.avgStudents();
                            break;
                        case 6: students.getStudentsReproved();
                            break;

                        default:
                            System.out.println("Opção inválida! Tente novamente.");
                            break;
                    }
                } while (act != 0);

    }
}
