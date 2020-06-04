FROM webdevops/liquibase:postgres

USER root

WORKDIR /liquibase

COPY ./entripoints/migrator.sh .
RUN chmod +x ./migrator.sh

COPY ./project/models/changeLog .

ENTRYPOINT ["./migrator.sh"]

CMD ["update"]