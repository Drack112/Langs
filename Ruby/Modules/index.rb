module Compartilhando
    class Carro
        include Compartilhando
        def metodo
            puts "carro"
        end
    end
end



novo_carro = Compartilhando::Carro.new
novo_carro.metodo
