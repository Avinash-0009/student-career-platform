
import { Link, Outlet, useNavigate } from "react-router-dom"
import { useAuth } from "../context/AuthContext"

function Layout() {
    const { logout } = useAuth()
    const navigate = useNavigate()
  
   return (
    <div className="app-layout">

      <aside className="sidebar">
        <h2>Career Platform</h2>

        <nav>
          <Link to="/dashboard">Dashboard</Link>
          <Link to="/profile">Profile</Link>
          <Link to="/projects">Projects</Link>
          <Link to="/skills">Skills</Link>
          <Link to="/applications">Applications</Link>
        </nav>

        <button
         onClick={() => {
         logout()
         navigate("/login")
         }}
         >
         Logout
       </button>
      </aside>

      <main className="main-content">
        <Outlet />
      </main>

    </div>
  )
}

export default Layout