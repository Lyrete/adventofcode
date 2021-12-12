data = []

File.readlines('data.txt', chomp:true).each do |line|
    data << line.split('').map(&:to_i)
end

def check_adjacent(x, y, data)
    if y < data.size - 1
        if data[y][x] >= data[y+1][x]
            return false
        end
    end
    
    if x < data[y].size - 1
        if data[y][x] >= data[y][x+1]
            return false
        end
    end
    if y > 0
        if data[y][x] >= data[y-1][x]
            return false
        end
    end
    if x > 0
        if data[y][x] >= data[y][x-1]
            return false
        end
    end

    return true
end

risk_level = 0

for y in 0...data.size
    for x in 0...data[0].size
        if check_adjacent x, y, data
            risk_level += data[y][x] + 1
        end
    end
end

puts "Task 1: "
p risk_level

puts "Task 2:"
p find_basin 0, 0, data


