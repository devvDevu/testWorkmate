# testWorkmate

endpoints:
localhost:8080/api/v1/create_task [POST]
request:
{
    "Title": "some_title"
}
response:
{
    "ID": 1,
    "Title": "some_title",
    "Status": "completed",
    "CreatedAt": "11:11:11 0000-00-00",
    "RunTime": "12:00:00"
}
localhost:8080/api/v1/get_task/{id} [GET]
response:
{
    "ID": 1,
    "Title": "some_title",
    "Status": "completed",
    "CreatedAt": "11:11:11 0000-00-00",
    "RunTime": "12:00:00"
}
localhost:8080/api/v1/delete_task/{id} [POST]
response:
{
    "ID": 1,
    "Title": "some_title",
    "Status": "completed",
    "CreatedAt": "11:11:11 0000-00-00",
    "RunTime": "12:00:00"
}
