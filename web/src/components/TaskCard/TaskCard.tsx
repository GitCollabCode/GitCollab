import React from 'react'
import { TaskCardType } from '../../constants/common'
import style from './TaskCard.module.css'

const TaskCard = ({ task }: { task: TaskCardType }) => {
  console.log(task)
  return (
    <>
      <p className={style.style}>{task.name}</p>
      <p>{task.description}</p>
      <p>{task.assignedTo}</p>
      <p>{task.assignedToImg}</p>
      <p>{task.languages}</p>
      <p>{task.progress}</p>
    </>
  )
}

export default TaskCard
