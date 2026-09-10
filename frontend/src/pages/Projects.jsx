import { useEffect, useState } from "react"
import { api } from "../services/api"

function Projects() {
  const [projects, setProjects] = useState([])
  const [loading, setLoading] = useState(true)
  const [message, setMessage] = useState("")

  const [form, setForm] = useState({
    name: "",
    description: "",
    technologies: "",
    github_url: "",
    live_url: "",
    status: "",
  })

  const [editingId, setEditingId] = useState(null)

  const getProjects = async () => {
    try {
      const data = await api.get("/projects")
      setProjects(data)
    } catch (error) {
      console.error(error)
      setMessage(error.message || "Failed to load projects")
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    getProjects()
  }, [])

  const handleChange = (e) => {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    })
  }

  const handleSubmit = async (e) => {
    e.preventDefault()

    try {
      if (editingId) {
        await api.put(`/projects/${editingId}`, form)
        setMessage("Project updated successfully")
      } else {
        await api.post("/projects", form)
        setMessage("Project created successfully")
      }

      setForm({
        name: "",
        description: "",
        technologies: "",
        github_url: "",
        live_url: "",
        status: "",
      })

      setEditingId(null)
      getProjects()
    } catch (error) {
      console.error(error)
      setMessage(error.message || "Failed to save project")
    }
  }

  const handleEdit = (project) => {
    setEditingId(project.id)

    setForm({
      name: project.name || "",
      description: project.description || "",
      technologies: project.technologies || "",
      github_url: project.github_url || "",
      live_url: project.live_url || "",
      status: project.status || "",
    })
  }

  const handleDelete = async (id) => {
    if (!window.confirm("Are you sure you want to delete this project?")) {
      return
    }

    try {
      await api.delete(`/projects/${id}`)
      setMessage("Project deleted successfully")
      getProjects()
    } catch (error) {
      console.error(error)
      setMessage(error.message || "Failed to delete project")
    }
  }

  const cancelEdit = () => {
    setEditingId(null)

    setForm({
      name: "",
      description: "",
      technologies: "",
      github_url: "",
      live_url: "",
      status: "",
    })
  }

  if (loading) {
    return (
      <div className="flex h-64 items-center justify-center">
        <div className="h-8 w-8 animate-spin rounded-full border-4 border-indigo-600 border-t-transparent"></div>
      </div>
    )
  }

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div className="mb-8 border-b border-slate-200 pb-5">
        <h1 className="text-2xl font-bold tracking-tight text-slate-900 sm:text-3xl">
          Projects
        </h1>

        <p className="mt-1 text-sm text-slate-500">
          Showcase technical builds, experiments, and deployments.
        </p>
      </div>

      {message && (
        <div
          className="mb-6 rounded-lg bg-indigo-50 border border-indigo-100 p-4 text-sm text-indigo-700 flex justify-between"
          role="status"
        >
          <span>{message}</span>

          <button
            type="button"
            onClick={() => setMessage("")}
            className="font-semibold hover:underline"
          >
            Dismiss
          </button>
        </div>
      )}

      <div className="grid grid-cols-1 lg:grid-cols-12 gap-8 items-start">
        {/* Project Form */}
        <div className="lg:col-span-4 bg-white p-6 rounded-xl border border-slate-200/80 shadow-sm sticky top-6">
          <h2 className="text-lg font-semibold text-slate-900 mb-4 pb-2 border-b border-slate-100">
            {editingId ? "Edit Project" : "Add Project"}
          </h2>

          <form onSubmit={handleSubmit} className="space-y-4">
            {/* Project Name */}
            <div>
              <label
                htmlFor="name"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                Project Name
              </label>

              <input
                id="name"
                name="name"
                value={form.name}
                onChange={handleChange}
                placeholder="e.g. Distributed Task Queue"
                required
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            {/* Description */}
            <div>
              <label
                htmlFor="description"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                Description
              </label>

              <textarea
                id="description"
                rows={3}
                name="description"
                value={form.description}
                onChange={handleChange}
                placeholder="What did you build and solve?"
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            {/* Technologies */}
            <div>
              <label
                htmlFor="technologies"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                Tech Stack (comma-separated)
              </label>

              <input
                id="technologies"
                name="technologies"
                value={form.technologies}
                onChange={handleChange}
                placeholder="Go, React, PostgreSQL, Docker"
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            {/* GitHub URL */}
            <div>
              <label
                htmlFor="github_url"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                GitHub URL
              </label>

              <input
                id="github_url"
                type="url"
                name="github_url"
                value={form.github_url}
                onChange={handleChange}
                placeholder="https://github.com/..."
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            {/* Live URL */}
            <div>
              <label
                htmlFor="live_url"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                Live URL
              </label>

              <input
                id="live_url"
                type="url"
                name="live_url"
                value={form.live_url}
                onChange={handleChange}
                placeholder="https://myproject.app"
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            {/* Status */}
            <div>
              <label
                htmlFor="status"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                Status
              </label>

              <select
                id="status"
                name="status"
                value={form.status}
                onChange={handleChange}
                required
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 bg-white focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              >
                <option value="">Select status</option>
                <option value="In Progress">In Progress</option>
                <option value="Completed">Completed</option>
                <option value="Archived">Archived</option>
              </select>
            </div>

            {/* Buttons */}
            <div className="flex gap-2 pt-2">
              <button
                type="submit"
                className="flex-1 bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-2 px-4 rounded-lg text-sm transition-all shadow-sm"
              >
                {editingId ? "Update Project" : "Create Project"}
              </button>

              {editingId && (
                <button
                  type="button"
                  onClick={cancelEdit}
                  className="bg-slate-100 hover:bg-slate-200 text-slate-700 font-medium py-2 px-4 rounded-lg text-sm"
                >
                  Cancel
                </button>
              )}
            </div>
          </form>
        </div>

        {/* Project Gallery Cards */}
        <div className="lg:col-span-8">
          {projects.length === 0 ? (
            <div className="text-center py-16 px-4 bg-white border border-dashed border-slate-300 rounded-xl">
              <p className="text-slate-500 text-sm">
                No projects added yet.
              </p>
            </div>
          ) : (
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {projects.map((project) => (
                <div
                  key={project.id}
                  className="bg-white p-5 rounded-xl border border-slate-200/80 shadow-sm flex flex-col justify-between hover:shadow-md transition-shadow"
                >
                  <div>
                    <div className="flex items-start justify-between gap-2 mb-2">
                      <h3 className="font-semibold text-slate-900 text-base">
                        {project.name}
                      </h3>

                      <span
                        className={`text-[11px] font-medium px-2 py-0.5 rounded-full border ${
                          project.status === "Completed"
                            ? "bg-emerald-50 text-emerald-700 border-emerald-200"
                            : project.status === "In Progress"
                            ? "bg-amber-50 text-amber-700 border-amber-200"
                            : "bg-slate-50 text-slate-600 border-slate-200"
                        }`}
                      >
                        {project.status}
                      </span>
                    </div>

                    <p className="text-slate-600 text-xs line-clamp-3 mb-4">
                      {project.description}
                    </p>

                    {/* Tech Badges */}
                    {project.technologies && (
                      <div className="flex flex-wrap gap-1.5 mb-4">
                        {project.technologies
                          .split(",")
                          .map((tech, i) => (
                            <span
                              key={i}
                              className="bg-slate-100 text-slate-700 text-[11px] px-2 py-0.5 rounded-md font-medium"
                            >
                              {tech.trim()}
                            </span>
                          ))}
                      </div>
                    )}
                  </div>

                  <div className="border-t border-slate-100 pt-3 flex items-center justify-between text-xs">
                    <div className="flex gap-2">
                      {project.github_url && (
                        <a
                          href={project.github_url}
                          target="_blank"
                          rel="noreferrer"
                          className="font-medium text-slate-700 hover:text-indigo-600"
                        >
                          Source ↗
                        </a>
                      )}

                      {project.live_url && (
                        <a
                          href={project.live_url}
                          target="_blank"
                          rel="noreferrer"
                          className="font-medium text-indigo-600 hover:text-indigo-800"
                        >
                          Demo ↗
                        </a>
                      )}
                    </div>

                    <div className="flex gap-3">
                      <button
                        type="button"
                        onClick={() => handleEdit(project)}
                        className="text-slate-500 hover:text-slate-900 font-medium"
                      >
                        Edit
                      </button>

                      <button
                        type="button"
                        onClick={() => handleDelete(project.id)}
                        className="text-rose-500 hover:text-rose-700 font-medium"
                      >
                        Delete
                      </button>
                    </div>
                  </div>
                </div>
              ))}
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default Projects