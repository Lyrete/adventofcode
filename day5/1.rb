def create_grid
    grid = Array.new(1000, '.')
    for i in 0...grid.length
        grid[i] = Array.new(1000, '.')
    end
    return grid
end

def parse_file(name)
    moves = []
    File.readlines(name).each do |line|
        line = line.strip
        move = line.split('->')
        move[0], move[1] = move[0].strip.split(','), move[1].strip.split(',')
        moves.append(move)
    end
    return moves
end

def draw_straightlines(grid, moves)
    for move in moves
        x1, x2 = move[0][1].to_i, move[1][1].to_i
        y1, y2 = move[0][0].to_i, move[1][0].to_i
        if x1 == x2 #Check if line is straight vertically
            y = [*y1...y2+1]
            if y1 > y2
                y = [*y2...y1+1]
            end
            for i in y
                if grid[x1][i] == '.'
                    grid[x1][i] = 1
                else
                    grid[x1][i] += 1
                end
            end
        end

        if y1 == y2
            x = [*x1...x2+1]
            if x1 > x2
                x = [*x2...x1+1]
            end
            for i in x
                if grid[i][y1].to_s != '.'                    
                    grid[i][y1] += 1
                else                    
                    grid[i][y1] = 1
                end
            end
        end
    end

    return grid
end

def find_overlaps(grid)
    count = 0
    for line in grid
        for number in line
            if number.to_i >= 2
                count += 1
            end
        end
    end
    return count
end

filename = 'data.txt'
grid = create_grid
moves = parse_file filename
grid = draw_straightlines grid,moves

puts "Task 1:"
puts find_overlaps grid


