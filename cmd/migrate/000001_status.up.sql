CREATE TABLE IF NOT EXISTS codes(
   code int primary key,
   phrase varchar(50) not null,
   title varchar(60) not null,
   content text not null,
   meaning varchar(120) not null
);

