def check_slope(input_data, x_move, y_move)
    x = 0
    y = 0
    #Data repeats width wise

    trees = 0

    while y <= input_data.length - 1 do    
        if x > 30
            #reset the x coordinate when it goes over the width
            x = x - 30 - 1
        end
        if input_data[y][x] == '#' # Check if we hit a tree and add it to total
            trees = trees + 1
        end
        x = x + x_move        
        y = y + y_move
    end

    return trees
end

data = File.readlines('data.txt')

value1 = check_slope data, 3, 1

value2 = value1 * (check_slope data, 1, 1) * (check_slope data, 5, 1) * (check_slope data, 7, 1) * (check_slope data, 1, 2)

puts "Task 1:"
puts "Trees hit: #{value1} times"

puts "Task 2:"
puts "Trees hit multiplied: #{value2} times"