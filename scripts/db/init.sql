create table if not exists contracts(
    id serial primary key, 
    sc_addr text unique,
    is_indexing boolean,
    last_tx_db text,
    last_indx_date timestamp);

create table if not exists transactions (
    id serial primary key,
    sc_id int,
    tx_addr text unique,
    from_addr text,
    to_addr text,
    tokens text,
    block_nr int,
    tx_index int,
    removed boolean,
    log_index int,
    log_name varchar);    
