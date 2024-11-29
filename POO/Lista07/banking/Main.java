package Lista07.banking;

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        Agency agency = new Agency();
        int option;

        do {
            System.out.println("\n----- MENU -----");
            System.out.println("1. Register Client");
            System.out.println("2. View Client Details");
            System.out.println("3. Check Balance");
            System.out.println("4. Withdraw from Account");
            System.out.println("5. Deposit to Account");
            System.out.println("0. Exit");
            System.out.print("Choose an option: ");
            option = sc.nextInt();

            switch (option) {
                case 1:
                    sc.nextLine(); // Clear buffer
                    System.out.print("Enter client name: ");
                    String name = sc.nextLine();

                    System.out.print("Enter account number: ");
                    int accountNumber = sc.nextInt();

                    System.out.print("Enter initial balance: ");
                    double balance = sc.nextDouble();

                    agency.registerClient(name, accountNumber, balance);
                    break;

                case 2:
                    System.out.print("Enter account number: ");
                    accountNumber = sc.nextInt();

                    Client client = agency.findClient(accountNumber);
                    if (client != null) {
                        client.displayDetails();
                    } else {
                        System.out.println("Account not found.");
                    }
                    break;

                case 3:
                    System.out.print("Enter account number: ");
                    accountNumber = sc.nextInt();
                    agency.checkBalance(accountNumber);
                    break;

                case 4:
                    System.out.print("Enter account number: ");
                    accountNumber = sc.nextInt();
                    System.out.print("Enter withdrawal amount: ");
                    double withdrawalAmount = sc.nextDouble();
                    agency.withdraw(accountNumber, withdrawalAmount);
                    break;

                case 5:
                    System.out.print("Enter account number: ");
                    accountNumber = sc.nextInt();
                    System.out.print("Enter deposit amount: ");
                    double depositAmount = sc.nextDouble();
                    agency.deposit(accountNumber, depositAmount);
                    break;

                case 0:
                    System.out.println("Exiting...");
                    break;

                default:
                    System.out.println("Invalid option! Please try again.");
            }
        } while (option != 0);

        sc.close();
    }
}
