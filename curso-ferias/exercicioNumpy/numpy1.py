import numpy as np
from PIL import Image


# 1 Crie uma matriz 1D com números de 0 a 9
matriz = np.fromfunction(lambda i,: (i + 1), (9,), dtype=int)

print(matriz)

# 2 Crie uma matriz booleana numpy 3×3 com ‘True’

mbool = np.full((3,3),True, dtype=bool)

print(mbool)
# 3 Crie uma matriz 5 x 5 com números inteiros aleatórios e substitua os
# pares por -1

mRandom = np.random.randint(0,100,size =(5,5))

mRandom[mRandom %  2 == 0 ] = -1

print(mRandom)


# 4 Crie uma imagem com ruídos 256x256 usando apenas 1 canal

mruidos = np.random.randint(0,256, size=(256,256), dtype=np.uint8)

imagem = Image.fromarray(mruidos, 'L') 
imagem.save('ruidos.png')

imagem.show()
