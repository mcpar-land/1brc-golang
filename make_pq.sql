.timer on
.echo on

create view measurements as select * from read_csv(
    '1brc.csv.gz',
    header = false,
    columns = {
        'station_name': 'VARCHAR',
        'measurement': 'double',
    },
    delim=';'
);

copy (select * from measurements) to '1brc.parquet' (FORMAT PARQUET);
