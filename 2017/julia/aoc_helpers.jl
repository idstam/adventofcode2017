function stringToIntArray(input)
    ret = ArrayInt32(1)
    for c in input
        i = parse(Int32, c)
        push!(ret, i)
    end
    return ret
end

function fileToStringArray(fileName)
    ret = Array{String,1}()
    open(fileName) do file

        for line in eachline(file)
            push!(ret, line)
        end
    end
    return ret
end
