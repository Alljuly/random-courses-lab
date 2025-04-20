import pandas as pd
import matplotlib.pyplot as plt


#(a) Leia o arquivo
df = pd.read_csv('diabetes.csv', sep=",") 


print(df.head())
#Pregnancies  Glucose  BloodPressure  SkinThickness  Insulin   
# BMI  DiabetesPedigreeFunction  Age  Outcome


#(b) Verfique os valores ausentes
print(df.isna())  # Não apontou valores ausentes
print(df.isna().sum())  # 0 valores ausentes por coluna


#(c) Selecione um subconjunto de colunnas

subConjunto = df[["BloodPressure","Age","Insulin"]]

print(subConjunto)



#(d) Plot o histograma de pelo menos duas colunas de sua escolha
#* Possível solução


col1 = df[["Age"]].value_counts().sort_index()

y = col1.index.get_level_values(0)

plt.figure(figsize=(20,12))
plt.bar(y.astype(str), col1.values)
plt.xlabel('Idade')
plt.ylabel('BMI')
plt.title('BMI por idade')
plt.legend()
plt.savefig('BMI-Idade', dpi=400)
plt.show()
