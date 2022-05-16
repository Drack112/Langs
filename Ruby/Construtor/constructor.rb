class Produto
    attr_accessor :nome, :preco

    def initialize(nome, preco)
        @nome = nome
        @preco = preco
    end

    def info
        puts "Produto: #{@nome} com preço de R$ #{@preco}"
    end
end

produto = Produto.new("Caneta", 12.90)

produto.info
