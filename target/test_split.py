test_file = open("test", 'r')

name_iterator = 0
splited_test = open("./tests/test" + str(name_iterator), "w+")

for line in test_file:
    if line != "$\n":
        splited_test.write(line)
    else:
        name_iterator += 1
        splited_test.close()
        splited_test = open("./tests/test" + str(name_iterator), "w+")

