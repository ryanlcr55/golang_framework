#!/bin/bash
set -e

readonly output_dir="internal/genapi"
readonly package="ports"
readonly service="service"

oapi-codegen -generate types -o "$output_dir/openapi_types.gen.go" -package "$package" "api/openapi/$service.yml"
oapi-codegen -generate gin -o "$output_dir/openapi_api.gen.go" -package "$package" "api/openapi/$service.yml"