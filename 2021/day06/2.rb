require './1.rb'

fishes = parse_file "data.txt"
fishes = pass_days 80, fishes
puts "Task 1:"
puts count_fishes fishes

more_days = 256 - 80
fishes = pass_days more_days, fishes

puts "Task 2:"
puts count_fishes fishes