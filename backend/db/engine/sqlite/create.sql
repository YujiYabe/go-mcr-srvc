
CREATE TABLE products (
	jan_code        integer  primary key,
	is_valid        bool     default '0'  not null,
	place           text     default ''  not null,
	name_ja         text     default ''  not null,
	name_en         text     default ''  not null,
	name_es         text     default ''  not null,
	name_zh         text     default ''  not null,
	name_ar         text     default ''  not null,
	information_ja  text     default ''  not null,
	information_en  text     default ''  not null,
	information_es  text     default ''  not null,
	information_zh  text     default ''  not null,
	information_ar  text     default ''  not null,
	recipe          text     default ''  not null,
	calorie	        integer  default '0'  not null,
	retail_price    integer  default '0'  not null,
	operation_cost  integer  default '0'  not null,
	energy_cost     integer  default '0'  not null,
	cost_price      integer  default '0'  not null,
	tag	            text     default ''  not null,
	allergy         text     default ''  not null
);

