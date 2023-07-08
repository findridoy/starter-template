import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import { RouterProvider, createBrowserRouter, redirect } from 'react-router-dom'
import Login, { LoginRouteAction } from './routes/login.tsx'
import { isLoggedIn } from './lib/auth.ts'
import Layout, { LayoutRouteAction } from './components/Layout.tsx'
import Dashboard from './routes/dashboard.tsx'
import UserIndex from './routes/users/index.tsx'
import { userList } from './repository/user.ts'

const router = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    loader: async ({request}) => {
        if (!await isLoggedIn()) return redirect('/login')
        if (new URL(request.url).pathname === "/") return redirect('/dashboard')
        return null
    },
    action: LayoutRouteAction,
    children: [
      {
        path: "/dashboard",
        element: <Dashboard/>,
        loader: async () => {
          if (!await isLoggedIn()) {
            return redirect("/login")
          }
          return null
        },
      },
      {
        path: "/users",
        element: <UserIndex/>,
        loader: async () => {
          if (!await isLoggedIn()) {
            return redirect("/login")
          }
          return await userList()
        },
      },
    ]
  },

  {
    path: "/login",
    element: <Login />,
    action: LoginRouteAction,
  }
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
