### to do service-user-investor


- Admin req
- [x] ~GET Log service
- [x] GET Service status

- [x] POST Register
    - [x] POST Check email
    - [x] POST Check Phone
- [x] POST Login


- Dashboard
- [ ] GET User Profile
- [ ] POST Update_avatar
- [ ] PUT Update User profile

- [ ] POST Logout

# Info

Make database

`migrate create -ext sql -dir database/migrations nama_file_migration`

Run Migrate

```
migrate -database "mysql://root@tcp(127.0.0.1:3306)/service_user_investor" -path database/migrations up
```
