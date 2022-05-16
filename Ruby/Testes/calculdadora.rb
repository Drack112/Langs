a = 4
b = 2

puts "O que deseja realizar?"

text = <<E0S
1 - Soma
2 - Subtração
E0S

puts text

operation = gets

case operation
when 1
    puts a + b
when 2
    puts a - b
end
