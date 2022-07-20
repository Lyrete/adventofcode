open("data.txt") do f
    global good_lines = 0
    while ! eof(f)
        
        line = readline(f)

        spaced_letter = false
        repeated_double = false
        for i in 1:length(line)
            char = line[i]
            if i > 1
                last_two = line[i-1:i]
                if occursin(last_two, line[i+1:end])
                    repeated_double = true
                end
            end
            if i > 2
                if line[i] === line[i-2]
                    spaced_letter = true
                end
            end
        end

        if !spaced_letter
            continue
        end

        if !repeated_double
            continue
        end

        println("$line")

        good_lines += 1
    end
end

println("Good lines:")
println("$good_lines")