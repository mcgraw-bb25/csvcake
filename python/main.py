import csv

filename = "../test.csv"

if __name__ == "__main__":
    
    with open(filename, "r") as csvin:
        reader = csv.DictReader(csvin)
        rows = [row for row in reader]

    age = sum([int(row["age"]) for row in rows])
    print("Total age: {}".format(age))
