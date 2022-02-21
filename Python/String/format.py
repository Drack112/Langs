age = 36
txt = "My name is John, and I am {}"
print(txt.format(age))

txt = "bruh"
print(f"My name is John, and I am {txt}")

price = 49
txt = "The price is {:.2f} dollars"
print(txt.format(price))

quantity = 3
itemno = 567
price = 49
myorder = "I want {} pieces of item number {} for {:.2f} dollars."
print(myorder.format(quantity, itemno, price))

quantity = 3
itemno = 567
price = 49
myorder = "I want {0} pieces of item number {1} for {2:.2f} dollars."
print(myorder.format(quantity, itemno, price))

myorder = "I have a {carname}, it is a {model}."
print(myorder.format(carname="Ford", model="Mustang"))
