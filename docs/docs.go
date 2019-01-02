// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-01-03 01:01:26.381791276 +0800 CST m=+0.030696399

package docs

import (
	"github.com/swaggo/swag"
)

var doc = `{
    "swagger": "2.0",
    "info": {
        "description": "云盘的 Api 服务.",
        "title": "云盘 Api 服务",
        "termsOfService": "https://github.com/zm-dev",
        "contact": {
            "name": "API Support",
            "url": "https://github.com/wq1019",
            "email": "2013855675@qq.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/api",
    "paths": {
        "/file/rename": {
            "put": {
                "description": "通过文件 ID 重命名文件",
                "consumes": [
                    "application/json",
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json",
                    "multipart/form-data"
                ],
                "tags": [
                    "文件"
                ],
                "summary": "重命名文件",
                "operationId": "rename-file",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "uint64",
                        "description": "文件 ID",
                        "name": "file_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "format": "uint64",
                        "description": "文件所属的目录 ID",
                        "name": "folder_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "format": "string",
                        "description": "新的文件名",
                        "name": "new_name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {},
                    "404": {
                        "description": "文件不存在\" | \"目录不存在",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/errors.GlobalError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.GlobalError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 10001
                },
                "inner_err": {
                    "type": "error"
                },
                "message": {
                    "type": "string",
                    "example": "文件不存在"
                },
                "service_name": {
                    "type": "string",
                    "example": "cloud_disk"
                },
                "status_code": {
                    "type": "integer",
                    "example": 500
                }
            }
        }
    }
}`

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swag.Register(swag.Name, &s{})
}
