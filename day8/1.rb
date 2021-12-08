require 'set'

numbers = []
outputs = []

def find_diff(str1, str2)
    a, b = str1.split(''), str2.split('')
    diff = a - b | b - a 
    return diff
end

def solve_display(display, output)
    definite_lengths = {2 => 1, 3 => 7, 4 => 4, 7 => 8}
    definite_numbers = Hash.new()
    display_values = ['.'] * 7

    #Find easy numbers
    for number in display
        if definite_lengths.key?(number.size)
            definite_numbers[definite_lengths[number.size]] = number
        end
    end

    #delete the values we solved this way
    definite_numbers.each_value do |value|
        display.delete(value)
    end    

    #7 and 1 have a difference of only one line so we can deduce the top line from this
    diff = find_diff definite_numbers[7], definite_numbers[1]
    display_values[0] = diff[0]

    #Now that we know the top line we can find the 9 as it is all the letters of 4 + top + one extra (bottom)
    for number in display
        missing1 = display_values[0] + definite_numbers[4]
        diff = find_diff missing1, number
        #check there's only one missing value
        if diff.size == 1
            display_values[6] = diff[0]
            definite_numbers[9] = number
            display.delete(number) #Delete the number so we don't have to loop over it again in the future
            break
        end
    end

    #Now that we know the 9 we can figure bottom left from difference of 9 and 8
    diff = find_diff definite_numbers[9], definite_numbers[8]
    display_values[4] = diff[0]

    #We can find 3 from the missing letter of 7 + bottom row
    for number in display
        missing1 = display_values[6] + definite_numbers[7]
        diff = find_diff missing1, number
        #check there's only one missing value
        if diff.size == 1
            display_values[3] = diff[0]
            definite_numbers[3] = number
            display.delete(number) #Delete the number so we don't have to loop over it again in the future
            break
        end
    end

    #We know 4 letters of 2 so we can now know for certain the top right (and find 2 at the same time)
    for number in display
        missing1 = display_values[0] + display_values[3] + display_values[4] + display_values[6]
        diff = find_diff missing1, number
        #check there's only one missing value
        if diff.size == 1
            display_values[2] = diff[0]
            definite_numbers[2] = number
            display.delete(number) #Delete the number so we don't have to loop over it again in the future
            break
        end
    end

    #Figure the missing letters from 1 and 8
    diff = find_diff definite_numbers[1], display_values[2]
    display_values[5] = diff[0]

    diff = find_diff definite_numbers[8], display_values.join('')    
    display_values[1] = diff[0]

    #add missing values to definite numbers
    definite_numbers[5] = display_values[0] + display_values[1] + display_values[3] + display_values[5] + display_values[6]
    definite_numbers[0] = display_values[0] + display_values[1] + display_values[2] + display_values[4] + display_values[5] + display_values[6]
    definite_numbers[6] = display_values[0] + display_values[1] + display_values[3] + display_values[5] + display_values[6] + display_values[4]
    
    return definite_numbers
end

def sum_easy(outputs)
    definite_lengths = {2 => 1, 3 => 7, 4 => 4, 7 => 8}
    count = 0
    for output in outputs
        for number in output
            if definite_lengths.key?(number.size)
                count += 1
            end
        end        
    end
    return count
end

#Helper function to print the display when needed
def print_display(display_values)
    i = 0
    while i < display_values.size do
        if i == 0 || i == 3 || i == 6
            puts " #{display_values[i]*4} "
        else
            puts "#{display_values[i]}    #{display_values[i+1]}"
            puts "#{display_values[i]}    #{display_values[i+1]}"
            i += 1
        end
        i += 1
    end
end

def sum_all(numbers, outputs)
    count = 0
    for i in 0...outputs.size
        output = ''
        values = solve_display numbers[i], outputs[i]
        for number in outputs[i]
            value = check_value number, values
            if value >= 0
                output += value.to_s
            end
        end
        count += output.to_i
    end
    return count
end

#helper since our hash doesn't necessarily have the output string in the right order (should've used arrays)
def check_value(str1, definite_numbers)
    definite_numbers.each do |key, value|
        diff = find_diff str1, value
        if diff.size == 0
            return key
        end
    end
    return -1
end

File.readlines('data.txt', chomp: true).each do |line|
    display, output = line.split(' | ')
    numbers = numbers.append(display.split(' '))
    outputs = outputs.append(output.split(' '))
end

puts "Task 1:"
p sum_easy outputs

puts "Task 2:"
p sum_all numbers, outputs





