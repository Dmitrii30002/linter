# Простой линтер
Данный линтер проверяет 4 правила:
1. Лог-сообщения должны начинаться со строчной буквы
2. Лог-сообщения должны быть только на английском языке
3. Лог-сообщения не должны содержать спецсимволы или эмодзи
4. Лог-сообщения не должны содержать потенциально чувствительные данные <br>
<br>
* язык разработки: Go 1.25.1
* поддержка log/slog
* Совместимость с golangci-lint
<br>

## Инструкция
Скопируйте репозиторий: <br>
``` bash
	git clone https://github.com/Dmitrii30002/linter.git
```
Для windows:
Используйте команду build_windows:
``` bash
	make build_windows
```
Пример работы:
``` bash
	./linter.exe ./pkg/analyze/testdata/src/tests
```
Для других систем:
``` bash
	make install
```
чтобы использовать линтер совместно с golanci-linter:
``` bash
	./golangci -c config.yml run ./...
```
## Тестирование
Для запуска тестов пропишите:
``` bash
	make test
```

## Пример работы
Проверим файл tests.go:
``` bash
	./linter.exe ./pkg/analyze/testdata/src/tests
```

Результатом будут следующие строчки:
```
D:\GO\linter\pkg\analyzer\testdata\src\tests\tests.go:11:2: log.Fatal: log messages must start with a lowercase letter (message: Starting server on port 8080)
D:\GO\linter\pkg\analyzer\testdata\src\tests\tests.go:12:2: slog.Error: log messages must start with a lowercase letter (message: Failed to connect to database)
D:\GO\linter\pkg\analyzer\testdata\src\tests\tests.go:20:2: log.Print: log message should be in English only (message: а)
D:\GO\linter\pkg\analyzer\testdata\src\tests\tests.go:21:2: log.Fatal: log message should be in English only (message: ш)
D:\GO\linter\pkg\analyzer\testdata\src\tests\tests.go:29:2: log.Print: log messages should not contain special characters or emojis (message: server started!🚀)
D:\GO\linter\pkg\analyzer\testdata\src\tests\tests.go:30:2: log.Fatal: log messages should not contain special characters or emojis (message: connection failed!!!)
D:\GO\linter\pkg\analyzer\testdata\src\tests\tests.go:31:2: log.Fatal: log messages should not contain special characters or emojis (message: warning: something went wrong...)   
D:\GO\linter\pkg\analyzer\testdata\src\tests\tests.go:43:2: log.Print: log message should not contain sensitive data (message: user password )
```

  


