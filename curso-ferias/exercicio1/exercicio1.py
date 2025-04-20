
import numpy as np

from PIL import Image

a = np.full((28, 28, 3), (0, 255, 0), dtype=np.uint8)

img = Image.fromarray(a, "RGB")

img.save("green.png", "PNG")