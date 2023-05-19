# Инструкция по запуску и примеры

Это руководство содержит инструкции по запуску и примеры для кода на Go из вашего проекта.

## Установка зависимостей

Перед запуском кода убедитесь, что все зависимости установлены. Вам понадобятся следующие зависимости:

- Go (версия 1.16 или выше)
- PostgreSQL

Выполните следующую команду для установки зависимостей Go:

```shell
go get github.com/lib/pq
go get github.com/radovskyb/watcher
go get gopkg.in/yaml.v
```

## Настройка базы данных
Прежде чем запустить код, убедитесь, что у вас есть база данных PostgreSQL. Затем отредактируйте следующую строку в коде, чтобы указать соответствующие параметры вашей базы данных:
```
db.DB, err = sql.Open("postgres", "user=postgres password=12345 dbname=postgres sslmode=disable")
```
## Конфигурация
Конфигурация для вашего приложения хранится в файле config.yaml. Убедитесь, что этот файл находится в том же каталоге, где находится исполняемый файл вашего приложения.

Пример содержимого файла config.yaml:
```
- path: "C:/Users/akrit/GolandProjects/awesomeProject/s2"
  commands:
    - "go build -o ./app C:/Users/akrit/GolandProjects/awesomeProject/main.go"
    - "go run ./app"
- path: "C:/Users/akrit/GolandProjects/awesomeProject/s1"
  commands:
    - "go test -v"
```   
Вы можете добавить любое количество конфигураций для отслеживания разных каталогов и выполнения различных команд при изменении файлов.
 
## Запуск приложения
Выполните следующую команду для сборки и запуска вашего приложения:
go build
./app-name

Замените app-name на имя исполняемого файла вашего приложения.

## Примеры
Вот несколько примеров для использования вашего приложения:
```
При изменении файлов в /path/to/directory, команда echo "File created or modified" будет выполнена, а также будет занесена запись в базу данных.

Если файлы в /path/to/another/directory будут созданы или изменены, будут выполнены команды echo "File created or modified in another directory" и echo "Performing another action", а также будет создана запись в базе данных.
```
Убедитесь, что указанные команды в конфигурации выполняются правильно на вашей системе.


    
