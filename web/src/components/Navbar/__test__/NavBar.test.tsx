import React from 'react'
import { fireEvent, render, screen, waitFor } from '@testing-library/react'

import NavBar from '../Navbar'
import Button from '../../Button/Button'

const mockedOnClick = jest.fn()

describe('NavBar', () => {
  window.matchMedia =
    window.matchMedia ||
    function () {
      return {
        matches: false,
        addListener: function () {},
        removeListener: function () {},
      }
    }

  it('Test NavBar renders', () => {
    render(<NavBar />)
    waitFor(() => expect(screen.findByTestId('logo')).toBeInTheDocument())
  })
})
