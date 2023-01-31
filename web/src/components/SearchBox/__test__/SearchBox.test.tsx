import React from 'react'
import { fireEvent, render, screen, waitFor } from '@testing-library/react'

import SearchBox from '../SearchBox'

const mockedOnClick = jest.fn()

type TestElement = Document | Element | Window | Node

//Used to test if an input has a specific input string
function hasInputValue(e: TestElement, inputValue: string) {
  return screen.getByDisplayValue(inputValue) === e
}

describe('SearchBox', () => {
  it('Test SearchBox render', () => {
    render(
      <SearchBox
        setFilterOption={(string: string, string2: any) =>
          mockedOnClick(string, string2)
        }
      />
    )
    expect(screen.getByTestId('search')).toBeInTheDocument()
  })

  it('Test SearchBox onclick', async () => {
    render(
      <SearchBox
        setFilterOption={(string: string, string2: any) =>
          mockedOnClick(string, string2)
        }
      />
    )

    const search = screen.getByTestId('search')

    fireEvent.click(search)
    fireEvent.change(search, { target: { value: 'project' } })
    fireEvent.keyPress(search, { key: 'Enter', code: 13, charCode: 13 })
    await waitFor(() => {
      expect(mockedOnClick).toHaveBeenCalled()
      expect(hasInputValue(search, 'project')).toBe(true)
    })
  })
})
