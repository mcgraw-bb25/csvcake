import csv

filename = "./test2.csv"

if __name__ == "__main__":
    
    with open(filename, "r") as csvin:
        reader = csv.DictReader(csvin)
        age = sum([int(row["age"]) for row in reader])

    print("Total age: {}".format(age))
