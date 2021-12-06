require './1.rb'

def draw_diagonals(grid, moves)
    for move in moves
        x1, x2 = move[0][1].to_i, move[1][1].to_i
        y1, y2 = move[0][0].to_i, move[1][0].to_i
        y = [*y1..y2]
        x = [*x1..x2]
        if y1 > y2 
            y = (y2..y1).to_a.reverse
        end
        if x1 > x2
            x = (x2..x1).to_a.reverse
        end
        if x.length == y.length
            for i in 0...x.length
                if grid[x[i]][y[i]] == '.'
                    grid[x[i]][y[i]] = 1
                else
                    grid[x[i]][y[i]] += 1
                end
            end
        end
    end

    return grid
end

def print_grid(grid)
    for line in grid
        for letter in line
            print letter
        end
        puts ""
    end
end

filename = 'data.txt'
grid = create_grid
moves = parse_file filename
grid = draw_straightlines grid,moves
grid = draw_diagonals grid, moves

puts "Task 2:"
puts find_overlaps grid


