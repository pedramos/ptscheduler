Create table if not exists trainee (
	id  integer primary key,
	name text,
	perweek integer,
	late integer
);

Create table if not exists schedule (
	id  integer primary key,
	traineeid integer,
	date text,
	foreign key(traineeid) references trainee(id)
);


Create table if not exists usernames (
	traineeid integer,
	username blob primary key,
	password blob,
	foreign key(traineeid) references trainee(id)
);

Create table if not exists availability (
	id  integer primary key,
	traineeid integer,
	startdate text,
	enddate text,
	foreign key(traineeid) references trainee(id)
);