# TC Kimlik Numarası Kontrolü ve Doğrulaması

Golang kullanılarak TC kontrolü ve doğrulama işlemini beraber yapmayı sağlayan bir lambda function.

## Kullanım

``` text
     GOOS= linux CGO_ENABLED = 0 go build main.go
```

TC kimlik numarası algoritmik doğrulama için [Barış Esen](https://github.com/barisesen)'in
oluşturduğu [tcverify](https://github.com/barisesen/tcverify) paketi kullanılmıştır.

-----------------

### Sample Event

```json
    {
      "id": "12312312312",
      "name": "Onur",
      "surname": "Harputluoğlu",
      "birthYear": "2002"
    } 
````