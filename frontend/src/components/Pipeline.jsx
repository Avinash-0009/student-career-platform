function Pipeline({ stats }) {
  const total = stats.total_applications

  const getPercentage = (value) => {
    if (total === 0) {
      return 0
    }

    return Math.round((value / total) * 100)
  }

  const stages = [
    {
      name: "Applied",
      value: stats.applied,
    },
    {
      name: "Assessment",
      value: stats.assessment,
    },
    {
      name: "Interview",
      value: stats.interview,
    },
    {
      name: "Offer",
      value: stats.offer,
    },
    {
      name: "Rejected",
      value: stats.rejected,
    },
  ]

  return (
    <div className="pipeline">

      <h2>Application Pipeline</h2>

      {stages.map((stage) => (
        <div className="pipeline-stage" key={stage.name}>

          <div className="pipeline-info">
            <span>{stage.name}</span>
            <span>
              {stage.value} ({getPercentage(stage.value)}%)
            </span>
          </div>

          <div className="pipeline-bar">
            <div
              className="pipeline-fill"
              style={{
                width: `${getPercentage(stage.value)}%`,
              }}
            />
          </div>

        </div>
      ))}

    </div>
  )
}

export default Pipeline