create table personalidades
(
    id       serial primary key,
    nome     varchar,
    historia varchar
);

INSERT INTO personalidades(nome, historia)
VALUES ('Deodato Petit Wertheimer',
        'Deodato Petit Wertheimer foi um m?dico e pol?tico brasileiro, seus primeiros anos de vida foram em S?o Paulo, mas logo mudou para Nova Friburgo no Estado do Rio de Janeiro e com 11 anos de idade ingressou no Col?gio Anchieta dos jesu?tas.'),
       ('Carmela Dutra',
        'Carmela Teles Leite Dutra foi a primeira-dama do Brasil, de 31 de janeiro de 1946 at? a sua morte, tendo sido a esposa de Eurico Gaspar Dutra, 16.? Presidente do Brasil. Era, carinhosamente, chamada de Dona Santinha, pela sua forte religiosidade, fazendo seu marido abrir uma capelinha no Pal?cio Guanabara.');