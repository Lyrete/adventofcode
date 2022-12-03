require 'set'

def parse_file(filename)
    highest_id = 0
    ids = []
    File.readlines(filename).each do |line|
        data = (line.strip.split(''))
        row_range = *(0..127)
        seat_range = *(0..7)
        for letter in data
            if letter == 'F' || letter == 'B'
                row_range = parse_letter letter, row_range
            end

            if letter == 'L' || letter == 'R'
                seat_range = parse_letter letter, seat_range
            end
        end

        id = row_range[0] * 8 + seat_range[0]
        ids.append(id)
        highest_id < id ? highest_id = id : false
    end
    ids = ids.sort

    my_seat = 0
    for i in 1..ids.length-1
        seat = ids[i]
        if seat-1 != ids[i-1]
            my_seat = seat-1
            break
        end
    end    
    return highest_id, my_seat
end

def parse_letter(letter, range)
    if letter == 'F' || letter == 'L'
        return range[0...range.length/2]
    elsif letter == 'B' || letter == 'R'
        return range[range.length/2...range.length]
    end
end

highest_id, my_seat = parse_file 'data.txt'
puts "Task 1:"
p highest_id
puts "Task 2:"
p my_seat
