open("data.txt") do f
    global good_lines = 0
    while ! eof(f)
        
        line = readline(f)
        bad_strings = ["ab", "cd", "pq", "xy"]
        bad_included = false
        for bad in bad_strings
            if occursin(bad, line)
                bad_included = true
            end
        end

        if bad_included
            continue
        end

        double_letter = false
        vowel_count = 0
        vowels = "aeiou"

        spaced_letter = false
        repeated_double = false
        for i in 1:length(line)
            char = line[i]
            if occursin(char, vowels)
                vowel_count += 1
            end
            if i > 1
                last_two = line[i-1:i]
                if occursin(last_two, line[i+1:end])
                    repeated_double = true
                end
                if line[i] === line[i-1] 
                    double_letter = true
                end
            end
            if i > 2
                if line[i] === line[i-2]
                    spaced_letter = true
                end
            end
        end
        
        if vowel_count < 3
            continue
        end

        if !double_letter
            continue
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