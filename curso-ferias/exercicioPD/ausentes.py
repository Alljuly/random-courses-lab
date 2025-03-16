import pandas as pd

# 1 Usando o dataset heart.csv
# 2 Leia o arquivo
csv_heart = pd.read_csv("heart.csv", sep=",")

print(csv_heart.head())
## Age  Sex     ChestPain  RestBP  Chol  Fbs  RestECG  MaxHR  
# ExAng  Oldpeak  Slope   Ca        Thal  AHD

# 3 Verifique valores ausentes
print(csv_heart.isna().sum()) 