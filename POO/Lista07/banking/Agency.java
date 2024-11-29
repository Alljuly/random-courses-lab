package Lista07.banking;

import java.util.ArrayList;

public class Agency {
    private ArrayList<Client> clients;

    public Agency() {
        this.clients = new ArrayList<>();
    }

    public void registerClient(String name, int accountNumber, double balance) {
        Client newClient = new Client(name, accountNumber, balance);
        clients.add(newClient);
        System.out.println("Client successfully registered!");
    }

    public Client findClient(int accountNumber) {
        for (Client client : clients) {
            if (client.accountNumber == accountNumber) {
                return client;
            }
        }
        return null; // Client not found
    }

    public void checkBalance(int accountNumber) {
        Client client = findClient(accountNumber);
        if (client != null) {
            client.displayDetails();
        } else {
            System.out.println("Account not found.");
        }
    }

    public void withdraw(int accountNumber, double amount) {
        Client client = findClient(accountNumber);
        if (client != null) {
            if (amount > client.balance) {
                System.out.println("Insufficient balance!");
            } else {
                client.balance -= amount;
                System.out.printf("Withdrawal successful! New balance: R$%.2f\n", client.balance);
            }
        } else {
            System.out.println("Account not found.");
        }
    }

    public void deposit(int accountNumber, double amount) {
        Client client = findClient(accountNumber);
        if (client != null) {
            client.balance += amount;
            System.out.printf("Deposit successful! New balance: R$%.2f\n", client.balance);
        } else {
            System.out.println("Account not found.");
        }
    }
}
