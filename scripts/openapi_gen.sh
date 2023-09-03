#!/bin/bash
set -e

readonly output_dir="internal"
readonly package="genapi"
readonly service="service"

oapi-codegen -generate types -o "../$output_dir/$package/openapi_types.gen.go" -package "$package" "../api/openapi/$service.yml"
oapi-codegen -package "$package"  -o "../$output_dir/$package/openapi_api.gen.go" --config=./gen_server_cfg.yaml  "../api/openapi/$service.yml"
