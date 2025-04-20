import seaborn as sns
import matplotlib.pyplot as plt

import pandas as pd

csv_heart = pd.read_csv("heart.csv", sep=",")

sns.histplot(csv_heart['Age'], bins=5, color='blue')  
plt.xlabel('Idade')
plt.ylabel('Frequência')
plt.title('Histograma de Idade (Seaborn)')
plt.savefig('Histograma de Idade (Seaborn)', dpi=300)
plt.show()

sns.histplot(csv_heart['ChestPain'], bins=3, kde=False, color='green')
plt.xlabel('Tipo de Dor no Peito')
plt.ylabel('Frequência')
plt.title('Histograma de ChestPain (Seaborn)')
plt.savefig('Histograma de ChestPain (Seaborn)', dpi = 300)
plt.show()

sns.countplot(data=csv_heart, x='Sex', palette='Set2')
plt.xlabel('Sexo')
plt.ylabel('Frequência')
plt.title('Distribuição de Sexo (Seaborn)')
plt.savefig('Distribuição de Sexo (Seaborn)', dpi=300)
plt.show()