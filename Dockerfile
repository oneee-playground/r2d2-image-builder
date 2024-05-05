FROM golang:1.22 AS build
WORKDIR /work
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o app 

FROM public.ecr.aws/lambda/provided:al2023
COPY --from=build /work/app ./app
ENTRYPOINT [ "./app" ]
