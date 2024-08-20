FROM apache/nifi:latest

# Cria o diretório para armazenar o driver JDBC (caso não exista)
RUN mkdir -p /opt/nifi/nifi-current/lib

# Copia o driver JDBC do MySQL para o diretório apropriado
COPY mysql-connector-java-8.0.30.jar /opt/nifi/nifi-current/lib/
COPY postgresql-42.7.3.jar /opt/nifi/nifi-current/lib/
