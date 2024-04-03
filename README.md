# Kartaca DevOPS Internship Assignment

Merhaba Ben Harun KURT, Bölümüm mekatronik müh. fakat, burada öğreneceğim bilgilerle farklı branşlarda,

    - Cloud Robotics
    - Robotics DevOPS  
    - ML
    - Edge Computing
  
gibi terimleri kullanmak istiyorum.

Aşağıda ki video'da projeye ait, çalışmasını izleyebilirsiniz.

[![Youtube Link](https://i9.ytimg.com/vi_webp/tNaWTzpV4yE/mqdefault.webp?v=660dcb7e&sqp=CNybt7AG&rs=AOn4CLDK9nJLdLfJiqnSjLw0kSd1WKyfHA)](https://www.youtube.com/watch?v=tNaWTzpV4yE)


Burada yaptığım adımlar 2 farklı OS'ta denenmiştir.
    - MacOS Venture
    - Ubuntu 22.04'tür

farklı 2 branch'ta mevcut olarak çalışan projeler bulunmaktadır.

macOS branch'ında ```elastic.datas``` oluşturularak verileri local'e kayıt edilmesi sağlanmaktadır.
Fakat, ubuntu22.04 branch'ında `elastic.datas` bu veriler docker volumes altında oluşturulmaktadır.

### Make Komutları

```Makefile``` altında bazı komut listeleri bulunmaktadır.

    - rm_files:
        `elasticsearch` tarafında üretilen bu dosyaların silinmesini sağlamaktadır. `elastic.datas` klasörünü silmektedir. (macOS tarafında çalışmaktadır.)
    - build:
        `docker compose up --build` demek yerine `make build` diyerek proje dizininde kısa bir komut oluşturmuş olduk.
    - start:
        proje dizininde`docker compose up -d` demek yerine make start ile projemizi başlatmaktayız. 

### Kullanılan Portlar

    Değerli gözetmen çalıştırdığınız da bilgisayarınızda ki portları bilmediğim için bunlardan birtanesinyle çakışma olursa eğerki lütfen değiştirir misiniz? 

    - elasticsearch 9200
    - goapp 5555
    - prometheus 9090
    - pythonapp 4444
    - grafana 3001
