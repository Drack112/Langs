class Carro

    attr_accessor :marca, :modelo

    def velocidade
        250
    end

    def info
        puts @marca
        puts @modelo
    end
end

carro = Carro.new()

carro.marca = "Ford"
carro.modelo = "Mustang"

puts carro.info
