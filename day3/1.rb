#read file into 2d array
data = File.readlines("data.txt").map do |line|
    line.strip.split("")
end

def find_binarystring(data)
    #flip the array
    data = data.transpose

    #Find the most common value (0 or 1) in every column (row now as we transposed the array)
    common_string = ''
    for i in 0...data.length
        column = data[i]
        ones, zeros = 0, 0
        for letter in column
            if letter == '1'
                ones += 1
            else
                zeros += 1
            end
        end
        if ones >= zeros
            common_value = '1'
        else
            common_value = '0'
        end
        common_string += common_value
    end

    #Basically get the opposite of the most common string (1 -> 0 and 0 -> 1)
    uncommon_string = ''
    for i in 0...common_string.length
        letter = common_string[i]
        if letter == '0'
            uncommon_string += '1'
        else
            uncommon_string += '0'
        end
    end

    return common_string, uncommon_string
end

value1, value2 = find_binarystring data

puts "Values multiplied: #{value1.to_i(2) * value2.to_i(2)}"