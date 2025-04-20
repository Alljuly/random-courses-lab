import pandas as pd

csv_heart = pd.read_csv("heart.csv", sep=",")

import matplotlib.pyplot as plt

plt.hist(csv_heart['Age'], bins=5, edgecolor='black')
plt.xlabel('Idade')
plt.ylabel('Frequência')
plt.title('Histograma de Idade (matplotlib)')
plt.savefig('Histograma de Idade (matplotlib)', dpi=300)
plt.show()

plt.hist(csv_heart['ChestPain'], bins=5, edgecolor='black')
plt.xlabel('Dor no peito')
plt.ylabel('Frequência')
plt.title('Histograma tipo de Dor no peito (matplotlib)')
plt.savefig('Histograma tipo de Dor no peito (matplotlib)', dpi=300)
plt.show()

plt.hist(csv_heart['Sex'], bins=5, edgecolor='black')
plt.xlabel('Sexo')
plt.ylabel('Frequência')
plt.title('Histograma Sexo')
plt.savefig('Histograma Sexo (matplotlib)', dpi=300)
plt.show()

