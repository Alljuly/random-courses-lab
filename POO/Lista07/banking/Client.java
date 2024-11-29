package Lista07.banking;

public class Client {
    String name;
    int accountNumber;
    double balance;

    public Client(String name, int accountNumber, double balance) {
        this.name = name;
        this.accountNumber = accountNumber;
        this.balance = balance;
    }

    public void displayDetails() {
        System.out.println("Name: " + name);
        System.out.println("Account Number: " + accountNumber);
        System.out.printf("Balance: R$%.2f\n", balance);
    }
}
