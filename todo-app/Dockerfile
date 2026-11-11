FROM nginx:alpine

ENV PORT=8080

COPY index.html /usr/share/nginx/html/index.html

RUN rm /etc/nginx/conf.d/default.conf

CMD ["/bin/sh", "-c", \
  "echo \"server { listen ${PORT}; root /usr/share/nginx/html; index index.html; }\" > /etc/nginx/conf.d/default.conf && exec nginx -g 'daemon off;'"]