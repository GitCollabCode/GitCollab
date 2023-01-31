import React from 'react'
import styles from './SearchBox.module.css'

const SearchBox = ({
  setFilterOption,
}: {
  setFilterOption: (value: string, filterName: string) => void
}) => {
  return (
    <input
      className={styles.searchBox}
      type="text"
      placeholder="Search"
      name="search"
      onChange={(e: React.ChangeEvent<HTMLInputElement>) =>
        setFilterOption(e.target.value, 'search')
      }
      data-testid="search"
    />
  )
}

export default SearchBox
