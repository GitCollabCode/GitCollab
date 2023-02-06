import React, { ReactNode } from 'react'

import styles from './Table.module.css'

export type rowType = {
  id: string
  name: string
  date: string
  total: number
  points: number
  percent: number
  status: string
}

const Table = ({ children }: { children: ReactNode }) => {
  return (
    <>
      <div className={styles.tableBackground}>
        <table>
          <thead></thead>
          <tbody>{children}</tbody>
        </table>
      </div>
    </>
  )
}

export default Table
