import { useEffect, useState } from "react"
import { api } from "../services/api"

function Profile() {
  const [profile, setProfile] = useState({
    university: "",
    degree: "",
    graduation_year: "",
    cgpa: "",
    github: "",
    linkedin: "",
    portfolio: "",
  })

  const [message, setMessage] = useState("")
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const getProfile = async () => {
      try {
        const data = await api.get("/profile")

        setProfile({
          university: data.university || "",
          degree: data.degree || "",
          graduation_year: data.graduation_year || "",
          cgpa: data.cgpa || "",
          github: data.github || "",
          linkedin: data.linkedin || "",
          portfolio: data.portfolio || "",
        })
      } catch (error) {
        console.error(error)
        setMessage(error.message || "Failed to load profile")
      } finally {
        setLoading(false)
      }
    }

    getProfile()
  }, [])

  const handleChange = (e) => {
    setProfile({
      ...profile,
      [e.target.name]: e.target.value,
    })
  }

  const handleSubmit = async (e) => {
    e.preventDefault()

    try {
      await api.put("/profile", {
        university: profile.university,
        degree: profile.degree,
        graduation_year: Number(profile.graduation_year),
        cgpa: Number(profile.cgpa),
        github: profile.github,
        linkedin: profile.linkedin,
        portfolio: profile.portfolio,
      })

      setMessage("Profile updated successfully")
    } catch (error) {
      console.error(error)
      setMessage(error.message || "Failed to update profile")
    }
  }

  if (loading) {
    return (
      <div className="flex h-64 items-center justify-center">
        <div className="h-8 w-8 animate-spin rounded-full border-4 border-indigo-600 border-t-transparent"></div>
      </div>
    )
  }

  return (
    <div className="max-w-4xl mx-auto px-4 sm:px-6 py-8">
      <div className="mb-8 border-b border-slate-200 pb-5">
        <h1 className="text-2xl font-bold tracking-tight text-slate-900 sm:text-3xl">
          Profile Settings
        </h1>

        <p className="mt-1 text-sm text-slate-500">
          Manage your educational background and public developer links.
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

      <form onSubmit={handleSubmit} className="space-y-6">
        {/* Education Section */}
        <div className="bg-white rounded-xl border border-slate-200/80 shadow-sm p-6">
          <h2 className="text-base font-semibold text-slate-900 mb-1">
            Academic Background
          </h2>

          <p className="text-xs text-slate-500 mb-5">
            Your degree and university credentials.
          </p>

          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label
                htmlFor="university"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                University
              </label>

              <input
                id="university"
                name="university"
                value={profile.university}
                onChange={handleChange}
                placeholder="e.g. Stanford University"
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            <div>
              <label
                htmlFor="degree"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                Degree / Major
              </label>

              <input
                id="degree"
                name="degree"
                value={profile.degree}
                onChange={handleChange}
                placeholder="e.g. B.Tech Computer Science"
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            <div>
              <label
                htmlFor="graduation_year"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                Graduation Year
              </label>

              <input
                id="graduation_year"
                type="number"
                name="graduation_year"
                value={profile.graduation_year}
                onChange={handleChange}
                placeholder="2027"
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            <div>
              <label
                htmlFor="cgpa"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                CGPA / Score
              </label>

              <input
                id="cgpa"
                type="number"
                step="0.01"
                name="cgpa"
                value={profile.cgpa}
                onChange={handleChange}
                placeholder="8.50"
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>
          </div>
        </div>

        {/* Social Links Section */}
        <div className="bg-white rounded-xl border border-slate-200/80 shadow-sm p-6">
          <h2 className="text-base font-semibold text-slate-900 mb-1">
            Online Presence
          </h2>

          <p className="text-xs text-slate-500 mb-5">
            Links to your code repositories and online resume.
          </p>

          <div className="space-y-4">
            <div>
              <label
                htmlFor="github"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                GitHub Profile
              </label>

              <input
                id="github"
                name="github"
                value={profile.github}
                onChange={handleChange}
                placeholder="https://github.com/username"
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            <div>
              <label
                htmlFor="linkedin"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                LinkedIn Profile
              </label>

              <input
                id="linkedin"
                name="linkedin"
                value={profile.linkedin}
                onChange={handleChange}
                placeholder="https://linkedin.com/in/username"
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>

            <div>
              <label
                htmlFor="portfolio"
                className="block text-xs font-semibold uppercase tracking-wider text-slate-600 mb-1"
              >
                Portfolio Site
              </label>

              <input
                id="portfolio"
                name="portfolio"
                value={profile.portfolio}
                onChange={handleChange}
                placeholder="https://yourportfolio.dev"
                className="w-full rounded-lg border border-slate-300 px-3.5 py-2 text-sm text-slate-800 placeholder-slate-400 focus:border-indigo-500 focus:outline-none focus:ring-2 focus:ring-indigo-500/20"
              />
            </div>
          </div>
        </div>

        <div className="flex justify-end">
          <button
            type="submit"
            className="bg-indigo-600 hover:bg-indigo-700 text-white font-medium py-2.5 px-6 rounded-lg text-sm transition-all shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2"
          >
            Save Changes
          </button>
        </div>
      </form>
    </div>
  )
}

export default Profile