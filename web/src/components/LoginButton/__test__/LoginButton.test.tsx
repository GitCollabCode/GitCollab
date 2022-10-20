import React from 'react'
import { render, screen } from '@testing-library/react'

import Login from '../LoginButton'

const mockedUsedNavigate = jest.fn()
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: () => mockedUsedNavigate,
}))

describe('Login', () => {
  it('Test Render', () => {
    render(<Login />)
    expect(screen.getByText('Login')).toBeInTheDocument()
    expect(screen.getByText('Login with GitHub')).toBeInTheDocument()
  })
})
