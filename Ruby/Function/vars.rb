def nomes(idade, *nomes)
    nomes.each {|n| puts n}
    puts "Idade = #{idade}"
end

nomes 42, "ricardo", "roberto"
