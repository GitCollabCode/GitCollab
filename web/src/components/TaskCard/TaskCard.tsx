import React from 'react'
import { TaskType } from '../../constants/common'
import style from './TaskCard.module.css'

const TaskCard = ({ task }: { task: TaskType }) => {
  console.log(task)
  return (
    <>
      <p className={style.style}>{task.task_title}</p>
      <p>{task.task_description}</p>
      <p>{task.skills}</p>
      <p>{task.task_status}</p>
      <></>
    </>
  )
}

export default TaskCard
