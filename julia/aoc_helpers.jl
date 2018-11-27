function stringToIntArray(input)
    ret = ArrayInt32(1)
    for c in input
        i = parse(Int32, c)
        push!(ret, i)
    end
    return ret
end
