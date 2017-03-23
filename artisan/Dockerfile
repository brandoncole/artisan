FROM artisan-compiler:0.0.1
COPY api/ /artisan/include/artisan
COPY bin/artisan /artisan/artisan
RUN ln -s /artisan/artisan /usr/local/bin/artisan
WORKDIR /artisan