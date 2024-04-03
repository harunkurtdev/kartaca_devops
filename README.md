# Kartaca DevOPS Internship Assignment

Merhaba Ben Harun KURT, Bölümüm mekatronik müh. fakat, burada öğreneceğim bilgilerle farklı branşlarda,

    - Cloud Robotics
    - Robotics DevOPS  
    - ML
    - Edge Computing
  
gibi terimleri kullanmak istiyorum.

Aşağıda ki video'da projeye ait, çalışmasını izleyebilirsiniz.

![](https://www.youtube.com/watch?v=tNaWTzpV4yE)

[![Youtube Link]
(https://www.google.com/imgres?q=youtube%20logo&imgurl=https%3A%2F%2Fupload.wikimedia.org%2Fwikipedia%2Fcommons%2Fthumb%2F0%2F09%2FYouTube_full-color_icon_%25282017%2529.svg%2F1280px-YouTube_full-color_icon_%25282017%2529.svg.png&imgrefurl=https%3A%2F%2Ftr.m.wikipedia.org%2Fwiki%2FDosya%3AYouTube_full-color_icon_(2017).svg&docid=A1LXF8ZNsnCheM&tbnid=4SdipOmISh0hFM&vet=12ahUKEwiA_b-dgaeFAxUPYPEDHfpEAWUQM3oECBsQAA..i&w=1280&h=886&hcb=2&ved=2ahUKEwiA_b-dgaeFAxUPYPEDHfpEAWUQM3oECBsQAA)]
(https://www.youtube.com/watch?v=tNaWTzpV4yE)  

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
