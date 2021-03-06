// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the API for SourcePods - git in the cloud.",
    "title": "SourcePods OpenAPI",
    "license": {
      "name": "Apache-2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/repositories": {
      "post": {
        "tags": [
          "repositories"
        ],
        "summary": "Create a new repository",
        "operationId": "createRepository",
        "parameters": [
          {
            "description": "The repository to create",
            "name": "newRepository",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "name"
              ],
              "properties": {
                "description": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "website": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The repository has been created and is returned to you",
            "schema": {
              "$ref": "#/definitions/repository"
            }
          },
          "422": {
            "description": "The new repository has not been created due to invalid input",
            "schema": {
              "$ref": "#/definitions/validationError"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/repositories/{owner}": {
      "get": {
        "tags": [
          "repositories"
        ],
        "summary": "Get a owner's repositories",
        "operationId": "getOwnerRepositories",
        "parameters": [
          {
            "type": "string",
            "description": "The owner's username",
            "name": "owner",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The repositories found by its owner name",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/repository"
              }
            }
          },
          "404": {
            "description": "The owner could not be found by this username",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/repositories/{owner}/{name}": {
      "get": {
        "tags": [
          "repositories"
        ],
        "summary": "Get a repository by owner name and its name",
        "operationId": "getRepository",
        "parameters": [
          {
            "type": "string",
            "description": "The owner's username",
            "name": "owner",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The repository's name",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The repository found by its owner and name",
            "schema": {
              "$ref": "#/definitions/repository"
            }
          },
          "404": {
            "description": "The owner and name combination could not be found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/repositories/{owner}/{name}/branches": {
      "get": {
        "tags": [
          "repositories"
        ],
        "summary": "Get all branches of a repository",
        "operationId": "getRepositoryBranches",
        "parameters": [
          {
            "type": "string",
            "description": "The owner's username",
            "name": "owner",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The repository's name",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The repository's branches",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/branch"
              }
            }
          },
          "404": {
            "description": "The owner and name combination could not be found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "tags": [
          "users"
        ],
        "summary": "List all users",
        "operationId": "listUsers",
        "responses": {
          "200": {
            "description": "An array of all users",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users/me": {
      "get": {
        "tags": [
          "users"
        ],
        "summary": "Get the current authenticated user",
        "operationId": "getUserMe",
        "responses": {
          "200": {
            "description": "The current authenticated user",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users/{username}": {
      "get": {
        "tags": [
          "users"
        ],
        "summary": "Get a user by their username",
        "operationId": "getUser",
        "parameters": [
          {
            "type": "string",
            "description": "The username of a user",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The user by their username",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "404": {
            "description": "The user is not found by their username",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "tags": [
          "users"
        ],
        "summary": "Update the user's information",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "The username of the user to update",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "The updated user",
            "name": "updatedUser",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "name"
              ],
              "properties": {
                "name": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The user has been updated",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "404": {
            "description": "The user could not be found by this username",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "422": {
            "description": "The updated user has invalid input",
            "schema": {
              "$ref": "#/definitions/validationError"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "branch": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "sha1": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "repository": {
      "type": "object",
      "required": [
        "id",
        "name"
      ],
      "properties": {
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "default_branch": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string",
          "format": "uuid",
          "readOnly": true
        },
        "name": {
          "type": "string"
        },
        "owner": {
          "type": "object",
          "$ref": "#/definitions/user"
        },
        "updated_at": {
          "type": "string",
          "format": "date-time"
        },
        "website": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "id",
        "username"
      ],
      "properties": {
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "email": {
          "type": "string",
          "format": "email"
        },
        "id": {
          "type": "string",
          "format": "uuid",
          "readOnly": true
        },
        "name": {
          "type": "string"
        },
        "updated_at": {
          "type": "string",
          "format": "date-time"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "validationError": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "errors": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "field": {
                "type": "string"
              },
              "message": {
                "type": "string"
              }
            }
          }
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the API for SourcePods - git in the cloud.",
    "title": "SourcePods OpenAPI",
    "license": {
      "name": "Apache-2.0",
      "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
    },
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/repositories": {
      "post": {
        "tags": [
          "repositories"
        ],
        "summary": "Create a new repository",
        "operationId": "createRepository",
        "parameters": [
          {
            "description": "The repository to create",
            "name": "newRepository",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "name"
              ],
              "properties": {
                "description": {
                  "type": "string"
                },
                "name": {
                  "type": "string"
                },
                "website": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The repository has been created and is returned to you",
            "schema": {
              "$ref": "#/definitions/repository"
            }
          },
          "422": {
            "description": "The new repository has not been created due to invalid input",
            "schema": {
              "$ref": "#/definitions/validationError"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/repositories/{owner}": {
      "get": {
        "tags": [
          "repositories"
        ],
        "summary": "Get a owner's repositories",
        "operationId": "getOwnerRepositories",
        "parameters": [
          {
            "type": "string",
            "description": "The owner's username",
            "name": "owner",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The repositories found by its owner name",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/repository"
              }
            }
          },
          "404": {
            "description": "The owner could not be found by this username",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/repositories/{owner}/{name}": {
      "get": {
        "tags": [
          "repositories"
        ],
        "summary": "Get a repository by owner name and its name",
        "operationId": "getRepository",
        "parameters": [
          {
            "type": "string",
            "description": "The owner's username",
            "name": "owner",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The repository's name",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The repository found by its owner and name",
            "schema": {
              "$ref": "#/definitions/repository"
            }
          },
          "404": {
            "description": "The owner and name combination could not be found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/repositories/{owner}/{name}/branches": {
      "get": {
        "tags": [
          "repositories"
        ],
        "summary": "Get all branches of a repository",
        "operationId": "getRepositoryBranches",
        "parameters": [
          {
            "type": "string",
            "description": "The owner's username",
            "name": "owner",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "The repository's name",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The repository's branches",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/branch"
              }
            }
          },
          "404": {
            "description": "The owner and name combination could not be found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "tags": [
          "users"
        ],
        "summary": "List all users",
        "operationId": "listUsers",
        "responses": {
          "200": {
            "description": "An array of all users",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user"
              }
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users/me": {
      "get": {
        "tags": [
          "users"
        ],
        "summary": "Get the current authenticated user",
        "operationId": "getUserMe",
        "responses": {
          "200": {
            "description": "The current authenticated user",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users/{username}": {
      "get": {
        "tags": [
          "users"
        ],
        "summary": "Get a user by their username",
        "operationId": "getUser",
        "parameters": [
          {
            "type": "string",
            "description": "The username of a user",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "The user by their username",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "404": {
            "description": "The user is not found by their username",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "patch": {
        "tags": [
          "users"
        ],
        "summary": "Update the user's information",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "The username of the user to update",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "The updated user",
            "name": "updatedUser",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "required": [
                "name"
              ],
              "properties": {
                "name": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "The user has been updated",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "404": {
            "description": "The user could not be found by this username",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "422": {
            "description": "The updated user has invalid input",
            "schema": {
              "$ref": "#/definitions/validationError"
            }
          },
          "default": {
            "description": "unexpected error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "branch": {
      "type": "object",
      "properties": {
        "name": {
          "type": "string"
        },
        "sha1": {
          "type": "string"
        },
        "type": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "message": {
          "type": "string"
        }
      }
    },
    "repository": {
      "type": "object",
      "required": [
        "id",
        "name"
      ],
      "properties": {
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "default_branch": {
          "type": "string"
        },
        "description": {
          "type": "string"
        },
        "id": {
          "type": "string",
          "format": "uuid",
          "readOnly": true
        },
        "name": {
          "type": "string"
        },
        "owner": {
          "type": "object",
          "$ref": "#/definitions/user"
        },
        "updated_at": {
          "type": "string",
          "format": "date-time"
        },
        "website": {
          "type": "string"
        }
      }
    },
    "user": {
      "type": "object",
      "required": [
        "id",
        "username"
      ],
      "properties": {
        "created_at": {
          "type": "string",
          "format": "date-time"
        },
        "email": {
          "type": "string",
          "format": "email"
        },
        "id": {
          "type": "string",
          "format": "uuid",
          "readOnly": true
        },
        "name": {
          "type": "string"
        },
        "updated_at": {
          "type": "string",
          "format": "date-time"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "validationError": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "errors": {
          "type": "array",
          "items": {
            "type": "object",
            "properties": {
              "field": {
                "type": "string"
              },
              "message": {
                "type": "string"
              }
            }
          }
        },
        "message": {
          "type": "string"
        }
      }
    }
  }
}`))
}
