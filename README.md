<p align="center">
<a href="https://dscvit.com">
	<img src="https://user-images.githubusercontent.com/30529572/92081025-fabe6f00-edb1-11ea-9169-4a8a61a5dd45.png" alt="DSC VIT"/>
</a>
	<h2 align="center"> LocalB Backend </h2>
	<h4 align="center"> A backend written for the localb project in Go with Fiber, Postgresql and using Clean architecture. </h4>
</p>

---
[![DOCS](https://img.shields.io/badge/Documentation-see%20docs-green?style=flat-square&logo=appveyor)](INSERT_LINK_FOR_DOCS_HERE) 
[![Join Us](https://img.shields.io/badge/Join%20Us-Developer%20Student%20Clubs-red)](https://dsc.community.dev/vellore-institute-of-technology/)

## Functionalities

### Client
- [x] Add a new business
- [x] Show all businesses
- [x] Show businesses for a city
- [x] Paginate all responses
- [ ] Report a business
- [x] Find by types

### Admin
- [x] Admin login
- [x] View all businesses to approve
- [x] Approve a business
- [x] Delete a business
- [ ] Paginate all responses

<br>


## Instructions to run

* Pre-requisites:
	-  Go >= 1.14
	-  Add .env to project root

* Example .env
 
| Key | Value |
|-----|-------|
| jwtSecret | secret_key |
| onServer | False |
| dbHost | localhost |
| dbPort | 5432 |
| dbUser | postgres |
| dbPass | password |
| dbName | dbname |
| sslmode | disable |

* Directions to execute

```bash
go mod download
go run main.go
```

## Contributors

<table>
<tr align="center">


<td>

Rithik Jain

<p align="center">
<img src = "https://avatars2.githubusercontent.com/u/12408595?s=460&u=8c49665f477bda73c00473dd3f5131156a5ecc31&v=4" width="150" height="150" alt="Your Name Here (Insert Your Image Link In Src">
</p>
<p align="center">
<a href = "https://github.com/rithikjain"><img src = "http://www.iconninja.com/files/241/825/211/round-collaboration-social-github-code-circle-network-icon.svg" width="36" height = "36" alt="GitHub"/></a>
<a href = "https://www.linkedin.com/in/rithik-jain-710b3a199/">
<img src = "http://www.iconninja.com/files/863/607/751/network-linkedin-social-connection-circular-circle-media-icon.svg" width="36" height="36" alt="LinkedIn"/>
</a>
</p>
</td>

</tr>
  </table>

<p align="center">
	Made with :heart: by <a href="https://dscvit.com">DSC VIT</a>
</p>

