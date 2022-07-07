# Simple Video chat app

> Video chat application build with Vue and Go.

## Prerequisites

- [Node](https://nodejs.org/en/download/)
- [Go](https://go.dev/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)

# Frontend setup

```sh
cd frontend
```

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```

### Type-Check, Compile and Minify for Production

```sh
npm run build
```

### Lint with [ESLint](https://eslint.org/)

```sh
npm run lint
```

# Backend Setup

- Create new PostgreSQL database;
- `cd backend`
- Then run this command with your PostgreSQL username and database name:

```sh
psql -U <username> -d <database name> -f db.sql
```

- Create your enviroment variables according to `.env-sample` and `config-sample.json`
- To start server run:

```sh
go run main.go
```

## Contributors

<a href="https://github.com/Seydulla"><img src="https://avatars.githubusercontent.com/u/74240235?v=4" width="50" style="margin-right:10px; border-radius:50%" /></a>
<a href="https://github.com/ResulShamuhammedov"><img src="https://avatars.githubusercontent.com/u/85562280?v=4" width="50" style="border-radius:50%" /></a>

## Features

- ✅ P2P video call;
- ⭕ Invite friends with email;
- ⭕ Chat;
- and whatever comes to our crazy minds😜
