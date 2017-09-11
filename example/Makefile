clean:
	@rm -rf ./cert; mkdir ./cert

rootCA:
	@openssl genrsa -out ./cert/RootCA.key 4096 && \
	openssl req -new -x509 -days 1826 -key ./cert/RootCA.key -out ./cert/RootCA.crt -subj '/C=CA/ST=Ontario/O=Test/CN=root';

intermediateCA:
	@openssl genrsa -out ./cert/IntermediateCA.key 4096 && \
	openssl req -new -key ./cert/IntermediateCA.key -out ./cert/IntermediateCA.csr -subj '/C=CA/ST=Ontario/O=Test/CN=intermediate' && \
	openssl x509 -req -days 1000 -in ./cert/IntermediateCA.csr -CA ./cert/RootCA.crt -CAkey ./cert/RootCA.key -CAcreateserial -CAserial ./cert/serial -out ./cert/IntermediateCA.crt && \
	openssl rsa -in ./cert/IntermediateCA.key -pubout -out ./cert/IntermediateCA.pub;

serverCert:
	@openssl genrsa -out ./cert/server.key 2048 && \
	openssl req -new -key ./cert/server.key -out ./cert/server.csr -subj '/C=CA/ST=Ontario/O=Test/CN=server' && \
	openssl x509 -req -in ./cert/server.csr -CA ./cert/IntermediateCA.crt -CAkey ./cert/IntermediateCA.key -set_serial 01 -out ./cert/server.crt -days 500 -sha256;

clientCert:
	@openssl genrsa -out ./cert/client.key 2048 && \
	openssl req -new -key ./cert/client.key -out ./cert/client.csr -subj '/C=CA/ST=Ontario/O=Test/CN=client' && \
	openssl x509 -req -in ./cert/client.csr -CA ./cert/IntermediateCA.crt -CAkey ./cert/IntermediateCA.key -set_serial 01 -out ./cert/client.crt -days 500 -sha256;

setup: clean rootCA intermediateCA serverCert clientCert
