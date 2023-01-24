CREATE TABLE libro(
	id int AUTO_INCREMENT,
	titulo varchar(64),
	editorial varchar(64),
	area varchar(64),
	PRIMARY KEY(id)
);
 

CREATE TABLE autor(
	id int AUTO_INCREMENT,
	nombre varchar(64),
	nacionalidad varchar(64),
	PRIMARY KEY(id)
)

 CREATE TABLE estudiante( 
 	id int AUTO_INCREMENT,
 	nombre varchar(64),
 	apellido varchar(64),
 	direccion varchar(64),
 	carrera varchar(64),
 	edad int,
 	PRIMARY KEY(id)
 )
 
 CREATE TABLE prestamo(
 	idEstudiante int,
 	idLibro int,
 	fechaPrestamo date,
 	fechaDevolucion date,
 	devuelto bool,
 	CONSTRAINT FK_prestamoEstudiante FOREIGN KEY(idEstudiante)
 		REFERENCES estudiante(id),
 	CONSTRAINT FK_prestamoLibro FOREIGN KEY(idLibro)
 		REFERENCES libro(id),
 	PRIMARY KEY(idEstudiante, idLibro)
 )
 
 CREATE TABLE libroAutor(
 	idAutor int,
 	idLibro int,
 	CONSTRAINT FK_autor FOREIGN KEY(idAutor)
 		REFERENCES autor(id),
 	CONSTRAINT FK_libro FOREIGN KEY(idLibro)
 		REFERENCES libro(id),
 	PRIMARY KEY(idAutor, idLibro)
 )

 