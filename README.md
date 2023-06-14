### to do service-user-investor


- Admin req

- [x]CORS CONFIG
- [x] Fix config: log and cors

- [] Setup Middleware Auth user(unix_id)
- [] Setup Middleware Auth admin(unix_id with os.Getenv("ADMIN_ID"))

- [x] ~GET Log service
    - [ ] Auth middleware admin, Uri(unix_id on env), 
- [x] GET Service status

- [x] POST deactive user by admin
    - [x] with uri unix_id admin, use fetch get to service-user-investor and save to db
- [ ] POST active admin
    - [ ] with midlleware auth admin

- [x] POST Register
    - [x] POST Check email
    - [x] POST Check Phone
- [x] POST Login


- Dashboard

- [ ] PUT Update User profile admin
    - Update data can be empty content
    - update with name, email, phone
- [ ] GET User Profile
    - with middleware on token

- [ ] GET User Profile
- [ ] POST Update_avatar
- [ ] PUT Update User profile

- [ ] POST Logout
    - Delete token 

### CI/CD Github Actions

- [] ~CI/CD Github Actions
- [] Push to Docker Hub
- [] Push to GCP registry