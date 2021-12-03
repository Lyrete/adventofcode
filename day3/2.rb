require "./1"

#read file into 2d array
data = File.readlines("data.txt").map do |line|
    line.strip.split("")
end

def find_ratings(data, i = 0, value_idx = 0)
    if data.length == 1 #When only one datapoint left return it
        return data
    end
        
    values = find_binarystring data #Get the common/uncommon values for current data
    reg = values[value_idx] #A bit of a janky solution to get common/uncommon value since our original function returned both at once

    new_data = []
    for row in data
        check = row.join("")
        if check[i] == reg[i] #if letter we're checking matches then add to data for next check
            new_data.append(row)
        end
    end

    i += 1
    find_ratings(new_data, i, value_idx)
end

o2_rating = find_ratings data, 0, 0
co2_rating = find_ratings data, 0, 1

puts "Result:"
puts "O2: #{o2_rating.join("")}, CO2: #{co2_rating.join("")}"
puts "Multiplied: #{o2_rating.join("").to_i(2) * co2_rating.join("").to_i(2)}"
