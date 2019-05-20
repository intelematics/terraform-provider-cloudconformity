FROM golang:1.12 as builder
COPY . /go/src/github.com/intelematics/terraform-provider-cloudconformity/
WORKDIR /go/src/github.com/intelematics/terraform-provider-cloudconformity/.
RUN CGO_ENABLED=0 go build -o terraform-provider-cloudconformity .

FROM hashicorp/terraform:light as release
RUN mkdir -p /root/.terraform.d/plugins/
COPY --from=builder /go/src/github.com/intelematics/terraform-provider-cloudconformity/terraform-provider-cloudconformity /root/.terraform.d/plugins/terraform-provider-cloudconformity