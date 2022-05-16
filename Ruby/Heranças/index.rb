class Automovel

    def self.tipo_cambio
        puts "Manual"
    end

    def acelera
        puts "Acelerando..."
    end
end

class Carro < Automovel
    def acelera
        puts "Verificando"
        super
    end
end

carro = Carro.new()

carro.acelera
