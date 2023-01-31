import React from 'react'
import { fireEvent, render, screen, waitFor } from '@testing-library/react'

import CheckboxFilter from '../CheckboxFilter'

const mockedOnClick = jest.fn()

describe('CheckboxFilter', () => {
  it('Test CheckboxFilter render', () => {
    render(
      <CheckboxFilter
        filterName="checkbox"
        setFilterOption={(string: string, string2: any) =>
          mockedOnClick(string, string2)
        }
        values={['string1', 'string2']}
      />
    )
    expect(screen.getByText('string1')).toBeInTheDocument()
    expect(screen.getByText('string2')).toBeInTheDocument()
  })

  it('Test CheckboxFilter onclick', async () => {
    render(
      <CheckboxFilter
        filterName="checkbox"
        setFilterOption={(string: string, string2: any) =>
          mockedOnClick(string, string2)
        }
        values={['string1', 'string2']}
      />
    )

    const option = screen.getByText('string1')

    fireEvent.click(option)
    await waitFor(() => expect(mockedOnClick).toHaveBeenCalled())
  })
})
