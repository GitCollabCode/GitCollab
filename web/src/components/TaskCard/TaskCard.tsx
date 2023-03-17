import React, { useContext, useState } from 'react'
import { TaskProgress, TaskType } from '../../constants/common'
import { UserLoginContext } from '../../context/userLoginContext/userLoginContext'
import style from './TaskCard.module.css'

const TaskCard = ({ task, owner }: { task: TaskType; owner: string }) => {
  const [taskStatus, setTaskStatus] = useState(task.task_status)
  const { isLoggedIn } = useContext(UserLoginContext)
  const getPills = (data: string[]) => {
    //eslint-disable-next-line
    let pills: JSX.Element[] = []
    data.forEach((element, index) => {
      pills.push(
        <div className={style.pill} key={index}>
          {element}
        </div>
      )
    })

    return pills
  }

  const getActionsSectionBasedOnTaskProgress = () => {
    switch (taskStatus) {
      case TaskProgress.TaskStatusUnassigned:
        return (
          <>
            <div className={[style.taskPill, style.red].join(' ')}>
              Unassigned
            </div>
            {isLoggedIn && localStorage.getItem('user') !== owner && (
              <button
                className={[style.taskPill, style.green].join(' ')}
                onClick={() => {
                  setTaskStatus(TaskProgress.TaskStatusAssigned)
                }}
              >
                Take this Task 10 points
              </button>
            )}
          </>
        )
      case TaskProgress.TaskStatusAssigned:
        return (
          <>
            <div className={[style.taskPill, style.yellow].join(' ')}>
              In Progress
            </div>
          </>
        )
      case TaskProgress.TaskStatusCompleted:
        return (
          <>
            <div className={[style.taskPill, style.green].join(' ')}>Done</div>
          </>
        )
      case TaskProgress.TaskStatusChangesRequested:
        return (
          <>
            <div className={[style.taskPill, style.yellow].join(' ')}>
              Changes Requested
            </div>
          </>
        )
      case TaskProgress.TaskStatusDismissed:
        return (
          <>
            <div className={[style.taskPill, style.red].join(' ')}>
              Task Dismissed
            </div>
          </>
        )
      case TaskProgress.TaskStatusReadyToMerge:
        return (
          <>
            <div className={[style.taskPill, style.yellow].join(' ')}>
              Ready to Merge
            </div>
          </>
        )
      case TaskProgress.TaskStatusApproved:
        return (
          <>
            <div className={[style.taskPill, style.yellow].join(' ')}>
              Approved
            </div>
          </>
        )
    }
  }

  const getSkillsOrActionsBasedOnTaskProgress = () => {
    if (task.task_status === TaskProgress.TaskStatusAssigned) {
      return (
        <>
          <button className={[style.taskButton, style.green].join(' ')}>
            Make Pull Request
          </button>
          <button className={[style.taskButton, style.red].join(' ')}>
            Drop
          </button>
        </>
      )
    } else if (task.task_status === TaskProgress.TaskStatusReadyToMerge) {
      return (
        <>
          <button className={[style.taskButton, style.green].join(' ')}>
            Merge
          </button>
          <button className={[style.taskButton, style.red].join(' ')}>
            Drop
          </button>
        </>
      )
    } else {
      return getPills(task.skills)
    }
  }

  return (
    <div className={style.card}>
      <div className={style.taskDescription}>
        <div className={style.title}>{task.task_title}</div>
        <p className={style.description}>{task.task_description}</p>
      </div>
      <div className={style.taskSkills}>
        <div className={style.over}>
          {getSkillsOrActionsBasedOnTaskProgress()}
        </div>
      </div>
      <div className={style.taskActions}>
        {getActionsSectionBasedOnTaskProgress()}
      </div>
    </div>
  )
}

export default TaskCard
