def my_function(*kids):
    # print("The youngest child is " + kids[2])
    for c in kids:
        print(f"Welcome {c}.")


my_function("Emil", "Tobias", "Linus")
