package Lista07.students;

import java.util.ArrayList;
import java.util.Scanner;
import java.util.stream.Collectors;

public class StudentManager {
    ArrayList<Student> students = new ArrayList<>();
    Scanner e = new Scanner(System.in);


    public void addNewStudent(){
        System.out.print("Nome: ");
        String name = e.nextLine();
        System.out.print("Número de Matrícula: ");
        String entry = e.nextLine();
        System.out.print("Endereço: ");
        String addres = e.nextLine();
        System.out.print("Nota 1: ");
        double[] n = new double[2];
        n[0] = e.nextDouble();
        System.out.print("Nota 2: ");
        n[1] = e.nextDouble();
        Student e = new Student(name, entry, addres, n);
        students.add(e);
    }
    public void avgStudent(){
        System.out.println("Digite a matricula do estudante:");
        String aux = e.nextLine();

        boolean contain = false;
        for (Student s : students){
            if (s.getEntry().equals(aux)) {
                contain = true;
                System.out.println("A média desse aluno é: " +s.studentAVGNote());
            }

        }
        if(!contain){
            System.out.println("Estudante com matrícula " + aux + " não encontrado.");
        }
    }
     public void getStudentByEntry(){
         System.out.println("Digite a matricula do estudante:");
         String aux = e.nextLine();

         boolean contain = false;
         for (Student s : students){
             if (s.getEntry().equals(aux)) {
                 contain = true;
                 System.out.println(s.toString());
             }
         }
         if(!contain){
             System.out.println("Estudante com matrícula " + aux + " não encontrado.");
         }
     }
     public void getStudentsByAddres(){
         System.out.println("Digite o endereço do estudante:");
         String aux = e.nextLine();

         boolean contain = false;
         for (Student s : students){
             if (s.getAddres().equals(aux)) {
                 contain = true;
                 System.out.println(s.toString());
             }
         }
         if(!contain){
             System.out.println("Não foram encontrados alunos cadastrados com o endereço " + aux);
         }

     }
     public void avgStudents(){
        double avg = students.stream().mapToDouble(Student::studentAVGNote).average().orElse(0.0);

         System.out.println("\nMÉDIA DA TURMA\n"+avg);
     }
     public void getStudentsReproved(){
        ArrayList<Student> reproved =  new ArrayList<>();

        for(Student s : students){
            if(s.getSituation().equals("Reproved")){
                reproved.add(s);
            }
        }

        if(!reproved.isEmpty()){
            System.out.println("___________ REPROVADOS ___________");
            String names = reproved.stream().map(Student::getName).collect(Collectors.joining(", "));
            System.out.println(names);
        } else{
            System.out.println("___________ REPROVADOS ___________\nNão há alunos reprovados nessa turma.");

        }

     }


}
