import React from 'react'
import styles from './CheckboxFilter.module.css'

const CheckboxFilter = ({
  filterName,
  setFilterOption,
  values,
}: {
  filterName: string
  setFilterOption: (value: string, filterName: string) => void
  values: string[]
}) => {
  const getFilterOptions = () => {
    //eslint-disable-next-line
    let options: JSX.Element[] = []
    values.forEach((element) => {
      options.push(
        <div className={styles.checkboxBox}>
          <label>
            <input
              className={styles.checkBox}
              type="checkbox"
              name={filterName}
              value={element}
              onClick={() => setFilterOption(element, filterName)}
              id={element}
            />
            {element}
          </label>
        </div>
      )
    })
    return options
  }

  return (
    <div className={styles.filterBox}>
      <div className={styles.filterTitle}>
        <p>{filterName}</p>
        <div className={styles.line} />
      </div>
      <div className={styles.options}>{getFilterOptions()}</div>
    </div>
  )
}

export default CheckboxFilter
