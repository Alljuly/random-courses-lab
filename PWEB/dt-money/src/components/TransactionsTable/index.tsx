export function TransactionsTable() {
  const transactions = [
    {
      id: 1,
      title: "Desenvolvimento de website",
      amount: 12000,
      type: "income",
      category: "Venda",
      date: "13/04/2022",
    },
    {
      id: 2,
      title: "Hamburguer",
      amount: 59,
      type: "outcome",
      category: "Alimentação",
      date: "10/04/2022",
    },
    {
      id: 3,
      title: "Aluguel do apartamento",
      amount: 1200,
      type: "outcome",
      category: "Casa",
      date: "27/03/2022",
    },
    {
      id: 4,
      title: "Computador",
      amount: 5400,
      type: "income",
      category: "Venda",
      date: "15/03/2022",
    },
  ];

  return (
    <div className="flex flex-col gap-4 justify-between my-4">
      <div className="grid grid-cols-4 gap-4 font-thin text-[#969CB2]">
        <div>Título</div>
        <div className="text-right">Preço</div>
        <div className="text-right">Categoria</div>
        <div className="text-right">Data</div>
      </div>

      {transactions.map((transaction) => (
        <div
          key={transaction.id}
          className="grid grid-cols-4 gap-4 bg-white rounded-md p-4 flex-wrap"
        >
          <div>{transaction.title}</div>
          <div
            className={`text-right ${
              transaction.type === "income"
                ? "text-[#33CC95]"
                : "text-[#E52E4D]"
            }`}
          >
            {transaction.type === "outcome" ? "-" : ""} R${" "}
            {transaction.amount.toFixed(2)}
          </div>
          <div className="text-right">{transaction.category}</div>
          <div className="text-right">{transaction.date}</div>
        </div>
      ))}
    </div>
  );
}
