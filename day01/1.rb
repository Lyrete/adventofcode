# Read file into array
data = File.read("data_1.txt").split

#init sum var
inc = 0

#loop through the array
for i in 1..(data.length()-1) do
    data1 = data[i].to_i
    data0 = data[i-1].to_i
    if data1 > data0 #Check if data point increased from previous
        inc = inc + 1
    end
end

# prints out the amount of increased values
puts "Value increased #{inc} times."

