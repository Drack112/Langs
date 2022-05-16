class Carro
    def velocidade_maxima
        250
    end

    def adiciona_marca(marca)
        @marca = marca
    end
end

BMW = Carro.new

puts BMW.adiciona_marca("Toyota")
puts BMW.velocidade_maxima
