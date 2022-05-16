def divide(a,b)
    raise "Divisão por 0 invalida" if b==0
    a / b
end

begin
    puts divide(8,0)
rescue Exception =>e
    puts "Erro ao dividir:" + e.message
end
