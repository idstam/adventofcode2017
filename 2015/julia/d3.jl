let
        visited = Dict("0,0" => 1)
        x = 0
        y = 0
        open("data.txt") do file
            for line in eachline(file)
                for c in line
                    if c == 'v'
                        y -= 1
                    end
                    if c == '^'
                        y += 1
                    end
                    if c == '<'
                        x -= 1
                    end
                    if c == '>'
                        x += 1
                    end
                    adress = string(x, ',', y)
                    v = get(visited, adress, 0)
                    visited[adress] = v +1

                end
            end
        end
        println(length(visited))
end
