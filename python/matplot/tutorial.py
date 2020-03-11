import matplotlib.pyplot as plt
import numpy as np

names = ['GroupA', 'GroupB', 'GroupC']
values = [1, 10, 100]

plt.figure(figsize=(9, 6))
plt.subplot(231)
plt.bar(names, values)
plt.subplot(232)
plt.scatter(names, values)
plt.subplot(233)
plt.plot(names, values)
plt.subplot(234)
plt.bar(names, values)
plt.subplot(235)
plt.scatter(names, values)
plt.subplot(236)
plt.plot(names, values)
plt.suptitle('Categorical Plotting')
plt.show()
