mylist = []

mylist.append(1)
mylist.append(2)
mylist.append(3)

print(mylist)

mylist = [87, 5, 2, 42, 5, 52]

mylist.sort()  # Organizar
print(mylist)

mylist.remove(87)
print(mylist)

mylist.reverse()
print(mylist)

mylist2 = [21, 4, 513, 51]

mylist.extend(mylist2)
print(mylist)

mylist3 = mylist + mylist2
mylist3.sort()
print(mylist3)

listOfWords = ["this", "is", "a", "list", "of", "words"]
items = [word[0] for word in listOfWords]
print(items)

fruits = [1, 4, 2, 9, 7, 8, 9, 3, 1]
x = fruits.count(9)
print(x)

fruits = [4, 55, 64, 32, 16, 32]
x = fruits.index(32)
print(x)
