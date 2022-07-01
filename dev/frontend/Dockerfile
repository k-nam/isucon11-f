FROM nginx:1.21 as runner

# drop symlinks
RUN unlink /var/log/nginx/access.log
RUN unlink /var/log/nginx/error.log

COPY ./dist /www/data
