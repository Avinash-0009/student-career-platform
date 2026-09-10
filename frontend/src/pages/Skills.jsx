import { useEffect, useState } from "react"
import { api } from "../services/api"

function Skills() {
  const [skills, setSkills] = useState([])
  const [loading, setLoading] = useState(true)
  const [message, setMessage] = useState("")

  const [form, setForm] = useState({
    name: "",
    proficiency: "",
  })

  const [editingId, setEditingId] = useState(null)

  const getSkills = async () => {
    try {
      const data = await api.get("/skills")
      setSkills(data)
    } catch (error) {
      console.error(error)
      setMessage(error.message || "Failed to load skills")
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    getSkills()
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
      const skillData = {
        name: form.name,
        proficiency: Number(form.proficiency),
      }

      if (editingId) {
        await api.put(`/skills/${editingId}`, skillData)
        setMessage("Skill updated successfully")
      } else {
        await api.post("/skills", skillData)
        setMessage("Skill added successfully")
      }

      setForm({
        name: "",
        proficiency: "",
      })

      setEditingId(null)
      getSkills()
    } catch (error) {
      console.error(error)
      setMessage(error.message || "Failed to save skill")
    }
  }

  const handleEdit = (skill) => {
    setEditingId(skill.id)

    setForm({
      name: skill.name || "",
      proficiency: skill.proficiency ?? "",
    })
  }

  const handleDelete = async (id) => {
    if (!window.confirm("Are you sure you want to delete this skill?")) {
      return
    }

    try {
      await api.delete(`/skills/${id}`)
      setMessage("Skill deleted successfully")
      getSkills()
    } catch (error) {
      console.error(error)
      setMessage(error.message || "Failed to delete skill")
    }
  }

  const cancelEdit = () => {
    setEditingId(null)

    setForm({
      name: "",
      proficiency: "",
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
    <div className="max-w-6xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <div className="mb-8 border-b border-slate-200 pb-5">
        <h1 className="text-2xl font-bold tracking-tight text-slate-900 sm:text-3xl">
          Skills Inventory
        </h1>

        <p className="mt-1 text-sm text-slate-500">
          Track and assess your core competencies and proficiencies.
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
        {/* Form */}
        <div className="lg:col-span-4 bg-white p-6 rounded-xl border border-slate-200/80 shadow-sm sticky top-6">
          <h2 className="text-lg font-semibold text-slate-900 mb-4 pb-2 border-b border-slate-100">
            {editingId ? "Edit Skill" : "Add New Skill"}
          </h2>

          <form onSubmit={handleSubmit} className="space-y-4">
            {/* Skill Name */}
            <div>
              <label
                htmlFor="skill-name"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                Skill Name
              </label>

              <input
                id="skill-name"
                name="name"
                value={form.name}
                onChange={handleChange}
                placeholder="e.g. Go, Kubernetes, React"
                required
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            {/* Proficiency */}
            <div>
              <div className="flex justify-between items-center mb-1">
                <label
                  htmlFor="skill-proficiency"
                  className="block text-xs font-semibold uppercase tracking-wider text-slate-600"
                >
                  Proficiency
                </label>

                <span className="text-sm font-semibold text-indigo-600">
                  {form.proficiency || 0}%
                </span>
              </div>

              <input
                id="skill-proficiency"
                type="range"
                name="proficiency"
                min="0"
                max="100"
                value={form.proficiency || 0}
                onChange={handleChange}
                aria-label="Skill proficiency"
                className="w-full h-2 bg-slate-200 rounded-lg appearance-none cursor-pointer accent-indigo-600"
              />
            </div>

            {/* Buttons */}
            <div className="flex gap-2 pt-2">
              <button
                type="submit"
                className="flex-1 bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-2 px-4 rounded-lg text-sm transition-all shadow-sm"
              >
                {editingId ? "Update Skill" : "Add Skill"}
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

        {/* Skills Cards */}
        <div className="lg:col-span-8">
          {skills.length === 0 ? (
            <div className="text-center py-16 px-4 bg-white border border-dashed border-slate-300 rounded-xl">
              <p className="text-slate-500 text-sm">
                No skills added yet.
              </p>
            </div>
          ) : (
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
              {skills.map((skill) => (
                <div
                  key={skill.id}
                  className="bg-white p-4 rounded-xl border border-slate-200/80 shadow-sm flex flex-col justify-between hover:shadow-md transition-shadow"
                >
                  <div>
                    <div className="flex items-center justify-between mb-2">
                      <span className="font-semibold text-slate-900 text-sm">
                        {skill.name}
                      </span>

                      <span className="text-xs font-bold text-slate-600">
                        {skill.proficiency}%
                      </span>
                    </div>

                    {/* Progress Track */}
                    <div className="w-full bg-slate-100 rounded-full h-2 overflow-hidden mb-3">
                      <div
                        className="bg-indigo-600 h-2 rounded-full transition-all duration-300"
                        style={{
                          width: `${skill.proficiency}%`,
                        }}
                      />
                    </div>
                  </div>

                  <div className="flex justify-end gap-3 pt-2 border-t border-slate-100 text-xs">
                    <button
                      type="button"
                      onClick={() => handleEdit(skill)}
                      className="text-slate-500 hover:text-slate-900 font-medium"
                    >
                      Edit
                    </button>

                    <button
                      type="button"
                      onClick={() => handleDelete(skill.id)}
                      className="text-rose-500 hover:text-rose-700 font-medium"
                    >
                      Delete
                    </button>
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

export default Skills