import React from 'react'
import { render, screen, waitFor } from '@testing-library/react'

import LandingPage from '../LandingPage'

describe('LandingPage', () => {
  window.matchMedia =
    window.matchMedia ||
    function () {
      return {
        matches: true,
        addListener: function () {},
        removeListener: function () {},
      }
    }

  it('Test LandingPage renders', () => {
    render(<LandingPage />)
    waitFor(() => expect(screen.findByTestId('logo')).toBeInTheDocument())
  })
})
