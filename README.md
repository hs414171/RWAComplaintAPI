
# GoFR API

The AVRWA_COMPLAINT API provides endpoints to manage complaints and workers within a Residential Welfare Association (RWA) system.

## Base URL

```http
  http://loacalhost:5000

```

## Run Locally

Clone the project

```bash
  git clone https://github.com/hs414171/RWAComplaintAPI.git
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

## Sequence Diagrams

### POST /addcomplaints
![App Screenshot](https://www.planttext.com/api/plantuml/svg/bPDDJyCm38Rl_HM-SPguxu3sqV4X8HXqWxFNUCWYIqlY5lRlIRgrIocJe9UeR-mdVXrduWIjL3kTb7xbd4iq0dFDPAGdPMHhiZtfjR3lP5qww-SMkr8Z6maRoO44ewiZO0otgzKIbe_P2YvHgRoBJZ0Nhb6eFUgmnAcBBx4-2j0eU47DAaHtvq6AWWbgQVUuC91LUKxERIkTWF-6p1ioP0DdRVYZyZXqT9V-2EwSZk6f8hjl1TWMrX_1WXjDhTZaTcPEJGfoznUNSCzEeDZ-3b9q1ZMqaf6clBqMc7aZ2dmBvfQ5BEDvJ-jIdVWN3gREyPlnsQ4Phcdtj4RMk67Dil-tmHuVKXBF-hExU9wyHwnxLnwnEb9l0iDEBC0eAcnUc2tRNJCgF-iPKDgIGatDL8_XWJxr7gPAXLuDYYBjA9x5Qw-IEiTQMEk7wPbSMHX72HcLHDKStyJ4o-6nVWC0)

### POST /addworkers
![App Screenshot](https://www.planttext.com/api/plantuml/svg/TL6xRiCm3Dpr5Vo18Bk78i-YKnF4HZeJ9WC8LII3efFsxoDRjkCCqKBeTCSxKgkeADfUdMbIjRLjq2jidMMlBr39ScDvWNw2_BB4tSQOB835ny0huAkgZd0yb1KiaUYUgGPhjJTK7jlbCjVj_rGI2z0JdD5PwjZP2NQeQD3f-623XxAjQuxnqZe3wJmRJkaAM1CDrvupXNWEqC7J3HQpFDlWBrQklLbF8twZ7wsJEA-ZW6--k__EdccQWHMjYuPTvFaSdmbbUjTeBn84JXoRu2CRzcHcEOPbhXASVl81)

## Use Case Diagram
![UseCaseDiagram](https://www.planttext.com/api/plantuml/svg/RP312i8m38RlUOgm-mxEdMLbWWyWPUUonbdisj5c1n7VtKsh5lJsvUVx4qWJMGMUHmkMJepiCPZXpAXDe9wDTw1wzW5bf80geanpvyFbbNGf4NeqRWq4mDfLRh24nsja6l46Y7EaNaN2GcYDhHydgupZuhcrtUI2dBHY-d6gtDcJfJ_EDdSEn5yX9gE6HP5ObhjqzirLnjkCwuMxZ4NM5LbNVBftW8QSNhxw0000)

## Authors

- [@hs414171](https://github.com/hs414171)

