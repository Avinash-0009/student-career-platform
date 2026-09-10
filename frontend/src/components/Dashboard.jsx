import { useEffect, useState } from "react"
import StatCard from "./StatCard"
import Pipeline from "./Pipeline"
import { api } from "../services/api"

function Dashboard() {
  const [stats, setStats] = useState(null)
  const [error, setError] = useState("")

  useEffect(() => {
    const getDashboard = async () => {
      try {
        const data = await api.get("/dashboard")
        setStats(data)
      } catch (error) {
        console.error(error)
        setError(error.message || "Failed to load dashboard")
      }
    }

    getDashboard()
  }, [])

  if (error) {
    return (
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        <div
          role="alert"
          className="rounded-xl border border-rose-200 bg-rose-50 p-5 text-sm text-rose-700"
        >
          {error}
        </div>
      </div>
    )
  }

  if (!stats) {
    return (
      <div className="flex h-64 items-center justify-center">
        <div className="h-8 w-8 animate-spin rounded-full border-4 border-indigo-600 border-t-transparent"></div>
      </div>
    )
  }

  return (
    <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">

      {/* Header */}
      <div className="mb-8 border-b border-slate-200 pb-5">
        <h1 className="text-2xl font-bold tracking-tight text-slate-900 sm:text-3xl">
          Welcome back 👋
        </h1>

        <p className="mt-1 text-sm text-slate-500">
          Here's an overview of your career progress.
        </p>
      </div>

      {/* Statistics */}
      <div className="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-4 mb-8">

        <StatCard
          title="Applications"
          value={stats.total_applications}
        />

        <StatCard
          title="Projects"
          value={stats.projects}
        />

        <StatCard
          title="Skills"
          value={stats.skills}
        />

        <StatCard
          title="Offers"
          value={stats.offer}
        />

      </div>

      {/* Main Dashboard */}
      <div className="grid grid-cols-1 lg:grid-cols-12 gap-6">

        {/* Pipeline */}
        <div className="lg:col-span-8 bg-white rounded-xl border border-slate-200/80 shadow-sm p-6">
          <Pipeline stats={stats} />
        </div>

        {/* Quick Summary */}
        <div className="lg:col-span-4 bg-white rounded-xl border border-slate-200/80 shadow-sm p-6">

          <h2 className="text-lg font-semibold text-slate-900 mb-1">
            Application Summary
          </h2>

          <p className="text-xs text-slate-500 mb-5">
            Current status of your job applications.
          </p>

          <div className="space-y-4">

            <div className="flex items-center justify-between">
              <span className="text-sm text-slate-600">
                Applied
              </span>

              <span className="text-sm font-semibold text-slate-900">
                {stats.applied}
              </span>
            </div>

            <div className="flex items-center justify-between">
              <span className="text-sm text-slate-600">
                Assessments
              </span>

              <span className="text-sm font-semibold text-slate-900">
                {stats.assessment}
              </span>
            </div>

            <div className="flex items-center justify-between">
              <span className="text-sm text-slate-600">
                Interviews
              </span>

              <span className="text-sm font-semibold text-slate-900">
                {stats.interview}
              </span>
            </div>

            <div className="flex items-center justify-between">
              <span className="text-sm text-slate-600">
                Offers
              </span>

              <span className="text-sm font-semibold text-slate-900">
                {stats.offer}
              </span>
            </div>

            <div className="flex items-center justify-between">
              <span className="text-sm text-slate-600">
                Rejected
              </span>

              <span className="text-sm font-semibold text-slate-900">
                {stats.rejected}
              </span>
            </div>

          </div>

          <div className="mt-6 pt-5 border-t border-slate-100">
            <div className="flex items-center justify-between">
              <span className="text-sm font-medium text-slate-700">
                Total Applications
              </span>

              <span className="text-xl font-bold text-indigo-600">
                {stats.total_applications}
              </span>
            </div>
          </div>

        </div>

      </div>

    </div>
  )
}

export default Dashboard