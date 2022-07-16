import hashlib

i = 1
secret = 'iwrupvqb'
i1 = 0

while True:
    test_str = secret + str(i)
    md5 = hashlib.md5(test_str.encode())

    hex = md5.hexdigest()

    print(hex)

    if(hex.startswith('00000') and i1 == 0):
        i1 = i

    if(hex.startswith('000000')):
        print('FOUND 5 zeroes -', i1)
        print('FOUND 6 zeroes -', i)
        break

    i += 1