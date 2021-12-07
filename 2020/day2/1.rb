def validate(line)
    validation, pw = line.split(':')
    range, letter = validation.split(' ')
    min, max = range.split('-').map(&:to_i)

    return pw.count(letter) >= min && pw.count(letter) <= max 
end

def check_pos(line)
    validation, pw = line.split(':')
    range, letter = validation.split(' ')
    pos1, pos2 = range.split('-').map(&:to_i)

    return (pw[pos1] == letter) ^ (pw[pos2] == letter)
end


data = File.readlines('data.txt', chomp:true)

valid1 = 0
valid2 = 0
for line in data
    if validate line 
        valid1 += 1
    end
    if check_pos line
        valid2 += 1
    end
end

puts "Task 1:"
p valid1

puts "Task 2:"
p valid2