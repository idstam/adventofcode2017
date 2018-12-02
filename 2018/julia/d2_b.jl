let
 tot2 = 0
 tot3 = 0
 open("d2data.txt") do file
   lines = readlines(file)

   for l1 in lines, l2 in lines
     if l1 == l2
       continue
     end
     diffCount = 0
     for i = 1:length(l1)
       if l1[i] != l2[i]
         diffCount += 1
       end
     end
     if diffCount == 1
       println(l1)
       println(l2)
     else
       #println(diffCount)
       #println(l1)
       #println(l2)
     end
   end
end


end
