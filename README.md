
# GoFR API

This is an API made in Golang using GOFR framework to create, search, update and delete complaints for the Resident Welfare Association.



## Run Locally

Clone the project

```bash
  git clone https://link-to-project
```

Go to the project directory

```bash
  cd RWAComplaintAPI
```

Install dependencies

```bash
  go install
```

Start the server

```bash
  go run main.go
```



## API Reference

### Get all complaints

```http
  GET /getcomplaints
```


### Get particular complaint

```http
  GET /complaint/{case_id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `case_id`      | `primitive.ObjectID` | **Required**. case_id of complaint to fetch |

### Add complaints
```http
  POST /addcomplaints
```
#### Body 
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Name of the complaitant |
| `houseno`      | `int32` | **Required**. House no of the complaitant |
| `complaint`      | `string` | **Required**. Description of the complaitant |
| `type`      | `string` | **Required**. Type of the complaitant |
| `allotedto`      | `primitive.ObjectID` |  emp_id of the worker |
| `caseid`      | `primitive.ObjectID` |  case_id of the complaint |

### Update complaints
```http
  PATCH /updatecomp/{case_id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `case_id`      | `primitive.ObjectID` | **Required**. case_id of complaint to update |

#### Body

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` |  Name to be updated |
| `houseno`      | `int32` |  House to be updated |
| `complaint`      | `string` |  Description to be updated |
| `type`      | `string` |  Type to be updated |
| `allotedto`      | `primitive.ObjectID` |  emp_id of the worker to be updated|

### Delete complaints

```http
  DELETE /delcomp/{case_id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `case_id`      | `primitive.ObjectID` | **Required**. case_id to be deleted |

### Get all workers
```http
  GET /getworkers
```

### Get particular worker
```http
  GET /worker/{emp_id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `emp_id`      | `primitive.ObjectID` | **Required**. emp_id to be searched |

### Add worker
```http
  POST /addworkers
```

#### Body 
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Name of the employee |
| `available`      | `bool` |  Avaialability of the worker |
| `expertise`      | `string` | **Required**. Expertise of the worker |
| `empid`      | `primitive.ObjectID` |  emp_id of the worker |
| `assignedcases`      | `[]primitive.ObjectID` |  array with assigned complaints to the worker  |

### Update worker

```http
  PATCH /updateworkers/{emp_id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `emp_id`      | `primitive.ObjectID` | **Required**. emp_id to be updated |

#### Body
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | **Required**. Name of the employee to be updated |
| `available`      | `bool` |  Avaialability of the worker to be updated |
| `expertise`      | `string` | **Required**. Expertise of the worker to be updated |
| `addcaseid`      | `primitive.ObjectID` |  case_id that needs to be added to assignedcases to be updated|
| `removecaseid`      | `primitive.ObjectID` |  case_id that needs to be removed from assignedcases to be updated |


### Delete worker

```http
  DELETE /delworker/{emp_id}
```
| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `emp_id`      | `primitive.ObjectID` | **Required**. emp_id to be deleted |



## Postman Collection

[![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/16926948-24c6a840-4ccb-4a8f-b161-36dcbcc4043b?action=collection%2Ffork&source=rip_markdown&collection-url=entityId%3D16926948-24c6a840-4ccb-4a8f-b161-36dcbcc4043b%26entityType%3Dcollection%26workspaceId%3D95e5319b-d400-4cb2-a57a-dc6ac6349380)
## Features

- CRUD Complaints 
- CRUD Workers
- Auto Worker assignment


## Authors

- [@hs414171](https://github.com/hs414171)

