#read file into 2d array
data = File.readlines("data.txt").map do |line|
    line.strip.split("")
end

#flip the array
data = data.transpose

common_string = ''

for i in 0...data.length
    column = data[i]
    common_value = column.group_by(&:itself).values.max_by(&:size).first
    common_string += common_value
end

uncommon_string = ''

for i in 0...common_string.length
    letter = common_string[i]
    if letter == '0'
        uncommon_string += '1'
    else
        uncommon_string += '0'
    end
end

puts common_string.to_i(2)
puts uncommon_string.to_i(2)

puts "Values multiplied: #{common_string.to_i(2) * uncommon_string.to_i(2)}"