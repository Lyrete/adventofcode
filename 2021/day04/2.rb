require './1.rb'

boards, draw = read_boards 'data.txt'

#We would be much more efficient if we deleted numbers from the draw array whenever we use one
while boards do
    winning_board, winning_number = check_boards boards, draw
    if boards.length == 1
        break
    end
    boards.delete_at(winning_board)
end

sum = find_sum boards[0]

puts "Task 2:"
puts sum * winning_number.to_i