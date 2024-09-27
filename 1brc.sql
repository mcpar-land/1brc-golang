.timer on
.echo on

select
    station_name,
    min(measurement) as min_measurement,
    avg(measurement) as mean_measurement,
    max(measurement) as max_measurement
from read_csv(
    '1brc.csv.gz',
    header = false,
    columns = {
        'station_name': 'VARCHAR',
        'measurement': 'double',
    },
    delim=';'
)
group by station_name
order by station_name;

select
    station_name,
    min(measurement) as min_measurement,
    avg(measurement) as mean_measurement,
    max(measurement) as max_measurement
from '1brc.parquet'
group by station_name
order by station_name;
