import React from 'react'
import { fireEvent, render, screen, waitFor } from '@testing-library/react'

import SideBar from '../Sidebar'

describe('SideBar', () => {
  it('Test SideBar closed render', () => {
    render(<SideBar />)
    expect(screen.getByTestId('hamburger')).toBeInTheDocument()
    expect(screen.getByTestId('hamburger')).not.toBeVisible()
  })

  it('Test SideBar openning and renders text', () => {
    const result = render(<SideBar />)

    const hamburgerButton = result.container.querySelector(
      '#react-burger-menu-btn'
    )
    if (hamburgerButton) {
      fireEvent.click(hamburgerButton)

      waitFor(() => {
        expect(screen.getByTestId('hamburger')).toBeVisible()
        expect(screen.findByText('projects')).toBeVisible()
        expect(screen.findByText('about')).toBeVisible()
        expect(screen.findByText('learn')).toBeVisible()
      })
    }
  })
})
