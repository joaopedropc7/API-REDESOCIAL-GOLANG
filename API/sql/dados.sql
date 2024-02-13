INSERT INTO usuarios (nome, nick, email, senha)
values 
("usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$.lY7eNk5D31zsASDhJJ5g.Zgh6/ef2JOptuPS2GcwAugDxCTB3EuC"),
("usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$.lY7eNk5D31zsASDhJJ5g.Zgh6/ef2JOptuPS2GcwAugDxCTB3EuC"),
("usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$.lY7eNk5D31zsASDhJJ5g.Zgh6/ef2JOptuPS2GcwAugDxCTB3EuC");

INSERT INTO seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);


insert into publicacoes(titulo, conteudo, autor_id)
values
("Publicacao do usuario 1", "Essa e a publicacao do usuario 1", 1),
("Publicacao do usuario 2", "Essa e a publicacao do usuario 2", 2),
("Publicacao do usuario 3", "Essa e a publicacao do usuario 3", 3);
