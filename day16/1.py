import numpy as np

bin_str = ''

hex_values = {
    '0' : '0000',
    '1' : '0001',
    '2' : '0010',
    '3' : '0011',
    '4' : '0100',
    '5' : '0101',
    '6' : '0110',
    '7' : '0111',
    '8' : '1000',
    '9' : '1001',
    'A' : '1010',
    'B' : '1011',
    'C' : '1100',
    'D' : '1101',
    'E' : '1110',
    'F' : '1111'
}

with open('data.txt') as f:
    for line in f:
        #Convert the line into binary
        hex_repr = line.strip()
        for letter in hex_repr:
            bin_str += hex_values[letter]

def read_next_chunk(bin_repr):
    if int(bin_repr, 2) == 0:
        return '',0
    ver = bin_repr[:3]
    bin_repr = bin_repr[3:]
    typing = bin_repr[:3]
    bin_repr = bin_repr[3:]
    subpacket_v = 0    
    numbers = []

    if int(typing, 2) != 4:
        lengthId = bin_repr[:1]
        bin_repr = bin_repr[1:]
        if int(lengthId, 2) == 0:
            subpackets_length = int(bin_repr[:15],2)
            bin_repr = bin_repr[15:]
            subpackets = bin_repr[:subpackets_length]            
            while len(subpackets) > 0:
                subpackets, sb_ver, sb_number = read_next_chunk(subpackets)
                subpacket_v += sb_ver
                numbers.append(sb_number)
            bin_repr = bin_repr[subpackets_length:]
        else:
            subpackets_amount = int(bin_repr[:11],2)
            bin_repr = bin_repr[11:]
            for i in range(subpackets_amount):
                bin_repr, sb_ver, sb_number = read_next_chunk(bin_repr)
                subpacket_v += sb_ver
                numbers.append(sb_number)

    
    # Check a literal number and return it along with the number and version sums
    if int(typing, 2) == 4:
        number = ''      
        while True:
            bit = bin_repr[:5]
            bin_repr = bin_repr[5:]
            number += bit[1:]
            if bit[0] == '0':
                break
        return bin_repr, int(ver,2) + subpacket_v, int(number, 2)
    
    #Return appropriate product of sum list depending of the type

    if int(typing, 2) == 0:
        return bin_repr, int(ver,2) + subpacket_v, sum(numbers)

    if int(typing, 2) == 1:
        return bin_repr, int(ver,2) + subpacket_v, np.prod(numbers)

    if int(typing, 2) == 2:
        return bin_repr, int(ver,2) + subpacket_v, min(numbers)

    if int(typing, 2) == 3:
        return bin_repr, int(ver,2) + subpacket_v, max(numbers)

    if int(typing, 2) == 5:
        return bin_repr, int(ver,2) + subpacket_v, int(numbers[0] > numbers[1])

    if int(typing, 2) == 6:
        return bin_repr, int(ver,2) + subpacket_v, int(numbers[0] < numbers[1])
    
    if int(typing, 2) == 7:
        return bin_repr, int(ver,2) + subpacket_v, int(numbers[0] == numbers[1])

    

total = 0
bin_str, ver, value = read_next_chunk(bin_str)
total += ver

print('Task 1:')
print(ver)
print('Task 2:')
print(value)