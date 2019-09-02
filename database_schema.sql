create table scanner_registration
(
	id serial not null constraint scanner_registration_pkey primary key,
	name varchar(255) not null constraint scanner_registration_name_key unique,
	description varchar(255),
	endpoint_url varchar(255) not null,
	default_flag boolean default false not null,
	enabled_flag boolean default false not null,
	deleted_flag boolean default false not null
);

insert into scanner_registration (name, endpoint_url, default_flag, enabled_flag) values ('MicroScanner', 'http://harbor-scanner-microscanner:8080/api/v1', true, true);
insert into scanner_registration (name, endpoint_url, default_flag, enabled_flag) values ('Anchore Engine', 'http://harbor-scanner-anchore:8080/api/v1', false, true);
insert into scanner_registration (name, endpoint_url, default_flag, enabled_flag) values ('Trivy', 'http://harbor-scanner-trivy:7070/api/v1', false, true);

drop table scanner_registration