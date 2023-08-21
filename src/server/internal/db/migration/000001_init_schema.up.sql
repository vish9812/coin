create table "user" (
	id bigint primary key generated always as identity,
	email text unique not null,
	"password" text not null,
	first_name text not null,
	last_name text,
	"created_at" timestamptz NOT NULL DEFAULT (now())
);

create table income (
	id bigint primary key generated always as identity,
	"name" text not null,
	amount bigint not null,
	target_saving bigint not null,
	received_at timestamptz,
	user_id bigint not null,
	constraint fk_income_user foreign key (user_id) references "user" (id) on delete cascade
);

create table category (
	id bigint primary key generated always as identity,
	"name" text not null,
	target_amount bigint,
	user_id bigint not null,
	constraint fk_category_user foreign key (user_id) references "user" (id) on delete cascade
);

create table goal (
	id bigint primary key generated always as identity,
	"name" text not null,
	target_amount bigint,
	user_id bigint not null,
	constraint fk_goal_user foreign key (user_id) references "user" (id) on delete cascade
);

create table expense (
	id bigint primary key generated always as identity,
	"name" text not null,
	amount bigint not null,
	category_id bigint,
	user_id bigint not null,
	constraint fk_expense_user foreign key (user_id) references "user" (id) on delete cascade,
	constraint fk_expense_category foreign key (category_id) references "category" (id) on delete set null
);

create table saving (
	id bigint primary key generated always as identity,
	"name" text not null,
	amount bigint not null,
	goal_id bigint,
	user_id bigint not null,
	constraint fk_expense_user foreign key (user_id) references "user" (id) on delete cascade,
	constraint fk_expense_goal foreign key (goal_id) references "goal" (id) on delete set null
);
