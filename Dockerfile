FROM golang:1.22.3-alpine3.19 AS builder

ARG VOLIO_GIT_DEPLOY_USERNAME=""
ARG VOLIO_GIT_DEPLOY_PWD=""

LABEL maintainer="NhokCrazy199 <phamkim.pr@gmail.com>"

WORKDIR /app

ENV BIN_DIR=/app/bin

RUN apk update && apk add --no-cache gcc build-base git

# Copy go.mod và go.sum trước để cache dependencies
COPY go.mod go.sum ./

# Loại bỏ các dòng replace nếu cần
RUN grep -v "replace\s.*=>.*" go.mod > tmpfile && mv tmpfile go.mod

COPY scripts/env.sh scripts/env.sh
RUN chmod +x scripts/env.sh && ./scripts/env.sh

# Tải dependencies vào cache
RUN go mod download


# Cài đặt delve
RUN go install github.com/go-delve/delve/cmd/dlv@latest

COPY scripts/build.sh scripts/build.sh

# Copy toàn bộ mã nguồn
RUN cp go.mod go.mod.backup && cp go.sum go.sum.backup
COPY . .
RUN cp go.mod.backup go.mod && cp go.sum.backup go.sum

# Đồng bộ go.mod với mã nguồn thực tế
RUN go mod tidy

COPY /docs /app/docs

# Build ứng dụng
RUN chmod +x scripts/build.sh && scripts/build.sh

# Stage 2
FROM alpine:latest

WORKDIR /app

ENV BIN_DIR=/app/bin
ENV SCRIPTS_DIR=/app/scripts
ENV LOG_DIR=/var/log/volio

RUN mkdir -p ${BIN_DIR} \
    mkdir -p ${SCRIPTS_DIR} \
    # mkdir -p ${DATA_DIR} \
    mkdir -p ${LOG_DIR} \
    && addgroup -S volio \
    && adduser -S volio -G volio \
    && chown volio:volio /app \
    && chown volio:volio ${LOG_DIR}

USER volio

COPY --chown=volio:volio --from=builder ${BIN_DIR} /app
COPY --chown=volio:volio --from=builder /go/bin/dlv /app
# COPY --chown=volio:volio --from=builder ${DATA_DIR} ${DATA_DIR}
COPY --chown=volio:volio --from=builder /app/docs /app/docs
COPY --chown=volio:volio --from=builder ${SCRIPTS_DIR} ${SCRIPTS_DIR}

RUN chmod +x ${SCRIPTS_DIR}/startup.sh

ENTRYPOINT ["/app/scripts/startup.sh"]