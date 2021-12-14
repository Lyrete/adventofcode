filename = "data.txt"

def check_win(board)
    check = Array.new(5, "X")

    #check if we can find a line that completes the bingo
    for line in board
        if (line - check).empty?
            return true
        end
    end

    #Do the same for columns
    for line in board.transpose
        if (line - check).empty?
            return true
        end
    end

    return false

end

def check_board(board, number)
    for y in 0...board.length do
        line = board[y]
        for x in 0...line.length do
            if line[x] == number
                board[y][x] = "X"
            end
        end
    end
    win = check_win board
    return board, win
end

def find_sum(board)
    sum = 0
    for line in board
        for number in line
            if number != 'X'
                sum += number.to_i
            end
        end
    end
    return sum
end

def read_boards(filename)    
    boards = []
    board = []
    draw = ''
    File.foreach(filename).with_index do |line, i|
        if i == 0
            draw = line.strip.split(",")
            next
        end
        
        if line.strip == ''
            if board.length > 0
                boards.append(board)
            end
            board = []
        else
            board.append(line.strip.split(" "))
        end
    end
    
    #appends last board to the list
    boards.append(board)

    return boards, draw
end



def check_boards(boards, draw)
    for number in draw do
        for i in 0...boards.length do
            board = boards[i]
            boards[i], win = check_board board, number
            if win
                return i, number
            end
        end
    end
end

boards, draw = read_boards "data.txt"
puts "Task 1:"
winning_board, winning_number = check_boards boards, draw
sum = find_sum boards[winning_board]
puts sum * winning_number.to_i


