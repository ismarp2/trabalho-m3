Desenvolvimento - M3
Alunos: Enrico Belo Ribeiro, Gabriel Carneiro Rodrigues e Ismar João Pedrini Neto.


Escopo:
	Nossa meta é desenvolver um serviço voltado para controle de livros em uma biblioteca. O mesmo será uma api, desse modo, tornando-o genérico para muitos clientes, já que com um mesmo back-end, diversos front-ends podem ser desenvolvidos. Desse modo, uma mesma infraestrutura de aplicação pode resolver o problema de diversos clientes.
	Para esse desenvolvimento, utilizaremos como linguagem GOlang, visando um código limpo e com um tempo de resposta rápido. Para persistência, optamos por MongoDB.

Problema:
	Em bibliotecas, o controle de livros aparenta ser algo simples, normalmente feito em cadernos ou até mesmo em excel, porém com o grande volume de locações o controle disso pode ser algo caótico.
	Imaginando um cenário no qual um cliente pretende devolver um livro alocado, cabe ao atendente efetuar a busca dentre dezenas ou até centenas de páginas onde estão registradas as anotações de locação, para poder dar entrada na devolução de um livro. Podemos supor que este trabalho além de complicado também é demorado.

Solução:
	Com o problema acima levantado, uma possível solução é um sistema de controle de locações de livros. Deste modo, através de uma busca por CPF, o trabalho do atendente acaba sendo muito mais rápido e efetivo.
	Nossa ideia consiste em desenvolver esse sistema, no qual teremos um cadastro de usuários, livros e locações. Dessa maneira, podemos solucionar esse problema de uma maneira simples, cobrindo as duas entidades principais desse problema e a relação entre elas.
