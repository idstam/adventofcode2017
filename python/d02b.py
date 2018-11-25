
with open('data.txt') as lines:
    check = 0
    for line in lines:

        min = 0
        max = 0
        onFirst = True

        for c in line.split('\t'):
            n = int(c)
            if onFirst:
                onFirst = False
                min = n
                max = n
            else:
                if min > n:
                    min = n
                if max < n:
                    max = n

        check += (max - min)

print(check)