import {
  BrowserRouter,
  Routes,
  Route,
  Navigate,
} from "react-router-dom"

import { useAuth } from "./context/AuthContext"

import Login from "./pages/Login"
import Dashboard from "./pages/Dashboard"
import Profile from "./pages/Profile"
import Projects from "./pages/Projects"
import Skills from "./pages/Skills"
import Applications from "./pages/Applications"
import Register from "./pages/Register"

import Layout from "./components/Layout"

function App() {
  const { isLoggedIn } = useAuth()

  return (
    <BrowserRouter>
      <Routes>

        <Route
          path="/login"
          element={
            isLoggedIn
              ? <Navigate to="/dashboard" />
              : <Login />
          }
        />
        <Route
  path="/register"
  element={
    isLoggedIn
      ? <Navigate to="/dashboard" />
      : <Register />
  }
/>

        <Route
          element={
            isLoggedIn
              ? <Layout />
              : <Navigate to="/login" />
          }
        >

          <Route
            path="/dashboard"
            element={<Dashboard />}
          />

          <Route
            path="/profile"
            element={<Profile />}
          />

          <Route
            path="/projects"
            element={<Projects />}
          />

          <Route
            path="/skills"
            element={<Skills />}
          />

          <Route
            path="/applications"
            element={<Applications />}
          />

        </Route>

        <Route
          path="*"
          element={
            <Navigate
              to={isLoggedIn ? "/dashboard" : "/login"}
            />
          }
        />

      </Routes>
    </BrowserRouter>
  )
}

export default App