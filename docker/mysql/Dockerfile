FROM mysql:5.7
COPY migration.sh /mysql/migration.sh
COPY schema.sql /mysql/schema.sql
RUN /mysql/migration.sh