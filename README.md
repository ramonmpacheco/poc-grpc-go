# POC-GRPC-GO
## Go Lang with GRPC

##

> Extensao usada no vs code:
> 
[VSCODE-PROTO3](https://github.com/zxh0/vscode-proto3)

> Instalando o compilador do protobuf

[SITE AQUI](https://grpc.io/docs/protoc-installation/)
>
~~~shell
# [exemplo]
$ apt install -y protobuf-compiler
~~~
>
>Instalando plugins do GO
>
[SITE AQUI](https://grpc.io/docs/languages/go/quickstart/)

```shell
# [exemplo]
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

> Complilando:

> Comando usado: 
> 
> ```protoc --proto_path=proto proto/*.proto --go_out=stub --go-grpc_out=stub```

> O Comando usado teve como base os pacotes em: 
~~~
> poc-grpc-go
    > /proto 
    > /stub
~~~

> O comando foi executando na dentro da pasta raiz: "poc-grpc-go"

> Possivel erro ao rodar o comando: **"The import path must contain at least one period ('.') or forward slash ('/') character."**
>
> Isso provavelmente se refere a prop "go_path" dentro do arquivo .proto
[ISSUE DISCUTIDA](https://github.com/techschool/pcbook-go/issues/3)

> Para saber mais sobre proto3
> 
> [DOC DEV GOOGLE](https://developers.google.com/protocol-buffers/docs/proto3)

> Cliente para conectar no grpc quando a reflection estiver abilitada 

> Baixar o binário [aqui](https://github.com/ktr0731/evans/releases)
> 
> descompactar o binario dentro de **/usr/local/evans**
```shell
# [exemplo]
$ sudo mkdir /usr/local/evans

# [dentro da pasta onde o arquivo baixado está]
$ sudo tar -C /usr/local/evans -xzf evans_linux_amd64.tar.gz

# [exporte o evans no seu bash]

# GRPC Evans
export PATH=$PATH:/usr/local/evans

# Recarrege sua variáveis
$ source ~/.{seu bash}
```

> Subindo o server GRPC:
> 
> Execute o arquivo **server.go**

```shell
# [exemplo]
$ go run server.go
```

> Em outra aba execute o evans:
```shell
$ evans -r -p 50051

# [resultado]
 ______
 |  ____|
 | |__    __   __   __ _   _ __    ___
 |  __|   \ \ / /  / _. | | '_ \  / __|
 | |____   \ V /  | (_| | | | | | \__ \
 |______|   \_/    \__,_| |_| |_| |___/

 more expressive universal gRPC client

controllergrpc.MathService@127.0.0.1:50051>
```

> Faça uma chamada ao serviço criado:
```shell
# [exemplo]
$ call Sum
```
> Responda as perguntas 
~~~shell
sum::a (TYPE_FLOAT) => 5
sum::b (TYPE_FLOAT) => 6
~~~
> Deve aparecer o resultado:
~~~shell
{
  "result": 11
}
~~~