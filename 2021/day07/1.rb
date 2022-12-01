#Read file
data = File.readlines('data.txt').sample.strip.split(',').map(&:to_i)

#Get the median
data = data.sort
mid = data[data.length/2]

avg = (data.reduce(:+).to_f / data.length) #Get avg for task 2


sum1 = 0
sum2 = 0
sum3 = 0
#Loop through array to get the fuel counts
for spot in data
    sum1 += (spot - mid).abs
    sum2 += (1..(spot - avg.floor).abs).sum
    sum3 += (1..(spot - avg.ceil).abs).sum
end

puts "Task 1:"
p sum1
puts "Task 2:"
p [sum2, sum3].min