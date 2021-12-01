# Read file into array
data = File.read("data_1.txt").split

#init sum var
inc = 0

#loop through the array
for i in 3..(data.length()-1) do
    #fetch values from the last 4 nodes
    data3 = data[i].to_i
    data2 = data[i-1].to_i
    data1 = data[i-2].to_i
    data0 = data[i-3].to_i
    
    #Calculate sums
    dataB = data3 + data2 + data1
    dataA = data2 + data1 + data0

    if dataB > dataA #Check if data point increased from previous
        inc = inc + 1
    end
end

# prints out the amount of increased values
puts "Value increased #{inc} times."