#### WB Tech: Онлайн игра "Грузчики"

Задача:
генерируется N случайных заданий ("название переносимых предметов", "вес"). Есть грузчики зарегистрировавшиеся на работу и получившие случайные свойства. Задача заказчика - выбрать нужный набор грузчиков и выполнить задания.

#### API
```
# Публичное:
/tasks - создание случайного набора заданий
/auth/register - регистрация (требуются логин, пароль и класс игрока - грузчик или заказчик)
/auth/login - вход (требуются логин и пароль)
```
```
# Для грузчика (/loaders): 
/me - показать характеристики (вес, деньги, пьянство, усталость)
/tasks - показать список выполненных заданий
```
```
# Для заказчика (/client):
/me - показать свои характеристики (деньги, зарегистрировавшиеся грузчики)
/tasks - показать список доступных заданий
/start - добавить грузчиков и начать выполнение задания (списываются деньги, рассчитывается выполнено задание или нет)
```

#### База данных (PostgreSQL): 
```

create table players (
    id SERIAL PRIMARY KEY,
    login varchar(255) UNIQUE,
    password varchar(255),
    class varchar(255)
);

create table loaders (
    p_id integer,
    capacity integer,
    is_drinking bool,
    fatigue integer,
    salary integer
);

create table clients (
    p_id integer,
    fund integer,
    in_game bool
);


create table tasks (
    id SERIAL PRIMARY KEY,
    name varchar(255),
    item varchar(255),
    weight int,
    available bool
);

create table tasks_archive (
    t_id int,
    p_id int
);
```
