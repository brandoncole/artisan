FROM artisan-compiler:0.0.1
COPY include /artisan/include
COPY bin/artisan /artisan/artisan
RUN ln -s /artisan/artisan /usr/local/bin/artisan
WORKDIR /artisan