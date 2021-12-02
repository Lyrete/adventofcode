#Init coordinates
x = 0
y = 0

#Used for task 2
y_2 = 0

File.open("data.txt").each_line do |line|
    data = line.split(" ")
    amount = data[1].to_i
    if data[0] == "forward"
        x += amount
        y_2 += amount * y #Use the y value calculated anyway already as the "aim"
    elsif data[0] == 'down'
        y += amount
    else
        y -= amount
    end
end

puts "X is #{x}"
puts "Y is #{y}"
puts "Y_2 is #{y_2}"
puts "Multiplied x*y = #{x*y}"
puts "Multiplied x*y_2 = #{x*y_2}"