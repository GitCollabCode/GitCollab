import React from 'react'
import { fireEvent, render, screen, waitFor } from '@testing-library/react'

import LandingPage from '../LandingPage'

describe('LandingPage', () => {
  it.skip('Test LandingPage renders', () => {
    render(<LandingPage />)
    expect(screen.findByText('GitCollab')).toBeVisible()
    expect(screen.findByText('projects')).toBeVisible()
    expect(screen.findByText('about')).toBeVisible()
    expect(screen.findByText('learn')).toBeVisible()
    expect(screen.findByText('login')).toBeVisible()
  })
})
