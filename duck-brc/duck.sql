-- name: Test 1
-- desc: Read some 5 records
SELECT * FROM READ_CSV_AUTO('measurements.txt') LIMIT 5;

-- name: Test 2
-- desc: Create a table
CREATE OR REPLACE TABLE measurements 
AS
  SELECT * FROM 
  READ_CSV('measurements.txt', 
    header=false, 
    columns= {'station_name':'VARCHAR','measurement':'double'}, delim=';') 
  LIMIT 2048;

-- name: Test 3
SELECT station_name, 
           MIN(measurement),
           AVG(measurement),
           MAX(measurement)
    FROM measurements 
    GROUP BY station_name;

-- name: Test 4
SELECT station_name, 
           MIN(measurement) AS min_measurement,
           CAST(AVG(measurement) AS DECIMAL(8, 1)) AS avg_measurement,
           MAX(measurement) AS max_measurement
    FROM measurements 
    GROUP BY station_name;

-- name: Test 5
-- Add timer and everything
.timer on
.mode ascii
.echo on

WITH src AS (SELECT station_name,
                    MIN(measurement) AS min_measurement,
                    CAST(AVG(measurement) AS DECIMAL(8,1)) AS mean_measurement,
                    MAX(measurement) AS max_measurement
            FROM measurements
            GROUP BY station_name)
    SELECT '{' ||
            ARRAY_TO_STRING(LIST_SORT(LIST(station_name || '=' || CONCAT_WS('/',min_measurement, mean_measurement, max_measurement))),', ') ||
            '}' AS "1BRC"
    FROM src;

.quit
