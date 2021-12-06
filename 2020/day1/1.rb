data = File.readlines('data.txt', chomp: true).map(&:to_i)

def find_sum(data, sum)
    checked = Hash.new()
    data.each_with_index do |number, i|
        if checked.has_key?(sum-number)
            return number, sum-number
        end
        checked[number] = i 
    end
    return false
end

def find_triple(data,sum)
    checked = Hash.new()
    data.each_with_index do |number, i|
        num2, num3 = find_sum data, sum-number
        if num2
            return number, num2, num3
        end
    end
end

value1, value2 = find_sum data.sort, 2020

puts "Our pair is:"
print value1.to_s + " * " + value2.to_s + " = "
p value1 * value2

value1, value2, value3 = find_triple data.sort, 2020
puts "The triples are:"
print value1.to_s + " * " + value2.to_s + " * " + value3.to_s + " = "
p value1 * value2 * value3