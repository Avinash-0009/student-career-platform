import { useEffect, useState } from "react"
import { api } from "../services/api"

const STATUS_CONFIG = {
  Applied: {
    bg: "bg-blue-50 text-blue-700 border-blue-200",
    dot: "bg-blue-500",
  },
  Assessment: {
    bg: "bg-amber-50 text-amber-700 border-amber-200",
    dot: "bg-amber-500",
  },
  Interview: {
    bg: "bg-purple-50 text-purple-700 border-purple-200",
    dot: "bg-purple-500",
  },
  Offer: {
    bg: "bg-emerald-50 text-emerald-700 border-emerald-200",
    dot: "bg-emerald-500",
  },
  Rejected: {
    bg: "bg-rose-50 text-rose-700 border-rose-200",
    dot: "bg-rose-500",
  },
}

function Applications() {
  const [applications, setApplications] = useState([])
  const [loading, setLoading] = useState(true)
  const [message, setMessage] = useState("")
  const [editingId, setEditingId] = useState(null)

  const [form, setForm] = useState({
    company: "",
    job_title: "",
    job_url: "",
    status: "",
    applied_date: "",
    deadline: "",
    notes: "",
  })

  const getApplications = async () => {
    try {
      const data = await api.get("/applications")
      setApplications(data)
    } catch (error) {
      console.error(error)
      setMessage(error.message || "Failed to load applications")
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    getApplications()
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
      const applicationData = {
        company: form.company,
        job_title: form.job_title,
        job_url: form.job_url,
        status: form.status,
        applied_date: form.applied_date
          ? new Date(form.applied_date).toISOString()
          : null,
        deadline: form.deadline
          ? new Date(form.deadline).toISOString()
          : null,
        notes: form.notes,
      }

      if (editingId) {
        await api.put(`/applications/${editingId}`, applicationData)
        setMessage("Application updated successfully")
      } else {
        await api.post("/applications", applicationData)
        setMessage("Application added successfully")
      }

      resetForm()
      getApplications()
    } catch (error) {
      console.error(error)
      setMessage(error.message || "Failed to save application")
    }
  }

  const handleEdit = (application) => {
    setEditingId(application.id)

    setForm({
      company: application.company || "",
      job_title: application.job_title || "",
      job_url: application.job_url || "",
      status: application.status || "",
      applied_date: application.applied_date
        ? application.applied_date.slice(0, 10)
        : "",
      deadline: application.deadline
        ? application.deadline.slice(0, 10)
        : "",
      notes: application.notes || "",
    })
  }

  const handleDelete = async (id) => {
    if (
      !window.confirm(
        "Are you sure you want to delete this application?"
      )
    ) {
      return
    }

    try {
      await api.delete(`/applications/${id}`)
      setMessage("Application deleted successfully")
      getApplications()
    } catch (error) {
      console.error(error)
      setMessage(error.message || "Failed to delete application")
    }
  }

  const resetForm = () => {
    setEditingId(null)

    setForm({
      company: "",
      job_title: "",
      job_url: "",
      status: "",
      applied_date: "",
      deadline: "",
      notes: "",
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
      {/* Page Header */}
      <div className="mb-8 border-b border-slate-200 pb-5">
        <h1 className="text-2xl font-bold tracking-tight text-slate-900 sm:text-3xl">
          Applications
        </h1>

        <p className="mt-1 text-sm text-slate-500">
          Track and manage your internship and job pipeline.
        </p>
      </div>

      {message && (
        <div
          className="mb-6 rounded-lg bg-indigo-50 border border-indigo-100 p-4 text-sm text-indigo-700 flex items-center justify-between"
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
        {/* Form Panel */}
        <div className="lg:col-span-5 bg-white p-6 rounded-xl border border-slate-200/80 shadow-sm sticky top-6">
          <h2 className="text-lg font-semibold text-slate-900 mb-4 pb-2 border-b border-slate-100">
            {editingId ? "Edit Application" : "New Application"}
          </h2>

          <form onSubmit={handleSubmit} className="space-y-4">
            {/* Company */}
            <div>
              <label
                htmlFor="company"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                Company
              </label>

              <input
                id="company"
                name="company"
                value={form.company}
                onChange={handleChange}
                placeholder="e.g. Google"
                required
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            {/* Job Title */}
            <div>
              <label
                htmlFor="job_title"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                Job Title
              </label>

              <input
                id="job_title"
                name="job_title"
                value={form.job_title}
                onChange={handleChange}
                placeholder="e.g. Software Engineer Intern"
                required
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            {/* Job URL */}
            <div>
              <label
                htmlFor="job_url"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                Job Posting URL
              </label>

              <input
                id="job_url"
                type="url"
                name="job_url"
                value={form.job_url}
                onChange={handleChange}
                placeholder="https://..."
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
                <option value="Applied">Applied</option>
                <option value="Assessment">Assessment</option>
                <option value="Interview">Interview</option>
                <option value="Offer">Offer</option>
                <option value="Rejected">Rejected</option>
              </select>
            </div>

            {/* Dates */}
            <div className="grid grid-cols-2 gap-3">
              <div>
                <label
                  htmlFor="applied_date"
                  className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
                >
                  Applied Date
                </label>

                <input
                  id="applied_date"
                  type="date"
                  name="applied_date"
                  value={form.applied_date}
                  onChange={handleChange}
                  className="w-full rounded-lg border border-slate-300 px-3 py-2 text-sm text-slate-800 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
                />
              </div>

              <div>
                <label
                  htmlFor="deadline"
                  className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
                >
                  Deadline
                </label>

                <input
                  id="deadline"
                  type="date"
                  name="deadline"
                  value={form.deadline}
                  onChange={handleChange}
                  className="w-full rounded-lg border border-slate-300 px-3 py-2 text-sm text-slate-800 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
                />
              </div>
            </div>

            {/* Notes */}
            <div>
              <label
                htmlFor="notes"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                Notes
              </label>

              <textarea
                id="notes"
                rows={3}
                name="notes"
                value={form.notes}
                onChange={handleChange}
                placeholder="Follow-ups, interview notes, contacts..."
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            {/* Buttons */}
            <div className="flex gap-2 pt-2">
              <button
                type="submit"
                className="flex-1 bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-2 px-4 rounded-lg text-sm transition-all shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
              >
                {editingId ? "Update Application" : "Save Application"}
              </button>

              {editingId && (
                <button
                  type="button"
                  onClick={resetForm}
                  className="bg-slate-100 hover:bg-slate-200 text-slate-700 font-medium py-2 px-4 rounded-lg text-sm transition-all"
                >
                  Cancel
                </button>
              )}
            </div>
          </form>
        </div>

        {/* Applications Stream */}
        <div className="lg:col-span-7">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-lg font-semibold text-slate-900">
              Your Records ({applications.length})
            </h2>
          </div>

          {applications.length === 0 ? (
            <div className="text-center py-16 px-4 bg-white border border-dashed border-slate-300 rounded-xl">
              <p className="text-slate-500 text-sm">
                No applications recorded yet. Fill out the form to add one.
              </p>
            </div>
          ) : (
            <div className="space-y-3">
              {applications.map((app) => {
                const badgeStyle = STATUS_CONFIG[app.status] || {
                  bg: "bg-slate-100 text-slate-700 border-slate-200",
                  dot: "bg-slate-400",
                }

                return (
                  <div
                    key={app.id}
                    className="bg-white p-5 rounded-xl border border-slate-200/80 shadow-sm hover:shadow-md transition-shadow"
                  >
                    <div className="flex items-start justify-between gap-4">
                      <div>
                        <h3 className="font-semibold text-slate-900 leading-tight">
                          {app.job_title}
                        </h3>

                        <p className="text-sm font-medium text-slate-600">
                          {app.company}
                        </p>
                      </div>

                      <span
                        className={`inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-semibold border ${badgeStyle.bg}`}
                      >
                        <span
                          className={`h-1.5 w-1.5 rounded-full ${badgeStyle.dot}`}
                        />

                        {app.status || "Unknown"}
                      </span>
                    </div>

                    {app.notes && (
                      <p className="mt-3 text-xs text-slate-600 bg-slate-50 p-2.5 rounded-md border border-slate-100 italic">
                        "{app.notes}"
                      </p>
                    )}

                    <div className="mt-4 pt-3 border-t border-slate-100 flex flex-wrap items-center justify-between gap-2 text-xs text-slate-500">
                      <div className="flex gap-3">
                        {app.applied_date && (
                          <span>
                            Applied:{" "}
                            <strong>
                              {new Date(
                                app.applied_date
                              ).toLocaleDateString()}
                            </strong>
                          </span>
                        )}

                        {app.deadline && (
                          <span className="text-rose-600 font-medium">
                            Due:{" "}
                            {new Date(
                              app.deadline
                            ).toLocaleDateString()}
                          </span>
                        )}
                      </div>

                      <div className="flex items-center gap-3">
                        {app.job_url && (
                          <a
                            href={app.job_url}
                            target="_blank"
                            rel="noreferrer"
                            className="text-indigo-600 hover:text-indigo-800 font-medium underline inline-flex items-center"
                          >
                            Listing ↗
                          </a>
                        )}

                        <button
                          type="button"
                          onClick={() => handleEdit(app)}
                          className="hover:text-slate-900 font-medium"
                        >
                          Edit
                        </button>

                        <button
                          type="button"
                          onClick={() => handleDelete(app.id)}
                          className="text-rose-500 hover:text-rose-700 font-medium"
                        >
                          Delete
                        </button>
                      </div>
                    </div>
                  </div>
                )
              })}
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default Applications