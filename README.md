# Hotel item service API  (GraphQL) #
Live at(until 28.Nov.2022. heroku has no more free plans): https://master-thesis-ci-cd.herokuapp.com/

Paste this token to section at bottom of the query window inside "REQUEST HEADERS", and You are ready to use service. 

```
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2NvdW50SWQiOiIxIiwiZXhwaXJlc19BdCI6IjIwMjItMDQtMjFUMTI6MDQ6MDMuMzY3MjE4MiswMjowMCIsImlzc3VlZEF0IjoxNjUwNDQ5MDQzfQ.5Rmy1KdyoJEsTrXOz0FhZpQy_AunAlQk0UkYSlCoBk4"
}
```

# Query and Mutation example #
```
query GetItems {
  getItems{
    id,
    typeId,
    lon,
    lat,
    brokenId,
    deleteStatus
  }
}
```

```
mutation CreateItem {
  createItem(Input: 
    {
      typeId: 2, 
      lon: 2.6, 
      lat: 3.2, 
      brokenId: 1
    })
}
```


