import java.util.Scanner;

public class Aula1 {

    public static void main(String[] args) {
        Scanner e = new Scanner(System.in);
        int opcao;

        do {
            System.out.println("\nEscolha uma questão para executar:");
            System.out.println("1 - Verificar se um numero é par ou ímpar");
            System.out.println("2 - Calculadora básica");
            System.out.println("3 - Calcular imposto de salario");
            System.out.println("4 - Área e perimetro de um retângulo");
            System.out.println("5 - Converter temperatura");
            System.out.println("0 - Sair");
            opcao = e.nextInt();

            switch (opcao) {
                case 1:
                    System.out.print("digite um numero: ");
                    int numero = e.nextInt();
                    System.out.println(parOuImpar(numero));
                    break;

                case 2:
                    System.out.print("digite o primeiro numero: ");
                    double num1 = e.nextDouble();
                    System.out.print("digite o segundo numero: ");
                    double num2 = e.nextDouble();
                    System.out.println("Escolha a operação: 1-soma, 2-sub, 3-mult, 4-divisao");
                    int operacao = e.nextInt();
                    calculadora(num1, num2, operacao);
                    break;

                case 3:
                    System.out.print("digite o salario: ");
                    double salario = e.nextDouble();
                    System.out.println("Imposto a ser pago: R$ " + calcularImposto(salario));
                    break;

                case 4:
                    System.out.println("1 - Calcular área do retangulo");
                    System.out.println("2 - Calcular perimetro do retangulo");
                    int escolhaRetangulo = e.nextInt();
                    System.out.print("digite a base do retangulo: ");
                    double base = e.nextDouble();
                    System.out.print("digite a altura do retangulo: ");
                    double altura = e.nextDouble();
                    if (escolhaRetangulo == 1) {
                        System.out.println("Área do retangulo: " + calcularArea(base, altura));
                    } else if (escolhaRetangulo == 2) {
                        System.out.println("perimetro do retangulo: " + calcularP(base, altura));
                    }
                    break;

                case 5:
                    System.out.println("1 - Converter F para C");
                    System.out.println("2 - Converter K para C");
                    int escolhaTemp = e.nextInt();
                    if (escolhaTemp == 1) {
                        System.out.print("digite a temperatura em F: ");
                        double fahrenheit = e.nextDouble();
                        System.out.println("Temperatura em C: " + calcularFparaC(fahrenheit));
                    } else if (escolhaTemp == 2) {
                        System.out.print("digite a temperatura em K: ");
                        double kelvin = e.nextDouble();
                        System.out.println("Temperatura em C: " + calcularKparaC(kelvin));
                    }
                    break;

                case 0:
                    System.out.println("Encerrando...");
                    break;

                default:
                    System.out.println("Opção inválida!");
                    break;
            }
        } while (opcao != 0);

        e.close();
    }

    // Questão 1
    public static String parOuImpar(int numero) {
        if (numero % 2 == 0) {
            return "O numero é par.";
        } else {
            return "O numero é ímpar.";
        }
    }

    // Questão 2
    public static void calculadora(double num1, double num2, int operacao) {
        switch (operacao) {
            case 1:
                System.out.println("Resultado da soma: " + (num1 + num2));
                break;
            case 2:
                System.out.println("Resultado da sub: " + (num1 - num2));
                break;
            case 3:
                System.out.println("Resultado da mult: " + (num1 * num2));
                break;
            case 4:
                if (num2 != 0) {
                    System.out.println("Resultado da divisao: " + (num1 / num2));
                } else {
                    System.out.println("Erro: divisao por zero!");
                }
                break;
            default:
                System.out.println("Operação inválida.");
                break;
        }
    }

    // Questão 3
    public static double calcularImposto(double salario) {
        if (salario <= 2000) {
            return 0;
        } else if (salario <= 3500) {
            return salario * 0.15;
        } else if (salario <= 5000) {
            return salario * 0.22;
        } else {
            return salario * 0.30;
        }
    }

    // Questão 4
    public static double calcularArea(double base, double altura) {
        return base * altura;
    }


    public static double calcularP(double base, double altura) {
        return 2 * (base + altura);
    }

    // Questão 5
    public static double calcularFparaC(double fahrenheit) {
        return 5 * (fahrenheit - 32) / 9;
    }

    public static double calcularKparaC(double kelvin) {
        return kelvin - 273;
    }
}
