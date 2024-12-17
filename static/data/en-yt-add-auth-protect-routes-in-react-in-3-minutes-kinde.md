# Add Auth & Protect Routes in React in 3 Minutes (Kinde)

https://www.youtube.com/watch?v=_EjOHdRihjA

- duration: 00:09:37
- language: en
- topics: kinde
- rank: 4.9
- description: quick ByteGrad video showing how to build Kinde into a React site
- year: 2024
- status: finished

## watchlog

09:37 - 2024-11-23

## notes

- using React with React router, going to protect a route
- kinde.com
- Try Kinde Free up to 10,500 monthly active users. No credit card required.
- I signed up for the free account with Google
- start a project from scratch
- answered some questions
- add application
- frontend and mobile for client-side
- Next.js would be backend-web
- picked "React"
- changed URL to that of my site: http://localhost:3663
- existing codebase
- `npm i @kinde-oss/kinde-auth-react`
- main.tsx

```
ReactDOM.createRoot(document.getElementById("root")!).render(
	<StoreProvider store={store}>
		<KindeProvider
			clientId="cdf4911c97a3411296e34c366068fc68"
			domain="https://edwardtanguay.kinde.com"
			redirectUri="http://localhost:3663"
			logoutUri="http://localhost:3663"
		>
			<RouterProvider router={router} />
		</KindeProvider>
	</StoreProvider>
);
```

- PageAbout.tsx

```
import {useKindeAuth} from "@kinde-oss/kinde-auth-react";
const { login, register } = useKindeAuth();
<button onClick={register} type="button">Register</button>
<button onClick={login} type="button">Log In</button>
```

- had an error, but then clicked on the link and "save"
- authentication: switched to Google
- click on users and see who has signed up
- hmm, user after logging in is authenticated
- const { login, register, isAuthenticated } = useKindeAuth();, isAuthenticated is always false
- so in the video, he is authenticate, but I'm not, although we have done the same setup
- create custom protected route component
- element <ProtectedRoute>
- now he fetches data with a bearer token

## VOCAB - SPANISH

```
check it out
compruébalo

security and storage
seguridad y almacenamiento

we want to add authentication
queremos agregar autenticación

```
