FROM tiangolo/uvicorn-gunicorn-fastapi:python3.9-slim AS builder

LABEL maintainer="Harun KURT <harunkurtdev> (https://github.com/harunkurtdev)"

ENV WORKDIRs=/app
WORKDIR ${WORKDIRs} 


FROM builder as dev


COPY requirements.txt ./
RUN --mount=type=cache,target=/root/.cache/pip \
    pip install -r requirements.txt

COPY . .

EXPOSE 4444


CMD ["python3", "__main__.py"]